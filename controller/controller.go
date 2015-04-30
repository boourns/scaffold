package controller

import (
	"flag"
)

type controller struct{}

func (c controller) Description() string {
	return "Generate JSON REST endpoints"
}

func (c controller) Generate(flags *flag.FlagSet) error {
	//out := bytes.NewBuffer(nil)
	//err := GenerateIndex(out, m)

	// todo save

	return nil
}

func (c controller) Details() string {
	return ""
}

var Scaffold = controller{}
