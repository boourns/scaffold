package sqlgen

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/boourns/scaffold/ast"
)

//go:embed "create_table.template"
var templatePath embed.FS
var createTableTemplate *template.Template

type CreateTableData struct {
	Model *ast.Model
}

func (ctd CreateTableData) Join(field []ast.Field) string {
	s := []string{}
	for _, f := range field {
		s = append(s, fmt.Sprintf("%s %s", f.Name, sqlType(f)))
	}
	return strings.Join(s, ",\n")
}

func CreateTable(m *ast.Model) string {
	var err error

	if createTableTemplate == nil {
		createTableTemplate = template.Must(template.ParseFS(templatePath, "create_table.template"))
	}

	data := CreateTableData{
		Model: m,
	}

	out := &bytes.Buffer{}
	if err = createTableTemplate.Execute(out, data); err != nil {
		panic(fmt.Sprintf("Error executing CreateTable template %s", err))
	}
	return out.String()
}
