/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add alias for current path",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		file, err := os.OpenFile(home+"/.bash_aliases", os.O_APPEND|os.O_WRONLY, 0644)
		cobra.CheckErr(err)
		defer file.Close()

		path, err := os.Getwd()
		cobra.CheckErr(err)

		_, err = file.WriteString(fmt.Sprintf("alias %s='cd %s'\n", args[0], path))
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
