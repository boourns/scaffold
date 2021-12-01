package util

import "testing"

func TestQuestionMarks(t *testing.T) {
	expected := "?,?,?,?,?,?,?,?,?,?"
	got := QuestionMarks(10)
	if expected != got {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}

func TestTransform(t *testing.T) {
	expected := []string{"<one>", "<two>", "<three>"}
	got := Transform("<", []string{"one", "two", "three"}, ">")
	if len(expected) != len(got) {
		t.Errorf("Expected []string of with len %d, got len %d", len(expected), len(got))
	}
	for i, v := range expected {
		if v != got[i] {
			t.Errorf("Expected %s, but got %s", v, got[i])
		}
	}
}
