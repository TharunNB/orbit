package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs [agent-name]",
	Short: "View logs for an agent runtime",
	Long:  `View the execution logs for a specified agent runtime. Use the -f or --follow flag to stream logs in real-time.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentName := args[0]
		follow, _ := cmd.Flags().GetBool("follow")

		// Resolve workspace path mismatch
		workspacePath := "./agents/" + agentName
		if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
			workspacePath = "./" + agentName
		}

		logFilePath := filepath.Join(workspacePath, "logs", "agent.log")

		// Check if log file exists
		if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
			fmt.Printf("❌ No logs found for agent '%s' (looked at: %s)\n", agentName, logFilePath)
			fmt.Println("Tip: Start the agent first with 'orbit run <agent-name>'.")
			return
		}

		fmt.Printf("\033[90m🛰️  Streaming logs for agent '%s' (press Ctrl+C to exit)...\033[0m\n\n", agentName)

		if follow {
			err := tailFile(logFilePath)
			if err != nil {
				log.Fatalf("❌ Error streaming logs: %v", err)
			}
		} else {
			content, err := os.ReadFile(logFilePath)
			if err != nil {
				log.Fatalf("❌ Failed to read log file: %v", err)
			}
			fmt.Print(string(content))
		}
	},
}

func tailFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Print existing content first
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}

	// Tail new content
	buffer := make([]byte, 2048)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			os.Stdout.Write(buffer[:n])
		}
		if err == io.EOF {
			time.Sleep(150 * time.Millisecond)
			continue
		}
		if err != nil {
			return err
		}
	}
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolP("follow", "f", false, "Stream log output in real-time")
}
