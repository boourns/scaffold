package sqlgen

import (
	"fmt"

	"scaffold/ast"
)

func sqlType(field ast.Field) string {
	override := field.Override("sqlType", "")
	if override != "" {
		return override
	}

	var sql string
	switch field.Type {
	case "string", "[]byte":
		sql = "VARCHAR(255)"
	case "int", "int64":
		sql = "INTEGER"
	case "bool":
		sql = "BOOLEAN"
	case "time.Time":
		sql = "DATETIME"
	default:
		panic(fmt.Sprintf("Don't know sql type for %s, field %v", field.Type, field))
	}

	if field.Name == "ID" {
		sql = fmt.Sprintf("%s PRIMARY KEY", sql)
	}
	return sql
}
