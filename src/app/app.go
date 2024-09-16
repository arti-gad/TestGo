package app

import (
	"testgo/config"

	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	SetLogrus(cfg.Log.Level)

	log.Info("Initializing handlers and routes...")
}
