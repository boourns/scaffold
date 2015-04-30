package sqlgen

import (
	"fmt"
	"github.com/boourns/scaffold/ast"
)

func sqlType(field ast.Field) string {
	var sql string
	switch field.Type {
	case "string", "[]byte":
		sql = "VARCHAR(255)"
	case "int", "int64":
		sql = "INT"
	case "bool":
		sql = "BOOLEAN"
	default:
		panic(fmt.Sprintf("Don't know sql type for %s", field.Type))
	}

	if field.Name == "ID" {
		sql = fmt.Sprintf("%s NOT NULL PRIMARY KEY", sql)
	}
	return sql
}
