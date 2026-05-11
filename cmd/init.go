/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"orbit/internal/tui"
	"orbit/internal/workspace"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [agent-name]",
	Short: "Initialize a new Orbit agent workspace",
	Run: func(cmd *cobra.Command, args []string) {

		agentName := args[0]

		//Select the model for the current agent
		selected := tui.Start()

		err := workspace.Initialize(agentName, selected.Name)
		if err != nil {
			fmt.Printf("Error inititalizing agent: %v\n", err)
			return
		}

		fmt.Printf("Succesfully initialized agent '%s' with model '%s' \n", agentName, selected.Name)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getModelURl(model string) string {
	models := map[string]string{
		"tinyllama": "https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF/resolve/main/tinyllama-1.1b-chat-v1.0.Q2_K.gguf",
	}

	return models[model]
}
