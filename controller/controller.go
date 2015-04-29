package controller

import (
	"bytes"
	"github.com/boourns/scaffold/ast"
)

type controller struct{}

func (c controller) Description() string {
	return "Generate JSON REST endpoints"
}

func (c controller) Generate(m *ast.Model) (error) {
	out := bytes.NewBuffer(nil)
	err := GenerateIndex(out, m)

	// todo save
	
	return err
}

var Scaffold = controller{}
