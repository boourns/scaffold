package model

import (
	"scaffold/util"
	"strings"
)

func fieldString(prefix string, list []string, suffix string) string {
	return strings.Join(util.Transform(prefix, list, suffix), ",")
}
