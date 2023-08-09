package slogloki

import (
	"log/slog"

	"github.com/afiskon/promtail-client/promtail"
)

var logLevelConverter = map[slog.Level]promtail.LogLevel{
	slog.LevelDebug: promtail.DEBUG,
	slog.LevelInfo:  promtail.INFO,
	slog.LevelWarn:  promtail.WARN,
	slog.LevelError: promtail.ERROR,
}
