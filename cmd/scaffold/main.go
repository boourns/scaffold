package main

import (
	"flag"
	"fmt"
	"github.com/boourns/scaffold/ast"
	"github.com/boourns/scaffold/controller"
	"github.com/boourns/scaffold/model"
	"io/ioutil"
)

var inFileName string
var outFileName string
var structName string
var scaffoldName string

var scaffolds = map[string]ast.Scaffold{
	"controller": controller.Scaffold,
	"model":      model.Scaffold,
}

func init() {
	flag.StringVar(&inFileName, "in", "", "input struct filename")
	flag.StringVar(&outFileName, "out", "", "output .go filename")
	flag.StringVar(&structName, "struct", "", "input struct name for scaffolding")
	flag.StringVar(&scaffoldName, "scaffold", "", "Scaffold name")
}

func main() {
	flag.Parse()

	if inFileName == "" || outFileName == "" || structName == "" || scaffoldName == "" {
		fmt.Println("Required parameters missing.")
		flag.PrintDefaults()
		return
	}

	scaffold, ok := scaffolds[scaffoldName]
	if !ok {
		fmt.Println("Invalid scaffold.  Valid scaffolds are:")
		for name, s := range scaffolds {
			fmt.Printf("%s: %s\n", name, s.Description())
		}
	}

	fmt.Printf("- Parsing %s:%s\n", inFileName, structName)
	model := ast.Parse(inFileName, structName)

	fmt.Printf("- Generating %s\n", scaffoldName)
	output, err := scaffold.Generate(model)
	if err != nil {
		fmt.Printf("Error generating %s: %s\n", scaffoldName, err)
	}

	fmt.Printf("- Saving as %s\n", outFileName)
	err = ioutil.WriteFile(outFileName, []byte(output), 0644)

	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
	}
}
