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
	if err := internal.Migrate(cfg); err != nil {
		slog.Error("Failed to migrate", "err", err)
		os.Exit(1)
	}

	slog.Info("Migrations applied successfully")
}
