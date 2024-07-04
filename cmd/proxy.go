/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// proxyCmd represents the sha command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "copies my proxy credentials to clipboard",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		username := os.Getenv("PROXY_USERNAME")
		password := os.Getenv("PROXY_PASSWORD")
		fmt.Println("Username:", username)
		clipboard.WriteAll(username)
		fmt.Println("Username copied to clipboard")
		time.Sleep(2 * time.Second)
		clipboard.WriteAll(password)
		fmt.Println("Password copied to clipboard")
	},

}

func init() {
	rootCmd.AddCommand(proxyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proxyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

