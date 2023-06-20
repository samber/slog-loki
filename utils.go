package slogloki

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
)

func appendAttrsToGroup(groups []string, actualAttrs []slog.Attr, newAttrs []slog.Attr) []slog.Attr {
	if len(groups) == 0 {
		return uniqAttrs(append(actualAttrs, newAttrs...))
	}

	for i := range actualAttrs {
		attr := actualAttrs[i]
		if attr.Key == groups[0] && attr.Value.Kind() == slog.KindGroup {
			actualAttrs[i] = slog.Group(groups[0], lo.ToAnySlice(appendAttrsToGroup(groups[1:], attr.Value.Group(), newAttrs))...)
			return actualAttrs
		}
	}

	return uniqAttrs(
		append(
			actualAttrs,
			slog.Group(
				groups[0],
				lo.ToAnySlice(appendAttrsToGroup(groups[1:], []slog.Attr{}, newAttrs))...,
			),
		),
	)
}

func uniqAttrs(attrs []slog.Attr) []slog.Attr {
	return uniqByLast(attrs, func(item slog.Attr) string {
		return item.Key
	})
}

func uniqByLast[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]int, len(collection))
	seenIndex := 0

	for _, item := range collection {
		key := iteratee(item)

		if index, ok := seen[key]; ok {
			result[index] = item
			continue
		}

		seen[key] = seenIndex
		seenIndex++
		result = append(result, item)
	}

	return result
}
