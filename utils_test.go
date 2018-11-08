package nborm

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	s := "BookDB__Tag"
	fmt.Println(toSnakeCase(s))
}
