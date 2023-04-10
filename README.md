
# slog: Loki handler

[![tag](https://img.shields.io/github/tag/samber/slog-loki.svg)](https://github.com/samber/slog-loki/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20.1-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/slog-loki?status.svg)](https://pkg.go.dev/github.com/samber/slog-loki)
![Build Status](https://github.com/samber/slog-loki/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/slog-loki)](https://goreportcard.com/report/github.com/samber/slog-loki)
[![Coverage](https://img.shields.io/codecov/c/github/samber/slog-loki)](https://codecov.io/gh/samber/slog-loki)
[![Contributors](https://img.shields.io/github/contributors/samber/slog-loki)](https://github.com/samber/slog-loki/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/slog-loki)](./LICENSE)

A [Loki](https://grafana.com/oss/loki/) Handler for [slog](https://pkg.go.dev/golang.org/x/exp/slog) Go library.

## 🚀 Install

```sh
go get github.com/samber/slog-loki
```

**Compatibility**: go >= 1.20.1

This library is v0 and follows SemVer strictly. On `slog` final release (go 1.21), this library will go v1.

No breaking changes will be made to exported APIs before v1.0.0.

## 💡 Usage

GoDoc: [https://pkg.go.dev/github.com/samber/slog-loki](https://pkg.go.dev/github.com/samber/slog-loki)

### Handler options

```go
type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// loki endpoint
	Endpoint string
	// log batching
	BatchWait          time.Duration
	BatchEntriesNumber int
}
```

Attributes will be injected in log payload.

Attributes added to records are not accepted.

### Example

```go
import (
	slogloki "github.com/samber/slog-loki"
	"golang.org/x/exp/slog"
)

func main() {
	endpoint := "localhost:3100"

	logger := slog.New(slogloki.Option{Level: slog.LevelDebug, Endpoint: endpoint}.NewLokiHandler())
    logger = logger.
        With("environment", "dev").
        With("release", "v1.0.0")

    // log error
    logger.Error("caramba!")

    // log user signup
    logger.Info("user registration")
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

![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
