package controller

import (
	"bytes"
	"github.com/boourns/scaffold/ast"
)

type controller struct{}

func (c controller) Description() string {
	return "Generate JSON REST endpoints"
}

func (c controller) Generate(m *ast.Model) (string, error) {
	out := bytes.NewBuffer(nil)
	err := GenerateIndex(out, m)
	return out.String(), err
}

var Scaffold = controller{}
