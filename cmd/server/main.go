package main

import (
	"log/slog"
	"os"
	"study-chat/internal/infra/service"

	"study-chat/internal"
)

func main() {
	cfg, err := service.LoadConfig()
	if err != nil {
		slog.Error("Could not load config", "err", err)
		os.Exit(1)
	}
	if err := internal.Run(cfg); err != nil {
		slog.Error("Failed to run server", "err", err)
		os.Exit(1)
	}
}
