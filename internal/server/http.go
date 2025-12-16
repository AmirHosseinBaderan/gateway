package server

import (
	"context"
	"fmt"
	"gateway/internal/config/base"
	"gateway/internal/logging"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type HttpServer struct {
	server *http.Server
}

func New(cfg *base.ServerConfig, handler http.Handler) *HttpServer {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(cfg.WriteTimeoutMs) * time.Millisecond,
	}

	return &HttpServer{
		server: srv,
	}
}

func (s *HttpServer) Start() error {
	logging.Logger.Info("HTTP server starting", zap.String("addr", s.server.Addr))
	return s.server.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	logging.Logger.Info("HTTP server shutting down")
	return s.server.Shutdown(ctx)
}
