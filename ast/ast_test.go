package ast

import "testing"

func TestAstModelParses(t *testing.T) {
	model := Parse("test/model.go", "User")

	if len(model.Fields) != 2 {
		t.Errorf("Expected 2 fields")
	}
}

func TestAstOverrides(t *testing.T) {
	model := Parse("test/model.go", "User")

	model.parseOverrides()

	field := model.Fields[1]
	field.parseOverrides()

	if str := field.Override("field", "wrong default"); str != "yolo" {
		t.Errorf("Overrides loading failed, returned %s overrides %#v", str, field.overrides)
	}
}
