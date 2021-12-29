package model

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/boourns/scaffold/ast"
)

type ModelViewData struct {
	Model *ast.Model
}

func (mvd ModelViewData) FieldString(prefix string) string {
	qualified := make([]string, len(mvd.Model.Fields))
	for i, f := range mvd.Model.Fields {
		qualified[i] = fmt.Sprintf("%s.%s", prefix, f.Name)
	}
	joined := strings.Join(qualified, ", ")
	return string(template.HTML(joined))
}
