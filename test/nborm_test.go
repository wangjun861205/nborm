package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/wangjun861205/nborm"
)

func TestNBorm(t *testing.T) {
	bookInfos := MakeBookInfoList()
	tag := NewTag()
	err := nborm.JoinQuery(&bookInfos, tag.Tag.Eq("文学"), nil, nil, bookInfos[0].Tag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(bookInfos))
}
