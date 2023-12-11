package approval_worker

import (
	"context"
	"github.com/deepaksinghvi/approval-process/approval_workflow"
	"github.com/deepaksinghvi/approval-process/common"
	"github.com/deepaksinghvi/approval-process/common/dto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func ApprovalWorkflowWorkerStarter() {
	c, err := client.Dial(client.Options{}) // local
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, common.ApprovalTaskQueueName, worker.Options{})

	// This approval_worker hosts both Workflow and Activity functions.
	w.RegisterWorkflowWithOptions(approval_workflow.ApprovalWorkflow, workflow.RegisterOptions{Name: "ApprovalWorkflow"})
	w.RegisterActivityWithOptions(approval_workflow.CartFinallization, activity.RegisterOptions{Name: "CartFinallization"})
	w.RegisterActivityWithOptions(approval_workflow.ManagerApproval, activity.RegisterOptions{Name: "ManagerApproval"})
	w.RegisterActivityWithOptions(approval_workflow.ITLeadApproval, activity.RegisterOptions{Name: "ITLeadApproval"})
	w.RegisterActivityWithOptions(approval_workflow.ProcurementManagerApproval, activity.RegisterOptions{Name: "ProcurementManagerApproval"})
	w.RegisterActivityWithOptions(approval_workflow.OrderInitiation, activity.RegisterOptions{Name: "OrderInitiation"})

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
	log.Info("Started Approval Workflow Worker")

}

func StartApprovalWorkflowExecution(input dto.Cart) (string, string, error) {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{}) // local
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
		return "", "", err
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        uuid.New().String(),
		TaskQueue: common.ApprovalTaskQueueName,
	}
	log.Info("Starting Approval Workflow for input: %v", input)
	log.Printf("Starting Approval Workflow for input: %v", input)

	we, err := c.ExecuteWorkflow(context.Background(), options, approval_workflow.ApprovalWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
		return "", "", err
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
	return we.GetID(), we.GetRunID(), nil
}

func QueryWorkflow(workflowID, runID, queryType string) (string, error) {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{}) // local
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
		return "", err
	}

	defer c.Close()
	response, err := c.QueryWorkflow(context.Background(), workflowID, runID, queryType)
	if err != nil {
		log.Printf("Unable to send singla to Workflow: %s with error: %s", workflowID, err.Error())
		return "", err
	}
	var result string
	err = response.Get(&result)
	if err != nil {
		log.Printf("Unable to send singla to Workflow: %s with error: %s", workflowID, err.Error())
		return "", err
	}
	log.Info("Received Query result. Result: " + result)
	log.Println("Received Query result. Result: " + result)
	return result, nil
}
