/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package stock

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/sjson"
)

// buyCmd represents the buy command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Buys one stock of each stock ticker provided from the specified brokerage",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// home, err := os.UserHomeDir()
		// cobra.CheckErr(err)

		switch args[0] {
		case "wf":
			runWellsBuy(args)
		default:
			fmt.Println("Invalid brokerage")
		}
	},
}

func runWellsBuy(args []string){
	link := viper.GetString("link_path") + "/wf_buy.lnk"
	macros := viper.GetString("macros_path") + "/wf_buy.json"

	fileContent, err := os.ReadFile(macros)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Modify the value in Commands[5] using sjson
	updatedJSON, err := sjson.SetBytes(fileContent, "Commands.5.Target", "return ['" + strings.Join(args[1:], "', '") + "'];")
	if err != nil {
		fmt.Println("Error updating JSON:", err)
		return
	}

	// Write the updated JSON back to the file
	err = os.WriteFile(macros, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = exec.Command("open", link).Start()
	cobra.CheckErr(err)
}

func init() {
	StockCmd.AddCommand(buyCmd)
	buyCmd.SetUsageTemplate("Usage:\n  jobin stock buy [brokerage] [...stock tickers]\n\n" +
			"Brokerages:\n  wf   Wells Fargo\n\n" +
			"Flags:\n  -h, --help   help for buy\n")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
