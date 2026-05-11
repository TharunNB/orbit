package config

type OrbitConfig struct {
	Name  string `yaml:"name"`
	Model struct {
		Provider string `yaml:"provider"`
		Name     string `yaml:"name"`
	} `yaml:"model"`
}
