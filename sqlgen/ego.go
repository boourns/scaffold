package sqlgen

import (
	"fmt"
	"io"
	"scaffold/ast"
)

//line create_table.sql.ego:1
 func createTableSQL(w io.Writer, m *ast.Model) error  {
//line create_table.sql.ego:2
_, _ = fmt.Fprintf(w, "\n\n")
//line create_table.sql.ego:8
_, _ = fmt.Fprintf(w, "\n\nCREATE TABLE ")
//line create_table.sql.ego:9
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line create_table.sql.ego:9
_, _ = fmt.Fprintf(w, " (\n  ")
//line create_table.sql.ego:10
 for i, field := range m.Fields { 
//line create_table.sql.ego:11
_, _ = fmt.Fprintf(w, "\n    ")
//line create_table.sql.ego:11
_, _ = fmt.Fprintf(w, "%v",  field.Name )
//line create_table.sql.ego:11
_, _ = fmt.Fprintf(w, " ")
//line create_table.sql.ego:11
_, _ = fmt.Fprintf(w, "%v",  sqlType(field) )
//line create_table.sql.ego:11
 if i < len(m.Fields)-1 { 
//line create_table.sql.ego:11
_, _ = fmt.Fprintf(w, ",")
//line create_table.sql.ego:11
 }
//line create_table.sql.ego:12
_, _ = fmt.Fprintf(w, "\n  ")
//line create_table.sql.ego:12
 } 
//line create_table.sql.ego:13
_, _ = fmt.Fprintf(w, "\n);\n\n")
return nil
}
