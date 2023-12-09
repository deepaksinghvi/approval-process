/*
Copyright © 2023 NAME HERE deepak.singhvi@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/deepaksinghvi/approval-process/api"
	_ "github.com/deepaksinghvi/approval-process/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// webappCmd represents the webapp command
var webappCmd = &cobra.Command{
	Use:   "webapp",
	Short: "Approval Process Web App",
	Long:  `Approval Process Web App to interact with Temporal Workflow Engine Worker.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Approval Process webapp")
		WebAppStarter()
	},
}

func WebAppStarter() {
	router := setupRouter()
	fmt.Println("Approval Workflow started, press ctrl+c to terminate...")
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Approval Workflow app apis
	router.POST("/procurement-approval", api.ProcurementApproval)
	router.GET("/procurement-approval/:workflow_id/:run_id", api.GetApprovalWorkflowStatus)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return router
}

func init() {
	rootCmd.AddCommand(webappCmd)
}
