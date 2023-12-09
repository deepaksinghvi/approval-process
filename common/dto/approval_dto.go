package dto

type Cart struct {
	ItemRequests    []ItemRequest `json:"item_requests"`
	RequestedByUser User          `json:"requested_by_user"`
	TotalAmount     float32       `json:"total_amount"`
}

type CartLocked struct {
	ItemRequests               []ItemRequest `json:"item_requests"`
	RequestedByUser            User          `json:"requested_by_user"`
	TotalAmount                float32       `json:"total_amount"`
	ManagerApproval            []Approval    `json:"manager_approval"`
	ProcurementManagerApproval Approval      `json:"procurement_manager_approval"`
	OrderInitiated             string        `json:"order_initiated"`
}

type Approval struct {
	Status       string `json:"status"`
	ApproverName string `json:"approver_name"`
}
type ItemRequest struct {
	ItemNo   int     `json:"item_no"`
	ItemName string  `json:"item_name"`
	Quantity int     `json:"quantity"` // default and max is always 1
	Price    float32 `json:"price"`
}

type User struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Designation string `json:"designation"` // Engineer, IT-Manager, ProcurementManager, Director
	ManagerID   string `json:"manager_id"`  // manager UserID
}
