package util

import (
	"fmt"
)

func StringJoin(fields []string, joiner string) string {
	out := ""
	for i, v := range fields {
		if i == len(fields)-1 {
			joiner = ""
		}
		out = fmt.Sprintf("%s%s%s", out, v, joiner)
	}
	return out
}

func QuestionMarks(num int) string {
	slice := []string{}
	for i := 0; i < num; i++ {
		slice = append(slice, "?")
	}
	return StringJoin(slice, ",")
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
