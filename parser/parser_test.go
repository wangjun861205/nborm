package parser

import (
	"log"
	"testing"
)

func TestSchema(t *testing.T) {
	schema, err := parse("db.json")
	if err != nil {
		log.Fatal(err)
	}
	err = create(schema, "wangjun", "Wt20110523", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}
}
