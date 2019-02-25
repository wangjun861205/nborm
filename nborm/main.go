package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wangjun861205/nborm"
)

func main() {
	defer nborm.CloseConns()
	var ignoreConstraintError bool
	flag.BoolVar(&ignoreConstraintError, "i", false, "ignore the errors of adding constraints to tables")
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("nborm error: require path and definition files")
		return
	}
	path := flag.Args()[0]
	os.Remove(filepath.Join(path, "modelMethods.go"))
	nborm.CleanSchemaCache()
	attrMap, err := nborm.ParseComment(path, flag.Args()[1:]...)
	if err != nil {
		panic(err)
	}
	for _, f := range flag.Args()[1:] {
		err := nborm.ParseAndCreate(filepath.Join(path, f), attrMap, ignoreConstraintError)
		if err != nil {
			panic(err)
		}
	}
	nborm.CreateMethodFile(path)
	nborm.CreateSchemaJSON(path)
}
