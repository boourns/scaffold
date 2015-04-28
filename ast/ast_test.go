package ast

import "testing"

func TestAstModelParses(t *testing.T) {
	model := Parse("test/model.go")

	if len(model.Fields) != 2 {
		t.Errorf("Expected 2 fields")
	}

	if model.Package != "user" {
		t.Errorf("Expected package name user, found %s", model.Package)
	}

	if model.Name != "User" {
		t.Errorf("Expected struct named User, found %s", model.Name)
	}
}

func TestAstOverrides(t *testing.T) {
	model := Parse("test/model.go")

	model.parseOverrides()

	field := model.Fields[1]
	field.parseOverrides()

	if str := field.Override("field", "wrong default"); str != "yolo" {
		t.Errorf("Overrides loading failed, returned %s overrides %#v", str, field.overrides)
	}
}
