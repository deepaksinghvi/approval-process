package approval_workflow

import (
	"fmt"
	"github.com/deepaksinghvi/approval-process/common"
	"github.com/deepaksinghvi/approval-process/common/dto"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"strings"
	"time"
)

func ApprovalWorkflow(ctx workflow.Context, cartInput dto.Cart) error {
	logger := workflow.GetLogger(ctx)
	currentState := "approval_workflow_submitted"
	err := workflow.SetQueryHandler(ctx, common.ApprovalWorkflowQueryType, func() (string, error) {
		return currentState, nil
	})
	if err != nil {
		currentState = "failed to register query handler"
		return err
	}
	cartOutput := dto.CartLocked{}

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        0, // unlimited retries
		NonRetryableErrorTypes: []string{"InvalidInputError", "ExternalIntegrationError"},
	}
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	logger.Info("Approval Workflow started")
	workflow.Sleep(ctx, time.Second*10) // approval_workflow can be put on sleep

	currentState = "cart_finalization_started"
	err = workflow.ExecuteActivity(ctx, CartFinallization, cartInput).Get(ctx, &cartOutput)
	if err != nil {
		logger.Error("CartFinallization failed", err.Error())
		return err
	}

	currentState = "manager_approval_started"
	var futures []workflow.Future

	managerLevel1Future := workflow.ExecuteActivity(ctx, ITLeadApproval, cartOutput, common.UserMap["itl1"])
	if err != nil {
		logger.Error("ITLeadApproval ITLead failed", err.Error())
		return err
	}
	futures = append(futures, managerLevel1Future)
	managerLevel2Future := workflow.ExecuteActivity(ctx, ManagerApproval, cartOutput, common.UserMap["itm1"])
	if err != nil {
		logger.Error("ManagerApproval ITManager failed", err.Error())
		return err
	}
	futures = append(futures, managerLevel2Future)

	approvalResult, err := getManagerApprovals(ctx, cartOutput, futures)
	if err != nil {
		logger.Error("Approval Workflow Manager Approval failed", err.Error())
		return err
	}
	approvedResult := dto.CartLocked{}
	if cartOutput.TotalAmount > 10000 {
		currentState = "procurement_manager_approval_started"
		err = workflow.ExecuteActivity(ctx, ProcurementManagerApproval, approvalResult, common.UserMap["pm1"]).Get(ctx, &approvedResult)
		if err != nil {
			logger.Error("ProcurementManagerApproval failed", err.Error())
			return err
		}
	}
	currentState = "approval_completed"

	err = workflow.ExecuteActivity(ctx, OrderInitiation, approvedResult).Get(ctx, &approvedResult)
	if err != nil {
		logger.Error("OrderInitiation failed", err.Error())
		return err
	}

	currentState = "order_initiation_" + strings.ToLower(approvedResult.OrderInitiated)

	logger.Info(fmt.Sprintf("Workflow result %v", approvedResult))
	return nil
}

func getManagerApprovals(ctx workflow.Context, cartOutput dto.CartLocked, futures []workflow.Future) (dto.CartLocked, error) {
	approvalResult := dto.CartLocked{
		ItemRequests:               cartOutput.ItemRequests,
		RequestedByUser:            cartOutput.RequestedByUser,
		TotalAmount:                cartOutput.TotalAmount,
		ManagerApproval:            nil,
		ProcurementManagerApproval: dto.Approval{},
		OrderInitiated:             "NOT_APPLICABLE",
	}

	for _, future := range futures {
		tempApprovalResult := dto.CartLocked{}
		err := future.Get(ctx, &tempApprovalResult)
		if err != nil {
			return approvalResult, err
		}
		for _, v := range tempApprovalResult.ManagerApproval {
			approvalResult.ManagerApproval = append(approvalResult.ManagerApproval, v)
		}
	}
	return approvalResult, nil
}
