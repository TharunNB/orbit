package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"orbit/internal/models"

	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List active agent runtimes",
	Long:  `Display a table of all active background Orbit agent runtimes with their PIDs, status, and start times.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemonPort := "54321"
		daemonURL := fmt.Sprintf("http://localhost:%s/status", daemonPort)

		resp, err := http.Get(daemonURL)
		if err != nil {
			fmt.Println("❌ Orbit Daemon is not running.")
			fmt.Println("Tip: Start an agent with 'orbit run <agent-name>' to auto-start the daemon.")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("❌ Daemon returned an unexpected status: %d\n", resp.StatusCode)
			return
		}

		var list []*models.RuntimeMetaData
		if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
			log.Fatalf("Failed to parse status response: %v", err)
		}

		if len(list) == 0 {
			fmt.Println("ℹ️  No active agent runtimes are currently running.")
			return
		}

		// Print premium-styled header and table
		fmt.Println()
		fmt.Printf("🛰️  \033[1m%-25s %-10s %-12s %-20s\033[0m\n", "AGENT NAME", "PID", "STATUS", "STARTED AT")
		fmt.Println("\033[90m----------------------------------------------------------------------\033[0m")
		for _, item := range list {
			startedStr := item.StartedAt.Format("2006-01-02 15:04:05")
			// Add nice green color for running status
			statusStr := fmt.Sprintf("\033[32m%s\033[0m", item.Status)
			fmt.Printf("   %-25s %-10d %-12s %-20s\n", item.Name, item.PID, statusStr, startedStr)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
}
