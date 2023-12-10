package approval_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/approval-process/common"
	"github.com/deepaksinghvi/approval-process/common/dto"
	"go.temporal.io/sdk/activity"
	"time"
)

func CartFinallization(ctx context.Context, input dto.Cart, userID string) (dto.Cart, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Adding Item to cart")
	totalAmount := float32(0.0)
	for _, v := range input.ItemRequests {
		totalAmount += v.Price * float32(v.Quantity)
	}
	cart := dto.Cart{
		ItemRequests:    input.ItemRequests,
		RequestedByUser: common.UserMap[userID],
		TotalAmount:     totalAmount,
	}
	time.Sleep(time.Second * 3)
	logger.Info(fmt.Sprintf("Cart Ready for approval %v", cart))
	return cart, nil
}
func ManagerApprovalActivity(ctx context.Context, cart dto.CartLocked, manager dto.User) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ManagerApprovalActivity Process Initiated")
	approval := dto.Approval{
		Status:       "APPROVED",
		ApproverName: manager.UserName,
	}
	time.Sleep(time.Second * 3)

	cart.ManagerApproval = append(cart.ManagerApproval, approval)
	logger.Info(fmt.Sprintf("ManagerApprovalActivity Process completed %v", cart))
	return cart, nil
}

func ProcurementManagerApprovalActivity(ctx context.Context, cart dto.CartLocked, procurementManager dto.User) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ProcurementManagerApprovalActivity Process Initiated")
	approval := dto.Approval{
		Status:       "APPROVED",
		ApproverName: procurementManager.UserName,
	}
	time.Sleep(time.Second * 3)
	cart.ManagerApproval = append(cart.ManagerApproval, approval)
	logger.Info(fmt.Sprintf("ProcurementManagerApprovalActivity Process completed %v", cart))
	return cart, nil
}

func OrderInitiatedActivity(ctx context.Context, cart dto.CartLocked) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderInitiatedActivity Process Initiated")
	cart.OrderInitiated = "STARTED"
	time.Sleep(time.Second * 3)
	logger.Info(fmt.Sprintf("OrderInitiatedActivity Process completed %v", cart))
	return cart, nil
}
