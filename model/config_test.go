package model

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	configFile := `[
		{
		  "type": "go",
		  "output": ".",
		  "suffix": "sql"
		},
		{
		  "type": "ts",
		  "output": "ts/project/model"
		}
	  ]`
	modelConfigs := parse([]byte(configFile))
	if len(modelConfigs) != 2 {
		t.Errorf("Expected 2 ModelConfig, got %d", len(modelConfigs))
	}
	mc := modelConfigs[0]
	if mc.ModelType != "go" || mc.OutputDir != "." || mc.FileSuffix != "sql" {
		t.Errorf("Go model Type parsed incorrectly %v", mc)
	}
	mc = modelConfigs[1]
	if mc.ModelType != "ts" || mc.OutputDir != "ts/project/model" || len(mc.FileSuffix) != 0 {
		t.Errorf("Go model Type parsed incorrectly %v", mc)
	}
}

func TestEnvironmentVariableAreExpanded(t *testing.T) {
	os.Setenv("SWEET", "dude")
	os.Setenv("DUDE", "sweet")
	expanded := string(expandEnvironmentVariable([]byte("${DUDE}/models/${SWEET}")))
	if expanded != "sweet/models/dude" {
		t.Errorf("sweet/models/dude, got %s", expanded)
	}
}

func TestGoModelMustSpecifySuffix(t *testing.T) {
	mc := ModelConfig{
		ModelType: "go",
		OutputDir: ".",
	}
	if mc.validate() {
		t.Error("Expected validate() to fail because FileSuffix wasn't specified")
	}
}

func TestTSModelValidWithoutSuffix(t *testing.T) {
	mc := ModelConfig{
		ModelType: "ts",
		OutputDir: ".",
	}
	if !mc.validate() {
		t.Error("Expected config to be valid")
	}
}

func TestModelConfigInvalidWithoutOutputDir(t *testing.T) {
	mc := ModelConfig{
		ModelType: "ts",
	}
	if mc.validate() {
		t.Error("Expected config to be invalid")
	}
}

func TestModelConfigInvalidWithoutModelType(t *testing.T) {
	mc := ModelConfig{
		OutputDir: ".",
	}
	if mc.validate() {
		t.Error("Expected config to be invalid")
	}
}
