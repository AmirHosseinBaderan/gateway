package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadUpstreams(path string) ([]Upstream, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	upstreams := make([]Upstream, 0)
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yml") || file.Name() == "settings.yml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("read config: %w", err)
		}

		var upstream Upstream
		if err := yaml.Unmarshal(data, &upstream); err != nil {
			return nil, fmt.Errorf("unmarshal yaml: %w", err)
		}

		if len(upstream.Items) == 0 {
			return nil, fmt.Errorf("empty upstream items")
		}
		upstreams = append(upstreams, upstream)
	}
	return upstreams, nil
}

func validate(cfg *Config) error {
	if cfg.App.Name == "" {
		return fmt.Errorf("app.name is required")
	}
	if cfg.Server.Port == 0 {
		return fmt.Errorf("server.port is required")
	}
	return nil
}
