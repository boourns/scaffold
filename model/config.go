package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type ModelConfig struct {
	ModelType  string `json:"type"`
	OutputDir  string `json:"output"`
	FileSuffix string `json:"suffix"`
}

func Load(path string) []ModelConfig {
	b := read(path)
	mcs := parse(b)
	for _, mc := range mcs {
		if !mc.validate() {
			return []ModelConfig{}
		}
	}
	return mcs
}

func (mf ModelConfig) Path() string {
	return string(expandEnvironmentVariable([]byte(mf.OutputDir)))
}

func (mf ModelConfig) validate() bool {
	if len(mf.OutputDir) == 0 || len(mf.ModelType) == 0 {
		log.Println("ModelType and OutputDir cannot be empty")
		return false
	}
	if strings.ToLower(mf.ModelType) == "go" && len(mf.FileSuffix) == 0 {
		log.Println("Go models must specify a FileSuffix to avoid overwriting original model file")
		return false
	}
	return true
}

func read(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("config file cannot be read %s\n", path)
	}
	return b
}

func parse(b []byte) []ModelConfig {
	var configs []ModelConfig
	if json.Unmarshal(b, &configs) != nil {
		log.Fatalln("Unable to parse config file")
	}
	return configs
}

func expandEnvironmentVariable(b []byte) []byte {
	//not happy about the []byte string madness in here.
	expanded := string(b)
	re := regexp.MustCompile(`\${[^\${}]+}`)
	matches := re.FindAll(b, -1)
	for _, m := range matches {
		envVar := string(m[2 : len(m)-1])
		val := os.Getenv(envVar)
		expanded = strings.Replace(expanded, string(m), val, -1)
	}
	return []byte(expanded)
}

func generate() {

}
