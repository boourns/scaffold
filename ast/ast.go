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
	parsedStruct    bool
	parsedOverrides bool
}

func Parse(filename string) *Model {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Printf("Oops! Can't parse the source: %v\n", err)
		return nil
	}

	model := &Model{}
	ast.Walk(model, f)

	return model
}

func (v *Model) Visit(node ast.Node) (w ast.Visitor) {
	if v.parsedStruct == false {
		switch t := node.(type) {
		case *ast.File:
			v.Package = t.Name.Name

		case *ast.TypeSpec:
			str, ok1 := t.Type.(*ast.StructType)

			if ok1 {
				v.Name = t.Name.Name

				for _, inp := range str.Fields.List {
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
				v.parsedStruct = true
			}
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

func (m *Model) FieldSliceInCamelCase() []string {
	out := []string{}
	for _, v := range m.Fields {
		out = append(out, v.NameInCamelCase())
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
