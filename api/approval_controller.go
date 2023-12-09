package api

import (
	"github.com/deepaksinghvi/approval-process/approval_worker"
	"github.com/deepaksinghvi/approval-process/common"
	"github.com/deepaksinghvi/approval-process/common/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ProcurementApproval godoc
// @Summary Procurement Approval For Cart Items
// @ID add-item-to-cart
// @Tags approval-process
// @Accept json
// @Produce json
// @Param cartInput body dto.Cart true "Cart Items"
// @Success 200 {object} common.WorkflowObject
// @Failure 500 {object} common.HTTPError
// @Router /procurement-approval [post]
func ProcurementApproval(c *gin.Context) {
	var cartInput dto.Cart
	// Call BindJSON to bind the received JSON to input.
	if err := c.ShouldBindJSON(&cartInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workflowID, runID, err := approval_worker.StartApprovalWorkflowExecution(cartInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
	}
	log.Infof("Cart Approval Initiated, WorkflowID: %s, RunID:%s", workflowID, runID)

	c.JSON(http.StatusOK, gin.H{"data": common.WorkflowObject{
		WorkflowID: workflowID,
		RunID:      runID,
	}})
}

// GetApprovalWorkflowStatus godoc
// @Summary Get Procurement Approval Workflow state
// @ID get-procurement-workflow-status
// @Tags approval-process
// @Accept json
// @Produce json
// @Param workflow_id path string true "Workflow ID"
// @Param run_id path string true "Run ID"
// @Success 200 {object} common.QueryResult
// @Failure 500 {object} common.HTTPError
// @Router /procurement-approval/{workflow_id}/{run_id} [get]
func GetApprovalWorkflowStatus(c *gin.Context) {
	workflowID := c.Param("workflow_id")
	runID := c.Param("run_id")
	log.Infof("GetApprovalWorkflowStatus received for, WorkflowID: %s, RunID:%s", workflowID, runID)
	result, err := approval_worker.QueryWorkflow(workflowID, runID, common.ApprovalWorkflowQueryType)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"workflow state": result})
}
