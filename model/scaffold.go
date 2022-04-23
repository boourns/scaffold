package model

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
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
var configFile string

var defaultConfig = ModelConfig{
	ModelType:  "go",
	OutputDir:  ".",
	FileSuffix: "sql",
}

func (c model) Generate(flags *flag.FlagSet) error {
	flags.StringVar(&inFileName, "in", "", "Input file name")
	flags.StringVar(&configFile, "config", "", "Config file path")
	flags.BoolVar(&showHelp, "h", false, "Show help")
	flags.Parse(os.Args[2:])

	if showHelp {
		fmt.Print(c.Details())
		flags.PrintDefaults()
		return nil
	}

	if inFileName == "" {
		inFileName = os.Getenv("GOFILE")
	}

	if inFileName == "" {
		return printError("Missing input file.")
	}

	configs := []ModelConfig{defaultConfig}
	if len(configFile) != 0 {
		configs = Load(configFile)
	}

	fmt.Println("- Parsing", inFileName)
	m := ast.Parse(inFileName)
	fmt.Println("- Found package", m.Package, "struct", m.Name)

	for _, mc := range configs {
		var err error
		switch mc.ModelType {
		case "go":
			err = writeGoFile(mc, m)
		case "ts":
			err = writeTSFile(mc, m)
		default:
			fmt.Printf("Ignoring unsupported model type: %s", mc.ModelType)
		}
		if err != nil {
			fmt.Printf("Error writing model %s of type %s [%s]", m.Name, mc.ModelType, err)
		}
	}
	return nil
}

func writeGoFile(config ModelConfig, m *ast.Model) error {
	out := bytes.NewBuffer(nil)
	modelTemplateGo(out, m)

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting file: %s", err)
	}

	outFileName := fmt.Sprintf("%s/%s.go", config.Path(), strings.ToLower(m.Name))
	if len(config.FileSuffix) != 0 {
		outFileName = fmt.Sprintf("%s/%s_%s.go", config.Path(), strings.ToLower(m.Name), strings.ToLower(config.FileSuffix))
	}

	fmt.Println("- Saving as", outFileName)
	err = ioutil.WriteFile(outFileName, formatted, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}

	return err
}

func writeTSFile(config ModelConfig, m *ast.Model) error {
	out := bytes.NewBuffer(nil)
	modelTemplateTS(out, m)

	outFileName := fmt.Sprintf("%s/%s.ts", config.Path(), strings.ToLower(m.Name))
	if len(config.FileSuffix) != 0 {
		outFileName = fmt.Sprintf("%s/%s_%s.go", config.Path(), strings.ToLower(m.Name), strings.ToLower(config.FileSuffix))
	}

	fmt.Println("- Saving as", outFileName)
	err := ioutil.WriteFile(outFileName, out.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
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
