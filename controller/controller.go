package controller

import (
	"github.com/boourns/scaffold/ast"
)

type controller struct{}

func (c controller) Description() string {
	return "Generate JSON REST endpoints"
}

func (c controller) Generate(m *ast.Model) (string, error) {
	return "yolo", nil
}

var Scaffold = controller{}
