/*
Copyright © 2023 Deepak Singhvi deepak.singhvi@gmail.com
*/
package main

import (
	"github.com/deepaksinghvi/approval-process/approval_worker"
	"github.com/deepaksinghvi/approval-process/cmd"
	"time"
)

func main() {
	//cmd.Execute()
	go approval_worker.ApprovalWorkflowWorkerStarter()
	time.Sleep(time.Second * 5)
	cmd.WebAppStarter()
}
