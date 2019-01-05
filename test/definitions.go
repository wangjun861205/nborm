package test

import (
	"github.com/wangjun861205/nborm"
)

//DB:nborm_test
//Tab:auth
//PrimaryKey:ID
type Auth struct {
	nborm.ModelStatus
	ID            nborm.IntField      `json:"id" auto_increment:"true"`
	Username      nborm.StringField   `json:"username"`
	Password      nborm.StringField   `json:"password"`
	Phone         nborm.StringField   `json:"phone" primary_key:"true"`
	Status        nborm.IntField      `json:"status" default_value:"d'1'"`
	SessionId     nborm.StringField   `json:"session_id" nullable:"true"`
	ExpireTime    nborm.DatetimeField `json:"expire_time" nullable:"true"`
	Email         nborm.StringField   `json:"email" unique:"true"`
	CreateTime    nborm.DatetimeField `json:"create_time" default_value:"@'CURRENT_TIMESTAMP'"`
	LastLoginTime nborm.DatetimeField `json:"last_login_time" nullable:"true"`
	IsSuperUser   nborm.BoolField     `json:"is_super_user" default_value:"b'false'"`
}

//DB:nborm_test
//Tab:book
//PrimaryKey:ID
//UniqueKey:UniqueCode
type Book struct {
	nborm.ModelStatus
	ID         nborm.IntField    `json:"id" auto_increment:"true"`
	Isbn       nborm.StringField `json:"isbn"`
	Volume     nborm.IntField    `json:"volume" default_value:"d'1'"`
	UniqueCode nborm.StringField `json:"unique_code"`
	Status     nborm.IntField    `json:"status" default_value:"d'1'"`
	BookInfo   nborm.ForeignKey  `json:"book_info" on:"Isbn = BookInfo.Isbn"`
}

//DB:nborm_test
//Tab:book_info
//PrimaryKey:ID
//UniqueKey:Isbn
type BookInfo struct {
	nborm.ModelStatus
	ID           nborm.IntField          `json:"id" auto_increment:"true"`
	Title        nborm.StringField       `json:"title" nullable:"true"`
	Price        nborm.IntField          `json:"price" nullable:"true"`
	Author       nborm.StringField       `json:"author" nullable:"true"`
	Publisher    nborm.StringField       `json:"publisher" nullable:"true"`
	Series       nborm.StringField       `json:"series" nullable:"true"`
	Isbn         nborm.StringField       `json:"isbn" primary_key:"true"`
	PublishDate  nborm.DateField         `json:"publish_date" nullable:"true"`
	Binding      nborm.StringField       `json:"binding" nullable:"true"`
	Format       nborm.StringField       `json:"format" nullable:"true"`
	Pages        nborm.IntField          `json:"pages" nullable:"true"`
	WordCount    nborm.IntField          `json:"word_count" nullable:"true"`
	ContentIntro nborm.StringField       `json:"content_intro" nullable:"true"`
	AuthorIntro  nborm.StringField       `json:"author_intro" nullable:"true"`
	Menu         nborm.StringField       `json:"menu" nullable:"true"`
	Image        nborm.StringField       `json:"image" nullable:"true"`
	IsOk         nborm.BoolField         `json:"is_ok" default_value:"b'true'"`
	Book         nborm.ReverseForeignKey `json:"book" on:"Isbn = Book.Isbn"`
	Tag          nborm.ManyToMany        `json:"tag" on:"Isbn = Tag.ID"`
}

//DB:nborm_test
//Tab:tag
//PrimaryKey:ID
type Tag struct {
	nborm.ModelStatus
	ID       nborm.IntField    `json:"id" auto_increment:"true" primary_key:"true"`
	Tag      nborm.StringField `json:"tag"`
	BookInfo nborm.ManyToMany  `json:"book_info" on:"ID = BookInfo.Isbn"`
}
