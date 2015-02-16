package ast

import (
	"encoding/json"
	"fmt"
)

type Field struct {
	Name            string
	Type            string
	Tag             string
	overrides       map[string]string
	overridesParsed bool
}

func (f *Field) Override(key string, def string) string {
	if !f.overridesParsed {
		f.parseOverrides()
	}
	val, ok := f.overrides[key]
	if !ok {
		val = def
	}
	return val
}

func (f *Field) parseOverrides() {
	f.overridesParsed = true
	fmt.Printf("tag %s", f.Tag)
	json.Unmarshal([]byte(f.Tag), &f.overrides)
}
