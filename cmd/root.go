package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	erp_clients "update-customer-image/erp-clients"
)
var erpClient *erp_clients.ErpClient

var rootCmd = &cobra.Command{
	Use: "help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is migrate user example")
	},
}

func init() {
	appID := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	baseURL := os.Getenv("BASE_URL")
	fmt.Printf(`[DB] Run with config, {"app_id":"%s", "app_secret":"%s", "base_url":"%s"`, appID, appSecret, baseURL)
	fmt.Println()
	erpClient = erp_clients.NewERPClient(appID, appSecret, baseURL)

	rootCmd.AddCommand(getUserCommand)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(attachFieCmd)
}

func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
