/*
Copyright © 2023 NAME HERE deepak.singhvi@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// approvalappCmd represents the approvalapp command
var approvalappCmd = &cobra.Command{
	Use:   "approvalapp",
	Short: "approvalapp",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("approvalapp called")
	},
}

func init() {
	rootCmd.AddCommand(approvalappCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// approvalappCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// approvalappCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
