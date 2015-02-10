package sql

import (
	"fmt"
	"github.com/boourns/avc/ast"
)

func Generate(model *ast.Model) {

}

func sqlType(goType string) string {
	switch goType {
	case "string", "[]byte":
		return "VARCHAR2"
	case "int", "int64":
		return "INT"
	case "bool":
		return "BOOLEAN"
	default:
		panic(fmt.Sprintf("Don't know sql type for %s", goType))
	}
}

func sqlCreateTable(model *ast.Model) string {
	sql := fmt.Sprintf("CREATE TABLE %s (\n", model.Name)
	comma := ","
	for index, field := range model.Fields {
		if index == len(model.Fields)-1 {
			comma = ""
		}
		sql += fmt.Sprintf("  %s AS %s%s\n", field.Name, sqlType(field.Type), comma)
	}
	sql += ")\n"
	return sql
}

func createTable(model *ast.Model) string {
	return ""
}
