package main

import (
	"gateway/internal/config/base"
	"gateway/internal/config/site"
	"gateway/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := base.Load("./config/settings.yml")
	if err != nil {
		log.Fatalf("load config: %v", err)
		return
	}

	sites, err := site.LoadSites(cfg.App.Upstream)
	if err != nil {
		log.Fatalf("load sites: %v", err)
	}
	log.Printf("load sites count : %v", len(sites))

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	httpServer := server.New(&cfg.Server, mux)

	log.Printf("starting %s (%s) on %s:%d",
		cfg.App.Name,
		cfg.App.Env,
		cfg.Server.Host,
		cfg.Server.Port)

	if err := httpServer.Start(); err != nil {
		log.Fatalf("start http server: %v", err)
	}
}
