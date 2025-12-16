package middleware

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(path string) ([]Middleware, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var middlewares []Middleware
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yml") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, err
		}

		var cfg Middleware
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
		middlewares = append(middlewares, cfg)
	}
	return middlewares, nil
}
