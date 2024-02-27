/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var verifyFlag bool

func readUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Remove newline character from the input
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input, nil
}

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "add, commit, and push all",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if(verifyFlag){
			out, err := exec.Command("git", "status").Output()
			cobra.CheckErr(err)

			
			fmt.Println(string(out))
			fmt.Print("Do you want to proceed? (Y/n): ")
			response, err := readUserInput()
			cobra.CheckErr(err)

			if response != "y" {
				fmt.Println("Exiting...")
				os.Exit(0)
			}
		}
		_, err := exec.Command("git", "add", ".").Output()
		cobra.CheckErr(err)

		_, errCommit := exec.Command("git", "commit", "-m", args[0]).Output()
		cobra.CheckErr(errCommit)

		output, errPush := exec.Command("git", "push").Output()
		cobra.CheckErr(errPush)
		fmt.Println(string(output))
	},
}

func init() {
	GitCmd.AddCommand(pushCmd)
	// Here you will define your flags and configuration settings.
	pushCmd.Flags().BoolVarP(&verifyFlag,"verify", "v", false, "Verify files")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
