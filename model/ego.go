package model
import (
"fmt"
"io"
"strings"
"github.com/boourns/scaffold/ast"
"github.com/boourns/scaffold/util"
"github.com/boourns/scaffold/sqlgen"
)
//line model.go.ego:1
 func modelTemplate(w io.Writer, m *ast.Model) error  {
//line model.go.ego:2
_, _ = fmt.Fprintf(w, "\n\n")
//line model.go.ego:9
_, _ = fmt.Fprintf(w, "\n\npackage ")
//line model.go.ego:10
_, _ = fmt.Fprintf(w, "%v",  strings.ToLower(m.Name) )
//line model.go.ego:11
_, _ = fmt.Fprintf(w, "\n\nimport (\n\t\"github.com/boourns/dbutil\"\n  \"database/sql\"\n  \"fmt\"\n)\n\nfunc sqlFieldsFor")
//line model.go.ego:18
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:18
_, _ = fmt.Fprintf(w, "() string {\n  return \"")
//line model.go.ego:19
_, _ = fmt.Fprintf(w, "%v",  util.StringJoin(m.FieldSlice(), ", ") )
//line model.go.ego:19
_, _ = fmt.Fprintf(w, "\"\n}\n\nfunc load")
//line model.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:22
_, _ = fmt.Fprintf(w, "(rows *sql.Rows) (*")
//line model.go.ego:22
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:22
_, _ = fmt.Fprintf(w, ", error) {\n\tret := ")
//line model.go.ego:23
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:23
_, _ = fmt.Fprintf(w, "{}\n\n\terr := rows.Scan(")
//line model.go.ego:25
_, _ = fmt.Fprintf(w, "%v",  fieldString("&ret.", m.FieldSlice(), "") )
//line model.go.ego:25
_, _ = fmt.Fprintf(w, ")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &ret, nil\n}\n\nfunc Where(tx dbutil.DBLike, where string, whereFields ...interface{}) ([]*")
//line model.go.ego:32
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:32
_, _ = fmt.Fprintf(w, ", error) {\n  ret := []*")
//line model.go.ego:33
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:33
_, _ = fmt.Fprintf(w, "{}\n  sql := fmt.Sprintf(\"SELECT %%s from ")
//line model.go.ego:34
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:34
_, _ = fmt.Fprintf(w, " WHERE %%s\", sqlFieldsFor")
//line model.go.ego:34
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:34
_, _ = fmt.Fprintf(w, "(), where)\n\trows, err := tx.Query(sql, whereFields)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n    item, err := load")
//line model.go.ego:40
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:40
_, _ = fmt.Fprintf(w, "(rows)\n    if err != nil {\n      return nil, err\n    }\n    ret = append(ret, item)\n\t}\n  return ret, nil\n}\n\nfunc (s *")
//line model.go.ego:49
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:49
_, _ = fmt.Fprintf(w, ") Update(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(fmt.Sprintf(\"UPDATE ")
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "(%%s) VALUES(")
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "%v",  util.QuestionMarks(len(m.FieldSlice())) )
//line model.go.ego:50
_, _ = fmt.Fprintf(w, ") WHERE ")
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:50
_, _ = fmt.Fprintf(w, ".ID = ?\", sqlFieldsFor")
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "%v", m.Name)
//line model.go.ego:50
_, _ = fmt.Fprintf(w, "()))\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n    params := []interface{}{")
//line model.go.ego:56
_, _ = fmt.Fprintf(w, "%v",  fieldString("s.", m.FieldSlice(), "") )
//line model.go.ego:56
_, _ = fmt.Fprintf(w, "}\n    params = append(params, s.ID)\n\n\t\t_, err = stmt.Exec(params)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n    return nil\n}\n\nfunc (s *")
//line model.go.ego:67
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:67
_, _ = fmt.Fprintf(w, ") Insert(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(\"INSERT INTO ")
//line model.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:68
_, _ = fmt.Fprintf(w, "(")
//line model.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  fieldString("", m.FieldSliceWithoutID(), "") )
//line model.go.ego:68
_, _ = fmt.Fprintf(w, ") VALUES(")
//line model.go.ego:68
_, _ = fmt.Fprintf(w, "%v",  util.QuestionMarks(len(m.FieldSliceWithoutID())) )
//line model.go.ego:68
_, _ = fmt.Fprintf(w, ")\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\t_, err = stmt.Exec(")
//line model.go.ego:73
_, _ = fmt.Fprintf(w, "%v",  fieldString("s.", m.FieldSliceWithoutID(), "") )
//line model.go.ego:73
_, _ = fmt.Fprintf(w, ")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t  return nil\n}\n\nfunc Create")
//line model.go.ego:80
_, _ = fmt.Fprintf(w, "%v",  m.Name )
//line model.go.ego:80
_, _ = fmt.Fprintf(w, "Table(tx dbutil.DBLike) error {\n\t\tstmt, err := tx.Prepare(`")
//line model.go.ego:81
_, _ = fmt.Fprintf(w, "%v",  sqlgen.CreateTable(m) )
//line model.go.ego:81
_, _ = fmt.Fprintf(w, "`)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\t_, err = stmt.Exec()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t  return nil\n}\n")
return nil
}
