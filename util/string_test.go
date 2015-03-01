package util

import (
	"testing"
)

func TestStringJoin(t *testing.T) {
	in := []string{"A", "B", "C"}

	out := StringJoin(in, ", ")

	if out != "A, B, C" {
		t.Errorf("String Join failed, returned %s", out)
	}
}
