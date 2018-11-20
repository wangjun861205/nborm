package nborm

import (
	"fmt"
	"log"
	"testing"
)

func TestDBTool(t *testing.T) {
	err := parseDB("test/definitions.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(schemaCache.databaseMap["bk_dalian"].tableMap["auth"].columns[0])
}
