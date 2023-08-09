package main

import (
	slogloki "github.com/samber/slog-loki"

	"log/slog"
)

func main() {
	endpoint := "http://localhost:3100"

	logger := slog.New(slogloki.Option{Level: slog.LevelDebug, Endpoint: endpoint}.NewLokiHandler())
	logger = logger.With("release", "v1.0.0")

	logger.Error("A message")
}
