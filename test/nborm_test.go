package test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/wangjun861205/nborm"
)

var tests = []struct {
	name string
	f    func() error
}{
	{"insert one", func() error {
		auth := NewAuth()
		auth.Username.Set("wangjun")
		auth.Password.Set("Wt20110523")
		auth.Phone.Set("13793148690")
		auth.Email.Set("wangjun20110523@gmail.com")
		return nborm.InsertOne(auth)
	}},
	{"insert many", func() error {
		bookInfos := MakeBookInfoList()
		for i := 0; i < 100; i++ {
			bi := NewBookInfo()
			bi.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfos = append(bookInfos, bi)
		}
		return nborm.InsertMul(&bookInfos)
	}},
	{"get first", func() error {
		auth := NewAuth()
		err := nborm.First(auth)
		if err != nil {
			return err
		}
		if !auth._isSync {
			return errors.New("_isSync is false, want true")
		}
		username, valid, null := auth.Username.Get()
		if !valid || null {
			return fmt.Errorf("username is %s, want 'wangjun'", username)
		}
		return nil
	}},
	{"get first no rows", func() error {
		book := NewBook()
		err := nborm.First(book)
		if err != nil {
			return err
		}
		if book._isSync {
			return errors.New("auth._isSync is true, want false")
		}
		return nil
	}},
	{"get one", func() error {
		bookInfo := NewBookInfo()
		bookInfo.Isbn.Set(fmt.Sprintf("%011d", 50))
		return nborm.GetOne(bookInfo)
	}},
	{"get multiple rows", func() error {
		bookInfos := MakeBookInfoList()
		for i := 20; i < 80; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfos = append(bookInfos, bookInfo)
		}
		err := nborm.GetMul(&bookInfos)
		if err != nil {
			return err
		}
		for _, bookInfo := range bookInfos[1:] {
			if !bookInfo._isSync {
				return fmt.Errorf("SyncFlag is false, want true")
			}
		}
		return nil
	}},
	{"join query one", func() error {
		book := NewBook()
		book.Isbn.Set(fmt.Sprintf("%011d", 20))
		book.UniqueCode.Set("dl000000001")
		if err := nborm.InsertOne(book); err != nil {
			return err
		}
		bookInfo := NewBookInfo()
		where := bookInfo.Isbn.Eq(fmt.Sprintf("%011d", 20))
		if err := nborm.JoinQueryOne(book, where, book.BookInfo); err != nil {
			return err
		}
		if !book._isSync {
			return errors.New("SyncFlag is false, want true")
		}
		isbn, _, _ := book.Isbn.Get()
		if isbn != fmt.Sprintf("%011d", 20) {
			return fmt.Errorf("isbn is %s, want %s", isbn, fmt.Sprintf("%011d", 20))
		}
		uniqueCode, _, _ := book.UniqueCode.Get()
		if uniqueCode != "dl000000001" {
			return fmt.Errorf("unique code is %s, want 'dl000000001'", uniqueCode)
		}

		return nil
	}},
	{"join query multiple", func() error {
		books := MakeBookList()
		for i := 0; i < 50; i++ {
			book := NewBook()
			book.Isbn.Set(fmt.Sprintf("%011d", i))
			book.UniqueCode.Set(fmt.Sprintf("dl%011d", i))
			books = append(books, book)
		}
		if err := nborm.InsertMul(&books); err != nil {
			return err
		}
		bookInfos := MakeBookInfoList()
		if err := nborm.JoinQuery(&bookInfos, nil, nil, nil, true, bookInfos[0].Book); err != nil {
			return err
		}
		if nborm.NumRes(&bookInfos) != 50 {
			return fmt.Errorf("number of results is %d, want 50", nborm.NumRes(&bookInfos))
		}
		return nil
	}},
	{"join query multiple and found rows", func() error {
		bookInfos := MakeBookInfoList()
		num, err := nborm.JoinQueryWithFoundRows(&bookInfos, nil, nil, nil, true, bookInfos[0].Book)
		if err != nil {
			return err
		}
		if num != 50 {
			return fmt.Errorf("number of results is %d, want 50", nborm.NumRes(&bookInfos))
		}
		return nil
	}},
	{"all", func() error {
		bookInfos := MakeBookInfoList()
		if err := nborm.All(&bookInfos, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&bookInfos) != 100 {
			return fmt.Errorf("number of results is %d, want 100", nborm.NumRes(&bookInfos))
		}
		return nil
	}},
	{"all and found rows", func() error {
		bookInfos := MakeBookInfoList()
		num, err := nborm.AllWithFoundRows(&bookInfos, nil, nil)
		if err != nil {
			return err
		}
		if num != 100 {
			return fmt.Errorf("number of results is %d, want 100", num)
		}
		return nil
	}},
	{"query one", func() error {
		bookInfo := NewBookInfo()
		if err := nborm.QueryOne(bookInfo, bookInfo.Isbn.Eq(fmt.Sprintf("%011d", 50))); err != nil {
			return err
		}
		if isbn, _, _ := bookInfo.Isbn.Get(); isbn != fmt.Sprintf("%011d", 50) {
			return fmt.Errorf("isbn is %s, want %s", isbn, fmt.Sprintf("%011d", 50))
		}
		return nil
	}},
	{"query", func() error {
		bookInfos := MakeBookInfoList()
		if err := nborm.Query(&bookInfos, bookInfos[0].Isbn.Eq(fmt.Sprintf("%011d", 50)), nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&bookInfos) != 1 {
			return fmt.Errorf("the number of results is %d, want 1", nborm.NumRes(&bookInfos))
		}
		if isbn, _, _ := bookInfos[1].Isbn.Get(); isbn != fmt.Sprintf("%011d", 50) {
			return fmt.Errorf("isbn is %s, want %s", isbn, fmt.Sprintf("%011d", 50))
		}
		return nil
	}},
	{"query with found rows", func() error {
		bookInfos := MakeBookInfoList()
		num, err := nborm.QueryWithFoundRows(&bookInfos, bookInfos[0].Isbn.Eq(fmt.Sprintf("%011d", 50)), nil, nil)
		if err != nil {
			return err
		}
		if num != 1 {
			return fmt.Errorf("the number of results is %d, want 1", num)
		}
		if isbn, _, _ := bookInfos[1].Isbn.Get(); isbn != fmt.Sprintf("%011d", 50) {
			return fmt.Errorf("isbn is %s, want %s", isbn, fmt.Sprintf("%011d", 50))
		}
		return nil
	}},
	{"insert or update one", func() error {
		bookInfo := NewBookInfo()
		bookInfo.Isbn.Set(fmt.Sprintf("%011d", 50))
		bookInfo.Author.Set("王君")
		if err := nborm.InsertOrUpdateOne(bookInfo); err != nil {
			return err
		}
		if !bookInfo._isSync {
			return errors.New("SyncFlag is false, want true")
		}
		return nil
	}},
	{"insert or update multiple", func() error {
		bookInfos := MakeBookInfoList()
		for i := 0; i < 10; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfo.Author.Set("王君")
			bookInfos = append(bookInfos, bookInfo)
		}
		if err := nborm.InsertOrUpdateMul(&bookInfos); err != nil {
			return err
		}
		for _, bookInfo := range bookInfos[1:] {
			if !bookInfo._isSync {
				return errors.New("SyncFlag is false, want true")
			}
		}
		return nil
	}},
	{"insert or get one", func() error {
		bookInfo := NewBookInfo()
		bookInfo.Isbn.Set(fmt.Sprintf("%011d", 50))
		if err := nborm.InsertOrGetOne(bookInfo); err != nil {
			return err
		}
		if !bookInfo._isSync {
			return errors.New("SyncFlag is false, want true")
		}
		if author, _, _ := bookInfo.Author.Get(); author != "王君" {
			return fmt.Errorf("author is %s, want 王君", author)
		}
		return nil
	}},
	{"insert or get multiple", func() error {
		bookInfos := MakeBookInfoList()
		for i := 0; i < 10; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfos = append(bookInfos, bookInfo)
		}
		if err := nborm.InsertOrGetMul(&bookInfos); err != nil {
			return err
		}
		for _, bookInfo := range bookInfos[1:] {
			if !bookInfo._isSync {
				return errors.New("SyncFlag is false, want true")
			}
			if author, _, _ := bookInfo.Author.Get(); author != "王君" {
				return fmt.Errorf("auhtor is %s, want 王君", author)
			}
		}
		return nil
	}},
	{"update one", func() error {
		bookInfo := NewBookInfo()
		bookInfo.Isbn.Set(fmt.Sprintf("%011d", 80))
		bookInfo.Binding.Set("平装")
		if err := nborm.UpdateOne(bookInfo); err != nil {
			return err
		}
		if !bookInfo._isSync {
			return errors.New("SyncFlag is false, want true")
		}
		return nil
	}},
	{"update multiple", func() error {
		bookInfos := MakeBookInfoList()
		for i := 10; i < 20; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfo.Format.Set("16开")
			bookInfos = append(bookInfos, bookInfo)
		}
		if err := nborm.UpdateMul(&bookInfos); err != nil {
			return err
		}
		for _, bookInfo := range bookInfos[1:] {
			if !bookInfo._isSync {
				return errors.New("SyncFlag is false, want true")
			}
		}
		return nil
	}},
	{"bulk update", func() error {
		bookInfo := NewBookInfo()
		return nborm.BulkUpdate(bookInfo, bookInfo.Format.Eq("16开"), bookInfo.Format.UpdateValue("32开", false))
	}},
	{"delete one", func() error {
		bookInfo := NewBookInfo()
		bookInfo.Isbn.Set(fmt.Sprintf("%011d", 0))
		if err := nborm.DeleteOne(bookInfo); err != nil {
			return err
		}
		if err := nborm.GetOne(bookInfo); err != sql.ErrNoRows {
			return fmt.Errorf("get %s, want sql.ErrNoRows error", err.Error())
		}
		return nil
	}},
	{"delete multiple", func() error {
		checkMap := make(map[string]bool)
		bookInfos := MakeBookInfoList()
		for i := 0; i < 30; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfos = append(bookInfos, bookInfo)
			checkMap[fmt.Sprintf("%011d", i)] = true
		}
		if err := nborm.DeleteMul(&bookInfos); err != nil {
			return err
		}
		nborm.ClsRes(&bookInfos)
		if err := nborm.All(&bookInfos, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&bookInfos) != 70 {
			return fmt.Errorf("the result number is %d, want 20", nborm.NumRes(&bookInfos))
		}
		for _, bookInfo := range bookInfos {
			isbn, _, _ := bookInfo.Isbn.Get()
			if checkMap[isbn] {
				return fmt.Errorf("isbn (%s) still exist", isbn)
			}
		}
		return nil
	}},
	{"bulk delete", func() error {
		bookInfo := NewBookInfo()
		if err := nborm.BulkDelete(bookInfo, bookInfo.Author.Eq("王君")); err != nil {
			return err
		}
		bookInfos := MakeBookInfoList()
		for _, bookInfo := range bookInfos {
			if author, _, _ := bookInfo.Author.Get(); author == "王君" {
				return errors.New("author ('王君') still exists")
			}
		}
		return nil
	}},
	{"delete all", func() error {
		books := MakeBookList()
		if err := nborm.All(&books, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&books) == 0 {
			return errors.New("the number of result set is 0, want non-zero value")
		}
		nborm.ClsRes(&books)
		book := NewBook()
		if err := nborm.DeleteAll(book); err != nil {
			return err
		}
		if err := nborm.All(&books, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&books) > 0 {
			return fmt.Errorf("the number of result set is %d, want 0", nborm.NumRes(&books))
		}
		return nil
	}},
	{"truncate table", func() error {
		auths := MakeAuthList()
		if err := nborm.All(&auths, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&auths) == 0 {
			return errors.New("the number of result set is 0, want non-zero value")
		}
		nborm.ClsRes(&auths)
		if err := nborm.TruncateTable(&auths); err != nil {
			return err
		}
		if err := nborm.All(&auths, nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&auths) > 0 {
			return fmt.Errorf("the number of result set is %d, want 0", nborm.NumRes(&auths))
		}
		return nil
	}},
	{"count", func() error {
		bookInfos := MakeBookInfoList()
		for i := 100; i < 150; i++ {
			bookInfo := NewBookInfo()
			bookInfo.Isbn.Set(fmt.Sprintf("%011d", i))
			bookInfo.Author.Set("不熊")
			bookInfos = append(bookInfos, bookInfo)
		}
		if err := nborm.InsertMul(&bookInfos); err != nil {
			return err
		}
		if num, err := nborm.Count(&bookInfos, bookInfos[0].Author.Eq("不熊")); err != nil {
			return err
		} else if num != 50 {
			return fmt.Errorf("the number of result set is %d, want 50", num)
		}
		return nil
	}},
	{"sort", func() error {
		checkList := make([]string, 0, 50)
		for i := 149; i >= 100; i-- {
			checkList = append(checkList, fmt.Sprintf("%011d", i))
		}
		bookInfos := MakeBookInfoList()
		if err := nborm.Query(&bookInfos, bookInfos[0].Author.Eq("不熊"), nil, nil); err != nil {
			return err
		}
		nborm.Sort(&bookInfos, bookInfos[0].Isbn.LessFunc(true))
		for i, bookInfo := range bookInfos[1:] {
			isbn, _, _ := bookInfo.Isbn.Get()
			if isbn != checkList[i] {
				return fmt.Errorf("isbn is %s, want %s", isbn, checkList[i])
			}
		}
		return nil
	}},
	{"distinct", func() error {
		bookInfos := MakeBookInfoList()
		if err := nborm.Query(&bookInfos, bookInfos[0].Author.Eq("不熊"), nil, nil); err != nil {
			return err
		}
		if nborm.NumRes(&bookInfos) < 2 {
			return errors.New("the number of result set is less than one, want larger than one")
		}
		nborm.Distinct(&bookInfos, bookInfos[0].Author)
		if nborm.NumRes(&bookInfos) != 1 {
			return fmt.Errorf("the number of result set is %d, want 1", nborm.NumRes(&bookInfos))
		}
		return nil
	}},
}

func TestInsert(t *testing.T) {
	defer func() {
		if err := nborm.DeleteAll(NewAuth()); err != nil {
			t.Error(err)
		}
		if err := nborm.DeleteAll(NewBook()); err != nil {
			t.Error(err)
		}
		if err := nborm.DeleteAll(NewBookInfo()); err != nil {
			t.Error(err)
		}
		if err := nborm.DeleteAll(NewTag()); err != nil {
			t.Error(err)
		}
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f(); err != nil {
				t.Error(err)
			}
		})
	}
}
