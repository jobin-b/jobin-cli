/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)
func convertSSHToHTTPS(sshURL string) (string, error) {
	if strings.HasPrefix(sshURL, "git@github.com:") {
		// Replace "git@github.com:" with "https://github.com/"
		httpsURL := "https://github.com/" + strings.TrimPrefix(sshURL, "git@github.com:")
		return httpsURL, nil
	}
	return "", fmt.Errorf("not a valid GitHub SSH URL. Try using --https flag")
}

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens remote github repository in browser",
	Long: `Opens the github repository in the default web browser. Assumes 
			that the remote uses ssh url. Use --https flag if the remote 
			uses https url.`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
		cobra.CheckErr(err)
		var remoteURL string
		if(cmd.Flag("https").Value.String() == "true"){
			if(!strings.HasPrefix(string(out), "https://")){
				cobra.CheckErr(fmt.Errorf("not a valid GitHub HTTPS URL. Try using --https flag"))
			}
			remoteURL = string(out)
		} else {
			remoteURL, err = convertSSHToHTTPS(string(out))
			cobra.CheckErr(err)
		}

		_, err = exec.Command("open", remoteURL).Output()
		cobra.CheckErr(err)
	},
}

func init() {
	GitCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.
	openCmd.Flags().Bool("https", false, "Repository uses https remote url")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
