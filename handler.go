package slogloki

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki/pkg/push"
	slogcommon "github.com/samber/slog-common"
)

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

	// When set to true, this handler sends record attributes as structured metadata. To send all record attributes as structured metadata,
	// use the RemoveAttrsConverter along with this option set to true.
	//
	// In Loki, labels are key-value pairs used to index and filter log streams.
	// They enable powerful querying but must remain in low cardinality to maintain performance.
	//
	// By default, LokiHandler.Handle sends all log record attributes as labels to Loki.
	// This works well for common, low-cardinality fields like service name, log level, environment, etc.
	//
	// However, in some cases, you may want to include high-cardinality metadata, such as request IDs, user IDs, or session tokens, for improved debugging and traceability.
	//
	// Starting with schema version 13, Loki introduced structured metadata, allowing high-cardinality attributes to be attached to log records without indexation.
	// This helps preserve Loki’s performance while improving traceability and logging capabilities.
	//
	// Note: Attributes added via LokiHandler.WithAttrs are always sent as labels, regardless of the value of this setting.
	// If set to false (default), the handler will not send record attributes as structured metadata.
	//
	// Learn more about structured metadata in Loki:
	// https://grafana.com/docs/loki/latest/get-started/labels/structured-metadata/
	HandleRecordsWithMetadata bool
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

	if o.Converter == nil {
		o.Converter = DefaultConverter
	}

	if o.AttrFromContext == nil {
		o.AttrFromContext = []func(ctx context.Context) []slog.Attr{}
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
	fromContext := slogcommon.ContextExtractor(ctx, h.option.AttrFromContext)

	attrs := h.option.Converter(h.option.AddSource, h.option.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record)

	if h.option.HandleRecordsWithMetadata {
		var m push.LabelsAdapter
		record.Attrs(func(attr slog.Attr) bool {
			m = append(m, push.LabelAdapter{
				Name:  attr.Key,
				Value: attr.Value.String(),
			})
			return true
		})
		return h.option.Client.HandleWithMetadata(attrs, record.Time, record.Message, m)
	}

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
	// https://cs.opensource.google/go/x/exp/+/46b07846:slog/handler.go;l=247
	if name == "" {
		return h
	}

	return &LokiHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
