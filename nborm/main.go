package main

import (
	"flag"
	"fmt"

	"github.com/wangjun861205/nborm"
)

func main() {
	defer nborm.CloseConns()
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("nborm error: require definition files")
		return
	}
	for _, f := range flag.Args() {
		err := nborm.ParseAndCreate(f)
		if err != nil {
			fmt.Println(err)
		}
	}
}
