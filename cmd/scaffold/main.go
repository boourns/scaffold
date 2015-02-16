package main

import (
	"flag"
	"fmt"
	"github.com/boourns/scaffold/ast"
)

var fileName string
var structName string
var scaffoldName string

func init() {
	flag.StringVar(&fileName, "file", "", "input struct filename")
	flag.StringVar(&structName, "struct", "", "input struct name")
	flag.StringVar(&scaffoldName, "scaffold", "", "Scaffold name")
}

func main() {
	model := ast.Parse(fileName, structName)
	fmt.Printf("%#v", model)
}
