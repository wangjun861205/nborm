package test

import (
	"github.com/wangjun861205/nborm"
)

func (model *Auth) DB() string {
	return "nborm_test"
}

func (model *Auth) Tab() string {
	return "auth"
}

func (model *Auth) PrimaryKey() []string {
	return []string{"id"}
}

func (model *Auth) UniqueKeys() [][]string {
	return [][]string{}
}

func (model *Auth) Keys() [][]string {
	return [][]string{}
}

func NewAuth() *Auth {
	model := new(Auth)
	nborm.InitModel(model)
	return model
}

type AuthSlice []*Auth

func (slice *AuthSlice) DB() string {
	return "nborm_test"
}

func (slice *AuthSlice) Tab() string {
	return "auth"
}

func (model *AuthSlice) PrimaryKey() []string {
	return []string{"id"}
}

func (model *AuthSlice) UniqueKeys() [][]string {
	return [][]string{}
}

func (model *AuthSlice) Keys() [][]string {
	return [][]string{}
}

func MakeAuthSlice(length, cap int) AuthSlice {
	slice := make(AuthSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
func (model *Book) DB() string {
	return "nborm_test"
}

func (model *Book) Tab() string {
	return "book"
}

func (model *Book) PrimaryKey() []string {
	return []string{"id"}
}

func (model *Book) UniqueKeys() [][]string {
	return [][]string{
		[]string{"unique_code"},
	}
}

func (model *Book) Keys() [][]string {
	return [][]string{}
}

func NewBook() *Book {
	model := new(Book)
	nborm.InitModel(model)
	return model
}

type BookSlice []*Book

func (slice *BookSlice) DB() string {
	return "nborm_test"
}

func (slice *BookSlice) Tab() string {
	return "book"
}

func (model *BookSlice) PrimaryKey() []string {
	return []string{"id"}
}

func (model *BookSlice) UniqueKeys() [][]string {
	return [][]string{
		[]string{"unique_code"},
	}
}

func (model *BookSlice) Keys() [][]string {
	return [][]string{}
}

func MakeBookSlice(length, cap int) BookSlice {
	slice := make(BookSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
func (model *BookInfo) DB() string {
	return "nborm_test"
}

func (model *BookInfo) Tab() string {
	return "book_info"
}

func (model *BookInfo) PrimaryKey() []string {
	return []string{"id"}
}

func (model *BookInfo) UniqueKeys() [][]string {
	return [][]string{
		[]string{"isbn"},
	}
}

func (model *BookInfo) Keys() [][]string {
	return [][]string{}
}

func NewBookInfo() *BookInfo {
	model := new(BookInfo)
	nborm.InitModel(model)
	return model
}

type BookInfoSlice []*BookInfo

func (slice *BookInfoSlice) DB() string {
	return "nborm_test"
}

func (slice *BookInfoSlice) Tab() string {
	return "book_info"
}

func (model *BookInfoSlice) PrimaryKey() []string {
	return []string{"id"}
}

func (model *BookInfoSlice) UniqueKeys() [][]string {
	return [][]string{
		[]string{"isbn"},
	}
}

func (model *BookInfoSlice) Keys() [][]string {
	return [][]string{}
}

func MakeBookInfoSlice(length, cap int) BookInfoSlice {
	slice := make(BookInfoSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
func (model *Tag) DB() string {
	return "nborm_test"
}

func (model *Tag) Tab() string {
	return "tag"
}

func (model *Tag) PrimaryKey() []string {
	return []string{"id"}
}

func (model *Tag) UniqueKeys() [][]string {
	return [][]string{}
}

func (model *Tag) Keys() [][]string {
	return [][]string{}
}

func NewTag() *Tag {
	model := new(Tag)
	nborm.InitModel(model)
	return model
}

type TagSlice []*Tag

func (slice *TagSlice) DB() string {
	return "nborm_test"
}

func (slice *TagSlice) Tab() string {
	return "tag"
}

func (model *TagSlice) PrimaryKey() []string {
	return []string{"id"}
}

func (model *TagSlice) UniqueKeys() [][]string {
	return [][]string{}
}

func (model *TagSlice) Keys() [][]string {
	return [][]string{}
}

func MakeTagSlice(length, cap int) TagSlice {
	slice := make(TagSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
