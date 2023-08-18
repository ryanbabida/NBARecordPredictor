package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port             string `yaml:"port"`
		TimeoutInSeconds int    `yaml:"timeoutInSeconds"`
	} `yaml:"server"`
	DataStore struct {
		Filepath string            `yaml:"filepath"`
		Files    map[string]string `yaml:"files,omitempty"`
	} `yaml:"datastore"`
}

func GetConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("GetConfig: unable to read file: %w", err)
	}

	config := Config{}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("GetConfig: unable to parse file: %w", err)
	}

	return &config, nil
}
