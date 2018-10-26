package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
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

type config struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Address   string   `json:"address"`
	Databases []string `json:"databases"`
	Package   string   `json:"package"`
}

func main() {
	defer nborm.CloseConns()
	pkg := flag.String("P", "", "pakcage name")
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	address := flag.String("a", "", "database address(<host>:<port>)")
	dbs := make(dbNames, 0, 8)
	flag.Var(&dbs, "d", "database name")
	cfgFile := flag.String("c", "", "json config file")
	flag.Parse()
	if *cfgFile != "" {
		f, err := os.Open(*cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		cfg := config{}
		err = json.Unmarshal(b, &cfg)
		if err != nil {
			log.Fatal(err)
		}
		username = &cfg.Username
		password = &cfg.Password
		address = &cfg.Address
		dbs = dbNames(cfg.Databases)
		pkg = &cfg.Package
	}
	if *pkg == "" || *username == "" || *password == "" || *address == "" || len(dbs) == 0 {
		log.Fatal("nborm error: invalid database dsn")
	}
	nborm.RegisterDB(*username, *password, *address, "information_schema")
	nborm.GetDBInfo(dbs...)
	nborm.MarshalDBInfo()
	nborm.GenDef(*pkg, *username, *password, *address)
}
