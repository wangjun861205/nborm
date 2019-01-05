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
	[][]string{}
}

func (model *Auth) Keys() [][]string {
	[][]string{}
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
	[][]string{}
}

func (model *AuthSlice) Keys() [][]string {
	[][]string{}
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
	[][]string{
		[]string{"unique_code"},
	}
}

func (model *Book) Keys() [][]string {
	[][]string{}
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
	[][]string{
		[]string{"unique_code"},
	}
}

func (model *BookSlice) Keys() [][]string {
	[][]string{}
}

func MakeBookSlice(length, cap int) BookSlice {
	slice := make(BookSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
func (model *BookInfo__Tag) DB() string {
	return "nborm_test"
}

func (model *BookInfo__Tag) Tab() string {
	return "book_info__tag"
}

func (model *BookInfo__Tag) PrimaryKey() []string {
	return []string{"id"}
}

func (model *BookInfo__Tag) UniqueKeys() [][]string {
	[][]string{
		[]string{"book_info__isbn", "tag__id"},
	}
}

func (model *BookInfo__Tag) Keys() [][]string {
	[][]string{}
}

func NewBookInfo__Tag() *BookInfo__Tag {
	model := new(BookInfo__Tag)
	nborm.InitModel(model)
	return model
}

type BookInfo__TagSlice []*BookInfo__Tag

func (slice *BookInfo__TagSlice) DB() string {
	return "nborm_test"
}

func (slice *BookInfo__TagSlice) Tab() string {
	return "book_info__tag"
}

func (model *BookInfo__TagSlice) PrimaryKey() []string {
	return []string{"id"}
}

func (model *BookInfo__TagSlice) UniqueKeys() [][]string {
	[][]string{
		[]string{"book_info__isbn", "tag__id"},
	}
}

func (model *BookInfo__TagSlice) Keys() [][]string {
	[][]string{}
}

func MakeBookInfo__TagSlice(length, cap int) BookInfo__TagSlice {
	slice := make(BookInfo__TagSlice, length, cap)
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
	[][]string{
		[]string{"isbn"},
	}
}

func (model *BookInfo) Keys() [][]string {
	[][]string{}
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
	[][]string{
		[]string{"isbn"},
	}
}

func (model *BookInfoSlice) Keys() [][]string {
	[][]string{}
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
	[][]string{}
}

func (model *Tag) Keys() [][]string {
	[][]string{}
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
	[][]string{}
}

func (model *TagSlice) Keys() [][]string {
	[][]string{}
}

func MakeTagSlice(length, cap int) TagSlice {
	slice := make(TagSlice, length, cap)
	nborm.InitSlice(&slice)
	return slice
}
