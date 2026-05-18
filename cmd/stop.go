package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [agent-name]",
	Short: "Stop an active agent runtime",
	Long:  `Send a request to the Orbit daemon to gracefully stop the background runtime process of a specified agent.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentName := args[0]
		daemonPort := "54321"
		daemonURL := fmt.Sprintf("http://localhost:%s/stop?name=%s", daemonPort, agentName)

		fmt.Printf("Stopping agent '%s'...\n", agentName)
		resp, err := http.Post(daemonURL, "application/json", nil)
		if err != nil {
			fmt.Println("❌ Orbit Daemon is not running.")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("✅ Success: Agent '%s' stopped successfully.\n", agentName)
		} else if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("❌ Error: Agent '%s' is not active or running.\n", agentName)
			fmt.Println("Tip: Use 'orbit ps' to check running agents.")
		} else {
			fmt.Printf("❌ Error: Failed to stop agent '%s'. Daemon returned status: %d\n", agentName, resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
