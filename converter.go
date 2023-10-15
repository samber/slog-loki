package slogloki

import (
	"fmt"
	"strconv"
	"strings"

	"log/slog"

	slogcommon "github.com/samber/slog-common"
)

func attrToLabelMap(base string, attrs []slog.Attr, labels *map[string]string) {
	for i := range attrs {
		attrToValue(base, attrs[i], labels)
	}
}

func attrToValue(base string, attr slog.Attr, labels *map[string]string) {
	k := attr.Key
	v := attr.Value
	kind := v.Kind()

	switch kind {
	case slog.KindAny:
		(*labels)[base+k] = slogcommon.AnyValueToString(v)
	case slog.KindLogValuer:
		(*labels)[base+k] = slogcommon.AnyValueToString(v)
	case slog.KindGroup:
		attrToLabelMap(base+k+".", v.Group(), labels)
	case slog.KindInt64:
		(*labels)[base+k] = fmt.Sprintf("%d", v.Int64())
	case slog.KindUint64:
		(*labels)[base+k] = fmt.Sprintf("%d", v.Uint64())
	case slog.KindFloat64:
		(*labels)[base+k] = fmt.Sprintf("%f", v.Float64())
	case slog.KindString:
		(*labels)[base+k] = v.String()
	case slog.KindBool:
		(*labels)[base+k] = strconv.FormatBool(v.Bool())
	case slog.KindDuration:
		(*labels)[base+k] = v.Duration().String()
	case slog.KindTime:
		(*labels)[base+k] = v.Time().UTC().String()
	default:
		(*labels)[base+k] = slogcommon.AnyValueToString(v)
	}
}

func mapToLabels(input map[string]string) string {
	labelsList := []string{}
	for k, v := range input {
		labelsList = append(labelsList, fmt.Sprintf(`%s="%s"`, k, v))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(labelsList, ", "))
}
