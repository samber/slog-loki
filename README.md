
# slog: Loki handler

[![tag](https://img.shields.io/github/tag/samber/slog-loki.svg)](https://github.com/samber/slog-loki/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/slog-loki?status.svg)](https://pkg.go.dev/github.com/samber/slog-loki)
![Build Status](https://github.com/samber/slog-loki/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/slog-loki)](https://goreportcard.com/report/github.com/samber/slog-loki)
[![Coverage](https://img.shields.io/codecov/c/github/samber/slog-loki)](https://codecov.io/gh/samber/slog-loki)
[![Contributors](https://img.shields.io/github/contributors/samber/slog-loki)](https://github.com/samber/slog-loki/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/slog-loki)](./LICENSE)

A [Loki](https://grafana.com/oss/loki/) Handler for [slog](https://pkg.go.dev/log/slog) Go library.

<div align="center">
  <hr>
  <sup><b>Sponsored by:</b></sup>
  <br>
  <a href="https://quickwit.io?utm_campaign=github_sponsorship&utm_medium=referral&utm_content=samber-slog-loki&utm_source=github">
    <div>
      <img src="https://github.com/samber/oops/assets/2951285/49aaaa2b-b8c6-4f21-909f-c12577bb6a2e" width="240" alt="Quickwit">
    </div>
    <div>
      Cloud-native search engine for observability - An OSS alternative to Splunk, Elasticsearch, Loki, and Tempo.
    </div>
  </a>
  <hr>
</div>

**See also:**

- [slog-multi](https://github.com/samber/slog-multi): `slog.Handler` chaining, fanout, routing, failover, load balancing...
- [slog-formatter](https://github.com/samber/slog-formatter): `slog` attribute formatting
- [slog-sampling](https://github.com/samber/slog-sampling): `slog` sampling policy
- [slog-mock](https://github.com/samber/slog-mock): `slog.Handler` for test purposes

**HTTP middlewares:**

- [slog-gin](https://github.com/samber/slog-gin): Gin middleware for `slog` logger
- [slog-echo](https://github.com/samber/slog-echo): Echo middleware for `slog` logger
- [slog-fiber](https://github.com/samber/slog-fiber): Fiber middleware for `slog` logger
- [slog-chi](https://github.com/samber/slog-chi): Chi middleware for `slog` logger
- [slog-http](https://github.com/samber/slog-http): `net/http` middleware for `slog` logger

**Loggers:**

- [slog-zap](https://github.com/samber/slog-zap): A `slog` handler for `Zap`
- [slog-zerolog](https://github.com/samber/slog-zerolog): A `slog` handler for `Zerolog`
- [slog-logrus](https://github.com/samber/slog-logrus): A `slog` handler for `Logrus`

**Log sinks:**

- [slog-datadog](https://github.com/samber/slog-datadog): A `slog` handler for `Datadog`
- [slog-betterstack](https://github.com/samber/slog-betterstack): A `slog` handler for `Betterstack`
- [slog-rollbar](https://github.com/samber/slog-rollbar): A `slog` handler for `Rollbar`
- [slog-loki](https://github.com/samber/slog-loki): A `slog` handler for `Loki`
- [slog-sentry](https://github.com/samber/slog-sentry): A `slog` handler for `Sentry`
- [slog-syslog](https://github.com/samber/slog-syslog): A `slog` handler for `Syslog`
- [slog-logstash](https://github.com/samber/slog-logstash): A `slog` handler for `Logstash`
- [slog-fluentd](https://github.com/samber/slog-fluentd): A `slog` handler for `Fluentd`
- [slog-graylog](https://github.com/samber/slog-graylog): A `slog` handler for `Graylog`
- [slog-quickwit](https://github.com/samber/slog-quickwit): A `slog` handler for `Quickwit`
- [slog-slack](https://github.com/samber/slog-slack): A `slog` handler for `Slack`
- [slog-telegram](https://github.com/samber/slog-telegram): A `slog` handler for `Telegram`
- [slog-mattermost](https://github.com/samber/slog-mattermost): A `slog` handler for `Mattermost`
- [slog-microsoft-teams](https://github.com/samber/slog-microsoft-teams): A `slog` handler for `Microsoft Teams`
- [slog-webhook](https://github.com/samber/slog-webhook): A `slog` handler for `Webhook`
- [slog-kafka](https://github.com/samber/slog-kafka): A `slog` handler for `Kafka`
- [slog-nats](https://github.com/samber/slog-nats): A `slog` handler for `NATS`
- [slog-parquet](https://github.com/samber/slog-parquet): A `slog` handler for `Parquet` + `Object Storage`
- [slog-channel](https://github.com/samber/slog-channel): A `slog` handler for Go channels

## 🚀 Install

```sh
go get github.com/samber/slog-loki/v3
```

**Compatibility**: go >= 1.21

No breaking changes will be made to exported APIs before v4.0.0.

## 💡 Usage

GoDoc: [https://pkg.go.dev/github.com/samber/slog-loki/v3](https://pkg.go.dev/github.com/samber/slog-loki/v3)

### Handler options

```go
type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// loki
	Client *loki.Client

	// optional: customize webhook event builder
	Converter Converter
	// optional: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
	
	// By default, LokiHandler.Handle sends record attributes as labels to Loki.
	// When set to true, this handler sends record attributes as structured metadata.
	// 
	// Combine with RemoveAttrsConverter to avoid sending attributes as labels.
	HandleRecordsWithMetadata bool
}
```

Attributes will be injected in log payload.

Other global parameters:

```go
slogloki.SourceKey = "source"
slogloki.ErrorKeys = []string{"error", "err"}
slogloki.SubAttributeSeparator = "__"
slogloki.AttributeKeyInvalidCharReplacement = "_"
```

### Example

Default behaviour (all attributes are treated as labels):

```go
import (
	"github.com/grafana/loki-client-go/loki"
	slogloki "github.com/samber/slog-loki/v3"
	"log/slog"
)

func main() {
	// setup loki client
	config, _ := loki.NewDefaultConfig("http://localhost:3100/loki/api/v1/push")
	config.TenantID = "xyz"
	client, _ := loki.New(config)

	logger := slog.New(slogloki.Option{Level: slog.LevelDebug, Client: client}.NewLokiHandler())
	logger = logger.
		With("environment", "dev").
		With("release", "v1.0.0")

	// log error
	logger.Error("caramba!")

	// log user signup
	logger.Info("user registration")

	// stop loki client and purge buffers
	client.Stop()
}
```

To send record attributes as structured metadata (instead of labels), use the `HandleRecordsWithMetadata` option along with the `RemoveAttrsConverter` converter:

```go
import (
	"github.com/grafana/loki-client-go/loki"
	slogloki "github.com/samber/slog-loki/v3"
	"log/slog"
)

func main() {
	// setup loki client
	config, _ := loki.NewDefaultConfig("http://localhost:3100/loki/api/v1/push")
	config.TenantID = "xyz"
	client, _ := loki.New(config)

	// With slogloki.RemoveAttrsConverter and HandleRecordsWithMetadata enabled, attributes are not sent as labels, thus
	// allowing to log high-cardinality metadata without impacting performance.
	o := slogloki.Option{
		HandleRecordsWithMetadata: true,
		Converter:                 slogloki.RemoveAttrsConverter,
		Level:                     slog.LevelDebug,
		Client:                    client,
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
```

> Note: Attributes added via `WithAttrs` are always sent as labels to Loki.

### Tracing

Import the samber/slog-otel library.

```go
import (
	slogloki "github.com/samber/slog-loki"
	slogotel "github.com/samber/slog-otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
	)
	tracer := tp.Tracer("hello/world")

	ctx, span := tracer.Start(context.Background(), "foo")
	defer span.End()

	span.AddEvent("bar")

	logger := slog.New(
		slogloki.Option{
			// ...
			AttrFromContext: []func(ctx context.Context) []slog.Attr{
				slogotel.ExtractOtelAttrFromContext([]string{"tracing"}, "trace_id", "span_id"),
			},
		}.NewLokiHandler(),
	)

	logger.ErrorContext(ctx, "a message")
}
```

## 🤝 Contributing

- Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/slog-loki)
- Fix [open issues](https://github.com/samber/slog-loki/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/slog-loki)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
