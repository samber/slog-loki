package slogloki

import (
	"context"
	"fmt"
	"time"

	"log/slog"

	"github.com/afiskon/promtail-client/promtail"
	slogcommon "github.com/samber/slog-common"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// loki endpoint
	Endpoint string
	// log batching
	BatchWait          time.Duration
	BatchEntriesNumber int

	// internal
	attrs  []slog.Attr
	groups []string
}

// @TODO: creating a promptail client at each `NewLokiHandler` call may lead to connection
// leak when chaining many operations: `logger.With(...).With(...).With(...).With(...)`
func (o Option) NewLokiHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	clients := map[slog.Level]promtail.Client{}
	for k, v := range LogLevels {
		conf := promtail.ClientConfig{
			PushURL:            o.Endpoint,
			BatchWait:          o.BatchWait,
			BatchEntriesNumber: o.BatchEntriesNumber,
			SendLevel:          v,
			PrintLevel:         promtail.DISABLE,
			Labels:             o.getLabels(k),
		}

		// Do not handle error here, because promtail method always returns `nil`.
		client, _ := promtail.NewClientJson(conf)
		clients[k] = client
	}

	return &LokiHandler{
		option:  o,
		clients: clients,
	}
}

func (o Option) getLabels(level slog.Level) string {
	labels := map[string]string{}
	attrToLabelMap("", append(o.attrs, slog.String("level", level.String())), &labels)
	return mapToLabels(labels)
}

var _ slog.Handler = (*LokiHandler)(nil)

type LokiHandler struct {
	option  Option
	clients map[slog.Level]promtail.Client
}

func (h *LokiHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *LokiHandler) Handle(ctx context.Context, record slog.Record) error {
	switch record.Level {
	case slog.LevelDebug:
		h.clients[slog.LevelDebug].Debugf(record.Message)
	case slog.LevelInfo:
		h.clients[slog.LevelInfo].Infof(record.Message)
	case slog.LevelWarn:
		h.clients[slog.LevelWarn].Warnf(record.Message)
	case slog.LevelError:
		h.clients[slog.LevelError].Errorf(record.Message)
	default:
		return fmt.Errorf("unknown log level")
	}
	return nil
}

func (h *LokiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	option := Option{
		Level:              h.option.Level,
		Endpoint:           h.option.Endpoint,
		BatchWait:          h.option.BatchWait,
		BatchEntriesNumber: h.option.BatchEntriesNumber,
		attrs:              slogcommon.AppendAttrsToGroup(h.option.groups, h.option.attrs, attrs...),
		groups:             h.option.groups,
	}

	// @TODO: creating a promptail client at each `NewLokiHandler` call may lead to connection
	// leak when chaining many operations: `logger.With(...).With(...).With(...).With(...)`
	return option.NewLokiHandler()
}

func (h *LokiHandler) WithGroup(name string) slog.Handler {
	option := Option{
		Level:              h.option.Level,
		Endpoint:           h.option.Endpoint,
		BatchWait:          h.option.BatchWait,
		BatchEntriesNumber: h.option.BatchEntriesNumber,
		attrs:              h.option.attrs,
		groups:             append(h.option.groups, name),
	}

	// @TODO: creating a promptail client at each `NewLokiHandler` call may lead to connection
	// leak when chaining many operations: `logger.With(...).With(...).With(...).With(...)`
	return option.NewLokiHandler()
}
