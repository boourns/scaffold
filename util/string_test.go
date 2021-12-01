package util

import "testing"

func TestQuestionMarks(t *testing.T) {
	expected := "?,?,?,?,?,?,?,?,?,?"
	got := QuestionMarks(10)
	if expected != QuestionMarks(3) {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}
