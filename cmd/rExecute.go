/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"orbit/internal/models"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// rExecuteCmd represents the rExecute command
var rExecuteCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute the agent runtime engine",
	RunE:  runExecute,
}

func init() {
	rExecuteCmd.Flags().StringP("name", "n", "", "Agent name to run")
	rExecuteCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rExecuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rExecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runExecute(cmd *cobra.Command, args []string) error {
	agentName, _ := cmd.Flags().GetString("name")

	fmt.Printf("Orbit Runtime Starting for agent: %s\n", agentName)

	config, err := loadConfig()
	if err != nil {
		return fmt.Errorf("failed to laod orbit.yaml: %w", err)
	}

	ollamaClient, err := api.ClientFromEnvironment()
	if err != nil {
		return fmt.Errorf("failed to create Ollama client: %w", err)
	}

	if err := checkModelReady(ollamaClient, config.Model.Name); err != nil {
		fmt.Printf("⚠️ Model check warning (continuing in offline/testing mode): %v\n", err)
	}

	return startAgentLoop(config, ollamaClient)
}

func loadConfig() (*models.OrbitConfig, error) {
	data, err := os.ReadFile("orbit.yaml")
	if err != nil {
		return nil, fmt.Errorf("orbit.yaml not found in directory")
	}

	var cfg models.OrbitConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid orbit.yaml: %w", err)
	}

	return &cfg, nil
}

func checkModelReady(client *api.Client, modelName string) error {
	ctx := context.Background()
	_, err := client.Show(ctx, &api.ShowRequest{Model: modelName})
	if err != nil {
		return fmt.Errorf("model '%s' is not ready. Run: ollama pull %s", modelName, modelName)
	}
	fmt.Printf("✅ Model '%s' is ready to use\n", modelName)
	return nil
}

func startAgentLoop(cfg *models.OrbitConfig, client *api.Client) error {
	fmt.Println("🔄 Agent runtime loop started. Press Ctrl+C to stop.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sigChan:
			fmt.Println("\nAgent shutting down gracefully...")
			return nil

		case <-ticker.C:
			fmt.Printf(" [%s] Agent is running...\n", cfg.Name)
		}
	}
}
