package nborm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configition struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	SchemaFile string `json:"schema_file"`
}

var config configition

func init() {
	f, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	if config.Username == "" {
		panic("nborm config error: required username")
	}
	if config.Password == "" {
		panic("nborm config error: required password")
	}
	if config.Host == "" {
		panic("nborm config error: required host")
	}
	if config.Port == 0 {
		panic("nborm config error: required port")
	}
	schemaFile, err := os.Open(config.SchemaFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("nborm.init() warnning: schema.json not exists, if you are using nborm command line tool, please ignore this warnning")
			return
		}
		panic(err)
	}
	defer schemaFile.Close()
	b, err = ioutil.ReadAll(schemaFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &SchemaCache)
	if err != nil {
		panic(err)
	}
}
