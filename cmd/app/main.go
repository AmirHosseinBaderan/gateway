package main

import (
	"gateway/internal/config/base"
	"gateway/internal/config/middleware"
	"gateway/internal/config/site"
	"gateway/internal/logging"
	"gateway/internal/server"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	// Initialize structured logger
	if err := logging.InitLogger(); err != nil {
		panic(err)
	}
	defer logging.Sync()

	cfg, err := base.Load("./config/settings.yml")
	if err != nil {
		logging.Logger.Fatal("Failed to load config", zap.Error(err))
		return
	}

	sites, err := site.LoadSites(cfg.App.Upstream)
	if err != nil {
		logging.Logger.Fatal("Failed to load sites", zap.Error(err))
	}
	logging.Logger.Info("Sites loaded successfully", zap.Int("count", len(sites)))

	var globalMiddleware []middleware.Middleware
	if cfg.MiddlewarePath != "" {
		globalMiddleware, err = middleware.Load(cfg.MiddlewarePath)
		if err != nil {
			logging.Logger.Fatal("Failed to load middleware", zap.Error(err))
		}
	}
	logging.Logger.Info("Middleware loaded successfully", zap.Int("count", len(globalMiddleware)))

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	httpServer := server.New(&cfg.Server, mux)

	logging.Logger.Info("Starting server",
		zap.String("name", cfg.App.Name),
		zap.String("env", cfg.App.Env),
		zap.String("host", cfg.Server.Host),
		zap.Int("port", cfg.Server.Port))

	if err := httpServer.Start(); err != nil {
		logging.Logger.Fatal("Failed to start HTTP server", zap.Error(err))
	}
}
