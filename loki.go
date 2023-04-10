package slogloki

import (
	"github.com/afiskon/promtail-client/promtail"
	"golang.org/x/exp/slog"
)

var logLevelConverter = map[slog.Level]promtail.LogLevel{
	slog.LevelDebug: promtail.DEBUG,
	slog.LevelInfo:  promtail.INFO,
	slog.LevelWarn:  promtail.WARN,
	slog.LevelError: promtail.ERROR,
}
