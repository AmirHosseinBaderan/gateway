package site

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadSites(path string) ([]Config, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var configs []Config
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yml") || file.Name() == "settings.yml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(path, file.Name()))
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
