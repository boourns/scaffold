package sqlgen

import (
	"testing"

	"github.com/boourns/scaffold/ast"
)

func TestCreateTable(t *testing.T) {
	model := &ast.Model{}

	model.Name = "User"
	model.Fields = []ast.Field{{Name: "Name", Type: "string"}, {Name: "Admin", Type: "bool", Tag: "`yolo`"}}
	sql := CreateTable(model)

	expected := `CREATE TABLE User (
Name VARCHAR(255),
Admin BOOLEAN
);`
	if sql != expected {
		t.Errorf("Generated create statement didn't match expected.\nexpected:\n%s\ngenerated:\n%s\n", expected, sql)
	}
}
