package main

import (
	"analytics/config"
	"analytics/internal/delivery/http"
)

func main() {
	cfg := config.LoadConfig()
	http.SetupRouter(cfg)
}
