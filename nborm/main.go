package main

import (
	"flag"
	"strings"

	"github.com/wangjun861205/nborm"
)

type dbNames []string

func (ns *dbNames) String() string {
	return strings.Join(*ns, ", ")
}

func (ns *dbNames) Set(name string) error {
	*ns = append(*ns, name)
	return nil
}

func main() {
	defer nborm.CloseConns()
	pkg := flag.String("P", "", "pakcage name")
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	address := flag.String("a", "", "database address(<host>:<port>)")
	dbs := make(dbNames, 0, 8)
	flag.Var(&dbs, "d", "database name")
	flag.Parse()
	if pkg == nil || username == nil || password == nil || address == nil || len(dbs) == 0 {
		panic("nborm error: invalid database dsn")
	}
	nborm.RegisterDB(*username, *password, *address, "information_schema")
	nborm.GetDBInfo(dbs...)
	nborm.MarshalDBInfo()
	nborm.GenDef(*pkg, *username, *password, *address)
}
