package model

import (
	"bytes"
	"github.com/boourns/scaffold/ast"
	"fmt"
	"io/ioutil"
	"strings"
)

type model struct{}

func (c model) Description() string {
	return "Generate JSON REST endpoints"
}

func (c model) Generate(m *ast.Model) (error) {
	out := bytes.NewBuffer(nil)
	err := modelTemplate(out, m)

	outFileName := fmt.Sprintf("%s-model.go", strings.ToLower(m.Name))

	fmt.Printf("- Saving as %s\n", outFileName)
	err = ioutil.WriteFile(outFileName, []byte(out.Bytes()), 0644)

	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
	}

	return err
}

var Scaffold = model{}
