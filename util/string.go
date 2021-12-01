package util

import (
	"fmt"
	"strings"
)

func QuestionMarks(num int) string {
	if num < 1 {
		return ""
	}
	slice := make([]string, num)
	slice[0] = "?"
	for i := 1; i < len(slice); i *= 2 {
		copy(slice[i:], slice[:i])
	}
	return strings.Join(slice, ",")
}

func TransformFunc(fields []string, trans func(string) string) []string {
	out := []string{}
	for _, v := range fields {
		out = append(out, trans(v))
	}
	return out
}

func Transform(prefix string, fields []string, suffix string) []string {
	return TransformFunc(fields, func(in string) string {
		return fmt.Sprintf("%s%s%s", prefix, in, suffix)
	})
}
