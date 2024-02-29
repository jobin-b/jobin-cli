/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package stock

import (
	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var StockCmd = &cobra.Command{
	Use:   "stock",
	Short: "Preform actions on stocks",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("stock called")
	// },
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
