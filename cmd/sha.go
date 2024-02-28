/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// shaCmd represents the sha command
var shaCmd = &cobra.Command{
	Use:   "sha",
	Short: "sha256 input",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sha := sha256.Sum256([]byte(args[0]))
		fmt.Println(hex.EncodeToString(sha[:]))
	},
}

func init() {
	rootCmd.AddCommand(shaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
