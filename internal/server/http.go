package server

import (
	"context"
	"fmt"
	"gateway/internal/config"
	"net/http"
	"time"
)

type HttpServer struct {
	server *http.Server
}

func New(cfg *config.ServerConfig, handler http.Handler) *HttpServer {
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
	return s.server.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
