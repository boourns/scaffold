package model

import (
	"github.com/boourns/scaffold/util"
)

func fieldString(prefix string, list []string, suffix string) string {
	return util.StringJoin(util.Transform(prefix, list, suffix), ",")
}
