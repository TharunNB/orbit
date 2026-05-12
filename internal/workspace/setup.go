package workspace

import (
	"orbit/internal/models"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Initialize(agentName string, modelName string) error {
	subDirs := []string{"workspace", "logs", "memory"}
	for _, d := range subDirs {
		path := filepath.Join(agentName, d)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	cfg := models.OrbitConfig{
		Name: agentName,
	}

	cfg.Model.Provider = "ollama"
	cfg.Model.Name = modelName

	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(agentName, "orbit.yaml"), yamlData, 0644)
}
