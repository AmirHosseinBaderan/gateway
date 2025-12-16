package site

import (
	"gateway/internal/config/base"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadSites(cfg base.Upstream) ([]Config, error) {
	files, err := os.ReadDir(cfg.ConfigPath)
	if err != nil {
		return nil, err
	}

	var configs []Config
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yml") || (slices.Contains(cfg.Servers, strings.Replace(file.Name(), ".yml", "", 1))) {
			continue
		}

		data, err := os.ReadFile(filepath.Join(cfg.ConfigPath, file.Name()))
		if err != nil {
			return nil, err
		}
		var config Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}
