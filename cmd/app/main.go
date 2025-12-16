package main

import (
	"gateway/internal/config"
	"log"
)

func main() {
	cfg, err := config.Load("./config/settings.yml")
	if err != nil {
		log.Fatalf("load config: %v", err)
		return
	}

	log.Printf("starting %s (%s) on %s:%d",
		cfg.App.Name,
		cfg.App.Env,
		cfg.Server.Host,
		cfg.Server.Port)

	select {}
}
