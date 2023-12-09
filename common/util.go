package common

import "github.com/deepaksinghvi/approval-process/common/dto"

const (
	ApprovalTaskQueueName     = "default"
	ApprovalWorkflowQueryType = "approval-workflow-state-query"
)

// Build Items

var ItemMap = map[int]dto.ItemRequest{
	laptop.ItemNo:  laptop,
	monitor.ItemNo: monitor,
	mouse.ItemNo:   mouse,
}

var laptop = dto.ItemRequest{
	ItemNo:   1,
	ItemName: "Laptop",
	Quantity: 1,
	Price:    100000,
}

var monitor = dto.ItemRequest{
	ItemNo:   2,
	ItemName: "Monitor",
	Quantity: 1,
	Price:    20000,
}

var mouse = dto.ItemRequest{
	ItemNo:   3,
	ItemName: "Mouse",
	Quantity: 1,
	Price:    6000,
}

// Build Users

var UserMap = map[string]dto.User{
	engineer1.UserID:          engineer1,
	engineer2.UserID:          engineer1,
	procurementManager.UserID: procurementManager,
	ItManager.UserID:          ItManager,
	director.UserID:           director,
}

var director = dto.User{
	UserID:      "d1",
	UserName:    "Director1",
	Designation: "Director",
	ManagerID:   "",
}

var itLead = dto.User{
	UserID:      "itl1",
	UserName:    "ITLead1",
	Designation: "IT-Lead",
	ManagerID:   "d1",
}

var ItManager = dto.User{
	UserID:      "itm1",
	UserName:    "Manager1",
	Designation: "IT-Manager",
	ManagerID:   "d1",
}

var procurementManager = dto.User{
	UserID:      "pm1",
	UserName:    "ProcurementManager1",
	Designation: "Procurement-Manager",
	ManagerID:   "d1",
}

var engineer2 = dto.User{
	UserID:      "user2",
	UserName:    "User2",
	Designation: "Engineer",
	ManagerID:   "itm1",
}

var engineer1 = dto.User{
	UserID:      "user1",
	UserName:    "User1",
	Designation: "Engineer",
	ManagerID:   "itm1",
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type WorkflowObject struct {
	WorkflowID string `json:"workflow_id"`
	RunID      string `json:"run_id"`
}

type State string
type QueryResult struct {
	State   State
	Content string
}
