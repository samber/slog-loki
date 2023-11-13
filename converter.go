package slogloki

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/prometheus/common/model"
	slogcommon "github.com/samber/slog-common"
)

var SourceKey = "source"
var ErrorKeys = []string{"error", "err"}

type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) model.LabelSet

func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) model.LabelSet {
	// aggregate all attributes
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	// developer formatters
	attrs = slogcommon.ReplaceError(attrs, ErrorKeys...)
	if addSource {
		attrs = append(attrs, slogcommon.Source(SourceKey, record))
	}
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)
	attrs = append(attrs, slog.String("level", record.Level.String()))

	// handler formatter
	output := slogcommon.AttrsToMap(attrs...)

	labelSet := model.LabelSet{}
	flatten("", output, labelSet)

	return labelSet
}

// https://stackoverflow.com/questions/64419565/how-to-efficiently-flatten-a-map
func flatten(prefix string, src map[string]any, dest model.LabelSet) {
	if len(prefix) > 0 {
		prefix += "."
	}
	for k, v := range src {
		switch child := v.(type) {
		case map[string]any:
			flatten(prefix+k, child, dest)
		case []any:
			for i := 0; i < len(child); i++ {
				dest[model.LabelName(prefix+k+"."+strconv.Itoa(i))] = model.LabelValue(fmt.Sprintf("%v", child[i]))
			}
		default:
			dest[model.LabelName(prefix+k)] = model.LabelValue(fmt.Sprintf("%v", v))
		}
	}
}
