package main

import (
	"flag"
	"fmt"
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
	if err := nborm.ParseComment(path); err != nil {
		panic(err)
	}
	for _, f := range flag.Args()[1:] {
		err := nborm.ParseAndCreate(filepath.Join(path, f))
		if err != nil {
			fmt.Println(err)
		}
	}
	nborm.CreateMethodFile(path)
}
