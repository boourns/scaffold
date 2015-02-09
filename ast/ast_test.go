package ast

import "testing"

func TestAstModelParses(t *testing.T) {
	model := Parse("test/model.go", "User")

	if len(model.Fields) != 2 {
		t.Errorf("Expected 2 fields")
	}

	/*if model.Fields[0] != Field{Name: "Name", Type: "string"} || model.Fields[1] != Field{Name: "Admin", Type: "bool", Tag: "`yolo`"} {
		fmt.Printf("model.Fields = %#v, unexpected", model.Fields)
	}*/
}
