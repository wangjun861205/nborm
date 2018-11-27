package test

import (
	"github.com/wangjun861205/nborm"
)

type Auth struct {
	Id            *nborm.IntField      `json:"id" auto_increment:"true"`
	Username      *nborm.StringField   `json:"username"`
	Password      *nborm.StringField   `json:"password"`
	Phone         *nborm.StringField   `json:"phone" primary_key:"true"`
	Status        *nborm.IntField      `json:"status" default_value:"d'1'"`
	SessionId     *nborm.StringField   `json:"session_id" nullable:"true"`
	ExpireTime    *nborm.DatetimeField `json:"expire_time" nullable:"true"`
	Email         *nborm.StringField   `json:"email" unique:"true"`
	CreateTime    *nborm.DatetimeField `json:"create_time" default_value:"@'CURRENT_TIMESTAMP'"`
	LastLoginTime *nborm.DatetimeField `json:"last_login_time" nullable:"true"`
	IsSuperUser   *nborm.BoolField     `json:"is_super_user" default_value:"b'false'"`
	_isSync       nborm.SyncFlag
}

func NewAuth() *Auth {
	auth := &Auth{}
	nborm.InitModel(auth)
	return auth
}

func (m *Auth) DB() string {
	return "bk_dalian"
}

func (m *Auth) Tab() string {
	return "auth"
}

type AuthList []*Auth

func MakeAuthList() AuthList {
	authList := make(AuthList, 0, 64)
	nborm.InitSlice(&authList)
	return authList
}

func (s *AuthList) DB() string {
	return "bk_dalian"
}

func (s *AuthList) Tab() string {
	return "auth"
}

type Book struct {
	Id         *nborm.IntField    `json:"id" auto_increment:"true"`
	Isbn       *nborm.StringField `json:"isbn"`
	Volume     *nborm.IntField    `json:"volume" default_value:"d'1'"`
	UniqueCode *nborm.StringField `json:"unique_code" primary_key:"true"`
	Status     *nborm.IntField    `json:"status" default_value:"d'1'"`
	BookInfo   *nborm.ForeignKey  `json:"book_info" source_column:"isbn" destination_column:"bk_dalian.book_info.isbn"`
	_isSync    nborm.SyncFlag
}

func NewBook() *Book {
	book := &Book{}
	nborm.InitModel(book)
	return book
}

func (m *Book) DB() string {
	return "bk_dalian"
}

func (m *Book) Tab() string {
	return "book"
}

type BookList []*Book

func MakeBookList() BookList {
	bookList := make(BookList, 0, 64)
	nborm.InitSlice(&bookList)
	return bookList
}

func (s *BookList) DB() string {
	return "bk_dalian"
}

func (s *BookList) Tab() string {
	return "book"
}

type BookInfo struct {
	Id           *nborm.IntField          `json:"id" auto_increment:"true"`
	Title        *nborm.StringField       `json:"title" nullable:"true"`
	Price        *nborm.IntField          `json:"price" nullable:"true"`
	Author       *nborm.StringField       `json:"author" nullable:"true"`
	Publisher    *nborm.StringField       `json:"publisher" nullable:"true"`
	Series       *nborm.StringField       `json:"series" nullable:"true"`
	Isbn         *nborm.StringField       `json:"isbn" primary_key:"true"`
	PublishDate  *nborm.DateField         `json:"publish_date" nullable:"true"`
	Binding      *nborm.StringField       `json:"binding" nullable:"true"`
	Format       *nborm.StringField       `json:"format" nullable:"true"`
	Pages        *nborm.IntField          `json:"pages" nullable:"true"`
	WordCount    *nborm.IntField          `json:"word_count" nullable:"true"`
	ContentIntro *nborm.StringField       `json:"content_intro" nullable:"true"`
	AuthorIntro  *nborm.StringField       `json:"author_intro" nullable:"true"`
	Menu         *nborm.StringField       `json:"menu" nullable:"true"`
	Image        *nborm.StringField       `json:"image" nullable:"true"`
	IsOk         *nborm.BoolField         `json:"is_ok" default_value:"b'true'"`
	Book         *nborm.ReverseForeignKey `json:"book" source_column:"isbn" destination_column:"bk_dalian.book.isbn"`
	Tag          *nborm.ManyToMany        `json:"tag" source_column:"isbn" middle_database:"bk_dalian" middle_table:"book_info__tag" middle_left_column:"book_info__isbn" middle_right_column:"tag__id" destination_column:"bk_dalian.tag.id"`
	_isSync      nborm.SyncFlag
}

func NewBookInfo() *BookInfo {
	bookInfo := &BookInfo{}
	nborm.InitModel(bookInfo)
	return bookInfo
}

func (m *BookInfo) DB() string {
	return "bk_dalian"
}

func (m *BookInfo) Tab() string {
	return "book_info"
}

type BookInfoList []*BookInfo

func MakeBookInfoList() BookInfoList {
	bookInfoList := make(BookInfoList, 0, 64)
	nborm.InitSlice(&bookInfoList)
	return bookInfoList
}

func (s *BookInfoList) DB() string {
	return "bk_dalian"
}

func (s *BookInfoList) Tab() string {
	return "book_info"
}

type Tag struct {
	Id       *nborm.IntField    `json:"id" primary_key:"true"`
	Tag      *nborm.StringField `json:"tag"`
	BookInfo *nborm.ManyToMany  `json:"book_info" source_column:"id" middle_database:"bk_dalian" middle_table:"book_info__tag" middle_left_column:"tag__id" middle_right_column:"book_info__isbn" destination_column:"bk_dalian.book_info.isbn"`
	_isSync  nborm.SyncFlag
}

func NewTag() *Tag {
	tag := &Tag{}
	nborm.InitModel(tag)
	return tag
}

func (m *Tag) DB() string {
	return "bk_dalian"
}

func (m *Tag) Tab() string {
	return "tag"
}

type TagList []*Tag

func MakeTagList() TagList {
	tagList := make(TagList, 0, 64)
	nborm.InitSlice(&tagList)
	return tagList
}

func (s *TagList) DB() string {
	return "bk_dalian"
}

func (s *TagList) Tab() string {
	return "tag"
}
