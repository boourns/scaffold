package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AstModel struct {
	Name            string
	Fields          []Field
	parseNextStruct bool
}

type Field struct {
	Name string
	Type string
	Tag  string
}

func Parse(filename string, name string) *AstModel {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Println("Oops! Can't parse the source: %vn", err)
		return nil
	}

	model := &AstModel{Name: name}
	ast.Walk(model, f)

	return model
}

func (v *AstModel) Visit(node ast.Node) (w ast.Visitor) {
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
					out.Tag = inp.Tag.Value
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
