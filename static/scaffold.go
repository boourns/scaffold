package static

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type scaffold struct{}

func (c scaffold) Description() string {
	return "HTTP Handler: Serve static files out of ./static"
}

var showHelp bool
var dirName string
var packageName string

func (c scaffold) Generate(flags *flag.FlagSet) error {
	flags.StringVar(&dirName, "dir", "static", "Directory name")
	flags.BoolVar(&showHelp, "h", false, "Show help")
	flags.StringVar(&packageName, "package", "main", "Package name for http handler")
	flags.Parse(os.Args[2:])

	if showHelp {
		fmt.Printf(c.Details())
		flags.PrintDefaults()
		return nil
	}

	fmt.Println("- Creating directory:", dirName)
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("Error creating directory: %s", err)
	}

	out := bytes.NewBuffer(nil)
	err = staticTemplate(out, dirName, packageName)
	if err != nil {
		return fmt.Errorf("Error generating static server template: %s", err)
	}

	outFileName := "static.go"

	fmt.Println("- Saving as", outFileName)
	err = ioutil.WriteFile(outFileName, []byte(out.Bytes()), 0644)

	if err != nil {
		return fmt.Errorf("Error writing file: %s", err)
	}

	fmt.Printf("\nInstructions for use:\n")
	fmt.Printf("Put static content in %s.\n", dirName)
	fmt.Printf("static.go adds an HTTP handler to serve the content from /%s.\n", dirName)
	fmt.Printf("Start HTTP server normally like:\n  http.ListenAndServe(\":3000\", nil)\n\n")

	return nil
}

var Scaffold = scaffold{}

func (c scaffold) Details() string {
	return ""
}

func printError(str string) error {
	fmt.Printf(Scaffold.Details())
	return fmt.Errorf(str)
}
