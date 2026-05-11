/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del [agent-name]",
	Short: "Remove agent and its workspace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentName := args[0]

		fmt.Printf("Are sure you want to delete '%s' and its workspace?")
		var confirm string
		fmt.Scanln(&confirm)

		if confirm != "y" && confirm != "Y" {
			fmt.Println("Cancelled Deletion.")
			return
		}

		err := os.RemoveAll(agentName)
		if err != nil {
			fmt.Printf("Error deleting agent: %v \n", err)
		}

		fmt.Printf("Agent '%s' succesfully removed.\n", agentName)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
