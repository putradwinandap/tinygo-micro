package main

import (
	"url/config"
	"url/internal/delivery/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting URL service...")

	cfg := config.LoadConfig()
	router := http.SetupRouter(cfg)

	log.Info("URL service is running on port 8002")
	router.Run(":8002")
}
