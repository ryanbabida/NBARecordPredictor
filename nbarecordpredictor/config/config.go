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
}

func GetConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("GetConfig: unable to read file: %w", err)
	}

	Config := Config{}

	if err := yaml.Unmarshal(data, &Config); err != nil {
		return nil, fmt.Errorf("GetConfig: unable to parse file: %w", err)
	}

	return &Config, nil
}
