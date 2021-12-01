package controller

import (
	"fmt"
	"io"
	"scaffold/ast"
)

//line controller.go.ego:1
func GenerateIndex(w io.Writer, m *ast.Model) error {
//line controller.go.ego:2
	_, _ = fmt.Fprintf(w, "\n\n")
//line controller.go.ego:8
	_, _ = fmt.Fprintf(w, "\n\npackage ")
//line controller.go.ego:9
	_, _ = fmt.Fprintf(w, "%v", m.Package)
//line controller.go.ego:10
	_, _ = fmt.Fprintf(w, "\n\nimport (\n  \"encoding/json\"\n  \"net/http\"\n)\n\n")
//line controller.go.ego:16
	controller := fmt.Sprintf("%sController", m.Name)
//line controller.go.ego:17
	_, _ = fmt.Fprintf(w, "\n\ntype ")
//line controller.go.ego:18
	_, _ = fmt.Fprintf(w, "%v", controller)
//line controller.go.ego:18
	_, _ = fmt.Fprintf(w, " struct{}\n\nfunc (m *")
//line controller.go.ego:20
	_, _ = fmt.Fprintf(w, "%v", controller)
//line controller.go.ego:20
	_, _ = fmt.Fprintf(w, ") Index() {\n\n}\n\n\n")
	return nil
}
