package main

import (
	"github.com/grafana/loki-client-go/loki"
	slogloki "github.com/samber/slog-loki/v3"

	"log/slog"
)

func main() {
	config, _ := loki.NewDefaultConfig("http://localhost:3100/loki/api/v1/push")
	config.TenantID = "xyz"
	client, _ := loki.New(config)

	logger := slog.New(slogloki.Option{Level: slog.LevelDebug, Client: client}.NewLokiHandler())
	logger = logger.With("release", "v1.0.0")

	logger.Error("A message")

	client.Stop()
}
