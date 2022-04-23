package ast

import "strings"

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
	f.overrides = make(map[string]string)
	fields := strings.Split(f.Tag, " ")

	for _, field := range fields {
		if field == "" {
			continue
		}
		kv := strings.Split(field, ":")
		if len(kv) != 2 {
			continue
		}
		key := strings.Trim(kv[0], " \"\t")
		value := strings.Trim(kv[1], " \"\t")
		f.overrides[key] = value
	}

}
