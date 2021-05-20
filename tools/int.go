package tools

import (
	"strconv"
	"strings"
)

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func IntJoin(elems []int, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return IntToString(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(IntToString(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(IntToString(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(IntToString(s))
	}
	return b.String()
}

func Int64ToString(el int64) string {
	return strconv.FormatInt(el, 10)
}
