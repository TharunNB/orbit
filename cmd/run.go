/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"orbit/internal/process"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [agent-name]",
	Short: "Run an agent",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentName := args[0]
		daemonPort := "54321"
		daemonURL := fmt.Sprintf("http://localhost:%s", daemonPort)

		_, err := http.Get(daemonURL + "/status")
		if err != nil {
			fmt.Println("Orbit Daemon not running. Starting it in the background")

			daemonProc := exec.Command(os.Args[0], "daemon", "start")

			daemonProc.SysProcAttr = &syscall.SysProcAttr{}

			process.DetachProcess(daemonProc.SysProcAttr)

			if err := daemonProc.Start(); err != nil {
				log.Fatalf("Failed to auto-start daemon: %v", err)
			}

			time.Sleep(1 * time.Second)
		}

		fmt.Printf("Sending run request for agent: %s...\n", agentName)
		resp, err := http.Post(fmt.Sprintf("%s/start?name=%s", daemonURL, agentName), "application/json", nil)

		if err != nil {
			log.Fatalf("Failed to communicate with daemon: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusAccepted {
			fmt.Printf("Success: Agent '%s' is now running in the background.\n", agentName)
			fmt.Println("Use 'orbit ps' to check status or 'orbit logs' to see output.")
		} else {
			fmt.Printf("Daemon returned error: %d\n", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
