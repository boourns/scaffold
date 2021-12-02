package sqlgen

import (
	"bytes"
	"fmt"
	"scaffold/ast"
)

func CreateTable(m *ast.Model) string {
	out := &bytes.Buffer{}

	err := createTableSQL(out, m)

	if err != nil {
		panic(fmt.Sprintf("Error running createTableSQL: %s", err))
	}

	return out.String()
}
