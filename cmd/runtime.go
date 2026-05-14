/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// runtimeCmd represents the runtime command
var runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "Agent runtime commands",
	Long:  `Manage and execute Orbit agent runtimes.`,
}

func init() {
	rootCmd.AddCommand(runtimeCmd)
	runtimeCmd.AddCommand(rExecuteCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runtimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runtimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
