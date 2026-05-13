package workspace

import (
	"fmt"
	"orbit/internal/models"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Initialize(agentName string, modelName string) error {
	if agentName == "" {
		return fmt.Errorf("agentName cannot be empty")
	}

	if err := os.MkdirAll(agentName, 0755); err != nil {
		return fmt.Errorf("failed to create agent directory: %w", err)
	}

	//creation of subdirectories
	subDirs := []string{"workspace", "logs", "memory"}
	for _, d := range subDirs {
		path := filepath.Join(agentName, d)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", d, err)
		}
	}

	//Creation of .orbit for logs
	dotOrbitPath := filepath.Join(agentName, ".orbit")
	if err := os.MkdirAll(dotOrbitPath, 0700); err != nil {
		return fmt.Errorf("failed to create .orbit directory: %w", err)
	}

	//orbit.yaml config 
	cfg := models.OrbitConfig{
		Name: agentName,
	}
	cfg.Model.Provider = "ollama"
	cfg.Model.Name = modelName

	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	configPath := filepath.Join(agentName, "orbit.yaml")
	if err := os.WriteFile(configPath, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to wrtie orbit.yaml: %w", err)
	}

	return nil
