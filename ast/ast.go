package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Model struct {
	Name            string
	Package         string
	Fields          []Field
	parseNextStruct bool
	parsedOverrides bool
}

func Parse(filename string, name string) *Model {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Println("Oops! Can't parse the source: %vn", err)
		return nil
	}

	model := &Model{Name: name}
	ast.Walk(model, f)

	return model
}

func (v *Model) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.Ident:
		if t.Name == v.Name {
			v.parseNextStruct = true
		}
	case *ast.StructType:
		if v.parseNextStruct {
			for _, inp := range t.Fields.List {
				var out Field
				if inp.Type != nil {
					typ, ok := (inp.Type).(*ast.Ident)
					if ok {
						out.Type = typ.Name
					}
				}
				if inp.Tag != nil {
					out.Tag = strings.Replace(inp.Tag.Value, "`", "", -1)
				}
				if len(inp.Names) == 1 {
					out.Name = inp.Names[0].Name
				} else {
					panic("Couldn't find field name")
				}
				v.Fields = append(v.Fields, out)
			}
			v.parseNextStruct = false
		}
	}

	return v
}

func (m *Model) FieldSlice() []string {
	out := []string{}
	for _, v := range m.Fields {
		out = append(out, v.Name)
	}
	return out
}

func (m *Model) FieldSliceWithoutID() []string {
	out := []string{}
	for _, v := range m.Fields {
		if v.Name != "ID" {
			out = append(out, v.Name)
		}
	}
	return out
}

func (m *Model) parseOverrides() {
	for _, f := range m.Fields {
		f.parseOverrides()
	}
}
