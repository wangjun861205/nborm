package test

import (
	"testing"

	"github.com/wangjun861205/nborm"
)

func TestFSTool(t *testing.T) {
	tempDir, err := nborm.NewTempDir("./", "nbfstool")
	if err != nil {
		t.Fatal(err)
	}
	tempFiles, err := tempDir.CopyFiles("./definitions.go", "./nborm_test.go", "./modelMethods.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range tempFiles {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
