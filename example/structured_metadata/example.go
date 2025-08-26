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

	// With structured metadata enabled, attributes are not sent as labels, thus
	// allowing to log high-cardinality metadata without impacting performance.
	o := slogloki.Option{
		HandleRecordAttrsAsStructuredMetadata: true,
		Level:                                 slog.LevelDebug,
		Client:                                client,
	}
	logger := slog.New(o.NewLokiHandler())

	// Attributes added via WithAttrs are always sent as labels to Loki.
	logger = logger.With("release", "v1.0.0")
	// This will send the "span_id", a high cardinality value, as structured metadata, not as a label.
	//
	// More about structured metadata in Loki:
	// https://grafana.com/docs/loki/latest/get-started/labels/structured-metadata/
	logger.Error("A message with structured metadata", slog.String("span_id", "1234567"))

	client.Stop()
}
