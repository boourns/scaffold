package model

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/boourns/scaffold/ast"
	"io/ioutil"
	"os"
	"strings"
)

type model struct{}

func (c model) Description() string {
	return "Micro ORM: SQL CreateTable, Insert, Update, Select, Delete"
}

var showHelp bool
var inFileName string

func (c model) Generate(flags *flag.FlagSet) error {
	flags.StringVar(&inFileName, "in", "", "Input file name")
	flags.BoolVar(&showHelp, "h", false, "Show help")
	flags.Parse(os.Args[2:])

	if showHelp {
		fmt.Printf(c.Details())
		flags.PrintDefaults()
		return nil
	}

	if inFileName == "" {
		return printError("Missing input file.")
	}

	fmt.Println("- Parsing", inFileName)
	m := ast.Parse(inFileName)
	fmt.Println("- Found package", m.Package, "struct", m.Name)

	out := bytes.NewBuffer(nil)
	err := modelTemplate(out, m)
	if err != nil {
		return fmt.Errorf("Error generating sql template: %s", err)
	}

	outFileName := fmt.Sprintf("%s_sql.go", strings.ToLower(m.Name))

	fmt.Println("- Saving as", outFileName)
	err = ioutil.WriteFile(outFileName, []byte(out.Bytes()), 0644)

	if err != nil {
		return fmt.Errorf("Error writing file: %s", err)
	}

	return err
}

var Scaffold = model{}

func (c model) Details() string {
	return ""
}

func printError(str string) error {
	fmt.Printf(Scaffold.Details())
	return fmt.Errorf(str)
}
