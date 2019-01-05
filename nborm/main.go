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
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("nborm error: require path and definition files")
		return
	}
	path := flag.Args()[0]
	os.Remove(filepath.Join(path, "modelMethods.go"))
	if err := nborm.ParseComment(path); err != nil {
		panic(err)
	}
	for _, f := range flag.Args()[1:] {
		err := nborm.ParseAndCreate(filepath.Join(path, f))
		if err != nil {
			panic(err)
		}
	}
	nborm.CreateMethodFile(path)
	nborm.CreateSchemaJSON(path)
}
