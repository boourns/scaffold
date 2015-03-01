package model
import (
"fmt"
"io"
"strings"
"github.com/boourns/scaffold/ast"
"github.com/boourns/scaffold/util"
)
//line model.go.ego:1
 func modelTemplate(w io.Writer, m *ast.Model) error  {
//line model.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n")
//line model.go.ego:8
_, _ = fmt.Fprintf(w, "\n\npackage ")
//line model.go.ego:9
_, _ = fmt.Fprintf(w, "%v",  strings.ToLower(m.Name) )
//line model.go.ego:10
_, _ = fmt.Fprintf(w, "\n\nimport (\n\t\"github.com/boourns/dbutil\"\n)\n\nfunc (m *")
//line model.go.ego:15
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:15
_, _ = fmt.Fprintf(w, ") sqlFields() string {\n  return \"")
//line model.go.ego:16
_, _ = fmt.Fprintf(w, "%v",  util.StringJoin(m.FieldSlice(), ", ") )
//line model.go.ego:16
_, _ = fmt.Fprintf(w, "\"\n}\n\nfunc load")
//line model.go.ego:19
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:19
_, _ = fmt.Fprintf(w, "(rows *sql.Rows) (*")
//line model.go.ego:19
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:19
_, _ = fmt.Fprintf(w, ", error) {\n\tret := ")
//line model.go.ego:20
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:20
_, _ = fmt.Fprintf(w, "{}\n\n\terr := rows.Scan(")
//line model.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  fieldString("&ret.", m.FieldSlice(), "") )
//line model.go.ego:22
_, _ = fmt.Fprintf(w, ")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &ret, nil\n}\n\nfunc Where(tx dbutil.DBLike, where string, whereFields ...interface{}) (*")
//line model.go.ego:29
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:29
_, _ = fmt.Fprintf(w, ", error) {\n  ret := []")
//line model.go.ego:30
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:30
_, _ = fmt.Fprintf(w, "{}\n  sql := fmt.Sprintf(\"SELECT %%s from ")
//line model.go.ego:31
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:31
_, _ = fmt.Fprintf(w, " WHERE %%s\", ")
//line model.go.ego:31
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:31
_, _ = fmt.Fprintf(w, ".sqlFields(), where)\n\trows, err := tx.Query(sql, whereFields)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n    item, err := load")
//line model.go.ego:37
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:37
_, _ = fmt.Fprintf(w, "(rows)\n    if err != nil {\n      return nil, err\n    }\n    ret = append(ret, item)\n\t}\n  return ret\n}\n\nfunc (s *")
//line model.go.ego:46
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:46
_, _ = fmt.Fprintf(w, ") Update(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(fmt.Sprintf(\"UPDATE ")
//line model.go.ego:47
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:47
_, _ = fmt.Fprintf(w, "(%%s) VALUES(")
//line model.go.ego:47
_, _ = fmt.Fprintf(w, "%v",  util.QuestionMarks(len(m.FieldSlice())) )
//line model.go.ego:47
_, _ = fmt.Fprintf(w, ") WHERE ")
//line model.go.ego:47
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:47
_, _ = fmt.Fprintf(w, ".ID = ?\", ")
//line model.go.ego:47
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:47
_, _ = fmt.Fprintf(w, ".sqlFields()))\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n    params := []interface{}{")
//line model.go.ego:53
_, _ = fmt.Fprintf(w, "%v",  fieldString("s.", m.FieldSlice(), "") )
//line model.go.ego:53
_, _ = fmt.Fprintf(w, "}\n    params = append(params, s.ID)\n\n\t\t_, err = stmt.Exec(params)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n}\n\nfunc (s *")
//line model.go.ego:62
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:62
_, _ = fmt.Fprintf(w, ") Insert(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(\"INSERT INTO ")
//line model.go.ego:63
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:63
_, _ = fmt.Fprintf(w, "(")
//line model.go.ego:63
_, _ = fmt.Fprintf(w, "%v",  fieldString("", m.FieldSliceWithoutID(), "") )
//line model.go.ego:63
_, _ = fmt.Fprintf(w, ") VALUES(")
//line model.go.ego:63
_, _ = fmt.Fprintf(w, "%v",  util.QuestionMarks(len(m.FieldSliceWithoutID())) )
//line model.go.ego:63
_, _ = fmt.Fprintf(w, ")\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\t_, err = stmt.Exec(")
//line model.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  fieldString("s.", m.FieldSliceWithoutID(), "") )
//line model.go.ego:68
_, _ = fmt.Fprintf(w, ")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t}\n\treturn nil\n}\n\n\n")
return nil
}
