package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/wangjun861205/nborm"
)

func TestNBorm(t *testing.T) {
	bookInfos := NewBookInfoList()
	err := nborm.QueryOne(bookInfos, bookInfos.Title.Contains("你好"))
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range bookInfos.List {
		tags := NewTagList()
		err := info.Tag.All(tags, nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, tag := range tags.List {
			fmt.Println(info.Title.Get())
			fmt.Println(tag.Tag.Get())
		}
	}
}
