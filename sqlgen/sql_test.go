package sqlgen

import (
	"log"
	"testing"

	"scaffold/ast"
)

func TestSqlCreateTable(t *testing.T) {
	model := &ast.Model{}

	model.Name = "User"
	model.Fields = []ast.Field{ast.Field{Name: "Name", Type: "string"}, ast.Field{Name: "Admin", Type: "bool", Tag: "`yolo`"}}

	log.Printf("sql %s", CreateTable(model))
}
