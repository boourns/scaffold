package model

import (
	"bytes"
	"github.com/boourns/scaffold/ast"
)

type model struct{}

func (c model) Description() string {
	return "Generate JSON REST endpoints"
}

func (c model) Generate(m *ast.Model) (string, error) {
	out := bytes.NewBuffer(nil)
	err := modelTemplate(out, m)
	return out.String(), err
}

var Scaffold = model{}
