package slogloki

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/grafana/loki-client-go/loki"
	slogcommon "github.com/samber/slog-common"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// loki
	Client *loki.Client

	// optional: customize webhook event builder
	Converter Converter

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

// Creating a Loki client at each `NewLokiHandler` call may lead to connection
// leak when chaining many operations: `logger.With(...).With(...).With(...).With(...)`
func (o Option) NewLokiHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Client == nil {
		panic(fmt.Errorf("missing *loki.Client"))
	}

	return &LokiHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

var _ slog.Handler = (*LokiHandler)(nil)

type LokiHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *LokiHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *LokiHandler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.option.Converter != nil {
		converter = h.option.Converter
	}

	attrs := converter(h.option.AddSource, h.option.ReplaceAttr, h.attrs, h.groups, &record)

	return h.option.Client.Handle(attrs, record.Time, record.Message)
}

func (h *LokiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LokiHandler{
		option: h.option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

func (h *LokiHandler) WithGroup(name string) slog.Handler {
	return &LokiHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
