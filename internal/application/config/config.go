package config

import (
	"gateway/internal/models"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "./config/config.json"

func LoadConfig() (*models.AppConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg *models.AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
