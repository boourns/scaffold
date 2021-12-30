package model

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/boourns/scaffold/ast"
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
		fmt.Print(c.Details())
		flags.PrintDefaults()
		return nil
	}

	if inFileName == "" {
		return printError("Missing input file.")
	}

	fmt.Println("- Parsing", inFileName)
	m := ast.Parse(inFileName)
	fmt.Println("- Found package", m.Package, "struct", m.Name)

	modelCode, err := createModel(m)
	if err != nil {
		return fmt.Errorf("error generating model file: %s", err)
	}

	formatted, err := format.Source(modelCode)
	if err != nil {
		fmt.Printf("error formatting file: %s\n", err)
		formatted = modelCode
	}

	outFileName := fmt.Sprintf("%s_sql.go", strings.ToLower(m.Name))

	fmt.Println("- Saving as", outFileName)
	err = ioutil.WriteFile(outFileName, formatted, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}

	return err
}

//go:embed "model.template"
var templatePath embed.FS
var modelTemplate *template.Template

func createModel(m *ast.Model) ([]byte, error) {
	if modelTemplate == nil {
		modelTemplate = template.Must(template.ParseFS(templatePath, "model.template"))
	}

	data := ModelViewData{
		Model: m,
	}

	out := &bytes.Buffer{}
	err := modelTemplate.Execute(out, data)
	return out.Bytes(), err
}

var Scaffold = model{}

func (c model) Details() string {
	return ""
}

func printError(str string) error {
	fmt.Println(Scaffold.Details())
	return fmt.Errorf(str)
}
