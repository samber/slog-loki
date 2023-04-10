package slogloki

import (
	"encoding"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
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
		(*labels)[base+k] = anyValueToString(v)
	case slog.KindLogValuer:
		(*labels)[base+k] = anyValueToString(v)
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
		(*labels)[base+k] = v.Time().String()
	default:
		(*labels)[base+k] = anyValueToString(v)
	}
}

func anyValueToString(v slog.Value) string {
	if tm, ok := v.Any().(encoding.TextMarshaler); ok {
		data, err := tm.MarshalText()
		if err != nil {
			return ""
		}

		return string(data)
	}

	return fmt.Sprintf("%+v", v.Any())
}

func mapToLabels(input map[string]string) string {
	labelsList := []string{}
	for k, v := range input {
		labelsList = append(labelsList, fmt.Sprintf(`%s="%s"`, k, v))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(labelsList, ", "))
}
