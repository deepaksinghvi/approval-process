package approval_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/approval-process/common/dto"
	"go.temporal.io/sdk/activity"
	"time"
)

func CartFinallization(ctx context.Context, input dto.Cart) (dto.Cart, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Adding Item to cart")
	totalAmount := float32(0.0)
	for _, v := range input.ItemRequests {
		totalAmount += v.Price * float32(v.Quantity)
	}
	cart := dto.Cart{
		ItemRequests:    input.ItemRequests,
		RequestedByUser: input.RequestedByUser,
		TotalAmount:     totalAmount,
	}
	time.Sleep(time.Second * 4)
	logger.Info(fmt.Sprintf("Cart Ready for approval %v", cart))
	return cart, nil
}
func ManagerApproval(ctx context.Context, cart dto.CartLocked, manager dto.User) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ManagerApproval Process Initiated")
	approval := dto.Approval{
		Status:       "APPROVED",
		ApproverName: manager.UserName,
	}
	time.Sleep(time.Second * 4)

	cart.ManagerApproval = append(cart.ManagerApproval, approval)
	logger.Info(fmt.Sprintf("ManagerApproval Process completed %v", cart))
	return cart, nil
}

func ITLeadApproval(ctx context.Context, cart dto.CartLocked, manager dto.User) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ITLeadApproval Process Initiated")
	approval := dto.Approval{
		Status:       "APPROVED",
		ApproverName: manager.UserName,
	}
	time.Sleep(time.Second * 3)

	cart.ManagerApproval = append(cart.ManagerApproval, approval)
	logger.Info(fmt.Sprintf("ManagerApproval Process completed %v", cart))
	return cart, nil
}

func ProcurementManagerApproval(ctx context.Context, cart dto.CartLocked, procurementManager dto.User) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ProcurementManagerApproval Process Initiated")
	approval := dto.Approval{
		Status:       "APPROVED",
		ApproverName: procurementManager.UserName,
	}
	time.Sleep(time.Second * 3)
	cart.ManagerApproval = append(cart.ManagerApproval, approval)
	logger.Info(fmt.Sprintf("ProcurementManagerApproval Process completed %v", cart))
	return cart, nil
}

func OrderInitiation(ctx context.Context, cart dto.CartLocked) (dto.CartLocked, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderInitiation Process Initiated")
	cart.OrderInitiated = "STARTED"
	time.Sleep(time.Second * 3)
	logger.Info(fmt.Sprintf("OrderInitiation Process completed %v", cart))
	return cart, nil
}
