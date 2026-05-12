/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"orbit/internal/daemon"
	"orbit/internal/process"

	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "A brief description of your command",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Orbit daemon",
	Run: func(cmd *cobra.Command, args []string) {
		reg := daemon.NewRegistry()
		mgr := &process.Manager{Updater: reg}
		srv := daemon.NewServer(reg, mgr)

		if err := srv.Start("54321"); err != nil {
			log.Fatalf("Daemon failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
