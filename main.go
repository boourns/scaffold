package main

import (
	"flag"
	"fmt"
	"os"
	"scaffold/ast"
	"scaffold/controller"
	"scaffold/model"
	"scaffold/static"
)

var flags *flag.FlagSet

var scaffolds = map[string]ast.Scaffold{
	"controller": controller.Scaffold,
	"model":      model.Scaffold,
	"static":     static.Scaffold,
}

func init() {
	flags = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Scaffold missing. Usage:", os.Args[0], "<scaffold name>\n")
		printValidScaffolds()

		flags.PrintDefaults()
		return
	}

	scaffoldName := os.Args[1]
	scaffold, ok := scaffolds[scaffoldName]
	if !ok {
		fmt.Println("Invalid scaffold: %s", scaffoldName)
		printValidScaffolds()
	}

	fmt.Printf("- Running %s\n", scaffoldName)

	err := scaffold.Generate(flags)

	if err != nil {
		fmt.Printf("Error generating %s: %s\n", scaffoldName, err)
	}
}

func printValidScaffolds() {
	fmt.Println("Valid scaffolds are:")

	for name, s := range scaffolds {
		fmt.Printf("%s: %s\n", name, s.Description())
	}
}
