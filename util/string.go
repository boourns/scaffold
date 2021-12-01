package util

import (
	"fmt"
	"strings"
)

func QuestionMarks(num int) string {
	slice := []string{}
	for i := 0; i < num; i++ {
		slice = append(slice, "?")
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
