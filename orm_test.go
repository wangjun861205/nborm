package nborm

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/wangjun861205/nbfmt"
)

type Auth struct {
	Id            *IntField
	Username      *StringField
	Password      *StringField
	Phone         *StringField
	Status        *IntField
	SessionId     *StringField
	ExpireTime    *DatetimeField
	Email         *StringField
	CreateTime    *DatetimeField
	LastLoginTime *DatetimeField
	IsSuperUser   *BoolField
}

func NewAuth() *Auth {
	return &Auth{
		Id:            NewIntField("id", false, true, false),
		Username:      NewStringField("username", false, false, false),
		Password:      NewStringField("password", false, false, false),
		Phone:         NewStringField("phone", true, false, false),
		Status:        NewIntField("status", false, false, false),
		SessionId:     NewStringField("session_id", false, false, false),
		ExpireTime:    NewDatetimeField("expire_time", false, false, false),
		Email:         NewStringField("email", false, false, true),
		CreateTime:    NewDatetimeField("create_time", false, false, false),
		LastLoginTime: NewDatetimeField("last_login_time", false, false, false),
		IsSuperUser:   NewBoolField("is_supper_user", false, false, false),
	}
}

var AuthExample = NewAuth()

func (m *Auth) DB() string {
	return "bk_dalian"
}

func (m *Auth) Tab() string {
	return "auth"
}

func (m *Auth) Fields() []Field {
	return []Field{m.Id, m.Username, m.Password, m.Phone, m.Status, m.SessionId, m.ExpireTime, m.Email, m.CreateTime, m.LastLoginTime, m.IsSuperUser}
}

type AuthList []*Auth

func (al *AuthList) DB() string {
	return "bk_dalian"
}

func (al *AuthList) Tab() string {
	return "auth"
}

func (al *AuthList) New() Model {
	na := NewAuth()
	*al = append(*al, na)
	return na
}

func (al *AuthList) Delete(i int) {
	switch i {
	case 0:
		*al = (*al)[1:]
	case al.Len() - 1:
		*al = (*al)[:al.Len()-1]
	default:
		*al = append((*al)[:i], (*al)[i+1:]...)
	}
}

func (al *AuthList) Index(i int) Model {
	return (*al)[i]
}

func (al *AuthList) Len() int {
	return len(*al)
}

func (al *AuthList) Swap(i, j int) {
	(*al)[i], (*al)[j] = (*al)[j], (*al)[i]
}

type Book struct {
	Id         *IntField
	Isbn       *StringField
	Volume     *IntField
	UniqueCode *StringField
	Status     *IntField
	BookInfo   *ForeignKey
}

func NewBook() *Book {
	book := &Book{
		Id:         NewIntField("id", false, true, false),
		Isbn:       NewStringField("isbn", false, false, false),
		Volume:     NewIntField("volumn", false, false, false),
		UniqueCode: NewStringField("unique_code", true, false, false),
		Status:     NewIntField("status", false, false, false),
	}
	book.BookInfo = NewForeignKey(BookInfoSample, BookInfoSample.Isbn, book.Isbn)
	return book
}

func (b *Book) DB() string {
	return "bk_dalian"
}

func (b *Book) Tab() string {
	return "book"
}

func (b *Book) Fields() []Field {
	return []Field{b.Id, b.Isbn, b.Volume, b.UniqueCode, b.Status}
}

var BookSample = NewBook()

type BookInfo struct {
	Id           *IntField
	Title        *StringField
	Price        *IntField
	Author       *StringField
	Publisher    *StringField
	Series       *StringField
	Isbn         *StringField
	PublishDate  *DateField
	Binding      *StringField
	Format       *StringField
	Pages        *IntField
	WordCount    *IntField
	ContentIntro *StringField
	AuthorIntro  *StringField
	Menu         *StringField
	Image        *StringField
}

func NewBookInfo() *BookInfo {
	bookInfo := &BookInfo{
		Id:           NewIntField("id", false, true, false),
		Title:        NewStringField("title", false, false, false),
		Price:        NewIntField("price", false, false, false),
		Author:       NewStringField("author", false, false, false),
		Publisher:    NewStringField("publisher", false, false, false),
		Series:       NewStringField("series", false, false, false),
		Isbn:         NewStringField("isbn", true, false, false),
		PublishDate:  NewDateField("publish_date", false, false, false),
		Binding:      NewStringField("binding", false, false, false),
		Format:       NewStringField("format", false, false, false),
		Pages:        NewIntField("pages", false, false, false),
		WordCount:    NewIntField("word_count", false, false, false),
		ContentIntro: NewStringField("content_intro", false, false, false),
		AuthorIntro:  NewStringField("author_intro", false, false, false),
		Menu:         NewStringField("menu", false, false, false),
		Image:        NewStringField("image", false, false, false),
	}
	return bookInfo
}

func (b *BookInfo) DB() string {
	return "bk_dalian"
}

func (b *BookInfo) Tab() string {
	return "book_info"
}

func (b *BookInfo) Fields() []Field {
	return []Field{b.Id, b.Title, b.Price, b.Author, b.Publisher, b.Series, b.Isbn, b.PublishDate, b.Binding, b.Format, b.Pages, b.WordCount,
		b.ContentIntro, b.AuthorIntro, b.Menu, b.Image}
}

var BookInfoSample = NewBookInfo()

type BookInfoList []*BookInfo

func (bil *BookInfoList) DB() string {
	return "bk_dalian"
}

func (bil *BookInfoList) Tab() string {
	return "book_info"
}

func (bil *BookInfoList) Len() int {
	return len(*bil)
}

func (bil *BookInfoList) New() Model {
	bi := NewBookInfo()
	*bil = append(*bil, bi)
	return bi
}

func (bil *BookInfoList) Index(i int) Model {
	return (*bil)[i]
}

func (bil *BookInfoList) Range(f func(int, Model) error) error {
	for i, bi := range *bil {
		err := f(i, bi)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bil *BookInfoList) Delete(i int) {
	switch i {
	case 0:
		*bil = (*bil)[1:]
	case bil.Len() - 1:
		*bil = (*bil)[:bil.Len()-1]
	default:
		*bil = append((*bil)[:i], (*bil)[i+1:]...)
	}
}

func (bil *BookInfoList) Swap(i, j int) {
	(*bil)[i], (*bil)[j] = (*bil)[j], (*bil)[i]
}

type Tag struct {
	Id        *IntField
	Tag       *StringField
	BookInfos *ManyToMany
}

func NewTag() *Tag {
	tag := &Tag{
		Id:  NewIntField("id", true, true, false),
		Tag: NewStringField("tag", false, false, true),
	}
	tag.BookInfos = NewManyToMany(tag, BookInfo__TagSample, BookInfoSample, tag.Id, BookInfo__TagSample.Tag__Id,
		BookInfo__TagSample.BookInfo__Isbn, BookInfoSample.Isbn)
	return tag
}

func (t *Tag) DB() string {
	return "bk_dalian"
}

func (t *Tag) Tab() string {
	return "tag"
}

func (t *Tag) Fields() []Field {
	return []Field{t.Id, t.Tag}
}

var TagSample = NewTag()

type BookInfo__Tag struct {
	Id             *IntField
	BookInfo__Isbn *StringField
	Tag__Id        *IntField
}

func NewBookInfo__Tag() *BookInfo__Tag {
	return &BookInfo__Tag{
		Id:             NewIntField("id", true, true, false),
		BookInfo__Isbn: NewStringField("book_info__isbn", false, false, false),
		Tag__Id:        NewIntField("tag__id", false, false, false),
	}
}

func (bit *BookInfo__Tag) DB() string {
	return "bk_dalian"
}

func (bit *BookInfo__Tag) Tab() string {
	return "book_info__tag"
}

func (bit *BookInfo__Tag) Fields() []Field {
	return []Field{bit.Id, bit.BookInfo__Isbn, bit.Tag__Id}
}

var BookInfo__TagSample = NewBookInfo__Tag()

func TestOrm(t *testing.T) {
	RegisterDB("wangjun", "Wt20110523", "127.0.0.1:12345", "information_schema")
	GetDBInfo("bk_dalian")
	f, err := os.OpenFile("dbinfo.json", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := json.MarshalIndent(dbInfo, "\t", "\t")
	if err != nil {
		panic(err)
	}
	jn, err := f.Write(b)
	if err != nil {
		panic(err)
	}
	f.Truncate(int64(jn))
	s, err := nbfmt.Fmt(template, map[string]interface{}{"pkgName": "nborm", "infos": dbInfo})
	if err != nil {
		panic(err)
	}
	df, err := os.OpenFile("definitions.go", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer df.Close()
	dn, err := df.WriteString(s)
	if err != nil {
		panic(err)
	}
	df.Truncate(int64(dn))
}
