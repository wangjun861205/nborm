package test

import (
	"encoding/json"
	"github.com/wangjun861205/nborm"
)

type Auth struct {
	Id            *nborm.IntField      `json:"id"`
	Username      *nborm.StringField   `json:"username"`
	Password      *nborm.StringField   `json:"password"`
	Phone         *nborm.StringField   `json:"phone"`
	Status        *nborm.IntField      `json:"status"`
	SessionId     *nborm.StringField   `json:"session_id"`
	ExpireTime    *nborm.DatetimeField `json:"expire_time"`
	Email         *nborm.StringField   `json:"email"`
	CreateTime    *nborm.DatetimeField `json:"create_time"`
	LastLoginTime *nborm.DatetimeField `json:"last_login_time"`
	IsSuperUser   *nborm.BoolField     `json:"is_super_user"`
	_isSync       bool
}

func NewAuth() *Auth {
	m := &Auth{}
	m.Id = nborm.NewIntField(m, "id", false, true, true)
	m.Username = nborm.NewStringField(m, "username", false, false, true)
	m.Password = nborm.NewStringField(m, "password", false, false, false)
	m.Phone = nborm.NewStringField(m, "phone", true, false, false)
	m.Status = nborm.NewIntField(m, "status", false, false, false)
	m.SessionId = nborm.NewStringField(m, "session_id", false, false, false)
	m.ExpireTime = nborm.NewDatetimeField(m, "expire_time", false, false, false)
	m.Email = nborm.NewStringField(m, "email", false, false, true)
	m.CreateTime = nborm.NewDatetimeField(m, "create_time", false, false, false)
	m.LastLoginTime = nborm.NewDatetimeField(m, "last_login_time", false, false, false)
	m.IsSuperUser = nborm.NewBoolField(m, "is_super_user", false, false, false)

	return m
}

func (m *Auth) DB() string {
	return "bk_dalian"
}

func (m *Auth) Tab() string {
	return "auth"
}

func (m *Auth) Fields() []nborm.Field {
	return []nborm.Field{
		m.Id,
		m.Username,
		m.Password,
		m.Phone,
		m.Status,
		m.SessionId,
		m.ExpireTime,
		m.Email,
		m.CreateTime,
		m.LastLoginTime,
		m.IsSuperUser,
	}
}

func (m *Auth) SetSync(b bool) {
	m._isSync = b
}

func (m *Auth) GetSync() bool {
	return m._isSync
}

type AuthList struct {
	*Auth
	List []*Auth
}

func NewAuthList() *AuthList {
	return &AuthList{NewAuth(), make([]*Auth, 0, 128)}
}

func (l *AuthList) New() nborm.Model {
	m := NewAuth()
	l.List = append(l.List, m)
	return m
}

func (l *AuthList) Len() int {
	return len(l.List)
}

func (l *AuthList) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func (l *AuthList) Index(i int) nborm.Model {
	return l.List[i]
}

func (l *AuthList) Delete(i int) {
	switch i {
	case 0:
		l.List = l.List[1:]
	case l.Len() - 1:
		l.List = l.List[:l.Len()-1]
	default:
		l.List = append(l.List[:i], l.List[i+1:]...)
	}
}

func (l *AuthList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

func (l *AuthList) Model() nborm.Model {
	return l.Auth
}

type Book struct {
	Id         *nborm.IntField    `json:"id"`
	Isbn       *nborm.StringField `json:"isbn"`
	Volume     *nborm.IntField    `json:"volume"`
	UniqueCode *nborm.StringField `json:"unique_code"`
	Status     *nborm.IntField    `json:"status"`
	BookInfo   *nborm.ForeignKey  `json:"book_info"`
	_isSync    bool
}

func NewBook() *Book {
	m := &Book{}
	m.Id = nborm.NewIntField(m, "id", false, true, true)
	m.Isbn = nborm.NewStringField(m, "isbn", false, false, false)
	m.Volume = nborm.NewIntField(m, "volume", false, false, false)
	m.UniqueCode = nborm.NewStringField(m, "unique_code", true, false, false)
	m.Status = nborm.NewIntField(m, "status", false, false, false)
	m.BookInfo = nborm.NewForeignKey(m.Isbn, func() nborm.Field { return NewBookInfo().Isbn })

	return m
}

func (m *Book) DB() string {
	return "bk_dalian"
}

func (m *Book) Tab() string {
	return "book"
}

func (m *Book) Fields() []nborm.Field {
	return []nborm.Field{
		m.Id,
		m.Isbn,
		m.Volume,
		m.UniqueCode,
		m.Status,
	}
}

func (m *Book) SetSync(b bool) {
	m._isSync = b
}

func (m *Book) GetSync() bool {
	return m._isSync
}

type BookList struct {
	*Book
	List []*Book
}

func NewBookList() *BookList {
	return &BookList{NewBook(), make([]*Book, 0, 128)}
}

func (l *BookList) New() nborm.Model {
	m := NewBook()
	l.List = append(l.List, m)
	return m
}

func (l *BookList) Len() int {
	return len(l.List)
}

func (l *BookList) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func (l *BookList) Index(i int) nborm.Model {
	return l.List[i]
}

func (l *BookList) Delete(i int) {
	switch i {
	case 0:
		l.List = l.List[1:]
	case l.Len() - 1:
		l.List = l.List[:l.Len()-1]
	default:
		l.List = append(l.List[:i], l.List[i+1:]...)
	}
}

func (l *BookList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

func (l *BookList) Model() nborm.Model {
	return l.Book
}

type BookInfo struct {
	Id           *nborm.IntField          `json:"id"`
	Title        *nborm.StringField       `json:"title"`
	Price        *nborm.IntField          `json:"price"`
	Author       *nborm.StringField       `json:"author"`
	Publisher    *nborm.StringField       `json:"publisher"`
	Series       *nborm.StringField       `json:"series"`
	Isbn         *nborm.StringField       `json:"isbn"`
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
	Book         *nborm.ReverseForeignKey `json:"book"`
	Tag          *nborm.ManyToMany        `json:"tag"`
	_isSync      bool
}

func NewBookInfo() *BookInfo {
	m := &BookInfo{}
	m.Id = nborm.NewIntField(m, "id", false, true, true)
	m.Title = nborm.NewStringField(m, "title", false, false, false)
	m.Price = nborm.NewIntField(m, "price", false, false, false)
	m.Author = nborm.NewStringField(m, "author", false, false, false)
	m.Publisher = nborm.NewStringField(m, "publisher", false, false, false)
	m.Series = nborm.NewStringField(m, "series", false, false, false)
	m.Isbn = nborm.NewStringField(m, "isbn", true, false, false)
	m.PublishDate = nborm.NewDateField(m, "publish_date", false, false, false)
	m.Binding = nborm.NewStringField(m, "binding", false, false, false)
	m.Format = nborm.NewStringField(m, "format", false, false, false)
	m.Pages = nborm.NewIntField(m, "pages", false, false, false)
	m.WordCount = nborm.NewIntField(m, "word_count", false, false, false)
	m.ContentIntro = nborm.NewStringField(m, "content_intro", false, false, false)
	m.AuthorIntro = nborm.NewStringField(m, "author_intro", false, false, false)
	m.Menu = nborm.NewStringField(m, "menu", false, false, false)
	m.Image = nborm.NewStringField(m, "image", false, false, false)
	m.IsOk = nborm.NewBoolField(m, "is_ok", false, false, false)
	m.Book = nborm.NewReverseForeignKey(m.Isbn, func() nborm.Field { return NewBook().Isbn })
	m.Tag = nborm.NewManyToMany(m.Isbn, func() (nborm.Field, nborm.Field) {
		m := NewBookInfo__Tag()
		return m.BookInfo__Isbn, m.Tag__Id
	},
		func() nborm.Field { return NewTag().Id })

	return m
}

func (m *BookInfo) DB() string {
	return "bk_dalian"
}

func (m *BookInfo) Tab() string {
	return "book_info"
}

func (m *BookInfo) Fields() []nborm.Field {
	return []nborm.Field{
		m.Id,
		m.Title,
		m.Price,
		m.Author,
		m.Publisher,
		m.Series,
		m.Isbn,
		m.PublishDate,
		m.Binding,
		m.Format,
		m.Pages,
		m.WordCount,
		m.ContentIntro,
		m.AuthorIntro,
		m.Menu,
		m.Image,
		m.IsOk,
	}
}

func (m *BookInfo) SetSync(b bool) {
	m._isSync = b
}

func (m *BookInfo) GetSync() bool {
	return m._isSync
}

type BookInfoList struct {
	*BookInfo
	List []*BookInfo
}

func NewBookInfoList() *BookInfoList {
	return &BookInfoList{NewBookInfo(), make([]*BookInfo, 0, 128)}
}

func (l *BookInfoList) New() nborm.Model {
	m := NewBookInfo()
	l.List = append(l.List, m)
	return m
}

func (l *BookInfoList) Len() int {
	return len(l.List)
}

func (l *BookInfoList) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func (l *BookInfoList) Index(i int) nborm.Model {
	return l.List[i]
}

func (l *BookInfoList) Delete(i int) {
	switch i {
	case 0:
		l.List = l.List[1:]
	case l.Len() - 1:
		l.List = l.List[:l.Len()-1]
	default:
		l.List = append(l.List[:i], l.List[i+1:]...)
	}
}

func (l *BookInfoList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

func (l *BookInfoList) Model() nborm.Model {
	return l.BookInfo
}

type BookInfo__Tag struct {
	Id             *nborm.IntField    `json:"id"`
	BookInfo__Isbn *nborm.StringField `json:"book_info__isbn"`
	Tag__Id        *nborm.IntField    `json:"tag__id"`
	_isSync        bool
}

func NewBookInfo__Tag() *BookInfo__Tag {
	m := &BookInfo__Tag{}
	m.Id = nborm.NewIntField(m, "id", true, true, false)
	m.BookInfo__Isbn = nborm.NewStringField(m, "book_info__isbn", false, false, false)
	m.Tag__Id = nborm.NewIntField(m, "tag__id", false, false, false)

	return m
}

func (m *BookInfo__Tag) DB() string {
	return "bk_dalian"
}

func (m *BookInfo__Tag) Tab() string {
	return "book_info__tag"
}

func (m *BookInfo__Tag) Fields() []nborm.Field {
	return []nborm.Field{
		m.Id,
		m.BookInfo__Isbn,
		m.Tag__Id,
	}
}

func (m *BookInfo__Tag) SetSync(b bool) {
	m._isSync = b
}

func (m *BookInfo__Tag) GetSync() bool {
	return m._isSync
}

type BookInfo__TagList struct {
	*BookInfo__Tag
	List []*BookInfo__Tag
}

func NewBookInfo__TagList() *BookInfo__TagList {
	return &BookInfo__TagList{NewBookInfo__Tag(), make([]*BookInfo__Tag, 0, 128)}
}

func (l *BookInfo__TagList) New() nborm.Model {
	m := NewBookInfo__Tag()
	l.List = append(l.List, m)
	return m
}

func (l *BookInfo__TagList) Len() int {
	return len(l.List)
}

func (l *BookInfo__TagList) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func (l *BookInfo__TagList) Index(i int) nborm.Model {
	return l.List[i]
}

func (l *BookInfo__TagList) Delete(i int) {
	switch i {
	case 0:
		l.List = l.List[1:]
	case l.Len() - 1:
		l.List = l.List[:l.Len()-1]
	default:
		l.List = append(l.List[:i], l.List[i+1:]...)
	}
}

func (l *BookInfo__TagList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

func (l *BookInfo__TagList) Model() nborm.Model {
	return l.BookInfo__Tag
}

type Tag struct {
	Id       *nborm.IntField    `json:"id"`
	Tag      *nborm.StringField `json:"tag"`
	BookInfo *nborm.ManyToMany  `json:"book_info"`
	_isSync  bool
}

func NewTag() *Tag {
	m := &Tag{}
	m.Id = nborm.NewIntField(m, "id", false, true, true)
	m.Tag = nborm.NewStringField(m, "tag", true, false, false)
	m.BookInfo = nborm.NewManyToMany(m.Id, func() (nborm.Field, nborm.Field) {
		m := NewBookInfo__Tag()
		return m.Tag__Id, m.BookInfo__Isbn
	},
		func() nborm.Field { return NewBookInfo().Isbn })

	return m
}

func (m *Tag) DB() string {
	return "bk_dalian"
}

func (m *Tag) Tab() string {
	return "tag"
}

func (m *Tag) Fields() []nborm.Field {
	return []nborm.Field{
		m.Id,
		m.Tag,
	}
}

func (m *Tag) SetSync(b bool) {
	m._isSync = b
}

func (m *Tag) GetSync() bool {
	return m._isSync
}

type TagList struct {
	*Tag
	List []*Tag
}

func NewTagList() *TagList {
	return &TagList{NewTag(), make([]*Tag, 0, 128)}
}

func (l *TagList) New() nborm.Model {
	m := NewTag()
	l.List = append(l.List, m)
	return m
}

func (l *TagList) Len() int {
	return len(l.List)
}

func (l *TagList) Swap(i, j int) {
	l.List[i], l.List[j] = l.List[j], l.List[i]
}

func (l *TagList) Index(i int) nborm.Model {
	return l.List[i]
}

func (l *TagList) Delete(i int) {
	switch i {
	case 0:
		l.List = l.List[1:]
	case l.Len() - 1:
		l.List = l.List[:l.Len()-1]
	default:
		l.List = append(l.List[:i], l.List[i+1:]...)
	}
}

func (l *TagList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

func (l *TagList) Model() nborm.Model {
	return l.Tag
}

func init() {
	nborm.RegisterDB("wangjun", "Wt20110523", "127.0.0.1:12345", "bk_dalian")
}
