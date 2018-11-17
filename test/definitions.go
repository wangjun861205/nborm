package test

import (
	"github.com/wangjun861205/nborm"
)

type Auth struct {
	Id            *nborm.IntField      `json:"id" auto_increment:"true"`
	Username      *nborm.StringField   `json:"username"`
	Password      *nborm.StringField   `json:"password"`
	Phone         *nborm.StringField   `json:"phone" primary_key:"true"`
	Status        *nborm.IntField      `json:"status"`
	SessionId     *nborm.StringField   `json:"session_id"`
	ExpireTime    *nborm.DatetimeField `json:"expire_time"`
	Email         *nborm.StringField   `json:"email" unique:"true"`
	CreateTime    *nborm.DatetimeField `json:"create_time"`
	LastLoginTime *nborm.DatetimeField `json:"last_login_time"`
	IsSuperUser   *nborm.BoolField     `json:"is_super_user"`
	IsSync        nborm.SyncFlag
}

func (m *Auth) DB() string {
	return "bk_dalian"
}

func (m *Auth) Tab() string {
	return "auth"
}

type AuthList []*Auth

func MakeAuthList() AuthList {
	l := make(AuthList, 0, 64)
	nborm.InitSlice(&l)
	return l
}

func (l *AuthList) DB() string {
	return "bk_dalian"
}

func (l *AuthList) Tab() string {
	return "auth"
}

type Book struct {
	Id         *nborm.IntField    `json:"id" auto_increment:"true"`
	Isbn       *nborm.StringField `json:"isbn"`
	Volume     *nborm.IntField    `json:"volume"`
	UniqueCode *nborm.StringField `json:"unique_code" primary_key:"true"`
	Status     *nborm.IntField    `json:"status"`
	BookInfo   *nborm.ForeignKey  `source_column:"isbn" destination_column:"bk_dalian.book_info.isbn" json:"book_info"`
	IsSync     nborm.SyncFlag
}

func NewBook() *Book {
	b := &Book{}
	nborm.InitModel(b)
	return b
}

func (m *Book) DB() string {
	return "bk_dalian"
}

func (m *Book) Tab() string {
	return "book"
}

type BookList []*Book

func MakeBookList() BookList {
	l := make(BookList, 0, 64)
	nborm.InitSlice(&l)
	return l
}

func (l *BookList) DB() string {
	return "bk_dalian"
}

func (m *BookList) Tab() string {
	return "book"
}

type BookInfo struct {
	Id           *nborm.IntField          `json:"id" auto_increment:"true"`
	Title        *nborm.StringField       `json:"title"`
	Price        *nborm.IntField          `json:"price"`
	Author       *nborm.StringField       `json:"author"`
	Publisher    *nborm.StringField       `json:"publisher"`
	Series       *nborm.StringField       `json:"series"`
	Isbn         *nborm.StringField       `json:"isbn" primary_key:"true"`
	PublishDate  *nborm.DateField         `json:"publish_date"`
	Binding      *nborm.StringField       `json:"binding"`
	Format       *nborm.StringField       `json:"format"`
	Pages        *nborm.IntField          `json:"pages"`
	WordCount    *nborm.IntField          `json:"word_count"`
	ContentIntro *nborm.StringField       `json:"content_intro"`
	AuthorIntro  *nborm.StringField       `json:"author_intro"`
	Menu         *nborm.StringField       `json:"menu"`
	Image        *nborm.StringField       `json:"image"`
	IsOk         *nborm.BoolField         `json:"is_ok"`
	Book         *nborm.ReverseForeignKey `json:"book" source_column:"isbn" destination_column:"bk_dalian.book.isbn"`
	Tag          *nborm.ManyToMany        `json:"tag" source_column:"isbn" middle_database:"bk_dalian" middle_table:"book_info__tag" middle_left_column:"book_info__isbn" middle_right_column:"tag__id" destination_column:"bk_dalian.tag.id"`
	IsSync       nborm.SyncFlag
}

func NewBookInfo() *BookInfo {
	m := &BookInfo{}
	nborm.InitModel(m)
	return m
}

func (m *BookInfo) DB() string {
	return "bk_dalian"
}

func (m *BookInfo) Tab() string {
	return "book_info"
}

type BookInfoList []*BookInfo

func MakeBookInfoList() BookInfoList {
	l := make(BookInfoList, 0, 64)
	nborm.InitSlice(&l)
	return l
}

func (l *BookInfoList) DB() string {
	return "bk_dalian"
}

func (l *BookInfoList) Tab() string {
	return "book_info"
}

type Tag struct {
	Id       *nborm.IntField    `json:"id" auto_increment:"true"`
	Tag      *nborm.StringField `json:"tag" primary_key:"true"`
	BookInfo *nborm.ManyToMany  `json:"book_info" source_column:"id" middle_database:"bk_dalian" middle_table:"book_info__tag" middle_left_column:"tag__id" middle_right_column:"book_info__isbn" destination_column:"bk_dalian.book_info.isbn"`
	IsSync   nborm.SyncFlag
}

func NewTag() *Tag {
	m := &Tag{}
	nborm.InitModel(m)
	return m
}

func (m *Tag) DB() string {
	return "bk_dalian"
}

func (m *Tag) Tab() string {
	return "tag"
}

type TagList []*Tag

func MakeTagList() TagList {
	l := make(TagList, 0, 64)
	nborm.InitSlice(&l)
	return l
}

func (l *TagList) DB() string {
	return "bk_dalian"
}

func (l *TagList) Tab() string {
	return "tag"
}
