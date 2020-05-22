package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wangjun861205/nborm"
	"reflect"
	"strings"
	"time"
)

func NewUser() *User {
	m := &User{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 1)
	m.UserCode.Init(m, "UserCode", "UserCode", "UserCode", "UserCode", 2)
	m.Name.Init(m, "Name", "Name", "Name", "Name", 3)
	m.Sex.Init(m, "Sex", "Sex", "Sex", "Sex", 4)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", "IdentityType", "IdentityType", 5)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", "IdentityNum", "IdentityNum", 6)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", "ExpirationDate", "ExpirationDate", 7)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", "UniversityCode", "UniversityCode", 8)
	m.UserType.Init(m, "UserType", "UserType", "UserType", "UserType", 9)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", 10)
	m.Type.Init(m, "Type", "Type", "Type", "Type", 11)
	m.Password.Init(m, "Password", "Password", "Password", "Password", 12)
	m.Phone.Init(m, "Phone", "Phone", "Phone", "Phone", 13)
	m.Email.Init(m, "Email", "Email", "Email", "Email", 14)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", "PictureURL", "PictureURL", 15)
	m.Question.Init(m, "Question", "Question", "Question", "Question", 16)
	m.Answer.Init(m, "Answer", "Answer", "Answer", "Answer", 17)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", "AvailableLogin", "AvailableLogin", 18)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 19)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 20)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 21)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 22)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 23)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 24)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 25)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 26)
	m.Nonego.Init(m, "Nonego", "Nonego", "Nonego", "Nonego", 27)
	m.InitRel()
	return m
}

func newSubUser(parent nborm.Model) *User {
	m := &User{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 1)
	m.UserCode.Init(m, "UserCode", "UserCode", "UserCode", "UserCode", 2)
	m.Name.Init(m, "Name", "Name", "Name", "Name", 3)
	m.Sex.Init(m, "Sex", "Sex", "Sex", "Sex", 4)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", "IdentityType", "IdentityType", 5)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", "IdentityNum", "IdentityNum", 6)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", "ExpirationDate", "ExpirationDate", 7)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", "UniversityCode", "UniversityCode", 8)
	m.UserType.Init(m, "UserType", "UserType", "UserType", "UserType", 9)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", 10)
	m.Type.Init(m, "Type", "Type", "Type", "Type", 11)
	m.Password.Init(m, "Password", "Password", "Password", "Password", 12)
	m.Phone.Init(m, "Phone", "Phone", "Phone", "Phone", 13)
	m.Email.Init(m, "Email", "Email", "Email", "Email", 14)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", "PictureURL", "PictureURL", 15)
	m.Question.Init(m, "Question", "Question", "Question", "Question", 16)
	m.Answer.Init(m, "Answer", "Answer", "Answer", "Answer", 17)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", "AvailableLogin", "AvailableLogin", 18)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 19)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 20)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 21)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 22)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 23)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 24)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 25)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 26)
	m.Nonego.Init(m, "Nonego", "Nonego", "Nonego", "Nonego", 27)
	return m
}

func (m *User) InitRel() {
	m.BasicInfo = newSubStudentbasicinfo(m)
	var relInfo0 *nborm.RelationInfo
	relInfo0 = relInfo0.Append("BasicInfo", m.BasicInfo, nborm.NewExpr("@=@", &m.IntelUserCode, &m.BasicInfo.IntelUserCode))
	m.AppendRelation(relInfo0)
	m.StudentClass = newSubClass(m)
	var relInfo1 *nborm.RelationInfo
	mm0 := newSubStudentbasicinfo(m)
	relInfo1 = relInfo1.Append("StudentClass", mm0, nborm.NewExpr("@=@", &m.IntelUserCode, &mm0.IntelUserCode))
	relInfo1 = relInfo1.Append("StudentClass", m.StudentClass, nborm.NewExpr("@=@", &mm0.Class, &m.StudentClass.ClassCode))
	m.AppendRelation(relInfo1)
	m.AddRelInited()
}

func (m *User) DB() string {
	return "*"
}

func (m *User) Tab() string {
	return "user"
}

func (m *User) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"Id", "Id", &m.Id, 0},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode, 1},
		{"UserCode", "UserCode", &m.UserCode, 2},
		{"Name", "Name", &m.Name, 3},
		{"Sex", "Sex", &m.Sex, 4},
		{"IdentityType", "IdentityType", &m.IdentityType, 5},
		{"IdentityNum", "IdentityNum", &m.IdentityNum, 6},
		{"ExpirationDate", "ExpirationDate", &m.ExpirationDate, 7},
		{"UniversityCode", "UniversityCode", &m.UniversityCode, 8},
		{"UserType", "UserType", &m.UserType, 9},
		{"EnrollmentStatus", "EnrollmentStatus", &m.EnrollmentStatus, 10},
		{"Type", "Type", &m.Type, 11},
		{"Password", "Password", &m.Password, 12},
		{"Phone", "Phone", &m.Phone, 13},
		{"Email", "Email", &m.Email, 14},
		{"PictureURL", "PictureURL", &m.PictureURL, 15},
		{"Question", "Question", &m.Question, 16},
		{"Answer", "Answer", &m.Answer, 17},
		{"AvailableLogin", "AvailableLogin", &m.AvailableLogin, 18},
		{"Operator", "Operator", &m.Operator, 19},
		{"InsertDatetime", "InsertDatetime", &m.InsertDatetime, 20},
		{"UpdateDatetime", "UpdateDatetime", &m.UpdateDatetime, 21},
		{"Status", "Status", &m.Status, 22},
		{"Remark1", "Remark1", &m.Remark1, 23},
		{"Remark2", "Remark2", &m.Remark2, 24},
		{"Remark3", "Remark3", &m.Remark3, 25},
		{"Remark4", "Remark4", &m.Remark4, 26},
		{"Nonego", "Nonego", &m.Nonego, 27},
	}
}

func (m *User) AutoIncField() nborm.Field {
	return &m.Id
}

func (m *User) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.Id,
	}
}

func (m *User) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.IntelUserCode,
		},
		{
			&m.IntelUserCode,
			&m.Status,
		},
	}
}
func (m User) MarshalJSON() ([]byte, error) {
	if !m.IsSynced() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	buffer.WriteString("{\n\"Aggs\": ")
	metaB, err := json.MarshalIndent(m.Meta, "", "\t")
	if err != nil {
		return nil, err
	}
	buffer.Write(metaB)
	if m.Id.IsValid() {
		buffer.WriteString(",\n\"Id\": ")
		IdB, err := json.MarshalIndent(m.Id, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdB)
	}
	if m.IntelUserCode.IsValid() {
		buffer.WriteString(",\n\"IntelUserCode\": ")
		IntelUserCodeB, err := json.MarshalIndent(m.IntelUserCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IntelUserCodeB)
	}
	if m.UserCode.IsValid() {
		buffer.WriteString(",\n\"UserCode\": ")
		UserCodeB, err := json.MarshalIndent(m.UserCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UserCodeB)
	}
	if m.Name.IsValid() {
		buffer.WriteString(",\n\"Name\": ")
		NameB, err := json.MarshalIndent(m.Name, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NameB)
	}
	if m.Sex.IsValid() {
		buffer.WriteString(",\n\"Sex\": ")
		SexB, err := json.MarshalIndent(m.Sex, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(SexB)
	}
	if m.IdentityType.IsValid() {
		buffer.WriteString(",\n\"IdentityType\": ")
		IdentityTypeB, err := json.MarshalIndent(m.IdentityType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdentityTypeB)
	}
	if m.IdentityNum.IsValid() {
		buffer.WriteString(",\n\"IdentityNum\": ")
		IdentityNumB, err := json.MarshalIndent(m.IdentityNum, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdentityNumB)
	}
	if m.ExpirationDate.IsValid() {
		buffer.WriteString(",\n\"ExpirationDate\": ")
		ExpirationDateB, err := json.MarshalIndent(m.ExpirationDate, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ExpirationDateB)
	}
	if m.UniversityCode.IsValid() {
		buffer.WriteString(",\n\"UniversityCode\": ")
		UniversityCodeB, err := json.MarshalIndent(m.UniversityCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UniversityCodeB)
	}
	if m.UserType.IsValid() {
		buffer.WriteString(",\n\"UserType\": ")
		UserTypeB, err := json.MarshalIndent(m.UserType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UserTypeB)
	}
	if m.EnrollmentStatus.IsValid() {
		buffer.WriteString(",\n\"EnrollmentStatus\": ")
		EnrollmentStatusB, err := json.MarshalIndent(m.EnrollmentStatus, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EnrollmentStatusB)
	}
	if m.Type.IsValid() {
		buffer.WriteString(",\n\"Type\": ")
		TypeB, err := json.MarshalIndent(m.Type, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(TypeB)
	}
	if m.Password.IsValid() {
		buffer.WriteString(",\n\"Password\": ")
		PasswordB, err := json.MarshalIndent(m.Password, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PasswordB)
	}
	if m.Phone.IsValid() {
		buffer.WriteString(",\n\"Phone\": ")
		PhoneB, err := json.MarshalIndent(m.Phone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PhoneB)
	}
	if m.Email.IsValid() {
		buffer.WriteString(",\n\"Email\": ")
		EmailB, err := json.MarshalIndent(m.Email, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EmailB)
	}
	if m.PictureURL.IsValid() {
		buffer.WriteString(",\n\"PictureURL\": ")
		PictureURLB, err := json.MarshalIndent(m.PictureURL, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PictureURLB)
	}
	if m.Question.IsValid() {
		buffer.WriteString(",\n\"Question\": ")
		QuestionB, err := json.MarshalIndent(m.Question, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(QuestionB)
	}
	if m.Answer.IsValid() {
		buffer.WriteString(",\n\"Answer\": ")
		AnswerB, err := json.MarshalIndent(m.Answer, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AnswerB)
	}
	if m.AvailableLogin.IsValid() {
		buffer.WriteString(",\n\"AvailableLogin\": ")
		AvailableLoginB, err := json.MarshalIndent(m.AvailableLogin, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AvailableLoginB)
	}
	if m.Operator.IsValid() {
		buffer.WriteString(",\n\"Operator\": ")
		OperatorB, err := json.MarshalIndent(m.Operator, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OperatorB)
	}
	if m.InsertDatetime.IsValid() {
		buffer.WriteString(",\n\"InsertDatetime\": ")
		InsertDatetimeB, err := json.MarshalIndent(m.InsertDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(InsertDatetimeB)
	}
	if m.UpdateDatetime.IsValid() {
		buffer.WriteString(",\n\"UpdateDatetime\": ")
		UpdateDatetimeB, err := json.MarshalIndent(m.UpdateDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateDatetimeB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.Remark1.IsValid() {
		buffer.WriteString(",\n\"Remark1\": ")
		Remark1B, err := json.MarshalIndent(m.Remark1, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark1B)
	}
	if m.Remark2.IsValid() {
		buffer.WriteString(",\n\"Remark2\": ")
		Remark2B, err := json.MarshalIndent(m.Remark2, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark2B)
	}
	if m.Remark3.IsValid() {
		buffer.WriteString(",\n\"Remark3\": ")
		Remark3B, err := json.MarshalIndent(m.Remark3, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark3B)
	}
	if m.Remark4.IsValid() {
		buffer.WriteString(",\n\"Remark4\": ")
		Remark4B, err := json.MarshalIndent(m.Remark4, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark4B)
	}
	if m.Nonego.IsValid() {
		buffer.WriteString(",\n\"Nonego\": ")
		NonegoB, err := json.MarshalIndent(m.Nonego, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NonegoB)
	}
	if m.BasicInfo != nil && m.BasicInfo.IsSynced() {
		buffer.WriteString(",\n\"BasicInfo\": ")
		BasicInfoB, err := json.MarshalIndent(m.BasicInfo, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BasicInfoB)
	}
	if m.StudentClass != nil && m.StudentClass.IsSynced() {
		buffer.WriteString(",\n\"StudentClass\": ")
		StudentClassB, err := json.MarshalIndent(m.StudentClass, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentClassB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type UserList struct {
	User   `json:"-"`
	dupMap map[string]int
	List   []*User
	Total  int
}

func (m *User) Collapse() {
	if m.BasicInfo != nil && m.BasicInfo.IsSynced() {
		m.BasicInfo.Collapse()
	}
	if m.StudentClass != nil && m.StudentClass.IsSynced() {
		m.StudentClass.Collapse()
	}
}

func NewUserList() *UserList {
	l := &UserList{
		User{},
		make(map[string]int),
		make([]*User, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 1)
	l.UserCode.Init(l, "UserCode", "UserCode", "UserCode", "UserCode", 2)
	l.Name.Init(l, "Name", "Name", "Name", "Name", 3)
	l.Sex.Init(l, "Sex", "Sex", "Sex", "Sex", 4)
	l.IdentityType.Init(l, "IdentityType", "IdentityType", "IdentityType", "IdentityType", 5)
	l.IdentityNum.Init(l, "IdentityNum", "IdentityNum", "IdentityNum", "IdentityNum", 6)
	l.ExpirationDate.Init(l, "ExpirationDate", "ExpirationDate", "ExpirationDate", "ExpirationDate", 7)
	l.UniversityCode.Init(l, "UniversityCode", "UniversityCode", "UniversityCode", "UniversityCode", 8)
	l.UserType.Init(l, "UserType", "UserType", "UserType", "UserType", 9)
	l.EnrollmentStatus.Init(l, "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", 10)
	l.Type.Init(l, "Type", "Type", "Type", "Type", 11)
	l.Password.Init(l, "Password", "Password", "Password", "Password", 12)
	l.Phone.Init(l, "Phone", "Phone", "Phone", "Phone", 13)
	l.Email.Init(l, "Email", "Email", "Email", "Email", 14)
	l.PictureURL.Init(l, "PictureURL", "PictureURL", "PictureURL", "PictureURL", 15)
	l.Question.Init(l, "Question", "Question", "Question", "Question", 16)
	l.Answer.Init(l, "Answer", "Answer", "Answer", "Answer", 17)
	l.AvailableLogin.Init(l, "AvailableLogin", "AvailableLogin", "AvailableLogin", "AvailableLogin", 18)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 19)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 20)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 21)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 22)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 23)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 24)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 25)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 26)
	l.Nonego.Init(l, "Nonego", "Nonego", "Nonego", "Nonego", 27)
	l.InitRel()
	return l
}

func newSubUserList(parent nborm.Model) *UserList {
	l := &UserList{
		User{},
		make(map[string]int),
		make([]*User, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 1)
	l.UserCode.Init(l, "UserCode", "UserCode", "UserCode", "UserCode", 2)
	l.Name.Init(l, "Name", "Name", "Name", "Name", 3)
	l.Sex.Init(l, "Sex", "Sex", "Sex", "Sex", 4)
	l.IdentityType.Init(l, "IdentityType", "IdentityType", "IdentityType", "IdentityType", 5)
	l.IdentityNum.Init(l, "IdentityNum", "IdentityNum", "IdentityNum", "IdentityNum", 6)
	l.ExpirationDate.Init(l, "ExpirationDate", "ExpirationDate", "ExpirationDate", "ExpirationDate", 7)
	l.UniversityCode.Init(l, "UniversityCode", "UniversityCode", "UniversityCode", "UniversityCode", 8)
	l.UserType.Init(l, "UserType", "UserType", "UserType", "UserType", 9)
	l.EnrollmentStatus.Init(l, "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", 10)
	l.Type.Init(l, "Type", "Type", "Type", "Type", 11)
	l.Password.Init(l, "Password", "Password", "Password", "Password", 12)
	l.Phone.Init(l, "Phone", "Phone", "Phone", "Phone", 13)
	l.Email.Init(l, "Email", "Email", "Email", "Email", 14)
	l.PictureURL.Init(l, "PictureURL", "PictureURL", "PictureURL", "PictureURL", 15)
	l.Question.Init(l, "Question", "Question", "Question", "Question", 16)
	l.Answer.Init(l, "Answer", "Answer", "Answer", "Answer", 17)
	l.AvailableLogin.Init(l, "AvailableLogin", "AvailableLogin", "AvailableLogin", "AvailableLogin", 18)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 19)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 20)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 21)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 22)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 23)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 24)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 25)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 26)
	l.Nonego.Init(l, "Nonego", "Nonego", "Nonego", "Nonego", 27)
	return l
}

func (l *UserList) NewModel() nborm.Model {
	m := &User{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 1)
	l.IntelUserCode.CopyStatus(&m.IntelUserCode)
	m.UserCode.Init(m, "UserCode", "UserCode", "UserCode", "UserCode", 2)
	l.UserCode.CopyStatus(&m.UserCode)
	m.Name.Init(m, "Name", "Name", "Name", "Name", 3)
	l.Name.CopyStatus(&m.Name)
	m.Sex.Init(m, "Sex", "Sex", "Sex", "Sex", 4)
	l.Sex.CopyStatus(&m.Sex)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", "IdentityType", "IdentityType", 5)
	l.IdentityType.CopyStatus(&m.IdentityType)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", "IdentityNum", "IdentityNum", 6)
	l.IdentityNum.CopyStatus(&m.IdentityNum)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", "ExpirationDate", "ExpirationDate", 7)
	l.ExpirationDate.CopyStatus(&m.ExpirationDate)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", "UniversityCode", "UniversityCode", 8)
	l.UniversityCode.CopyStatus(&m.UniversityCode)
	m.UserType.Init(m, "UserType", "UserType", "UserType", "UserType", 9)
	l.UserType.CopyStatus(&m.UserType)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", "EnrollmentStatus", 10)
	l.EnrollmentStatus.CopyStatus(&m.EnrollmentStatus)
	m.Type.Init(m, "Type", "Type", "Type", "Type", 11)
	l.Type.CopyStatus(&m.Type)
	m.Password.Init(m, "Password", "Password", "Password", "Password", 12)
	l.Password.CopyStatus(&m.Password)
	m.Phone.Init(m, "Phone", "Phone", "Phone", "Phone", 13)
	l.Phone.CopyStatus(&m.Phone)
	m.Email.Init(m, "Email", "Email", "Email", "Email", 14)
	l.Email.CopyStatus(&m.Email)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", "PictureURL", "PictureURL", 15)
	l.PictureURL.CopyStatus(&m.PictureURL)
	m.Question.Init(m, "Question", "Question", "Question", "Question", 16)
	l.Question.CopyStatus(&m.Question)
	m.Answer.Init(m, "Answer", "Answer", "Answer", "Answer", 17)
	l.Answer.CopyStatus(&m.Answer)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", "AvailableLogin", "AvailableLogin", 18)
	l.AvailableLogin.CopyStatus(&m.AvailableLogin)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 19)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 20)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 21)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 22)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 23)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 24)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 25)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 26)
	l.Remark4.CopyStatus(&m.Remark4)
	m.Nonego.Init(m, "Nonego", "Nonego", "Nonego", "Nonego", 27)
	l.Nonego.CopyStatus(&m.Nonego)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *UserList) SetTotal(total int) {
	l.Total = total
}

func (l *UserList) GetTotal() int {
	return l.Total
}

func (l *UserList) Len() int {
	return len(l.List)
}

func (l *UserList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l UserList) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, 1024)
	bs = append(bs, []byte("{")...)
	ListB, err := json.MarshalIndent(l.List, "", "\t")
	if err != nil {
		return nil, err
	}
	ListB = append([]byte("\"List\": "), ListB...)
	bs = append(bs, ListB...)
	bs = append(bs, []byte(", ")...)
	TotalB, err := json.MarshalIndent(l.Total, "", "\t")
	if err != nil {
		return nil, err
	}
	TotalB = append([]byte("\"Total\": "), TotalB...)
	bs = append(bs, TotalB...)
	bs = append(bs, []byte("}")...)
	return bs, nil
}

func (l *UserList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.User)
}

func (l *UserList) UnmarshalMeta(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &l.User)
}

func (l *UserList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List[idx].BasicInfo = l.List[l.Len()-1].BasicInfo
		l.List[idx].StudentClass = l.List[l.Len()-1].StudentClass
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *UserList) Filter(f func(m *User) bool) []*User {
	ll := make([]*User, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *UserList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.Id.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Id.AnyValue()))
	}
	if lastModel.IntelUserCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IntelUserCode.AnyValue()))
	}
	if lastModel.UserCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UserCode.AnyValue()))
	}
	if lastModel.Name.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Name.AnyValue()))
	}
	if lastModel.Sex.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Sex.AnyValue()))
	}
	if lastModel.IdentityType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IdentityType.AnyValue()))
	}
	if lastModel.IdentityNum.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IdentityNum.AnyValue()))
	}
	if lastModel.ExpirationDate.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ExpirationDate.AnyValue()))
	}
	if lastModel.UniversityCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UniversityCode.AnyValue()))
	}
	if lastModel.UserType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UserType.AnyValue()))
	}
	if lastModel.EnrollmentStatus.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.EnrollmentStatus.AnyValue()))
	}
	if lastModel.Type.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Type.AnyValue()))
	}
	if lastModel.Password.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Password.AnyValue()))
	}
	if lastModel.Phone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Phone.AnyValue()))
	}
	if lastModel.Email.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Email.AnyValue()))
	}
	if lastModel.PictureURL.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.PictureURL.AnyValue()))
	}
	if lastModel.Question.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Question.AnyValue()))
	}
	if lastModel.Answer.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Answer.AnyValue()))
	}
	if lastModel.AvailableLogin.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AvailableLogin.AnyValue()))
	}
	if lastModel.Operator.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Operator.AnyValue()))
	}
	if lastModel.InsertDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.InsertDatetime.AnyValue()))
	}
	if lastModel.UpdateDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateDatetime.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.Remark1.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark1.AnyValue()))
	}
	if lastModel.Remark2.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark2.AnyValue()))
	}
	if lastModel.Remark3.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark3.AnyValue()))
	}
	if lastModel.Remark4.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark4.AnyValue()))
	}
	if lastModel.Nonego.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Nonego.AnyValue()))
	}
	builder.WriteString(lastModel.AggCheckDup())
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *UserList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

func (m *User) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (l *UserList) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

type UserCacheElem struct {
	hashValue  string
	model      *User
	modifyTime time.Time
}

type UserListCacheElem struct {
	hashValue  string
	list       *UserList
	modifyTime time.Time
}

type UserCacheManager struct {
	container map[string]*UserCacheElem
	query     chan string
	in        chan *UserCacheElem
	out       chan *UserCacheElem
}

type UserListCacheManager struct {
	container map[string]*UserListCacheElem
	query     chan string
	in        chan *UserListCacheElem
	out       chan *UserListCacheElem
}

func newUserCacheManager() *UserCacheManager {
	return &UserCacheManager{
		make(map[string]*UserCacheElem),
		make(chan string),
		make(chan *UserCacheElem),
		make(chan *UserCacheElem),
	}
}

func newUserListCacheManager() *UserListCacheManager {
	return &UserListCacheManager{
		make(map[string]*UserListCacheElem),
		make(chan string),
		make(chan *UserListCacheElem),
		make(chan *UserListCacheElem),
	}
}

func (mgr *UserCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *UserListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var UserCache = newUserCacheManager()

var UserListCache = newUserListCacheManager()

func (m *User) GetCache(hashVal string, timeout time.Duration) bool {
	UserCache.query <- hashVal
	elem := <-UserCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *User) SetCache(hashValue string) {
	UserCache.in <- &UserCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *UserList) GetListCache(hashValue string, timeout time.Duration) bool {
	UserListCache.query <- hashValue
	elem := <-UserListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *UserList) SetListCache(hashValue string) {
	UserListCache.in <- &UserListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func (m *User) FromQuery(query interface{}) (*User, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				m.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					m.Id.AndWhereEq(fval0.Interface())
				case "!=":
					m.Id.AndWhereNeq(fval0.Interface())
				case ">":
					m.Id.AndWhereGt(fval0.Interface())
				case ">=":
					m.Id.AndWhereGte(fval0.Interface())
				case "<":
					m.Id.AndWhereLt(fval0.Interface())
				case "<=":
					m.Id.AndWhereLte(fval0.Interface())
				case "llike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					m.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					m.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("IntelUserCode")
	if exists {
		fval1 := val.FieldByName("IntelUserCode")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				m.IntelUserCode.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					m.IntelUserCode.AndWhereEq(fval1.Interface())
				case "!=":
					m.IntelUserCode.AndWhereNeq(fval1.Interface())
				case ">":
					m.IntelUserCode.AndWhereGt(fval1.Interface())
				case ">=":
					m.IntelUserCode.AndWhereGte(fval1.Interface())
				case "<":
					m.IntelUserCode.AndWhereLt(fval1.Interface())
				case "<=":
					m.IntelUserCode.AndWhereLte(fval1.Interface())
				case "llike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					m.IntelUserCode.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("UserCode")
	if exists {
		fval2 := val.FieldByName("UserCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				m.UserCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					m.UserCode.AndWhereEq(fval2.Interface())
				case "!=":
					m.UserCode.AndWhereNeq(fval2.Interface())
				case ">":
					m.UserCode.AndWhereGt(fval2.Interface())
				case ">=":
					m.UserCode.AndWhereGte(fval2.Interface())
				case "<":
					m.UserCode.AndWhereLt(fval2.Interface())
				case "<=":
					m.UserCode.AndWhereLte(fval2.Interface())
				case "llike":
					m.UserCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					m.UserCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					m.UserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					m.UserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					m.UserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					m.UserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					m.UserCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("Name")
	if exists {
		fval3 := val.FieldByName("Name")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				m.Name.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					m.Name.AndWhereEq(fval3.Interface())
				case "!=":
					m.Name.AndWhereNeq(fval3.Interface())
				case ">":
					m.Name.AndWhereGt(fval3.Interface())
				case ">=":
					m.Name.AndWhereGte(fval3.Interface())
				case "<":
					m.Name.AndWhereLt(fval3.Interface())
				case "<=":
					m.Name.AndWhereLte(fval3.Interface())
				case "llike":
					m.Name.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					m.Name.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					m.Name.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					m.Name.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					m.Name.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					m.Name.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					m.Name.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("Sex")
	if exists {
		fval4 := val.FieldByName("Sex")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				m.Sex.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					m.Sex.AndWhereEq(fval4.Interface())
				case "!=":
					m.Sex.AndWhereNeq(fval4.Interface())
				case ">":
					m.Sex.AndWhereGt(fval4.Interface())
				case ">=":
					m.Sex.AndWhereGte(fval4.Interface())
				case "<":
					m.Sex.AndWhereLt(fval4.Interface())
				case "<=":
					m.Sex.AndWhereLte(fval4.Interface())
				case "llike":
					m.Sex.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					m.Sex.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					m.Sex.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					m.Sex.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					m.Sex.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					m.Sex.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					m.Sex.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("IdentityType")
	if exists {
		fval5 := val.FieldByName("IdentityType")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				m.IdentityType.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					m.IdentityType.AndWhereEq(fval5.Interface())
				case "!=":
					m.IdentityType.AndWhereNeq(fval5.Interface())
				case ">":
					m.IdentityType.AndWhereGt(fval5.Interface())
				case ">=":
					m.IdentityType.AndWhereGte(fval5.Interface())
				case "<":
					m.IdentityType.AndWhereLt(fval5.Interface())
				case "<=":
					m.IdentityType.AndWhereLte(fval5.Interface())
				case "llike":
					m.IdentityType.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					m.IdentityType.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					m.IdentityType.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					m.IdentityType.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					m.IdentityType.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					m.IdentityType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					m.IdentityType.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("IdentityNum")
	if exists {
		fval6 := val.FieldByName("IdentityNum")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				m.IdentityNum.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					m.IdentityNum.AndWhereEq(fval6.Interface())
				case "!=":
					m.IdentityNum.AndWhereNeq(fval6.Interface())
				case ">":
					m.IdentityNum.AndWhereGt(fval6.Interface())
				case ">=":
					m.IdentityNum.AndWhereGte(fval6.Interface())
				case "<":
					m.IdentityNum.AndWhereLt(fval6.Interface())
				case "<=":
					m.IdentityNum.AndWhereLte(fval6.Interface())
				case "llike":
					m.IdentityNum.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					m.IdentityNum.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					m.IdentityNum.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					m.IdentityNum.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					m.IdentityNum.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					m.IdentityNum.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					m.IdentityNum.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("ExpirationDate")
	if exists {
		fval7 := val.FieldByName("ExpirationDate")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				m.ExpirationDate.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					m.ExpirationDate.AndWhereEq(fval7.Interface())
				case "!=":
					m.ExpirationDate.AndWhereNeq(fval7.Interface())
				case ">":
					m.ExpirationDate.AndWhereGt(fval7.Interface())
				case ">=":
					m.ExpirationDate.AndWhereGte(fval7.Interface())
				case "<":
					m.ExpirationDate.AndWhereLt(fval7.Interface())
				case "<=":
					m.ExpirationDate.AndWhereLte(fval7.Interface())
				case "llike":
					m.ExpirationDate.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					m.ExpirationDate.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					m.ExpirationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					m.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					m.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					m.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					m.ExpirationDate.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("UniversityCode")
	if exists {
		fval8 := val.FieldByName("UniversityCode")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				m.UniversityCode.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					m.UniversityCode.AndWhereEq(fval8.Interface())
				case "!=":
					m.UniversityCode.AndWhereNeq(fval8.Interface())
				case ">":
					m.UniversityCode.AndWhereGt(fval8.Interface())
				case ">=":
					m.UniversityCode.AndWhereGte(fval8.Interface())
				case "<":
					m.UniversityCode.AndWhereLt(fval8.Interface())
				case "<=":
					m.UniversityCode.AndWhereLte(fval8.Interface())
				case "llike":
					m.UniversityCode.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					m.UniversityCode.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					m.UniversityCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					m.UniversityCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					m.UniversityCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					m.UniversityCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					m.UniversityCode.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("UserType")
	if exists {
		fval9 := val.FieldByName("UserType")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				m.UserType.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					m.UserType.AndWhereEq(fval9.Interface())
				case "!=":
					m.UserType.AndWhereNeq(fval9.Interface())
				case ">":
					m.UserType.AndWhereGt(fval9.Interface())
				case ">=":
					m.UserType.AndWhereGte(fval9.Interface())
				case "<":
					m.UserType.AndWhereLt(fval9.Interface())
				case "<=":
					m.UserType.AndWhereLte(fval9.Interface())
				case "llike":
					m.UserType.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					m.UserType.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					m.UserType.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					m.UserType.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					m.UserType.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					m.UserType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					m.UserType.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("EnrollmentStatus")
	if exists {
		fval10 := val.FieldByName("EnrollmentStatus")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				m.EnrollmentStatus.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					m.EnrollmentStatus.AndWhereEq(fval10.Interface())
				case "!=":
					m.EnrollmentStatus.AndWhereNeq(fval10.Interface())
				case ">":
					m.EnrollmentStatus.AndWhereGt(fval10.Interface())
				case ">=":
					m.EnrollmentStatus.AndWhereGte(fval10.Interface())
				case "<":
					m.EnrollmentStatus.AndWhereLt(fval10.Interface())
				case "<=":
					m.EnrollmentStatus.AndWhereLte(fval10.Interface())
				case "llike":
					m.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					m.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					m.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					m.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					m.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					m.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					m.EnrollmentStatus.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("Type")
	if exists {
		fval11 := val.FieldByName("Type")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				m.Type.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					m.Type.AndWhereEq(fval11.Interface())
				case "!=":
					m.Type.AndWhereNeq(fval11.Interface())
				case ">":
					m.Type.AndWhereGt(fval11.Interface())
				case ">=":
					m.Type.AndWhereGte(fval11.Interface())
				case "<":
					m.Type.AndWhereLt(fval11.Interface())
				case "<=":
					m.Type.AndWhereLte(fval11.Interface())
				case "llike":
					m.Type.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					m.Type.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					m.Type.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					m.Type.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					m.Type.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					m.Type.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					m.Type.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("Password")
	if exists {
		fval12 := val.FieldByName("Password")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				m.Password.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					m.Password.AndWhereEq(fval12.Interface())
				case "!=":
					m.Password.AndWhereNeq(fval12.Interface())
				case ">":
					m.Password.AndWhereGt(fval12.Interface())
				case ">=":
					m.Password.AndWhereGte(fval12.Interface())
				case "<":
					m.Password.AndWhereLt(fval12.Interface())
				case "<=":
					m.Password.AndWhereLte(fval12.Interface())
				case "llike":
					m.Password.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					m.Password.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					m.Password.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					m.Password.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					m.Password.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					m.Password.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					m.Password.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("Phone")
	if exists {
		fval13 := val.FieldByName("Phone")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				m.Phone.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					m.Phone.AndWhereEq(fval13.Interface())
				case "!=":
					m.Phone.AndWhereNeq(fval13.Interface())
				case ">":
					m.Phone.AndWhereGt(fval13.Interface())
				case ">=":
					m.Phone.AndWhereGte(fval13.Interface())
				case "<":
					m.Phone.AndWhereLt(fval13.Interface())
				case "<=":
					m.Phone.AndWhereLte(fval13.Interface())
				case "llike":
					m.Phone.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					m.Phone.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					m.Phone.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					m.Phone.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					m.Phone.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					m.Phone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					m.Phone.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("Email")
	if exists {
		fval14 := val.FieldByName("Email")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				m.Email.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					m.Email.AndWhereEq(fval14.Interface())
				case "!=":
					m.Email.AndWhereNeq(fval14.Interface())
				case ">":
					m.Email.AndWhereGt(fval14.Interface())
				case ">=":
					m.Email.AndWhereGte(fval14.Interface())
				case "<":
					m.Email.AndWhereLt(fval14.Interface())
				case "<=":
					m.Email.AndWhereLte(fval14.Interface())
				case "llike":
					m.Email.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					m.Email.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					m.Email.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					m.Email.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					m.Email.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					m.Email.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					m.Email.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("PictureURL")
	if exists {
		fval15 := val.FieldByName("PictureURL")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				m.PictureURL.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					m.PictureURL.AndWhereEq(fval15.Interface())
				case "!=":
					m.PictureURL.AndWhereNeq(fval15.Interface())
				case ">":
					m.PictureURL.AndWhereGt(fval15.Interface())
				case ">=":
					m.PictureURL.AndWhereGte(fval15.Interface())
				case "<":
					m.PictureURL.AndWhereLt(fval15.Interface())
				case "<=":
					m.PictureURL.AndWhereLte(fval15.Interface())
				case "llike":
					m.PictureURL.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					m.PictureURL.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					m.PictureURL.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					m.PictureURL.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					m.PictureURL.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					m.PictureURL.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					m.PictureURL.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("Question")
	if exists {
		fval16 := val.FieldByName("Question")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				m.Question.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					m.Question.AndWhereEq(fval16.Interface())
				case "!=":
					m.Question.AndWhereNeq(fval16.Interface())
				case ">":
					m.Question.AndWhereGt(fval16.Interface())
				case ">=":
					m.Question.AndWhereGte(fval16.Interface())
				case "<":
					m.Question.AndWhereLt(fval16.Interface())
				case "<=":
					m.Question.AndWhereLte(fval16.Interface())
				case "llike":
					m.Question.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					m.Question.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					m.Question.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					m.Question.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					m.Question.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					m.Question.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					m.Question.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("Answer")
	if exists {
		fval17 := val.FieldByName("Answer")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				m.Answer.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					m.Answer.AndWhereEq(fval17.Interface())
				case "!=":
					m.Answer.AndWhereNeq(fval17.Interface())
				case ">":
					m.Answer.AndWhereGt(fval17.Interface())
				case ">=":
					m.Answer.AndWhereGte(fval17.Interface())
				case "<":
					m.Answer.AndWhereLt(fval17.Interface())
				case "<=":
					m.Answer.AndWhereLte(fval17.Interface())
				case "llike":
					m.Answer.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					m.Answer.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					m.Answer.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					m.Answer.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					m.Answer.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					m.Answer.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					m.Answer.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("AvailableLogin")
	if exists {
		fval18 := val.FieldByName("AvailableLogin")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				m.AvailableLogin.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					m.AvailableLogin.AndWhereEq(fval18.Interface())
				case "!=":
					m.AvailableLogin.AndWhereNeq(fval18.Interface())
				case ">":
					m.AvailableLogin.AndWhereGt(fval18.Interface())
				case ">=":
					m.AvailableLogin.AndWhereGte(fval18.Interface())
				case "<":
					m.AvailableLogin.AndWhereLt(fval18.Interface())
				case "<=":
					m.AvailableLogin.AndWhereLte(fval18.Interface())
				case "llike":
					m.AvailableLogin.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					m.AvailableLogin.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					m.AvailableLogin.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					m.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					m.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					m.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					m.AvailableLogin.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Operator")
	if exists {
		fval19 := val.FieldByName("Operator")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				m.Operator.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					m.Operator.AndWhereEq(fval19.Interface())
				case "!=":
					m.Operator.AndWhereNeq(fval19.Interface())
				case ">":
					m.Operator.AndWhereGt(fval19.Interface())
				case ">=":
					m.Operator.AndWhereGte(fval19.Interface())
				case "<":
					m.Operator.AndWhereLt(fval19.Interface())
				case "<=":
					m.Operator.AndWhereLte(fval19.Interface())
				case "llike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					m.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					m.Operator.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval20 := val.FieldByName("InsertDatetime")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				m.InsertDatetime.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					m.InsertDatetime.AndWhereEq(fval20.Interface())
				case "!=":
					m.InsertDatetime.AndWhereNeq(fval20.Interface())
				case ">":
					m.InsertDatetime.AndWhereGt(fval20.Interface())
				case ">=":
					m.InsertDatetime.AndWhereGte(fval20.Interface())
				case "<":
					m.InsertDatetime.AndWhereLt(fval20.Interface())
				case "<=":
					m.InsertDatetime.AndWhereLte(fval20.Interface())
				case "llike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					m.InsertDatetime.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval21 := val.FieldByName("UpdateDatetime")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				m.UpdateDatetime.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					m.UpdateDatetime.AndWhereEq(fval21.Interface())
				case "!=":
					m.UpdateDatetime.AndWhereNeq(fval21.Interface())
				case ">":
					m.UpdateDatetime.AndWhereGt(fval21.Interface())
				case ">=":
					m.UpdateDatetime.AndWhereGte(fval21.Interface())
				case "<":
					m.UpdateDatetime.AndWhereLt(fval21.Interface())
				case "<=":
					m.UpdateDatetime.AndWhereLte(fval21.Interface())
				case "llike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					m.UpdateDatetime.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("Status")
	if exists {
		fval22 := val.FieldByName("Status")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				m.Status.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					m.Status.AndWhereEq(fval22.Interface())
				case "!=":
					m.Status.AndWhereNeq(fval22.Interface())
				case ">":
					m.Status.AndWhereGt(fval22.Interface())
				case ">=":
					m.Status.AndWhereGte(fval22.Interface())
				case "<":
					m.Status.AndWhereLt(fval22.Interface())
				case "<=":
					m.Status.AndWhereLte(fval22.Interface())
				case "llike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					m.Status.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					m.Status.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("Remark1")
	if exists {
		fval23 := val.FieldByName("Remark1")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				m.Remark1.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					m.Remark1.AndWhereEq(fval23.Interface())
				case "!=":
					m.Remark1.AndWhereNeq(fval23.Interface())
				case ">":
					m.Remark1.AndWhereGt(fval23.Interface())
				case ">=":
					m.Remark1.AndWhereGte(fval23.Interface())
				case "<":
					m.Remark1.AndWhereLt(fval23.Interface())
				case "<=":
					m.Remark1.AndWhereLte(fval23.Interface())
				case "llike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					m.Remark1.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("Remark2")
	if exists {
		fval24 := val.FieldByName("Remark2")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				m.Remark2.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					m.Remark2.AndWhereEq(fval24.Interface())
				case "!=":
					m.Remark2.AndWhereNeq(fval24.Interface())
				case ">":
					m.Remark2.AndWhereGt(fval24.Interface())
				case ">=":
					m.Remark2.AndWhereGte(fval24.Interface())
				case "<":
					m.Remark2.AndWhereLt(fval24.Interface())
				case "<=":
					m.Remark2.AndWhereLte(fval24.Interface())
				case "llike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					m.Remark2.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("Remark3")
	if exists {
		fval25 := val.FieldByName("Remark3")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				m.Remark3.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					m.Remark3.AndWhereEq(fval25.Interface())
				case "!=":
					m.Remark3.AndWhereNeq(fval25.Interface())
				case ">":
					m.Remark3.AndWhereGt(fval25.Interface())
				case ">=":
					m.Remark3.AndWhereGte(fval25.Interface())
				case "<":
					m.Remark3.AndWhereLt(fval25.Interface())
				case "<=":
					m.Remark3.AndWhereLte(fval25.Interface())
				case "llike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					m.Remark3.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("Remark4")
	if exists {
		fval26 := val.FieldByName("Remark4")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				m.Remark4.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					m.Remark4.AndWhereEq(fval26.Interface())
				case "!=":
					m.Remark4.AndWhereNeq(fval26.Interface())
				case ">":
					m.Remark4.AndWhereGt(fval26.Interface())
				case ">=":
					m.Remark4.AndWhereGte(fval26.Interface())
				case "<":
					m.Remark4.AndWhereLt(fval26.Interface())
				case "<=":
					m.Remark4.AndWhereLte(fval26.Interface())
				case "llike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					m.Remark4.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("Nonego")
	if exists {
		fval27 := val.FieldByName("Nonego")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				m.Nonego.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					m.Nonego.AndWhereEq(fval27.Interface())
				case "!=":
					m.Nonego.AndWhereNeq(fval27.Interface())
				case ">":
					m.Nonego.AndWhereGt(fval27.Interface())
				case ">=":
					m.Nonego.AndWhereGte(fval27.Interface())
				case "<":
					m.Nonego.AndWhereLt(fval27.Interface())
				case "<=":
					m.Nonego.AndWhereLte(fval27.Interface())
				case "llike":
					m.Nonego.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					m.Nonego.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					m.Nonego.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					m.Nonego.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					m.Nonego.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					m.Nonego.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					m.Nonego.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	return m, nil
}
func (l *UserList) FromQuery(query interface{}) (*UserList, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				l.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					l.Id.AndWhereEq(fval0.Interface())
				case "!=":
					l.Id.AndWhereNeq(fval0.Interface())
				case ">":
					l.Id.AndWhereGt(fval0.Interface())
				case ">=":
					l.Id.AndWhereGte(fval0.Interface())
				case "<":
					l.Id.AndWhereLt(fval0.Interface())
				case "<=":
					l.Id.AndWhereLte(fval0.Interface())
				case "llike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					l.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					l.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("IntelUserCode")
	if exists {
		fval1 := val.FieldByName("IntelUserCode")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				l.IntelUserCode.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					l.IntelUserCode.AndWhereEq(fval1.Interface())
				case "!=":
					l.IntelUserCode.AndWhereNeq(fval1.Interface())
				case ">":
					l.IntelUserCode.AndWhereGt(fval1.Interface())
				case ">=":
					l.IntelUserCode.AndWhereGte(fval1.Interface())
				case "<":
					l.IntelUserCode.AndWhereLt(fval1.Interface())
				case "<=":
					l.IntelUserCode.AndWhereLte(fval1.Interface())
				case "llike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					l.IntelUserCode.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("UserCode")
	if exists {
		fval2 := val.FieldByName("UserCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				l.UserCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					l.UserCode.AndWhereEq(fval2.Interface())
				case "!=":
					l.UserCode.AndWhereNeq(fval2.Interface())
				case ">":
					l.UserCode.AndWhereGt(fval2.Interface())
				case ">=":
					l.UserCode.AndWhereGte(fval2.Interface())
				case "<":
					l.UserCode.AndWhereLt(fval2.Interface())
				case "<=":
					l.UserCode.AndWhereLte(fval2.Interface())
				case "llike":
					l.UserCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					l.UserCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					l.UserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					l.UserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					l.UserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					l.UserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					l.UserCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("Name")
	if exists {
		fval3 := val.FieldByName("Name")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				l.Name.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					l.Name.AndWhereEq(fval3.Interface())
				case "!=":
					l.Name.AndWhereNeq(fval3.Interface())
				case ">":
					l.Name.AndWhereGt(fval3.Interface())
				case ">=":
					l.Name.AndWhereGte(fval3.Interface())
				case "<":
					l.Name.AndWhereLt(fval3.Interface())
				case "<=":
					l.Name.AndWhereLte(fval3.Interface())
				case "llike":
					l.Name.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					l.Name.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					l.Name.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					l.Name.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					l.Name.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					l.Name.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					l.Name.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("Sex")
	if exists {
		fval4 := val.FieldByName("Sex")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				l.Sex.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					l.Sex.AndWhereEq(fval4.Interface())
				case "!=":
					l.Sex.AndWhereNeq(fval4.Interface())
				case ">":
					l.Sex.AndWhereGt(fval4.Interface())
				case ">=":
					l.Sex.AndWhereGte(fval4.Interface())
				case "<":
					l.Sex.AndWhereLt(fval4.Interface())
				case "<=":
					l.Sex.AndWhereLte(fval4.Interface())
				case "llike":
					l.Sex.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					l.Sex.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					l.Sex.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					l.Sex.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					l.Sex.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					l.Sex.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					l.Sex.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("IdentityType")
	if exists {
		fval5 := val.FieldByName("IdentityType")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				l.IdentityType.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					l.IdentityType.AndWhereEq(fval5.Interface())
				case "!=":
					l.IdentityType.AndWhereNeq(fval5.Interface())
				case ">":
					l.IdentityType.AndWhereGt(fval5.Interface())
				case ">=":
					l.IdentityType.AndWhereGte(fval5.Interface())
				case "<":
					l.IdentityType.AndWhereLt(fval5.Interface())
				case "<=":
					l.IdentityType.AndWhereLte(fval5.Interface())
				case "llike":
					l.IdentityType.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					l.IdentityType.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					l.IdentityType.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					l.IdentityType.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					l.IdentityType.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					l.IdentityType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					l.IdentityType.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("IdentityNum")
	if exists {
		fval6 := val.FieldByName("IdentityNum")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				l.IdentityNum.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					l.IdentityNum.AndWhereEq(fval6.Interface())
				case "!=":
					l.IdentityNum.AndWhereNeq(fval6.Interface())
				case ">":
					l.IdentityNum.AndWhereGt(fval6.Interface())
				case ">=":
					l.IdentityNum.AndWhereGte(fval6.Interface())
				case "<":
					l.IdentityNum.AndWhereLt(fval6.Interface())
				case "<=":
					l.IdentityNum.AndWhereLte(fval6.Interface())
				case "llike":
					l.IdentityNum.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					l.IdentityNum.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					l.IdentityNum.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					l.IdentityNum.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					l.IdentityNum.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					l.IdentityNum.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					l.IdentityNum.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("ExpirationDate")
	if exists {
		fval7 := val.FieldByName("ExpirationDate")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				l.ExpirationDate.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					l.ExpirationDate.AndWhereEq(fval7.Interface())
				case "!=":
					l.ExpirationDate.AndWhereNeq(fval7.Interface())
				case ">":
					l.ExpirationDate.AndWhereGt(fval7.Interface())
				case ">=":
					l.ExpirationDate.AndWhereGte(fval7.Interface())
				case "<":
					l.ExpirationDate.AndWhereLt(fval7.Interface())
				case "<=":
					l.ExpirationDate.AndWhereLte(fval7.Interface())
				case "llike":
					l.ExpirationDate.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					l.ExpirationDate.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					l.ExpirationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					l.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					l.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					l.ExpirationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					l.ExpirationDate.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("UniversityCode")
	if exists {
		fval8 := val.FieldByName("UniversityCode")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				l.UniversityCode.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					l.UniversityCode.AndWhereEq(fval8.Interface())
				case "!=":
					l.UniversityCode.AndWhereNeq(fval8.Interface())
				case ">":
					l.UniversityCode.AndWhereGt(fval8.Interface())
				case ">=":
					l.UniversityCode.AndWhereGte(fval8.Interface())
				case "<":
					l.UniversityCode.AndWhereLt(fval8.Interface())
				case "<=":
					l.UniversityCode.AndWhereLte(fval8.Interface())
				case "llike":
					l.UniversityCode.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					l.UniversityCode.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					l.UniversityCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					l.UniversityCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					l.UniversityCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					l.UniversityCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					l.UniversityCode.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("UserType")
	if exists {
		fval9 := val.FieldByName("UserType")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				l.UserType.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					l.UserType.AndWhereEq(fval9.Interface())
				case "!=":
					l.UserType.AndWhereNeq(fval9.Interface())
				case ">":
					l.UserType.AndWhereGt(fval9.Interface())
				case ">=":
					l.UserType.AndWhereGte(fval9.Interface())
				case "<":
					l.UserType.AndWhereLt(fval9.Interface())
				case "<=":
					l.UserType.AndWhereLte(fval9.Interface())
				case "llike":
					l.UserType.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					l.UserType.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					l.UserType.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					l.UserType.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					l.UserType.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					l.UserType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					l.UserType.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("EnrollmentStatus")
	if exists {
		fval10 := val.FieldByName("EnrollmentStatus")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				l.EnrollmentStatus.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					l.EnrollmentStatus.AndWhereEq(fval10.Interface())
				case "!=":
					l.EnrollmentStatus.AndWhereNeq(fval10.Interface())
				case ">":
					l.EnrollmentStatus.AndWhereGt(fval10.Interface())
				case ">=":
					l.EnrollmentStatus.AndWhereGte(fval10.Interface())
				case "<":
					l.EnrollmentStatus.AndWhereLt(fval10.Interface())
				case "<=":
					l.EnrollmentStatus.AndWhereLte(fval10.Interface())
				case "llike":
					l.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					l.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					l.EnrollmentStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					l.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					l.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					l.EnrollmentStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					l.EnrollmentStatus.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("Type")
	if exists {
		fval11 := val.FieldByName("Type")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				l.Type.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					l.Type.AndWhereEq(fval11.Interface())
				case "!=":
					l.Type.AndWhereNeq(fval11.Interface())
				case ">":
					l.Type.AndWhereGt(fval11.Interface())
				case ">=":
					l.Type.AndWhereGte(fval11.Interface())
				case "<":
					l.Type.AndWhereLt(fval11.Interface())
				case "<=":
					l.Type.AndWhereLte(fval11.Interface())
				case "llike":
					l.Type.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					l.Type.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					l.Type.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					l.Type.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					l.Type.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					l.Type.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					l.Type.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("Password")
	if exists {
		fval12 := val.FieldByName("Password")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				l.Password.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					l.Password.AndWhereEq(fval12.Interface())
				case "!=":
					l.Password.AndWhereNeq(fval12.Interface())
				case ">":
					l.Password.AndWhereGt(fval12.Interface())
				case ">=":
					l.Password.AndWhereGte(fval12.Interface())
				case "<":
					l.Password.AndWhereLt(fval12.Interface())
				case "<=":
					l.Password.AndWhereLte(fval12.Interface())
				case "llike":
					l.Password.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					l.Password.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					l.Password.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					l.Password.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					l.Password.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					l.Password.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					l.Password.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("Phone")
	if exists {
		fval13 := val.FieldByName("Phone")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				l.Phone.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					l.Phone.AndWhereEq(fval13.Interface())
				case "!=":
					l.Phone.AndWhereNeq(fval13.Interface())
				case ">":
					l.Phone.AndWhereGt(fval13.Interface())
				case ">=":
					l.Phone.AndWhereGte(fval13.Interface())
				case "<":
					l.Phone.AndWhereLt(fval13.Interface())
				case "<=":
					l.Phone.AndWhereLte(fval13.Interface())
				case "llike":
					l.Phone.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					l.Phone.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					l.Phone.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					l.Phone.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					l.Phone.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					l.Phone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					l.Phone.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("Email")
	if exists {
		fval14 := val.FieldByName("Email")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				l.Email.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					l.Email.AndWhereEq(fval14.Interface())
				case "!=":
					l.Email.AndWhereNeq(fval14.Interface())
				case ">":
					l.Email.AndWhereGt(fval14.Interface())
				case ">=":
					l.Email.AndWhereGte(fval14.Interface())
				case "<":
					l.Email.AndWhereLt(fval14.Interface())
				case "<=":
					l.Email.AndWhereLte(fval14.Interface())
				case "llike":
					l.Email.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					l.Email.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					l.Email.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					l.Email.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					l.Email.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					l.Email.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					l.Email.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("PictureURL")
	if exists {
		fval15 := val.FieldByName("PictureURL")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				l.PictureURL.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					l.PictureURL.AndWhereEq(fval15.Interface())
				case "!=":
					l.PictureURL.AndWhereNeq(fval15.Interface())
				case ">":
					l.PictureURL.AndWhereGt(fval15.Interface())
				case ">=":
					l.PictureURL.AndWhereGte(fval15.Interface())
				case "<":
					l.PictureURL.AndWhereLt(fval15.Interface())
				case "<=":
					l.PictureURL.AndWhereLte(fval15.Interface())
				case "llike":
					l.PictureURL.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					l.PictureURL.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					l.PictureURL.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					l.PictureURL.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					l.PictureURL.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					l.PictureURL.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					l.PictureURL.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("Question")
	if exists {
		fval16 := val.FieldByName("Question")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				l.Question.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					l.Question.AndWhereEq(fval16.Interface())
				case "!=":
					l.Question.AndWhereNeq(fval16.Interface())
				case ">":
					l.Question.AndWhereGt(fval16.Interface())
				case ">=":
					l.Question.AndWhereGte(fval16.Interface())
				case "<":
					l.Question.AndWhereLt(fval16.Interface())
				case "<=":
					l.Question.AndWhereLte(fval16.Interface())
				case "llike":
					l.Question.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					l.Question.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					l.Question.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					l.Question.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					l.Question.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					l.Question.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					l.Question.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("Answer")
	if exists {
		fval17 := val.FieldByName("Answer")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				l.Answer.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					l.Answer.AndWhereEq(fval17.Interface())
				case "!=":
					l.Answer.AndWhereNeq(fval17.Interface())
				case ">":
					l.Answer.AndWhereGt(fval17.Interface())
				case ">=":
					l.Answer.AndWhereGte(fval17.Interface())
				case "<":
					l.Answer.AndWhereLt(fval17.Interface())
				case "<=":
					l.Answer.AndWhereLte(fval17.Interface())
				case "llike":
					l.Answer.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					l.Answer.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					l.Answer.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					l.Answer.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					l.Answer.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					l.Answer.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					l.Answer.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("AvailableLogin")
	if exists {
		fval18 := val.FieldByName("AvailableLogin")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				l.AvailableLogin.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					l.AvailableLogin.AndWhereEq(fval18.Interface())
				case "!=":
					l.AvailableLogin.AndWhereNeq(fval18.Interface())
				case ">":
					l.AvailableLogin.AndWhereGt(fval18.Interface())
				case ">=":
					l.AvailableLogin.AndWhereGte(fval18.Interface())
				case "<":
					l.AvailableLogin.AndWhereLt(fval18.Interface())
				case "<=":
					l.AvailableLogin.AndWhereLte(fval18.Interface())
				case "llike":
					l.AvailableLogin.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					l.AvailableLogin.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					l.AvailableLogin.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					l.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					l.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					l.AvailableLogin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					l.AvailableLogin.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Operator")
	if exists {
		fval19 := val.FieldByName("Operator")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				l.Operator.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					l.Operator.AndWhereEq(fval19.Interface())
				case "!=":
					l.Operator.AndWhereNeq(fval19.Interface())
				case ">":
					l.Operator.AndWhereGt(fval19.Interface())
				case ">=":
					l.Operator.AndWhereGte(fval19.Interface())
				case "<":
					l.Operator.AndWhereLt(fval19.Interface())
				case "<=":
					l.Operator.AndWhereLte(fval19.Interface())
				case "llike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					l.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					l.Operator.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval20 := val.FieldByName("InsertDatetime")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				l.InsertDatetime.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					l.InsertDatetime.AndWhereEq(fval20.Interface())
				case "!=":
					l.InsertDatetime.AndWhereNeq(fval20.Interface())
				case ">":
					l.InsertDatetime.AndWhereGt(fval20.Interface())
				case ">=":
					l.InsertDatetime.AndWhereGte(fval20.Interface())
				case "<":
					l.InsertDatetime.AndWhereLt(fval20.Interface())
				case "<=":
					l.InsertDatetime.AndWhereLte(fval20.Interface())
				case "llike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					l.InsertDatetime.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval21 := val.FieldByName("UpdateDatetime")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				l.UpdateDatetime.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					l.UpdateDatetime.AndWhereEq(fval21.Interface())
				case "!=":
					l.UpdateDatetime.AndWhereNeq(fval21.Interface())
				case ">":
					l.UpdateDatetime.AndWhereGt(fval21.Interface())
				case ">=":
					l.UpdateDatetime.AndWhereGte(fval21.Interface())
				case "<":
					l.UpdateDatetime.AndWhereLt(fval21.Interface())
				case "<=":
					l.UpdateDatetime.AndWhereLte(fval21.Interface())
				case "llike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					l.UpdateDatetime.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("Status")
	if exists {
		fval22 := val.FieldByName("Status")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				l.Status.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					l.Status.AndWhereEq(fval22.Interface())
				case "!=":
					l.Status.AndWhereNeq(fval22.Interface())
				case ">":
					l.Status.AndWhereGt(fval22.Interface())
				case ">=":
					l.Status.AndWhereGte(fval22.Interface())
				case "<":
					l.Status.AndWhereLt(fval22.Interface())
				case "<=":
					l.Status.AndWhereLte(fval22.Interface())
				case "llike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					l.Status.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					l.Status.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("Remark1")
	if exists {
		fval23 := val.FieldByName("Remark1")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				l.Remark1.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					l.Remark1.AndWhereEq(fval23.Interface())
				case "!=":
					l.Remark1.AndWhereNeq(fval23.Interface())
				case ">":
					l.Remark1.AndWhereGt(fval23.Interface())
				case ">=":
					l.Remark1.AndWhereGte(fval23.Interface())
				case "<":
					l.Remark1.AndWhereLt(fval23.Interface())
				case "<=":
					l.Remark1.AndWhereLte(fval23.Interface())
				case "llike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					l.Remark1.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("Remark2")
	if exists {
		fval24 := val.FieldByName("Remark2")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				l.Remark2.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					l.Remark2.AndWhereEq(fval24.Interface())
				case "!=":
					l.Remark2.AndWhereNeq(fval24.Interface())
				case ">":
					l.Remark2.AndWhereGt(fval24.Interface())
				case ">=":
					l.Remark2.AndWhereGte(fval24.Interface())
				case "<":
					l.Remark2.AndWhereLt(fval24.Interface())
				case "<=":
					l.Remark2.AndWhereLte(fval24.Interface())
				case "llike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					l.Remark2.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("Remark3")
	if exists {
		fval25 := val.FieldByName("Remark3")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				l.Remark3.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					l.Remark3.AndWhereEq(fval25.Interface())
				case "!=":
					l.Remark3.AndWhereNeq(fval25.Interface())
				case ">":
					l.Remark3.AndWhereGt(fval25.Interface())
				case ">=":
					l.Remark3.AndWhereGte(fval25.Interface())
				case "<":
					l.Remark3.AndWhereLt(fval25.Interface())
				case "<=":
					l.Remark3.AndWhereLte(fval25.Interface())
				case "llike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					l.Remark3.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("Remark4")
	if exists {
		fval26 := val.FieldByName("Remark4")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				l.Remark4.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					l.Remark4.AndWhereEq(fval26.Interface())
				case "!=":
					l.Remark4.AndWhereNeq(fval26.Interface())
				case ">":
					l.Remark4.AndWhereGt(fval26.Interface())
				case ">=":
					l.Remark4.AndWhereGte(fval26.Interface())
				case "<":
					l.Remark4.AndWhereLt(fval26.Interface())
				case "<=":
					l.Remark4.AndWhereLte(fval26.Interface())
				case "llike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					l.Remark4.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("Nonego")
	if exists {
		fval27 := val.FieldByName("Nonego")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				l.Nonego.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					l.Nonego.AndWhereEq(fval27.Interface())
				case "!=":
					l.Nonego.AndWhereNeq(fval27.Interface())
				case ">":
					l.Nonego.AndWhereGt(fval27.Interface())
				case ">=":
					l.Nonego.AndWhereGte(fval27.Interface())
				case "<":
					l.Nonego.AndWhereLt(fval27.Interface())
				case "<=":
					l.Nonego.AndWhereLte(fval27.Interface())
				case "llike":
					l.Nonego.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					l.Nonego.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					l.Nonego.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					l.Nonego.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					l.Nonego.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					l.Nonego.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					l.Nonego.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	return l, nil
}
func NewStudentbasicinfo() *Studentbasicinfo {
	m := &Studentbasicinfo{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 2)
	m.Class.Init(m, "Class", "Class", "Class", "Class", 3)
	m.OtherName.Init(m, "OtherName", "OtherName", "OtherName", "OtherName", 4)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", "NameInPinyin", "NameInPinyin", 5)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", "EnglishName", "EnglishName", 6)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", "CountryCode", "CountryCode", 7)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", "NationalityCode", "NationalityCode", 8)
	m.Birthday.Init(m, "Birthday", "Birthday", "Birthday", "Birthday", 9)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", "PoliticalCode", "PoliticalCode", 10)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", "QQAcct", "QQAcct", 11)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", "WeChatAcct", "WeChatAcct", 12)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", "BankCardNumber", "BankCardNumber", 13)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", "AccountBankCode", "AccountBankCode", 14)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", "MaritalCode", "MaritalCode", 16)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", 17)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", 18)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", "Hobbies", "Hobbies", 19)
	m.Creed.Init(m, "Creed", "Creed", "Creed", "Creed", 20)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", 21)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", "FamilyAddress", "FamilyAddress", 22)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", "DetailAddress", "DetailAddress", 23)
	m.PostCode.Init(m, "PostCode", "PostCode", "PostCode", "PostCode", 24)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", "HomePhone", "HomePhone", 25)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", 26)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 27)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", 28)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", "MidSchoolName", "MidSchoolName", 29)
	m.Referee.Init(m, "Referee", "Referee", "Referee", "Referee", 30)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", "RefereeDuty", "RefereeDuty", 31)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", "RefereePhone", "RefereePhone", 32)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", 33)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", "AdmissionYear", "AdmissionYear", 35)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", 36)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", "StudentOrigin", "StudentOrigin", 37)
	m.BizType.Init(m, "BizType", "BizType", "BizType", "BizType", 38)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", "TaskCode", "TaskCode", 39)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", "ApproveStatus", "ApproveStatus", 40)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 41)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 42)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 43)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 44)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", "StudentStatus", "StudentStatus", 45)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", "IsAuth", "IsAuth", 46)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 47)
	m.Zone.Init(m, "Zone", "Zone", "Zone", "Zone", 48)
	m.Building.Init(m, "Building", "Building", "Building", "Building", 49)
	m.Unit.Init(m, "Unit", "Unit", "Unit", "Unit", 50)
	m.Room.Init(m, "Room", "Room", "Room", "Room", 51)
	m.Bed.Init(m, "Bed", "Bed", "Bed", "Bed", 52)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", "StatusSort", "StatusSort", 53)
	m.Height.Init(m, "Height", "Height", "Height", "Height", 54)
	m.Weight.Init(m, "Weight", "Weight", "Weight", "Weight", 55)
	m.FootSize.Init(m, "FootSize", "FootSize", "FootSize", "FootSize", 56)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", "ClothSize", "ClothSize", 57)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", "HeadSize", "HeadSize", 58)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 59)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 60)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 61)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 62)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", "IsPayment", "IsPayment", 63)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", "IsCheckIn", "IsCheckIn", 64)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", 65)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", "OriginAreaName", "OriginAreaName", 66)
	m.InitRel()
	return m
}

func newSubStudentbasicinfo(parent nborm.Model) *Studentbasicinfo {
	m := &Studentbasicinfo{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 2)
	m.Class.Init(m, "Class", "Class", "Class", "Class", 3)
	m.OtherName.Init(m, "OtherName", "OtherName", "OtherName", "OtherName", 4)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", "NameInPinyin", "NameInPinyin", 5)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", "EnglishName", "EnglishName", 6)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", "CountryCode", "CountryCode", 7)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", "NationalityCode", "NationalityCode", 8)
	m.Birthday.Init(m, "Birthday", "Birthday", "Birthday", "Birthday", 9)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", "PoliticalCode", "PoliticalCode", 10)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", "QQAcct", "QQAcct", 11)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", "WeChatAcct", "WeChatAcct", 12)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", "BankCardNumber", "BankCardNumber", 13)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", "AccountBankCode", "AccountBankCode", 14)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", "MaritalCode", "MaritalCode", 16)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", 17)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", 18)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", "Hobbies", "Hobbies", 19)
	m.Creed.Init(m, "Creed", "Creed", "Creed", "Creed", 20)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", 21)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", "FamilyAddress", "FamilyAddress", 22)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", "DetailAddress", "DetailAddress", 23)
	m.PostCode.Init(m, "PostCode", "PostCode", "PostCode", "PostCode", 24)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", "HomePhone", "HomePhone", 25)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", 26)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 27)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", 28)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", "MidSchoolName", "MidSchoolName", 29)
	m.Referee.Init(m, "Referee", "Referee", "Referee", "Referee", 30)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", "RefereeDuty", "RefereeDuty", 31)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", "RefereePhone", "RefereePhone", 32)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", 33)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", "AdmissionYear", "AdmissionYear", 35)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", 36)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", "StudentOrigin", "StudentOrigin", 37)
	m.BizType.Init(m, "BizType", "BizType", "BizType", "BizType", 38)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", "TaskCode", "TaskCode", 39)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", "ApproveStatus", "ApproveStatus", 40)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 41)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 42)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 43)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 44)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", "StudentStatus", "StudentStatus", 45)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", "IsAuth", "IsAuth", 46)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 47)
	m.Zone.Init(m, "Zone", "Zone", "Zone", "Zone", 48)
	m.Building.Init(m, "Building", "Building", "Building", "Building", 49)
	m.Unit.Init(m, "Unit", "Unit", "Unit", "Unit", 50)
	m.Room.Init(m, "Room", "Room", "Room", "Room", 51)
	m.Bed.Init(m, "Bed", "Bed", "Bed", "Bed", 52)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", "StatusSort", "StatusSort", 53)
	m.Height.Init(m, "Height", "Height", "Height", "Height", 54)
	m.Weight.Init(m, "Weight", "Weight", "Weight", "Weight", 55)
	m.FootSize.Init(m, "FootSize", "FootSize", "FootSize", "FootSize", 56)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", "ClothSize", "ClothSize", 57)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", "HeadSize", "HeadSize", 58)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 59)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 60)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 61)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 62)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", "IsPayment", "IsPayment", 63)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", "IsCheckIn", "IsCheckIn", 64)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", 65)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", "OriginAreaName", "OriginAreaName", 66)
	return m
}

func (m *Studentbasicinfo) InitRel() {
	m.User = newSubUser(m)
	var relInfo0 *nborm.RelationInfo
	relInfo0 = relInfo0.Append("User", m.User, nborm.NewExpr("@=@", &m.IntelUserCode, &m.User.IntelUserCode))
	m.AppendRelation(relInfo0)
	m.StudentClass = newSubClass(m)
	var relInfo1 *nborm.RelationInfo
	relInfo1 = relInfo1.Append("StudentClass", m.StudentClass, nborm.NewExpr("@=@", &m.Class, &m.StudentClass.ClassCode))
	m.AppendRelation(relInfo1)
	m.AddRelInited()
}

func (m *Studentbasicinfo) DB() string {
	return "*"
}

func (m *Studentbasicinfo) Tab() string {
	return "studentbasicinfo"
}

func (m *Studentbasicinfo) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"Id", "Id", &m.Id, 0},
		{"RecordId", "RecordId", &m.RecordId, 1},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode, 2},
		{"Class", "Class", &m.Class, 3},
		{"OtherName", "OtherName", &m.OtherName, 4},
		{"NameInPinyin", "NameInPinyin", &m.NameInPinyin, 5},
		{"EnglishName", "EnglishName", &m.EnglishName, 6},
		{"CountryCode", "CountryCode", &m.CountryCode, 7},
		{"NationalityCode", "NationalityCode", &m.NationalityCode, 8},
		{"Birthday", "Birthday", &m.Birthday, 9},
		{"PoliticalCode", "PoliticalCode", &m.PoliticalCode, 10},
		{"QQAcct", "QQAcct", &m.QQAcct, 11},
		{"WeChatAcct", "WeChatAcct", &m.WeChatAcct, 12},
		{"BankCardNumber", "BankCardNumber", &m.BankCardNumber, 13},
		{"AccountBankCode", "AccountBankCode", &m.AccountBankCode, 14},
		{"AllPowerfulCardNum", "AllPowerfulCardNum", &m.AllPowerfulCardNum, 15},
		{"MaritalCode", "MaritalCode", &m.MaritalCode, 16},
		{"OriginAreaCode", "OriginAreaCode", &m.OriginAreaCode, 17},
		{"StudentAreaCode", "StudentAreaCode", &m.StudentAreaCode, 18},
		{"Hobbies", "Hobbies", &m.Hobbies, 19},
		{"Creed", "Creed", &m.Creed, 20},
		{"TrainTicketinterval", "TrainTicketinterval", &m.TrainTicketinterval, 21},
		{"FamilyAddress", "FamilyAddress", &m.FamilyAddress, 22},
		{"DetailAddress", "DetailAddress", &m.DetailAddress, 23},
		{"PostCode", "PostCode", &m.PostCode, 24},
		{"HomePhone", "HomePhone", &m.HomePhone, 25},
		{"EnrollmentDate", "EnrollmentDate", &m.EnrollmentDate, 26},
		{"GraduationDate", "GraduationDate", &m.GraduationDate, 27},
		{"MidSchoolAddress", "MidSchoolAddress", &m.MidSchoolAddress, 28},
		{"MidSchoolName", "MidSchoolName", &m.MidSchoolName, 29},
		{"Referee", "Referee", &m.Referee, 30},
		{"RefereeDuty", "RefereeDuty", &m.RefereeDuty, 31},
		{"RefereePhone", "RefereePhone", &m.RefereePhone, 32},
		{"AdmissionTicketNo", "AdmissionTicketNo", &m.AdmissionTicketNo, 33},
		{"CollegeEntranceExamScores", "CollegeEntranceExamScores", &m.CollegeEntranceExamScores, 34},
		{"AdmissionYear", "AdmissionYear", &m.AdmissionYear, 35},
		{"ForeignLanguageCode", "ForeignLanguageCode", &m.ForeignLanguageCode, 36},
		{"StudentOrigin", "StudentOrigin", &m.StudentOrigin, 37},
		{"BizType", "BizType", &m.BizType, 38},
		{"TaskCode", "TaskCode", &m.TaskCode, 39},
		{"ApproveStatus", "ApproveStatus", &m.ApproveStatus, 40},
		{"Operator", "Operator", &m.Operator, 41},
		{"InsertDatetime", "InsertDatetime", &m.InsertDatetime, 42},
		{"UpdateDatetime", "UpdateDatetime", &m.UpdateDatetime, 43},
		{"Status", "Status", &m.Status, 44},
		{"StudentStatus", "StudentStatus", &m.StudentStatus, 45},
		{"IsAuth", "IsAuth", &m.IsAuth, 46},
		{"Campus", "Campus", &m.Campus, 47},
		{"Zone", "Zone", &m.Zone, 48},
		{"Building", "Building", &m.Building, 49},
		{"Unit", "Unit", &m.Unit, 50},
		{"Room", "Room", &m.Room, 51},
		{"Bed", "Bed", &m.Bed, 52},
		{"StatusSort", "StatusSort", &m.StatusSort, 53},
		{"Height", "Height", &m.Height, 54},
		{"Weight", "Weight", &m.Weight, 55},
		{"FootSize", "FootSize", &m.FootSize, 56},
		{"ClothSize", "ClothSize", &m.ClothSize, 57},
		{"HeadSize", "HeadSize", &m.HeadSize, 58},
		{"Remark1", "Remark1", &m.Remark1, 59},
		{"Remark2", "Remark2", &m.Remark2, 60},
		{"Remark3", "Remark3", &m.Remark3, 61},
		{"Remark4", "Remark4", &m.Remark4, 62},
		{"IsPayment", "IsPayment", &m.IsPayment, 63},
		{"isCheckIn", "IsCheckIn", &m.IsCheckIn, 64},
		{"GetMilitaryTC", "GetMilitaryTC", &m.GetMilitaryTC, 65},
		{"OriginAreaName", "OriginAreaName", &m.OriginAreaName, 66},
	}
}

func (m *Studentbasicinfo) AutoIncField() nborm.Field {
	return &m.Id
}

func (m *Studentbasicinfo) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.Id,
	}
}

func (m *Studentbasicinfo) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.IntelUserCode,
		},
	}
}
func (m Studentbasicinfo) MarshalJSON() ([]byte, error) {
	if !m.IsSynced() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	buffer.WriteString("{\n\"Aggs\": ")
	metaB, err := json.MarshalIndent(m.Meta, "", "\t")
	if err != nil {
		return nil, err
	}
	buffer.Write(metaB)
	if m.Id.IsValid() {
		buffer.WriteString(",\n\"Id\": ")
		IdB, err := json.MarshalIndent(m.Id, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdB)
	}
	if m.RecordId.IsValid() {
		buffer.WriteString(",\n\"RecordId\": ")
		RecordIdB, err := json.MarshalIndent(m.RecordId, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RecordIdB)
	}
	if m.IntelUserCode.IsValid() {
		buffer.WriteString(",\n\"IntelUserCode\": ")
		IntelUserCodeB, err := json.MarshalIndent(m.IntelUserCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IntelUserCodeB)
	}
	if m.Class.IsValid() {
		buffer.WriteString(",\n\"Class\": ")
		ClassB, err := json.MarshalIndent(m.Class, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClassB)
	}
	if m.OtherName.IsValid() {
		buffer.WriteString(",\n\"OtherName\": ")
		OtherNameB, err := json.MarshalIndent(m.OtherName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OtherNameB)
	}
	if m.NameInPinyin.IsValid() {
		buffer.WriteString(",\n\"NameInPinyin\": ")
		NameInPinyinB, err := json.MarshalIndent(m.NameInPinyin, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NameInPinyinB)
	}
	if m.EnglishName.IsValid() {
		buffer.WriteString(",\n\"EnglishName\": ")
		EnglishNameB, err := json.MarshalIndent(m.EnglishName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EnglishNameB)
	}
	if m.CountryCode.IsValid() {
		buffer.WriteString(",\n\"CountryCode\": ")
		CountryCodeB, err := json.MarshalIndent(m.CountryCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CountryCodeB)
	}
	if m.NationalityCode.IsValid() {
		buffer.WriteString(",\n\"NationalityCode\": ")
		NationalityCodeB, err := json.MarshalIndent(m.NationalityCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NationalityCodeB)
	}
	if m.Birthday.IsValid() {
		buffer.WriteString(",\n\"Birthday\": ")
		BirthdayB, err := json.MarshalIndent(m.Birthday, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BirthdayB)
	}
	if m.PoliticalCode.IsValid() {
		buffer.WriteString(",\n\"PoliticalCode\": ")
		PoliticalCodeB, err := json.MarshalIndent(m.PoliticalCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PoliticalCodeB)
	}
	if m.QQAcct.IsValid() {
		buffer.WriteString(",\n\"QQAcct\": ")
		QQAcctB, err := json.MarshalIndent(m.QQAcct, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(QQAcctB)
	}
	if m.WeChatAcct.IsValid() {
		buffer.WriteString(",\n\"WeChatAcct\": ")
		WeChatAcctB, err := json.MarshalIndent(m.WeChatAcct, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(WeChatAcctB)
	}
	if m.BankCardNumber.IsValid() {
		buffer.WriteString(",\n\"BankCardNumber\": ")
		BankCardNumberB, err := json.MarshalIndent(m.BankCardNumber, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BankCardNumberB)
	}
	if m.AccountBankCode.IsValid() {
		buffer.WriteString(",\n\"AccountBankCode\": ")
		AccountBankCodeB, err := json.MarshalIndent(m.AccountBankCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AccountBankCodeB)
	}
	if m.AllPowerfulCardNum.IsValid() {
		buffer.WriteString(",\n\"AllPowerfulCardNum\": ")
		AllPowerfulCardNumB, err := json.MarshalIndent(m.AllPowerfulCardNum, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AllPowerfulCardNumB)
	}
	if m.MaritalCode.IsValid() {
		buffer.WriteString(",\n\"MaritalCode\": ")
		MaritalCodeB, err := json.MarshalIndent(m.MaritalCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(MaritalCodeB)
	}
	if m.OriginAreaCode.IsValid() {
		buffer.WriteString(",\n\"OriginAreaCode\": ")
		OriginAreaCodeB, err := json.MarshalIndent(m.OriginAreaCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OriginAreaCodeB)
	}
	if m.StudentAreaCode.IsValid() {
		buffer.WriteString(",\n\"StudentAreaCode\": ")
		StudentAreaCodeB, err := json.MarshalIndent(m.StudentAreaCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentAreaCodeB)
	}
	if m.Hobbies.IsValid() {
		buffer.WriteString(",\n\"Hobbies\": ")
		HobbiesB, err := json.MarshalIndent(m.Hobbies, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(HobbiesB)
	}
	if m.Creed.IsValid() {
		buffer.WriteString(",\n\"Creed\": ")
		CreedB, err := json.MarshalIndent(m.Creed, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CreedB)
	}
	if m.TrainTicketinterval.IsValid() {
		buffer.WriteString(",\n\"TrainTicketinterval\": ")
		TrainTicketintervalB, err := json.MarshalIndent(m.TrainTicketinterval, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(TrainTicketintervalB)
	}
	if m.FamilyAddress.IsValid() {
		buffer.WriteString(",\n\"FamilyAddress\": ")
		FamilyAddressB, err := json.MarshalIndent(m.FamilyAddress, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(FamilyAddressB)
	}
	if m.DetailAddress.IsValid() {
		buffer.WriteString(",\n\"DetailAddress\": ")
		DetailAddressB, err := json.MarshalIndent(m.DetailAddress, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(DetailAddressB)
	}
	if m.PostCode.IsValid() {
		buffer.WriteString(",\n\"PostCode\": ")
		PostCodeB, err := json.MarshalIndent(m.PostCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PostCodeB)
	}
	if m.HomePhone.IsValid() {
		buffer.WriteString(",\n\"HomePhone\": ")
		HomePhoneB, err := json.MarshalIndent(m.HomePhone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(HomePhoneB)
	}
	if m.EnrollmentDate.IsValid() {
		buffer.WriteString(",\n\"EnrollmentDate\": ")
		EnrollmentDateB, err := json.MarshalIndent(m.EnrollmentDate, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EnrollmentDateB)
	}
	if m.GraduationDate.IsValid() {
		buffer.WriteString(",\n\"GraduationDate\": ")
		GraduationDateB, err := json.MarshalIndent(m.GraduationDate, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GraduationDateB)
	}
	if m.MidSchoolAddress.IsValid() {
		buffer.WriteString(",\n\"MidSchoolAddress\": ")
		MidSchoolAddressB, err := json.MarshalIndent(m.MidSchoolAddress, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(MidSchoolAddressB)
	}
	if m.MidSchoolName.IsValid() {
		buffer.WriteString(",\n\"MidSchoolName\": ")
		MidSchoolNameB, err := json.MarshalIndent(m.MidSchoolName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(MidSchoolNameB)
	}
	if m.Referee.IsValid() {
		buffer.WriteString(",\n\"Referee\": ")
		RefereeB, err := json.MarshalIndent(m.Referee, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RefereeB)
	}
	if m.RefereeDuty.IsValid() {
		buffer.WriteString(",\n\"RefereeDuty\": ")
		RefereeDutyB, err := json.MarshalIndent(m.RefereeDuty, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RefereeDutyB)
	}
	if m.RefereePhone.IsValid() {
		buffer.WriteString(",\n\"RefereePhone\": ")
		RefereePhoneB, err := json.MarshalIndent(m.RefereePhone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RefereePhoneB)
	}
	if m.AdmissionTicketNo.IsValid() {
		buffer.WriteString(",\n\"AdmissionTicketNo\": ")
		AdmissionTicketNoB, err := json.MarshalIndent(m.AdmissionTicketNo, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AdmissionTicketNoB)
	}
	if m.CollegeEntranceExamScores.IsValid() {
		buffer.WriteString(",\n\"CollegeEntranceExamScores\": ")
		CollegeEntranceExamScoresB, err := json.MarshalIndent(m.CollegeEntranceExamScores, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CollegeEntranceExamScoresB)
	}
	if m.AdmissionYear.IsValid() {
		buffer.WriteString(",\n\"AdmissionYear\": ")
		AdmissionYearB, err := json.MarshalIndent(m.AdmissionYear, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AdmissionYearB)
	}
	if m.ForeignLanguageCode.IsValid() {
		buffer.WriteString(",\n\"ForeignLanguageCode\": ")
		ForeignLanguageCodeB, err := json.MarshalIndent(m.ForeignLanguageCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ForeignLanguageCodeB)
	}
	if m.StudentOrigin.IsValid() {
		buffer.WriteString(",\n\"StudentOrigin\": ")
		StudentOriginB, err := json.MarshalIndent(m.StudentOrigin, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentOriginB)
	}
	if m.BizType.IsValid() {
		buffer.WriteString(",\n\"BizType\": ")
		BizTypeB, err := json.MarshalIndent(m.BizType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BizTypeB)
	}
	if m.TaskCode.IsValid() {
		buffer.WriteString(",\n\"TaskCode\": ")
		TaskCodeB, err := json.MarshalIndent(m.TaskCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(TaskCodeB)
	}
	if m.ApproveStatus.IsValid() {
		buffer.WriteString(",\n\"ApproveStatus\": ")
		ApproveStatusB, err := json.MarshalIndent(m.ApproveStatus, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ApproveStatusB)
	}
	if m.Operator.IsValid() {
		buffer.WriteString(",\n\"Operator\": ")
		OperatorB, err := json.MarshalIndent(m.Operator, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OperatorB)
	}
	if m.InsertDatetime.IsValid() {
		buffer.WriteString(",\n\"InsertDatetime\": ")
		InsertDatetimeB, err := json.MarshalIndent(m.InsertDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(InsertDatetimeB)
	}
	if m.UpdateDatetime.IsValid() {
		buffer.WriteString(",\n\"UpdateDatetime\": ")
		UpdateDatetimeB, err := json.MarshalIndent(m.UpdateDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateDatetimeB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.StudentStatus.IsValid() {
		buffer.WriteString(",\n\"StudentStatus\": ")
		StudentStatusB, err := json.MarshalIndent(m.StudentStatus, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentStatusB)
	}
	if m.IsAuth.IsValid() {
		buffer.WriteString(",\n\"IsAuth\": ")
		IsAuthB, err := json.MarshalIndent(m.IsAuth, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IsAuthB)
	}
	if m.Campus.IsValid() {
		buffer.WriteString(",\n\"Campus\": ")
		CampusB, err := json.MarshalIndent(m.Campus, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CampusB)
	}
	if m.Zone.IsValid() {
		buffer.WriteString(",\n\"Zone\": ")
		ZoneB, err := json.MarshalIndent(m.Zone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ZoneB)
	}
	if m.Building.IsValid() {
		buffer.WriteString(",\n\"Building\": ")
		BuildingB, err := json.MarshalIndent(m.Building, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BuildingB)
	}
	if m.Unit.IsValid() {
		buffer.WriteString(",\n\"Unit\": ")
		UnitB, err := json.MarshalIndent(m.Unit, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UnitB)
	}
	if m.Room.IsValid() {
		buffer.WriteString(",\n\"Room\": ")
		RoomB, err := json.MarshalIndent(m.Room, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RoomB)
	}
	if m.Bed.IsValid() {
		buffer.WriteString(",\n\"Bed\": ")
		BedB, err := json.MarshalIndent(m.Bed, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(BedB)
	}
	if m.StatusSort.IsValid() {
		buffer.WriteString(",\n\"StatusSort\": ")
		StatusSortB, err := json.MarshalIndent(m.StatusSort, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusSortB)
	}
	if m.Height.IsValid() {
		buffer.WriteString(",\n\"Height\": ")
		HeightB, err := json.MarshalIndent(m.Height, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(HeightB)
	}
	if m.Weight.IsValid() {
		buffer.WriteString(",\n\"Weight\": ")
		WeightB, err := json.MarshalIndent(m.Weight, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(WeightB)
	}
	if m.FootSize.IsValid() {
		buffer.WriteString(",\n\"FootSize\": ")
		FootSizeB, err := json.MarshalIndent(m.FootSize, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(FootSizeB)
	}
	if m.ClothSize.IsValid() {
		buffer.WriteString(",\n\"ClothSize\": ")
		ClothSizeB, err := json.MarshalIndent(m.ClothSize, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClothSizeB)
	}
	if m.HeadSize.IsValid() {
		buffer.WriteString(",\n\"HeadSize\": ")
		HeadSizeB, err := json.MarshalIndent(m.HeadSize, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(HeadSizeB)
	}
	if m.Remark1.IsValid() {
		buffer.WriteString(",\n\"Remark1\": ")
		Remark1B, err := json.MarshalIndent(m.Remark1, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark1B)
	}
	if m.Remark2.IsValid() {
		buffer.WriteString(",\n\"Remark2\": ")
		Remark2B, err := json.MarshalIndent(m.Remark2, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark2B)
	}
	if m.Remark3.IsValid() {
		buffer.WriteString(",\n\"Remark3\": ")
		Remark3B, err := json.MarshalIndent(m.Remark3, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark3B)
	}
	if m.Remark4.IsValid() {
		buffer.WriteString(",\n\"Remark4\": ")
		Remark4B, err := json.MarshalIndent(m.Remark4, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark4B)
	}
	if m.IsPayment.IsValid() {
		buffer.WriteString(",\n\"IsPayment\": ")
		IsPaymentB, err := json.MarshalIndent(m.IsPayment, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IsPaymentB)
	}
	if m.IsCheckIn.IsValid() {
		buffer.WriteString(",\n\"IsCheckIn\": ")
		IsCheckInB, err := json.MarshalIndent(m.IsCheckIn, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IsCheckInB)
	}
	if m.GetMilitaryTC.IsValid() {
		buffer.WriteString(",\n\"GetMilitaryTC\": ")
		GetMilitaryTCB, err := json.MarshalIndent(m.GetMilitaryTC, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GetMilitaryTCB)
	}
	if m.OriginAreaName.IsValid() {
		buffer.WriteString(",\n\"OriginAreaName\": ")
		OriginAreaNameB, err := json.MarshalIndent(m.OriginAreaName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OriginAreaNameB)
	}
	if m.User != nil && m.User.IsSynced() {
		buffer.WriteString(",\n\"User\": ")
		UserB, err := json.MarshalIndent(m.User, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UserB)
	}
	if m.StudentClass != nil && m.StudentClass.IsSynced() {
		buffer.WriteString(",\n\"StudentClass\": ")
		StudentClassB, err := json.MarshalIndent(m.StudentClass, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentClassB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type StudentbasicinfoList struct {
	Studentbasicinfo `json:"-"`
	dupMap           map[string]int
	List             []*Studentbasicinfo
	Total            int
}

func (m *Studentbasicinfo) Collapse() {
	if m.User != nil && m.User.IsSynced() {
		m.User.Collapse()
	}
	if m.StudentClass != nil && m.StudentClass.IsSynced() {
		m.StudentClass.Collapse()
	}
}

func NewStudentbasicinfoList() *StudentbasicinfoList {
	l := &StudentbasicinfoList{
		Studentbasicinfo{},
		make(map[string]int),
		make([]*Studentbasicinfo, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 2)
	l.Class.Init(l, "Class", "Class", "Class", "Class", 3)
	l.OtherName.Init(l, "OtherName", "OtherName", "OtherName", "OtherName", 4)
	l.NameInPinyin.Init(l, "NameInPinyin", "NameInPinyin", "NameInPinyin", "NameInPinyin", 5)
	l.EnglishName.Init(l, "EnglishName", "EnglishName", "EnglishName", "EnglishName", 6)
	l.CountryCode.Init(l, "CountryCode", "CountryCode", "CountryCode", "CountryCode", 7)
	l.NationalityCode.Init(l, "NationalityCode", "NationalityCode", "NationalityCode", "NationalityCode", 8)
	l.Birthday.Init(l, "Birthday", "Birthday", "Birthday", "Birthday", 9)
	l.PoliticalCode.Init(l, "PoliticalCode", "PoliticalCode", "PoliticalCode", "PoliticalCode", 10)
	l.QQAcct.Init(l, "QQAcct", "QQAcct", "QQAcct", "QQAcct", 11)
	l.WeChatAcct.Init(l, "WeChatAcct", "WeChatAcct", "WeChatAcct", "WeChatAcct", 12)
	l.BankCardNumber.Init(l, "BankCardNumber", "BankCardNumber", "BankCardNumber", "BankCardNumber", 13)
	l.AccountBankCode.Init(l, "AccountBankCode", "AccountBankCode", "AccountBankCode", "AccountBankCode", 14)
	l.AllPowerfulCardNum.Init(l, "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.MaritalCode.Init(l, "MaritalCode", "MaritalCode", "MaritalCode", "MaritalCode", 16)
	l.OriginAreaCode.Init(l, "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", 17)
	l.StudentAreaCode.Init(l, "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", 18)
	l.Hobbies.Init(l, "Hobbies", "Hobbies", "Hobbies", "Hobbies", 19)
	l.Creed.Init(l, "Creed", "Creed", "Creed", "Creed", 20)
	l.TrainTicketinterval.Init(l, "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", 21)
	l.FamilyAddress.Init(l, "FamilyAddress", "FamilyAddress", "FamilyAddress", "FamilyAddress", 22)
	l.DetailAddress.Init(l, "DetailAddress", "DetailAddress", "DetailAddress", "DetailAddress", 23)
	l.PostCode.Init(l, "PostCode", "PostCode", "PostCode", "PostCode", 24)
	l.HomePhone.Init(l, "HomePhone", "HomePhone", "HomePhone", "HomePhone", 25)
	l.EnrollmentDate.Init(l, "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", 26)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 27)
	l.MidSchoolAddress.Init(l, "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolName.Init(l, "MidSchoolName", "MidSchoolName", "MidSchoolName", "MidSchoolName", 29)
	l.Referee.Init(l, "Referee", "Referee", "Referee", "Referee", 30)
	l.RefereeDuty.Init(l, "RefereeDuty", "RefereeDuty", "RefereeDuty", "RefereeDuty", 31)
	l.RefereePhone.Init(l, "RefereePhone", "RefereePhone", "RefereePhone", "RefereePhone", 32)
	l.AdmissionTicketNo.Init(l, "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.CollegeEntranceExamScores.Init(l, "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.AdmissionYear.Init(l, "AdmissionYear", "AdmissionYear", "AdmissionYear", "AdmissionYear", 35)
	l.ForeignLanguageCode.Init(l, "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.StudentOrigin.Init(l, "StudentOrigin", "StudentOrigin", "StudentOrigin", "StudentOrigin", 37)
	l.BizType.Init(l, "BizType", "BizType", "BizType", "BizType", 38)
	l.TaskCode.Init(l, "TaskCode", "TaskCode", "TaskCode", "TaskCode", 39)
	l.ApproveStatus.Init(l, "ApproveStatus", "ApproveStatus", "ApproveStatus", "ApproveStatus", 40)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 41)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 42)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 43)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 44)
	l.StudentStatus.Init(l, "StudentStatus", "StudentStatus", "StudentStatus", "StudentStatus", 45)
	l.IsAuth.Init(l, "IsAuth", "IsAuth", "IsAuth", "IsAuth", 46)
	l.Campus.Init(l, "Campus", "Campus", "Campus", "Campus", 47)
	l.Zone.Init(l, "Zone", "Zone", "Zone", "Zone", 48)
	l.Building.Init(l, "Building", "Building", "Building", "Building", 49)
	l.Unit.Init(l, "Unit", "Unit", "Unit", "Unit", 50)
	l.Room.Init(l, "Room", "Room", "Room", "Room", 51)
	l.Bed.Init(l, "Bed", "Bed", "Bed", "Bed", 52)
	l.StatusSort.Init(l, "StatusSort", "StatusSort", "StatusSort", "StatusSort", 53)
	l.Height.Init(l, "Height", "Height", "Height", "Height", 54)
	l.Weight.Init(l, "Weight", "Weight", "Weight", "Weight", 55)
	l.FootSize.Init(l, "FootSize", "FootSize", "FootSize", "FootSize", 56)
	l.ClothSize.Init(l, "ClothSize", "ClothSize", "ClothSize", "ClothSize", 57)
	l.HeadSize.Init(l, "HeadSize", "HeadSize", "HeadSize", "HeadSize", 58)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 59)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 60)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 61)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 62)
	l.IsPayment.Init(l, "IsPayment", "IsPayment", "IsPayment", "IsPayment", 63)
	l.IsCheckIn.Init(l, "isCheckIn", "IsCheckIn", "IsCheckIn", "IsCheckIn", 64)
	l.GetMilitaryTC.Init(l, "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", 65)
	l.OriginAreaName.Init(l, "OriginAreaName", "OriginAreaName", "OriginAreaName", "OriginAreaName", 66)
	l.InitRel()
	return l
}

func newSubStudentbasicinfoList(parent nborm.Model) *StudentbasicinfoList {
	l := &StudentbasicinfoList{
		Studentbasicinfo{},
		make(map[string]int),
		make([]*Studentbasicinfo, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 2)
	l.Class.Init(l, "Class", "Class", "Class", "Class", 3)
	l.OtherName.Init(l, "OtherName", "OtherName", "OtherName", "OtherName", 4)
	l.NameInPinyin.Init(l, "NameInPinyin", "NameInPinyin", "NameInPinyin", "NameInPinyin", 5)
	l.EnglishName.Init(l, "EnglishName", "EnglishName", "EnglishName", "EnglishName", 6)
	l.CountryCode.Init(l, "CountryCode", "CountryCode", "CountryCode", "CountryCode", 7)
	l.NationalityCode.Init(l, "NationalityCode", "NationalityCode", "NationalityCode", "NationalityCode", 8)
	l.Birthday.Init(l, "Birthday", "Birthday", "Birthday", "Birthday", 9)
	l.PoliticalCode.Init(l, "PoliticalCode", "PoliticalCode", "PoliticalCode", "PoliticalCode", 10)
	l.QQAcct.Init(l, "QQAcct", "QQAcct", "QQAcct", "QQAcct", 11)
	l.WeChatAcct.Init(l, "WeChatAcct", "WeChatAcct", "WeChatAcct", "WeChatAcct", 12)
	l.BankCardNumber.Init(l, "BankCardNumber", "BankCardNumber", "BankCardNumber", "BankCardNumber", 13)
	l.AccountBankCode.Init(l, "AccountBankCode", "AccountBankCode", "AccountBankCode", "AccountBankCode", 14)
	l.AllPowerfulCardNum.Init(l, "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.MaritalCode.Init(l, "MaritalCode", "MaritalCode", "MaritalCode", "MaritalCode", 16)
	l.OriginAreaCode.Init(l, "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", 17)
	l.StudentAreaCode.Init(l, "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", 18)
	l.Hobbies.Init(l, "Hobbies", "Hobbies", "Hobbies", "Hobbies", 19)
	l.Creed.Init(l, "Creed", "Creed", "Creed", "Creed", 20)
	l.TrainTicketinterval.Init(l, "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", 21)
	l.FamilyAddress.Init(l, "FamilyAddress", "FamilyAddress", "FamilyAddress", "FamilyAddress", 22)
	l.DetailAddress.Init(l, "DetailAddress", "DetailAddress", "DetailAddress", "DetailAddress", 23)
	l.PostCode.Init(l, "PostCode", "PostCode", "PostCode", "PostCode", 24)
	l.HomePhone.Init(l, "HomePhone", "HomePhone", "HomePhone", "HomePhone", 25)
	l.EnrollmentDate.Init(l, "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", 26)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 27)
	l.MidSchoolAddress.Init(l, "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolName.Init(l, "MidSchoolName", "MidSchoolName", "MidSchoolName", "MidSchoolName", 29)
	l.Referee.Init(l, "Referee", "Referee", "Referee", "Referee", 30)
	l.RefereeDuty.Init(l, "RefereeDuty", "RefereeDuty", "RefereeDuty", "RefereeDuty", 31)
	l.RefereePhone.Init(l, "RefereePhone", "RefereePhone", "RefereePhone", "RefereePhone", 32)
	l.AdmissionTicketNo.Init(l, "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.CollegeEntranceExamScores.Init(l, "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.AdmissionYear.Init(l, "AdmissionYear", "AdmissionYear", "AdmissionYear", "AdmissionYear", 35)
	l.ForeignLanguageCode.Init(l, "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.StudentOrigin.Init(l, "StudentOrigin", "StudentOrigin", "StudentOrigin", "StudentOrigin", 37)
	l.BizType.Init(l, "BizType", "BizType", "BizType", "BizType", 38)
	l.TaskCode.Init(l, "TaskCode", "TaskCode", "TaskCode", "TaskCode", 39)
	l.ApproveStatus.Init(l, "ApproveStatus", "ApproveStatus", "ApproveStatus", "ApproveStatus", 40)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 41)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 42)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 43)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 44)
	l.StudentStatus.Init(l, "StudentStatus", "StudentStatus", "StudentStatus", "StudentStatus", 45)
	l.IsAuth.Init(l, "IsAuth", "IsAuth", "IsAuth", "IsAuth", 46)
	l.Campus.Init(l, "Campus", "Campus", "Campus", "Campus", 47)
	l.Zone.Init(l, "Zone", "Zone", "Zone", "Zone", 48)
	l.Building.Init(l, "Building", "Building", "Building", "Building", 49)
	l.Unit.Init(l, "Unit", "Unit", "Unit", "Unit", 50)
	l.Room.Init(l, "Room", "Room", "Room", "Room", 51)
	l.Bed.Init(l, "Bed", "Bed", "Bed", "Bed", 52)
	l.StatusSort.Init(l, "StatusSort", "StatusSort", "StatusSort", "StatusSort", 53)
	l.Height.Init(l, "Height", "Height", "Height", "Height", 54)
	l.Weight.Init(l, "Weight", "Weight", "Weight", "Weight", 55)
	l.FootSize.Init(l, "FootSize", "FootSize", "FootSize", "FootSize", 56)
	l.ClothSize.Init(l, "ClothSize", "ClothSize", "ClothSize", "ClothSize", 57)
	l.HeadSize.Init(l, "HeadSize", "HeadSize", "HeadSize", "HeadSize", 58)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 59)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 60)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 61)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 62)
	l.IsPayment.Init(l, "IsPayment", "IsPayment", "IsPayment", "IsPayment", 63)
	l.IsCheckIn.Init(l, "isCheckIn", "IsCheckIn", "IsCheckIn", "IsCheckIn", 64)
	l.GetMilitaryTC.Init(l, "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", 65)
	l.OriginAreaName.Init(l, "OriginAreaName", "OriginAreaName", "OriginAreaName", "OriginAreaName", 66)
	return l
}

func (l *StudentbasicinfoList) NewModel() nborm.Model {
	m := &Studentbasicinfo{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.RecordId.CopyStatus(&m.RecordId)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", "IntelUserCode", "IntelUserCode", 2)
	l.IntelUserCode.CopyStatus(&m.IntelUserCode)
	m.Class.Init(m, "Class", "Class", "Class", "Class", 3)
	l.Class.CopyStatus(&m.Class)
	m.OtherName.Init(m, "OtherName", "OtherName", "OtherName", "OtherName", 4)
	l.OtherName.CopyStatus(&m.OtherName)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", "NameInPinyin", "NameInPinyin", 5)
	l.NameInPinyin.CopyStatus(&m.NameInPinyin)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", "EnglishName", "EnglishName", 6)
	l.EnglishName.CopyStatus(&m.EnglishName)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", "CountryCode", "CountryCode", 7)
	l.CountryCode.CopyStatus(&m.CountryCode)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", "NationalityCode", "NationalityCode", 8)
	l.NationalityCode.CopyStatus(&m.NationalityCode)
	m.Birthday.Init(m, "Birthday", "Birthday", "Birthday", "Birthday", 9)
	l.Birthday.CopyStatus(&m.Birthday)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", "PoliticalCode", "PoliticalCode", 10)
	l.PoliticalCode.CopyStatus(&m.PoliticalCode)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", "QQAcct", "QQAcct", 11)
	l.QQAcct.CopyStatus(&m.QQAcct)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", "WeChatAcct", "WeChatAcct", 12)
	l.WeChatAcct.CopyStatus(&m.WeChatAcct)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", "BankCardNumber", "BankCardNumber", 13)
	l.BankCardNumber.CopyStatus(&m.BankCardNumber)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", "AccountBankCode", "AccountBankCode", 14)
	l.AccountBankCode.CopyStatus(&m.AccountBankCode)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.AllPowerfulCardNum.CopyStatus(&m.AllPowerfulCardNum)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", "MaritalCode", "MaritalCode", 16)
	l.MaritalCode.CopyStatus(&m.MaritalCode)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", "OriginAreaCode", 17)
	l.OriginAreaCode.CopyStatus(&m.OriginAreaCode)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", "StudentAreaCode", 18)
	l.StudentAreaCode.CopyStatus(&m.StudentAreaCode)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", "Hobbies", "Hobbies", 19)
	l.Hobbies.CopyStatus(&m.Hobbies)
	m.Creed.Init(m, "Creed", "Creed", "Creed", "Creed", 20)
	l.Creed.CopyStatus(&m.Creed)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", "TrainTicketinterval", 21)
	l.TrainTicketinterval.CopyStatus(&m.TrainTicketinterval)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", "FamilyAddress", "FamilyAddress", 22)
	l.FamilyAddress.CopyStatus(&m.FamilyAddress)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", "DetailAddress", "DetailAddress", 23)
	l.DetailAddress.CopyStatus(&m.DetailAddress)
	m.PostCode.Init(m, "PostCode", "PostCode", "PostCode", "PostCode", 24)
	l.PostCode.CopyStatus(&m.PostCode)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", "HomePhone", "HomePhone", 25)
	l.HomePhone.CopyStatus(&m.HomePhone)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", "EnrollmentDate", 26)
	l.EnrollmentDate.CopyStatus(&m.EnrollmentDate)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 27)
	l.GraduationDate.CopyStatus(&m.GraduationDate)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolAddress.CopyStatus(&m.MidSchoolAddress)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", "MidSchoolName", "MidSchoolName", 29)
	l.MidSchoolName.CopyStatus(&m.MidSchoolName)
	m.Referee.Init(m, "Referee", "Referee", "Referee", "Referee", 30)
	l.Referee.CopyStatus(&m.Referee)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", "RefereeDuty", "RefereeDuty", 31)
	l.RefereeDuty.CopyStatus(&m.RefereeDuty)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", "RefereePhone", "RefereePhone", 32)
	l.RefereePhone.CopyStatus(&m.RefereePhone)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.AdmissionTicketNo.CopyStatus(&m.AdmissionTicketNo)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.CollegeEntranceExamScores.CopyStatus(&m.CollegeEntranceExamScores)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", "AdmissionYear", "AdmissionYear", 35)
	l.AdmissionYear.CopyStatus(&m.AdmissionYear)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.ForeignLanguageCode.CopyStatus(&m.ForeignLanguageCode)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", "StudentOrigin", "StudentOrigin", 37)
	l.StudentOrigin.CopyStatus(&m.StudentOrigin)
	m.BizType.Init(m, "BizType", "BizType", "BizType", "BizType", 38)
	l.BizType.CopyStatus(&m.BizType)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", "TaskCode", "TaskCode", 39)
	l.TaskCode.CopyStatus(&m.TaskCode)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", "ApproveStatus", "ApproveStatus", 40)
	l.ApproveStatus.CopyStatus(&m.ApproveStatus)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 41)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 42)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 43)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 44)
	l.Status.CopyStatus(&m.Status)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", "StudentStatus", "StudentStatus", 45)
	l.StudentStatus.CopyStatus(&m.StudentStatus)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", "IsAuth", "IsAuth", 46)
	l.IsAuth.CopyStatus(&m.IsAuth)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 47)
	l.Campus.CopyStatus(&m.Campus)
	m.Zone.Init(m, "Zone", "Zone", "Zone", "Zone", 48)
	l.Zone.CopyStatus(&m.Zone)
	m.Building.Init(m, "Building", "Building", "Building", "Building", 49)
	l.Building.CopyStatus(&m.Building)
	m.Unit.Init(m, "Unit", "Unit", "Unit", "Unit", 50)
	l.Unit.CopyStatus(&m.Unit)
	m.Room.Init(m, "Room", "Room", "Room", "Room", 51)
	l.Room.CopyStatus(&m.Room)
	m.Bed.Init(m, "Bed", "Bed", "Bed", "Bed", 52)
	l.Bed.CopyStatus(&m.Bed)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", "StatusSort", "StatusSort", 53)
	l.StatusSort.CopyStatus(&m.StatusSort)
	m.Height.Init(m, "Height", "Height", "Height", "Height", 54)
	l.Height.CopyStatus(&m.Height)
	m.Weight.Init(m, "Weight", "Weight", "Weight", "Weight", 55)
	l.Weight.CopyStatus(&m.Weight)
	m.FootSize.Init(m, "FootSize", "FootSize", "FootSize", "FootSize", 56)
	l.FootSize.CopyStatus(&m.FootSize)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", "ClothSize", "ClothSize", 57)
	l.ClothSize.CopyStatus(&m.ClothSize)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", "HeadSize", "HeadSize", 58)
	l.HeadSize.CopyStatus(&m.HeadSize)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 59)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 60)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 61)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 62)
	l.Remark4.CopyStatus(&m.Remark4)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", "IsPayment", "IsPayment", 63)
	l.IsPayment.CopyStatus(&m.IsPayment)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", "IsCheckIn", "IsCheckIn", 64)
	l.IsCheckIn.CopyStatus(&m.IsCheckIn)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", "GetMilitaryTC", 65)
	l.GetMilitaryTC.CopyStatus(&m.GetMilitaryTC)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", "OriginAreaName", "OriginAreaName", 66)
	l.OriginAreaName.CopyStatus(&m.OriginAreaName)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentbasicinfoList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentbasicinfoList) GetTotal() int {
	return l.Total
}

func (l *StudentbasicinfoList) Len() int {
	return len(l.List)
}

func (l *StudentbasicinfoList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentbasicinfoList) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, 1024)
	bs = append(bs, []byte("{")...)
	ListB, err := json.MarshalIndent(l.List, "", "\t")
	if err != nil {
		return nil, err
	}
	ListB = append([]byte("\"List\": "), ListB...)
	bs = append(bs, ListB...)
	bs = append(bs, []byte(", ")...)
	TotalB, err := json.MarshalIndent(l.Total, "", "\t")
	if err != nil {
		return nil, err
	}
	TotalB = append([]byte("\"Total\": "), TotalB...)
	bs = append(bs, TotalB...)
	bs = append(bs, []byte("}")...)
	return bs, nil
}

func (l *StudentbasicinfoList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.Studentbasicinfo)
}

func (l *StudentbasicinfoList) UnmarshalMeta(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &l.Studentbasicinfo)
}

func (l *StudentbasicinfoList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List[idx].User = l.List[l.Len()-1].User
		l.List[idx].StudentClass = l.List[l.Len()-1].StudentClass
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *StudentbasicinfoList) Filter(f func(m *Studentbasicinfo) bool) []*Studentbasicinfo {
	ll := make([]*Studentbasicinfo, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *StudentbasicinfoList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.Id.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Id.AnyValue()))
	}
	if lastModel.RecordId.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RecordId.AnyValue()))
	}
	if lastModel.IntelUserCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IntelUserCode.AnyValue()))
	}
	if lastModel.Class.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Class.AnyValue()))
	}
	if lastModel.OtherName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.OtherName.AnyValue()))
	}
	if lastModel.NameInPinyin.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.NameInPinyin.AnyValue()))
	}
	if lastModel.EnglishName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.EnglishName.AnyValue()))
	}
	if lastModel.CountryCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.CountryCode.AnyValue()))
	}
	if lastModel.NationalityCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.NationalityCode.AnyValue()))
	}
	if lastModel.Birthday.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Birthday.AnyValue()))
	}
	if lastModel.PoliticalCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.PoliticalCode.AnyValue()))
	}
	if lastModel.QQAcct.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.QQAcct.AnyValue()))
	}
	if lastModel.WeChatAcct.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.WeChatAcct.AnyValue()))
	}
	if lastModel.BankCardNumber.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.BankCardNumber.AnyValue()))
	}
	if lastModel.AccountBankCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AccountBankCode.AnyValue()))
	}
	if lastModel.AllPowerfulCardNum.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AllPowerfulCardNum.AnyValue()))
	}
	if lastModel.MaritalCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.MaritalCode.AnyValue()))
	}
	if lastModel.OriginAreaCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.OriginAreaCode.AnyValue()))
	}
	if lastModel.StudentAreaCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.StudentAreaCode.AnyValue()))
	}
	if lastModel.Hobbies.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Hobbies.AnyValue()))
	}
	if lastModel.Creed.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Creed.AnyValue()))
	}
	if lastModel.TrainTicketinterval.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.TrainTicketinterval.AnyValue()))
	}
	if lastModel.FamilyAddress.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.FamilyAddress.AnyValue()))
	}
	if lastModel.DetailAddress.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.DetailAddress.AnyValue()))
	}
	if lastModel.PostCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.PostCode.AnyValue()))
	}
	if lastModel.HomePhone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.HomePhone.AnyValue()))
	}
	if lastModel.EnrollmentDate.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.EnrollmentDate.AnyValue()))
	}
	if lastModel.GraduationDate.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.GraduationDate.AnyValue()))
	}
	if lastModel.MidSchoolAddress.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.MidSchoolAddress.AnyValue()))
	}
	if lastModel.MidSchoolName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.MidSchoolName.AnyValue()))
	}
	if lastModel.Referee.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Referee.AnyValue()))
	}
	if lastModel.RefereeDuty.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RefereeDuty.AnyValue()))
	}
	if lastModel.RefereePhone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RefereePhone.AnyValue()))
	}
	if lastModel.AdmissionTicketNo.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AdmissionTicketNo.AnyValue()))
	}
	if lastModel.CollegeEntranceExamScores.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.CollegeEntranceExamScores.AnyValue()))
	}
	if lastModel.AdmissionYear.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AdmissionYear.AnyValue()))
	}
	if lastModel.ForeignLanguageCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ForeignLanguageCode.AnyValue()))
	}
	if lastModel.StudentOrigin.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.StudentOrigin.AnyValue()))
	}
	if lastModel.BizType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.BizType.AnyValue()))
	}
	if lastModel.TaskCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.TaskCode.AnyValue()))
	}
	if lastModel.ApproveStatus.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ApproveStatus.AnyValue()))
	}
	if lastModel.Operator.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Operator.AnyValue()))
	}
	if lastModel.InsertDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.InsertDatetime.AnyValue()))
	}
	if lastModel.UpdateDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateDatetime.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.StudentStatus.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.StudentStatus.AnyValue()))
	}
	if lastModel.IsAuth.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IsAuth.AnyValue()))
	}
	if lastModel.Campus.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Campus.AnyValue()))
	}
	if lastModel.Zone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Zone.AnyValue()))
	}
	if lastModel.Building.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Building.AnyValue()))
	}
	if lastModel.Unit.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Unit.AnyValue()))
	}
	if lastModel.Room.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Room.AnyValue()))
	}
	if lastModel.Bed.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Bed.AnyValue()))
	}
	if lastModel.StatusSort.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.StatusSort.AnyValue()))
	}
	if lastModel.Height.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Height.AnyValue()))
	}
	if lastModel.Weight.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Weight.AnyValue()))
	}
	if lastModel.FootSize.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.FootSize.AnyValue()))
	}
	if lastModel.ClothSize.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ClothSize.AnyValue()))
	}
	if lastModel.HeadSize.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.HeadSize.AnyValue()))
	}
	if lastModel.Remark1.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark1.AnyValue()))
	}
	if lastModel.Remark2.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark2.AnyValue()))
	}
	if lastModel.Remark3.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark3.AnyValue()))
	}
	if lastModel.Remark4.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark4.AnyValue()))
	}
	if lastModel.IsPayment.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IsPayment.AnyValue()))
	}
	if lastModel.IsCheckIn.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.IsCheckIn.AnyValue()))
	}
	if lastModel.GetMilitaryTC.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.GetMilitaryTC.AnyValue()))
	}
	if lastModel.OriginAreaName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.OriginAreaName.AnyValue()))
	}
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *StudentbasicinfoList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

func (m *Studentbasicinfo) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (l *StudentbasicinfoList) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

type StudentbasicinfoCacheElem struct {
	hashValue  string
	model      *Studentbasicinfo
	modifyTime time.Time
}

type StudentbasicinfoListCacheElem struct {
	hashValue  string
	list       *StudentbasicinfoList
	modifyTime time.Time
}

type StudentbasicinfoCacheManager struct {
	container map[string]*StudentbasicinfoCacheElem
	query     chan string
	in        chan *StudentbasicinfoCacheElem
	out       chan *StudentbasicinfoCacheElem
}

type StudentbasicinfoListCacheManager struct {
	container map[string]*StudentbasicinfoListCacheElem
	query     chan string
	in        chan *StudentbasicinfoListCacheElem
	out       chan *StudentbasicinfoListCacheElem
}

func newStudentbasicinfoCacheManager() *StudentbasicinfoCacheManager {
	return &StudentbasicinfoCacheManager{
		make(map[string]*StudentbasicinfoCacheElem),
		make(chan string),
		make(chan *StudentbasicinfoCacheElem),
		make(chan *StudentbasicinfoCacheElem),
	}
}

func newStudentbasicinfoListCacheManager() *StudentbasicinfoListCacheManager {
	return &StudentbasicinfoListCacheManager{
		make(map[string]*StudentbasicinfoListCacheElem),
		make(chan string),
		make(chan *StudentbasicinfoListCacheElem),
		make(chan *StudentbasicinfoListCacheElem),
	}
}

func (mgr *StudentbasicinfoCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *StudentbasicinfoListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var StudentbasicinfoCache = newStudentbasicinfoCacheManager()

var StudentbasicinfoListCache = newStudentbasicinfoListCacheManager()

func (m *Studentbasicinfo) GetCache(hashVal string, timeout time.Duration) bool {
	StudentbasicinfoCache.query <- hashVal
	elem := <-StudentbasicinfoCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *Studentbasicinfo) SetCache(hashValue string) {
	StudentbasicinfoCache.in <- &StudentbasicinfoCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *StudentbasicinfoList) GetListCache(hashValue string, timeout time.Duration) bool {
	StudentbasicinfoListCache.query <- hashValue
	elem := <-StudentbasicinfoListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *StudentbasicinfoList) SetListCache(hashValue string) {
	StudentbasicinfoListCache.in <- &StudentbasicinfoListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func (m *Studentbasicinfo) FromQuery(query interface{}) (*Studentbasicinfo, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				m.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					m.Id.AndWhereEq(fval0.Interface())
				case "!=":
					m.Id.AndWhereNeq(fval0.Interface())
				case ">":
					m.Id.AndWhereGt(fval0.Interface())
				case ">=":
					m.Id.AndWhereGte(fval0.Interface())
				case "<":
					m.Id.AndWhereLt(fval0.Interface())
				case "<=":
					m.Id.AndWhereLte(fval0.Interface())
				case "llike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					m.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					m.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("RecordId")
	if exists {
		fval1 := val.FieldByName("RecordId")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				m.RecordId.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					m.RecordId.AndWhereEq(fval1.Interface())
				case "!=":
					m.RecordId.AndWhereNeq(fval1.Interface())
				case ">":
					m.RecordId.AndWhereGt(fval1.Interface())
				case ">=":
					m.RecordId.AndWhereGte(fval1.Interface())
				case "<":
					m.RecordId.AndWhereLt(fval1.Interface())
				case "<=":
					m.RecordId.AndWhereLte(fval1.Interface())
				case "llike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					m.RecordId.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("IntelUserCode")
	if exists {
		fval2 := val.FieldByName("IntelUserCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				m.IntelUserCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					m.IntelUserCode.AndWhereEq(fval2.Interface())
				case "!=":
					m.IntelUserCode.AndWhereNeq(fval2.Interface())
				case ">":
					m.IntelUserCode.AndWhereGt(fval2.Interface())
				case ">=":
					m.IntelUserCode.AndWhereGte(fval2.Interface())
				case "<":
					m.IntelUserCode.AndWhereLt(fval2.Interface())
				case "<=":
					m.IntelUserCode.AndWhereLte(fval2.Interface())
				case "llike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					m.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					m.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					m.IntelUserCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("Class")
	if exists {
		fval3 := val.FieldByName("Class")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				m.Class.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					m.Class.AndWhereEq(fval3.Interface())
				case "!=":
					m.Class.AndWhereNeq(fval3.Interface())
				case ">":
					m.Class.AndWhereGt(fval3.Interface())
				case ">=":
					m.Class.AndWhereGte(fval3.Interface())
				case "<":
					m.Class.AndWhereLt(fval3.Interface())
				case "<=":
					m.Class.AndWhereLte(fval3.Interface())
				case "llike":
					m.Class.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					m.Class.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					m.Class.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					m.Class.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					m.Class.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					m.Class.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					m.Class.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("OtherName")
	if exists {
		fval4 := val.FieldByName("OtherName")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				m.OtherName.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					m.OtherName.AndWhereEq(fval4.Interface())
				case "!=":
					m.OtherName.AndWhereNeq(fval4.Interface())
				case ">":
					m.OtherName.AndWhereGt(fval4.Interface())
				case ">=":
					m.OtherName.AndWhereGte(fval4.Interface())
				case "<":
					m.OtherName.AndWhereLt(fval4.Interface())
				case "<=":
					m.OtherName.AndWhereLte(fval4.Interface())
				case "llike":
					m.OtherName.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					m.OtherName.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					m.OtherName.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					m.OtherName.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					m.OtherName.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					m.OtherName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					m.OtherName.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("NameInPinyin")
	if exists {
		fval5 := val.FieldByName("NameInPinyin")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				m.NameInPinyin.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					m.NameInPinyin.AndWhereEq(fval5.Interface())
				case "!=":
					m.NameInPinyin.AndWhereNeq(fval5.Interface())
				case ">":
					m.NameInPinyin.AndWhereGt(fval5.Interface())
				case ">=":
					m.NameInPinyin.AndWhereGte(fval5.Interface())
				case "<":
					m.NameInPinyin.AndWhereLt(fval5.Interface())
				case "<=":
					m.NameInPinyin.AndWhereLte(fval5.Interface())
				case "llike":
					m.NameInPinyin.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					m.NameInPinyin.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					m.NameInPinyin.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					m.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					m.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					m.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					m.NameInPinyin.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("EnglishName")
	if exists {
		fval6 := val.FieldByName("EnglishName")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				m.EnglishName.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					m.EnglishName.AndWhereEq(fval6.Interface())
				case "!=":
					m.EnglishName.AndWhereNeq(fval6.Interface())
				case ">":
					m.EnglishName.AndWhereGt(fval6.Interface())
				case ">=":
					m.EnglishName.AndWhereGte(fval6.Interface())
				case "<":
					m.EnglishName.AndWhereLt(fval6.Interface())
				case "<=":
					m.EnglishName.AndWhereLte(fval6.Interface())
				case "llike":
					m.EnglishName.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					m.EnglishName.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					m.EnglishName.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					m.EnglishName.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					m.EnglishName.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					m.EnglishName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					m.EnglishName.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("CountryCode")
	if exists {
		fval7 := val.FieldByName("CountryCode")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				m.CountryCode.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					m.CountryCode.AndWhereEq(fval7.Interface())
				case "!=":
					m.CountryCode.AndWhereNeq(fval7.Interface())
				case ">":
					m.CountryCode.AndWhereGt(fval7.Interface())
				case ">=":
					m.CountryCode.AndWhereGte(fval7.Interface())
				case "<":
					m.CountryCode.AndWhereLt(fval7.Interface())
				case "<=":
					m.CountryCode.AndWhereLte(fval7.Interface())
				case "llike":
					m.CountryCode.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					m.CountryCode.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					m.CountryCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					m.CountryCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					m.CountryCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					m.CountryCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					m.CountryCode.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("NationalityCode")
	if exists {
		fval8 := val.FieldByName("NationalityCode")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				m.NationalityCode.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					m.NationalityCode.AndWhereEq(fval8.Interface())
				case "!=":
					m.NationalityCode.AndWhereNeq(fval8.Interface())
				case ">":
					m.NationalityCode.AndWhereGt(fval8.Interface())
				case ">=":
					m.NationalityCode.AndWhereGte(fval8.Interface())
				case "<":
					m.NationalityCode.AndWhereLt(fval8.Interface())
				case "<=":
					m.NationalityCode.AndWhereLte(fval8.Interface())
				case "llike":
					m.NationalityCode.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					m.NationalityCode.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					m.NationalityCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					m.NationalityCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					m.NationalityCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					m.NationalityCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					m.NationalityCode.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("Birthday")
	if exists {
		fval9 := val.FieldByName("Birthday")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				m.Birthday.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					m.Birthday.AndWhereEq(fval9.Interface())
				case "!=":
					m.Birthday.AndWhereNeq(fval9.Interface())
				case ">":
					m.Birthday.AndWhereGt(fval9.Interface())
				case ">=":
					m.Birthday.AndWhereGte(fval9.Interface())
				case "<":
					m.Birthday.AndWhereLt(fval9.Interface())
				case "<=":
					m.Birthday.AndWhereLte(fval9.Interface())
				case "llike":
					m.Birthday.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					m.Birthday.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					m.Birthday.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					m.Birthday.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					m.Birthday.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					m.Birthday.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					m.Birthday.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("PoliticalCode")
	if exists {
		fval10 := val.FieldByName("PoliticalCode")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				m.PoliticalCode.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					m.PoliticalCode.AndWhereEq(fval10.Interface())
				case "!=":
					m.PoliticalCode.AndWhereNeq(fval10.Interface())
				case ">":
					m.PoliticalCode.AndWhereGt(fval10.Interface())
				case ">=":
					m.PoliticalCode.AndWhereGte(fval10.Interface())
				case "<":
					m.PoliticalCode.AndWhereLt(fval10.Interface())
				case "<=":
					m.PoliticalCode.AndWhereLte(fval10.Interface())
				case "llike":
					m.PoliticalCode.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					m.PoliticalCode.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					m.PoliticalCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					m.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					m.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					m.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					m.PoliticalCode.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("QQAcct")
	if exists {
		fval11 := val.FieldByName("QQAcct")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				m.QQAcct.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					m.QQAcct.AndWhereEq(fval11.Interface())
				case "!=":
					m.QQAcct.AndWhereNeq(fval11.Interface())
				case ">":
					m.QQAcct.AndWhereGt(fval11.Interface())
				case ">=":
					m.QQAcct.AndWhereGte(fval11.Interface())
				case "<":
					m.QQAcct.AndWhereLt(fval11.Interface())
				case "<=":
					m.QQAcct.AndWhereLte(fval11.Interface())
				case "llike":
					m.QQAcct.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					m.QQAcct.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					m.QQAcct.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					m.QQAcct.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					m.QQAcct.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					m.QQAcct.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					m.QQAcct.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("WeChatAcct")
	if exists {
		fval12 := val.FieldByName("WeChatAcct")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				m.WeChatAcct.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					m.WeChatAcct.AndWhereEq(fval12.Interface())
				case "!=":
					m.WeChatAcct.AndWhereNeq(fval12.Interface())
				case ">":
					m.WeChatAcct.AndWhereGt(fval12.Interface())
				case ">=":
					m.WeChatAcct.AndWhereGte(fval12.Interface())
				case "<":
					m.WeChatAcct.AndWhereLt(fval12.Interface())
				case "<=":
					m.WeChatAcct.AndWhereLte(fval12.Interface())
				case "llike":
					m.WeChatAcct.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					m.WeChatAcct.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					m.WeChatAcct.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					m.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					m.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					m.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					m.WeChatAcct.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("BankCardNumber")
	if exists {
		fval13 := val.FieldByName("BankCardNumber")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				m.BankCardNumber.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					m.BankCardNumber.AndWhereEq(fval13.Interface())
				case "!=":
					m.BankCardNumber.AndWhereNeq(fval13.Interface())
				case ">":
					m.BankCardNumber.AndWhereGt(fval13.Interface())
				case ">=":
					m.BankCardNumber.AndWhereGte(fval13.Interface())
				case "<":
					m.BankCardNumber.AndWhereLt(fval13.Interface())
				case "<=":
					m.BankCardNumber.AndWhereLte(fval13.Interface())
				case "llike":
					m.BankCardNumber.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					m.BankCardNumber.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					m.BankCardNumber.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					m.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					m.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					m.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					m.BankCardNumber.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("AccountBankCode")
	if exists {
		fval14 := val.FieldByName("AccountBankCode")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				m.AccountBankCode.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					m.AccountBankCode.AndWhereEq(fval14.Interface())
				case "!=":
					m.AccountBankCode.AndWhereNeq(fval14.Interface())
				case ">":
					m.AccountBankCode.AndWhereGt(fval14.Interface())
				case ">=":
					m.AccountBankCode.AndWhereGte(fval14.Interface())
				case "<":
					m.AccountBankCode.AndWhereLt(fval14.Interface())
				case "<=":
					m.AccountBankCode.AndWhereLte(fval14.Interface())
				case "llike":
					m.AccountBankCode.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					m.AccountBankCode.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					m.AccountBankCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					m.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					m.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					m.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					m.AccountBankCode.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("AllPowerfulCardNum")
	if exists {
		fval15 := val.FieldByName("AllPowerfulCardNum")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				m.AllPowerfulCardNum.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					m.AllPowerfulCardNum.AndWhereEq(fval15.Interface())
				case "!=":
					m.AllPowerfulCardNum.AndWhereNeq(fval15.Interface())
				case ">":
					m.AllPowerfulCardNum.AndWhereGt(fval15.Interface())
				case ">=":
					m.AllPowerfulCardNum.AndWhereGte(fval15.Interface())
				case "<":
					m.AllPowerfulCardNum.AndWhereLt(fval15.Interface())
				case "<=":
					m.AllPowerfulCardNum.AndWhereLte(fval15.Interface())
				case "llike":
					m.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					m.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					m.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					m.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					m.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					m.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					m.AllPowerfulCardNum.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("MaritalCode")
	if exists {
		fval16 := val.FieldByName("MaritalCode")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				m.MaritalCode.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					m.MaritalCode.AndWhereEq(fval16.Interface())
				case "!=":
					m.MaritalCode.AndWhereNeq(fval16.Interface())
				case ">":
					m.MaritalCode.AndWhereGt(fval16.Interface())
				case ">=":
					m.MaritalCode.AndWhereGte(fval16.Interface())
				case "<":
					m.MaritalCode.AndWhereLt(fval16.Interface())
				case "<=":
					m.MaritalCode.AndWhereLte(fval16.Interface())
				case "llike":
					m.MaritalCode.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					m.MaritalCode.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					m.MaritalCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					m.MaritalCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					m.MaritalCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					m.MaritalCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					m.MaritalCode.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("OriginAreaCode")
	if exists {
		fval17 := val.FieldByName("OriginAreaCode")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				m.OriginAreaCode.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					m.OriginAreaCode.AndWhereEq(fval17.Interface())
				case "!=":
					m.OriginAreaCode.AndWhereNeq(fval17.Interface())
				case ">":
					m.OriginAreaCode.AndWhereGt(fval17.Interface())
				case ">=":
					m.OriginAreaCode.AndWhereGte(fval17.Interface())
				case "<":
					m.OriginAreaCode.AndWhereLt(fval17.Interface())
				case "<=":
					m.OriginAreaCode.AndWhereLte(fval17.Interface())
				case "llike":
					m.OriginAreaCode.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					m.OriginAreaCode.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					m.OriginAreaCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					m.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					m.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					m.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					m.OriginAreaCode.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("StudentAreaCode")
	if exists {
		fval18 := val.FieldByName("StudentAreaCode")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				m.StudentAreaCode.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					m.StudentAreaCode.AndWhereEq(fval18.Interface())
				case "!=":
					m.StudentAreaCode.AndWhereNeq(fval18.Interface())
				case ">":
					m.StudentAreaCode.AndWhereGt(fval18.Interface())
				case ">=":
					m.StudentAreaCode.AndWhereGte(fval18.Interface())
				case "<":
					m.StudentAreaCode.AndWhereLt(fval18.Interface())
				case "<=":
					m.StudentAreaCode.AndWhereLte(fval18.Interface())
				case "llike":
					m.StudentAreaCode.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					m.StudentAreaCode.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					m.StudentAreaCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					m.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					m.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					m.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					m.StudentAreaCode.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Hobbies")
	if exists {
		fval19 := val.FieldByName("Hobbies")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				m.Hobbies.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					m.Hobbies.AndWhereEq(fval19.Interface())
				case "!=":
					m.Hobbies.AndWhereNeq(fval19.Interface())
				case ">":
					m.Hobbies.AndWhereGt(fval19.Interface())
				case ">=":
					m.Hobbies.AndWhereGte(fval19.Interface())
				case "<":
					m.Hobbies.AndWhereLt(fval19.Interface())
				case "<=":
					m.Hobbies.AndWhereLte(fval19.Interface())
				case "llike":
					m.Hobbies.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					m.Hobbies.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					m.Hobbies.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					m.Hobbies.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					m.Hobbies.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					m.Hobbies.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					m.Hobbies.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("Creed")
	if exists {
		fval20 := val.FieldByName("Creed")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				m.Creed.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					m.Creed.AndWhereEq(fval20.Interface())
				case "!=":
					m.Creed.AndWhereNeq(fval20.Interface())
				case ">":
					m.Creed.AndWhereGt(fval20.Interface())
				case ">=":
					m.Creed.AndWhereGte(fval20.Interface())
				case "<":
					m.Creed.AndWhereLt(fval20.Interface())
				case "<=":
					m.Creed.AndWhereLte(fval20.Interface())
				case "llike":
					m.Creed.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					m.Creed.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					m.Creed.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					m.Creed.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					m.Creed.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					m.Creed.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					m.Creed.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("TrainTicketinterval")
	if exists {
		fval21 := val.FieldByName("TrainTicketinterval")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				m.TrainTicketinterval.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					m.TrainTicketinterval.AndWhereEq(fval21.Interface())
				case "!=":
					m.TrainTicketinterval.AndWhereNeq(fval21.Interface())
				case ">":
					m.TrainTicketinterval.AndWhereGt(fval21.Interface())
				case ">=":
					m.TrainTicketinterval.AndWhereGte(fval21.Interface())
				case "<":
					m.TrainTicketinterval.AndWhereLt(fval21.Interface())
				case "<=":
					m.TrainTicketinterval.AndWhereLte(fval21.Interface())
				case "llike":
					m.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					m.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					m.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					m.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					m.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					m.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					m.TrainTicketinterval.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("FamilyAddress")
	if exists {
		fval22 := val.FieldByName("FamilyAddress")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				m.FamilyAddress.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					m.FamilyAddress.AndWhereEq(fval22.Interface())
				case "!=":
					m.FamilyAddress.AndWhereNeq(fval22.Interface())
				case ">":
					m.FamilyAddress.AndWhereGt(fval22.Interface())
				case ">=":
					m.FamilyAddress.AndWhereGte(fval22.Interface())
				case "<":
					m.FamilyAddress.AndWhereLt(fval22.Interface())
				case "<=":
					m.FamilyAddress.AndWhereLte(fval22.Interface())
				case "llike":
					m.FamilyAddress.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					m.FamilyAddress.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					m.FamilyAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					m.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					m.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					m.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					m.FamilyAddress.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("DetailAddress")
	if exists {
		fval23 := val.FieldByName("DetailAddress")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				m.DetailAddress.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					m.DetailAddress.AndWhereEq(fval23.Interface())
				case "!=":
					m.DetailAddress.AndWhereNeq(fval23.Interface())
				case ">":
					m.DetailAddress.AndWhereGt(fval23.Interface())
				case ">=":
					m.DetailAddress.AndWhereGte(fval23.Interface())
				case "<":
					m.DetailAddress.AndWhereLt(fval23.Interface())
				case "<=":
					m.DetailAddress.AndWhereLte(fval23.Interface())
				case "llike":
					m.DetailAddress.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					m.DetailAddress.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					m.DetailAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					m.DetailAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					m.DetailAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					m.DetailAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					m.DetailAddress.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("PostCode")
	if exists {
		fval24 := val.FieldByName("PostCode")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				m.PostCode.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					m.PostCode.AndWhereEq(fval24.Interface())
				case "!=":
					m.PostCode.AndWhereNeq(fval24.Interface())
				case ">":
					m.PostCode.AndWhereGt(fval24.Interface())
				case ">=":
					m.PostCode.AndWhereGte(fval24.Interface())
				case "<":
					m.PostCode.AndWhereLt(fval24.Interface())
				case "<=":
					m.PostCode.AndWhereLte(fval24.Interface())
				case "llike":
					m.PostCode.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					m.PostCode.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					m.PostCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					m.PostCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					m.PostCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					m.PostCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					m.PostCode.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("HomePhone")
	if exists {
		fval25 := val.FieldByName("HomePhone")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				m.HomePhone.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					m.HomePhone.AndWhereEq(fval25.Interface())
				case "!=":
					m.HomePhone.AndWhereNeq(fval25.Interface())
				case ">":
					m.HomePhone.AndWhereGt(fval25.Interface())
				case ">=":
					m.HomePhone.AndWhereGte(fval25.Interface())
				case "<":
					m.HomePhone.AndWhereLt(fval25.Interface())
				case "<=":
					m.HomePhone.AndWhereLte(fval25.Interface())
				case "llike":
					m.HomePhone.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					m.HomePhone.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					m.HomePhone.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					m.HomePhone.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					m.HomePhone.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					m.HomePhone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					m.HomePhone.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("EnrollmentDate")
	if exists {
		fval26 := val.FieldByName("EnrollmentDate")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				m.EnrollmentDate.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					m.EnrollmentDate.AndWhereEq(fval26.Interface())
				case "!=":
					m.EnrollmentDate.AndWhereNeq(fval26.Interface())
				case ">":
					m.EnrollmentDate.AndWhereGt(fval26.Interface())
				case ">=":
					m.EnrollmentDate.AndWhereGte(fval26.Interface())
				case "<":
					m.EnrollmentDate.AndWhereLt(fval26.Interface())
				case "<=":
					m.EnrollmentDate.AndWhereLte(fval26.Interface())
				case "llike":
					m.EnrollmentDate.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					m.EnrollmentDate.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					m.EnrollmentDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					m.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					m.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					m.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					m.EnrollmentDate.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("GraduationDate")
	if exists {
		fval27 := val.FieldByName("GraduationDate")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				m.GraduationDate.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					m.GraduationDate.AndWhereEq(fval27.Interface())
				case "!=":
					m.GraduationDate.AndWhereNeq(fval27.Interface())
				case ">":
					m.GraduationDate.AndWhereGt(fval27.Interface())
				case ">=":
					m.GraduationDate.AndWhereGte(fval27.Interface())
				case "<":
					m.GraduationDate.AndWhereLt(fval27.Interface())
				case "<=":
					m.GraduationDate.AndWhereLte(fval27.Interface())
				case "llike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					m.GraduationDate.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	ftyp28, exists := typ.FieldByName("MidSchoolAddress")
	if exists {
		fval28 := val.FieldByName("MidSchoolAddress")
		fop28, ok := ftyp28.Tag.Lookup("op")
		for fval28.Kind() == reflect.Ptr && !fval28.IsNil() {
			fval28 = fval28.Elem()
		}
		if fval28.Kind() != reflect.Ptr {
			if !ok {
				m.MidSchoolAddress.AndWhereEq(fval28.Interface())
			} else {
				switch fop28 {
				case "=":
					m.MidSchoolAddress.AndWhereEq(fval28.Interface())
				case "!=":
					m.MidSchoolAddress.AndWhereNeq(fval28.Interface())
				case ">":
					m.MidSchoolAddress.AndWhereGt(fval28.Interface())
				case ">=":
					m.MidSchoolAddress.AndWhereGte(fval28.Interface())
				case "<":
					m.MidSchoolAddress.AndWhereLt(fval28.Interface())
				case "<=":
					m.MidSchoolAddress.AndWhereLte(fval28.Interface())
				case "llike":
					m.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%%%s", fval28.String()))
				case "rlike":
					m.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%s%%", fval28.String()))
				case "alike":
					m.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval28.String()))
				case "nllike":
					m.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval28.String()))
				case "nrlike":
					m.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval28.String()))
				case "nalike":
					m.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval28.String()))
				case "in":
					m.MidSchoolAddress.AndWhereIn(fval28.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop28)
				}
			}
		}
	}
	ftyp29, exists := typ.FieldByName("MidSchoolName")
	if exists {
		fval29 := val.FieldByName("MidSchoolName")
		fop29, ok := ftyp29.Tag.Lookup("op")
		for fval29.Kind() == reflect.Ptr && !fval29.IsNil() {
			fval29 = fval29.Elem()
		}
		if fval29.Kind() != reflect.Ptr {
			if !ok {
				m.MidSchoolName.AndWhereEq(fval29.Interface())
			} else {
				switch fop29 {
				case "=":
					m.MidSchoolName.AndWhereEq(fval29.Interface())
				case "!=":
					m.MidSchoolName.AndWhereNeq(fval29.Interface())
				case ">":
					m.MidSchoolName.AndWhereGt(fval29.Interface())
				case ">=":
					m.MidSchoolName.AndWhereGte(fval29.Interface())
				case "<":
					m.MidSchoolName.AndWhereLt(fval29.Interface())
				case "<=":
					m.MidSchoolName.AndWhereLte(fval29.Interface())
				case "llike":
					m.MidSchoolName.AndWhereLike(fmt.Sprintf("%%%s", fval29.String()))
				case "rlike":
					m.MidSchoolName.AndWhereLike(fmt.Sprintf("%s%%", fval29.String()))
				case "alike":
					m.MidSchoolName.AndWhereLike(fmt.Sprintf("%%%s%%", fval29.String()))
				case "nllike":
					m.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%%%s", fval29.String()))
				case "nrlike":
					m.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%s%%", fval29.String()))
				case "nalike":
					m.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval29.String()))
				case "in":
					m.MidSchoolName.AndWhereIn(fval29.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop29)
				}
			}
		}
	}
	ftyp30, exists := typ.FieldByName("Referee")
	if exists {
		fval30 := val.FieldByName("Referee")
		fop30, ok := ftyp30.Tag.Lookup("op")
		for fval30.Kind() == reflect.Ptr && !fval30.IsNil() {
			fval30 = fval30.Elem()
		}
		if fval30.Kind() != reflect.Ptr {
			if !ok {
				m.Referee.AndWhereEq(fval30.Interface())
			} else {
				switch fop30 {
				case "=":
					m.Referee.AndWhereEq(fval30.Interface())
				case "!=":
					m.Referee.AndWhereNeq(fval30.Interface())
				case ">":
					m.Referee.AndWhereGt(fval30.Interface())
				case ">=":
					m.Referee.AndWhereGte(fval30.Interface())
				case "<":
					m.Referee.AndWhereLt(fval30.Interface())
				case "<=":
					m.Referee.AndWhereLte(fval30.Interface())
				case "llike":
					m.Referee.AndWhereLike(fmt.Sprintf("%%%s", fval30.String()))
				case "rlike":
					m.Referee.AndWhereLike(fmt.Sprintf("%s%%", fval30.String()))
				case "alike":
					m.Referee.AndWhereLike(fmt.Sprintf("%%%s%%", fval30.String()))
				case "nllike":
					m.Referee.AndWhereNotLike(fmt.Sprintf("%%%s", fval30.String()))
				case "nrlike":
					m.Referee.AndWhereNotLike(fmt.Sprintf("%s%%", fval30.String()))
				case "nalike":
					m.Referee.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval30.String()))
				case "in":
					m.Referee.AndWhereIn(fval30.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop30)
				}
			}
		}
	}
	ftyp31, exists := typ.FieldByName("RefereeDuty")
	if exists {
		fval31 := val.FieldByName("RefereeDuty")
		fop31, ok := ftyp31.Tag.Lookup("op")
		for fval31.Kind() == reflect.Ptr && !fval31.IsNil() {
			fval31 = fval31.Elem()
		}
		if fval31.Kind() != reflect.Ptr {
			if !ok {
				m.RefereeDuty.AndWhereEq(fval31.Interface())
			} else {
				switch fop31 {
				case "=":
					m.RefereeDuty.AndWhereEq(fval31.Interface())
				case "!=":
					m.RefereeDuty.AndWhereNeq(fval31.Interface())
				case ">":
					m.RefereeDuty.AndWhereGt(fval31.Interface())
				case ">=":
					m.RefereeDuty.AndWhereGte(fval31.Interface())
				case "<":
					m.RefereeDuty.AndWhereLt(fval31.Interface())
				case "<=":
					m.RefereeDuty.AndWhereLte(fval31.Interface())
				case "llike":
					m.RefereeDuty.AndWhereLike(fmt.Sprintf("%%%s", fval31.String()))
				case "rlike":
					m.RefereeDuty.AndWhereLike(fmt.Sprintf("%s%%", fval31.String()))
				case "alike":
					m.RefereeDuty.AndWhereLike(fmt.Sprintf("%%%s%%", fval31.String()))
				case "nllike":
					m.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%%%s", fval31.String()))
				case "nrlike":
					m.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%s%%", fval31.String()))
				case "nalike":
					m.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval31.String()))
				case "in":
					m.RefereeDuty.AndWhereIn(fval31.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop31)
				}
			}
		}
	}
	ftyp32, exists := typ.FieldByName("RefereePhone")
	if exists {
		fval32 := val.FieldByName("RefereePhone")
		fop32, ok := ftyp32.Tag.Lookup("op")
		for fval32.Kind() == reflect.Ptr && !fval32.IsNil() {
			fval32 = fval32.Elem()
		}
		if fval32.Kind() != reflect.Ptr {
			if !ok {
				m.RefereePhone.AndWhereEq(fval32.Interface())
			} else {
				switch fop32 {
				case "=":
					m.RefereePhone.AndWhereEq(fval32.Interface())
				case "!=":
					m.RefereePhone.AndWhereNeq(fval32.Interface())
				case ">":
					m.RefereePhone.AndWhereGt(fval32.Interface())
				case ">=":
					m.RefereePhone.AndWhereGte(fval32.Interface())
				case "<":
					m.RefereePhone.AndWhereLt(fval32.Interface())
				case "<=":
					m.RefereePhone.AndWhereLte(fval32.Interface())
				case "llike":
					m.RefereePhone.AndWhereLike(fmt.Sprintf("%%%s", fval32.String()))
				case "rlike":
					m.RefereePhone.AndWhereLike(fmt.Sprintf("%s%%", fval32.String()))
				case "alike":
					m.RefereePhone.AndWhereLike(fmt.Sprintf("%%%s%%", fval32.String()))
				case "nllike":
					m.RefereePhone.AndWhereNotLike(fmt.Sprintf("%%%s", fval32.String()))
				case "nrlike":
					m.RefereePhone.AndWhereNotLike(fmt.Sprintf("%s%%", fval32.String()))
				case "nalike":
					m.RefereePhone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval32.String()))
				case "in":
					m.RefereePhone.AndWhereIn(fval32.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop32)
				}
			}
		}
	}
	ftyp33, exists := typ.FieldByName("AdmissionTicketNo")
	if exists {
		fval33 := val.FieldByName("AdmissionTicketNo")
		fop33, ok := ftyp33.Tag.Lookup("op")
		for fval33.Kind() == reflect.Ptr && !fval33.IsNil() {
			fval33 = fval33.Elem()
		}
		if fval33.Kind() != reflect.Ptr {
			if !ok {
				m.AdmissionTicketNo.AndWhereEq(fval33.Interface())
			} else {
				switch fop33 {
				case "=":
					m.AdmissionTicketNo.AndWhereEq(fval33.Interface())
				case "!=":
					m.AdmissionTicketNo.AndWhereNeq(fval33.Interface())
				case ">":
					m.AdmissionTicketNo.AndWhereGt(fval33.Interface())
				case ">=":
					m.AdmissionTicketNo.AndWhereGte(fval33.Interface())
				case "<":
					m.AdmissionTicketNo.AndWhereLt(fval33.Interface())
				case "<=":
					m.AdmissionTicketNo.AndWhereLte(fval33.Interface())
				case "llike":
					m.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%%%s", fval33.String()))
				case "rlike":
					m.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%s%%", fval33.String()))
				case "alike":
					m.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%%%s%%", fval33.String()))
				case "nllike":
					m.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%%%s", fval33.String()))
				case "nrlike":
					m.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%s%%", fval33.String()))
				case "nalike":
					m.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval33.String()))
				case "in":
					m.AdmissionTicketNo.AndWhereIn(fval33.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop33)
				}
			}
		}
	}
	ftyp34, exists := typ.FieldByName("CollegeEntranceExamScores")
	if exists {
		fval34 := val.FieldByName("CollegeEntranceExamScores")
		fop34, ok := ftyp34.Tag.Lookup("op")
		for fval34.Kind() == reflect.Ptr && !fval34.IsNil() {
			fval34 = fval34.Elem()
		}
		if fval34.Kind() != reflect.Ptr {
			if !ok {
				m.CollegeEntranceExamScores.AndWhereEq(fval34.Interface())
			} else {
				switch fop34 {
				case "=":
					m.CollegeEntranceExamScores.AndWhereEq(fval34.Interface())
				case "!=":
					m.CollegeEntranceExamScores.AndWhereNeq(fval34.Interface())
				case ">":
					m.CollegeEntranceExamScores.AndWhereGt(fval34.Interface())
				case ">=":
					m.CollegeEntranceExamScores.AndWhereGte(fval34.Interface())
				case "<":
					m.CollegeEntranceExamScores.AndWhereLt(fval34.Interface())
				case "<=":
					m.CollegeEntranceExamScores.AndWhereLte(fval34.Interface())
				case "llike":
					m.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%%%s", fval34.String()))
				case "rlike":
					m.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%s%%", fval34.String()))
				case "alike":
					m.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%%%s%%", fval34.String()))
				case "nllike":
					m.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%%%s", fval34.String()))
				case "nrlike":
					m.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%s%%", fval34.String()))
				case "nalike":
					m.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval34.String()))
				case "in":
					m.CollegeEntranceExamScores.AndWhereIn(fval34.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop34)
				}
			}
		}
	}
	ftyp35, exists := typ.FieldByName("AdmissionYear")
	if exists {
		fval35 := val.FieldByName("AdmissionYear")
		fop35, ok := ftyp35.Tag.Lookup("op")
		for fval35.Kind() == reflect.Ptr && !fval35.IsNil() {
			fval35 = fval35.Elem()
		}
		if fval35.Kind() != reflect.Ptr {
			if !ok {
				m.AdmissionYear.AndWhereEq(fval35.Interface())
			} else {
				switch fop35 {
				case "=":
					m.AdmissionYear.AndWhereEq(fval35.Interface())
				case "!=":
					m.AdmissionYear.AndWhereNeq(fval35.Interface())
				case ">":
					m.AdmissionYear.AndWhereGt(fval35.Interface())
				case ">=":
					m.AdmissionYear.AndWhereGte(fval35.Interface())
				case "<":
					m.AdmissionYear.AndWhereLt(fval35.Interface())
				case "<=":
					m.AdmissionYear.AndWhereLte(fval35.Interface())
				case "llike":
					m.AdmissionYear.AndWhereLike(fmt.Sprintf("%%%s", fval35.String()))
				case "rlike":
					m.AdmissionYear.AndWhereLike(fmt.Sprintf("%s%%", fval35.String()))
				case "alike":
					m.AdmissionYear.AndWhereLike(fmt.Sprintf("%%%s%%", fval35.String()))
				case "nllike":
					m.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%%%s", fval35.String()))
				case "nrlike":
					m.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%s%%", fval35.String()))
				case "nalike":
					m.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval35.String()))
				case "in":
					m.AdmissionYear.AndWhereIn(fval35.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop35)
				}
			}
		}
	}
	ftyp36, exists := typ.FieldByName("ForeignLanguageCode")
	if exists {
		fval36 := val.FieldByName("ForeignLanguageCode")
		fop36, ok := ftyp36.Tag.Lookup("op")
		for fval36.Kind() == reflect.Ptr && !fval36.IsNil() {
			fval36 = fval36.Elem()
		}
		if fval36.Kind() != reflect.Ptr {
			if !ok {
				m.ForeignLanguageCode.AndWhereEq(fval36.Interface())
			} else {
				switch fop36 {
				case "=":
					m.ForeignLanguageCode.AndWhereEq(fval36.Interface())
				case "!=":
					m.ForeignLanguageCode.AndWhereNeq(fval36.Interface())
				case ">":
					m.ForeignLanguageCode.AndWhereGt(fval36.Interface())
				case ">=":
					m.ForeignLanguageCode.AndWhereGte(fval36.Interface())
				case "<":
					m.ForeignLanguageCode.AndWhereLt(fval36.Interface())
				case "<=":
					m.ForeignLanguageCode.AndWhereLte(fval36.Interface())
				case "llike":
					m.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%%%s", fval36.String()))
				case "rlike":
					m.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%s%%", fval36.String()))
				case "alike":
					m.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval36.String()))
				case "nllike":
					m.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval36.String()))
				case "nrlike":
					m.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval36.String()))
				case "nalike":
					m.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval36.String()))
				case "in":
					m.ForeignLanguageCode.AndWhereIn(fval36.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop36)
				}
			}
		}
	}
	ftyp37, exists := typ.FieldByName("StudentOrigin")
	if exists {
		fval37 := val.FieldByName("StudentOrigin")
		fop37, ok := ftyp37.Tag.Lookup("op")
		for fval37.Kind() == reflect.Ptr && !fval37.IsNil() {
			fval37 = fval37.Elem()
		}
		if fval37.Kind() != reflect.Ptr {
			if !ok {
				m.StudentOrigin.AndWhereEq(fval37.Interface())
			} else {
				switch fop37 {
				case "=":
					m.StudentOrigin.AndWhereEq(fval37.Interface())
				case "!=":
					m.StudentOrigin.AndWhereNeq(fval37.Interface())
				case ">":
					m.StudentOrigin.AndWhereGt(fval37.Interface())
				case ">=":
					m.StudentOrigin.AndWhereGte(fval37.Interface())
				case "<":
					m.StudentOrigin.AndWhereLt(fval37.Interface())
				case "<=":
					m.StudentOrigin.AndWhereLte(fval37.Interface())
				case "llike":
					m.StudentOrigin.AndWhereLike(fmt.Sprintf("%%%s", fval37.String()))
				case "rlike":
					m.StudentOrigin.AndWhereLike(fmt.Sprintf("%s%%", fval37.String()))
				case "alike":
					m.StudentOrigin.AndWhereLike(fmt.Sprintf("%%%s%%", fval37.String()))
				case "nllike":
					m.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%%%s", fval37.String()))
				case "nrlike":
					m.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%s%%", fval37.String()))
				case "nalike":
					m.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval37.String()))
				case "in":
					m.StudentOrigin.AndWhereIn(fval37.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop37)
				}
			}
		}
	}
	ftyp38, exists := typ.FieldByName("BizType")
	if exists {
		fval38 := val.FieldByName("BizType")
		fop38, ok := ftyp38.Tag.Lookup("op")
		for fval38.Kind() == reflect.Ptr && !fval38.IsNil() {
			fval38 = fval38.Elem()
		}
		if fval38.Kind() != reflect.Ptr {
			if !ok {
				m.BizType.AndWhereEq(fval38.Interface())
			} else {
				switch fop38 {
				case "=":
					m.BizType.AndWhereEq(fval38.Interface())
				case "!=":
					m.BizType.AndWhereNeq(fval38.Interface())
				case ">":
					m.BizType.AndWhereGt(fval38.Interface())
				case ">=":
					m.BizType.AndWhereGte(fval38.Interface())
				case "<":
					m.BizType.AndWhereLt(fval38.Interface())
				case "<=":
					m.BizType.AndWhereLte(fval38.Interface())
				case "llike":
					m.BizType.AndWhereLike(fmt.Sprintf("%%%s", fval38.String()))
				case "rlike":
					m.BizType.AndWhereLike(fmt.Sprintf("%s%%", fval38.String()))
				case "alike":
					m.BizType.AndWhereLike(fmt.Sprintf("%%%s%%", fval38.String()))
				case "nllike":
					m.BizType.AndWhereNotLike(fmt.Sprintf("%%%s", fval38.String()))
				case "nrlike":
					m.BizType.AndWhereNotLike(fmt.Sprintf("%s%%", fval38.String()))
				case "nalike":
					m.BizType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval38.String()))
				case "in":
					m.BizType.AndWhereIn(fval38.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop38)
				}
			}
		}
	}
	ftyp39, exists := typ.FieldByName("TaskCode")
	if exists {
		fval39 := val.FieldByName("TaskCode")
		fop39, ok := ftyp39.Tag.Lookup("op")
		for fval39.Kind() == reflect.Ptr && !fval39.IsNil() {
			fval39 = fval39.Elem()
		}
		if fval39.Kind() != reflect.Ptr {
			if !ok {
				m.TaskCode.AndWhereEq(fval39.Interface())
			} else {
				switch fop39 {
				case "=":
					m.TaskCode.AndWhereEq(fval39.Interface())
				case "!=":
					m.TaskCode.AndWhereNeq(fval39.Interface())
				case ">":
					m.TaskCode.AndWhereGt(fval39.Interface())
				case ">=":
					m.TaskCode.AndWhereGte(fval39.Interface())
				case "<":
					m.TaskCode.AndWhereLt(fval39.Interface())
				case "<=":
					m.TaskCode.AndWhereLte(fval39.Interface())
				case "llike":
					m.TaskCode.AndWhereLike(fmt.Sprintf("%%%s", fval39.String()))
				case "rlike":
					m.TaskCode.AndWhereLike(fmt.Sprintf("%s%%", fval39.String()))
				case "alike":
					m.TaskCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval39.String()))
				case "nllike":
					m.TaskCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval39.String()))
				case "nrlike":
					m.TaskCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval39.String()))
				case "nalike":
					m.TaskCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval39.String()))
				case "in":
					m.TaskCode.AndWhereIn(fval39.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop39)
				}
			}
		}
	}
	ftyp40, exists := typ.FieldByName("ApproveStatus")
	if exists {
		fval40 := val.FieldByName("ApproveStatus")
		fop40, ok := ftyp40.Tag.Lookup("op")
		for fval40.Kind() == reflect.Ptr && !fval40.IsNil() {
			fval40 = fval40.Elem()
		}
		if fval40.Kind() != reflect.Ptr {
			if !ok {
				m.ApproveStatus.AndWhereEq(fval40.Interface())
			} else {
				switch fop40 {
				case "=":
					m.ApproveStatus.AndWhereEq(fval40.Interface())
				case "!=":
					m.ApproveStatus.AndWhereNeq(fval40.Interface())
				case ">":
					m.ApproveStatus.AndWhereGt(fval40.Interface())
				case ">=":
					m.ApproveStatus.AndWhereGte(fval40.Interface())
				case "<":
					m.ApproveStatus.AndWhereLt(fval40.Interface())
				case "<=":
					m.ApproveStatus.AndWhereLte(fval40.Interface())
				case "llike":
					m.ApproveStatus.AndWhereLike(fmt.Sprintf("%%%s", fval40.String()))
				case "rlike":
					m.ApproveStatus.AndWhereLike(fmt.Sprintf("%s%%", fval40.String()))
				case "alike":
					m.ApproveStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval40.String()))
				case "nllike":
					m.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval40.String()))
				case "nrlike":
					m.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval40.String()))
				case "nalike":
					m.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval40.String()))
				case "in":
					m.ApproveStatus.AndWhereIn(fval40.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop40)
				}
			}
		}
	}
	ftyp41, exists := typ.FieldByName("Operator")
	if exists {
		fval41 := val.FieldByName("Operator")
		fop41, ok := ftyp41.Tag.Lookup("op")
		for fval41.Kind() == reflect.Ptr && !fval41.IsNil() {
			fval41 = fval41.Elem()
		}
		if fval41.Kind() != reflect.Ptr {
			if !ok {
				m.Operator.AndWhereEq(fval41.Interface())
			} else {
				switch fop41 {
				case "=":
					m.Operator.AndWhereEq(fval41.Interface())
				case "!=":
					m.Operator.AndWhereNeq(fval41.Interface())
				case ">":
					m.Operator.AndWhereGt(fval41.Interface())
				case ">=":
					m.Operator.AndWhereGte(fval41.Interface())
				case "<":
					m.Operator.AndWhereLt(fval41.Interface())
				case "<=":
					m.Operator.AndWhereLte(fval41.Interface())
				case "llike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval41.String()))
				case "rlike":
					m.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval41.String()))
				case "alike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval41.String()))
				case "nllike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval41.String()))
				case "nrlike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval41.String()))
				case "nalike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval41.String()))
				case "in":
					m.Operator.AndWhereIn(fval41.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop41)
				}
			}
		}
	}
	ftyp42, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval42 := val.FieldByName("InsertDatetime")
		fop42, ok := ftyp42.Tag.Lookup("op")
		for fval42.Kind() == reflect.Ptr && !fval42.IsNil() {
			fval42 = fval42.Elem()
		}
		if fval42.Kind() != reflect.Ptr {
			if !ok {
				m.InsertDatetime.AndWhereEq(fval42.Interface())
			} else {
				switch fop42 {
				case "=":
					m.InsertDatetime.AndWhereEq(fval42.Interface())
				case "!=":
					m.InsertDatetime.AndWhereNeq(fval42.Interface())
				case ">":
					m.InsertDatetime.AndWhereGt(fval42.Interface())
				case ">=":
					m.InsertDatetime.AndWhereGte(fval42.Interface())
				case "<":
					m.InsertDatetime.AndWhereLt(fval42.Interface())
				case "<=":
					m.InsertDatetime.AndWhereLte(fval42.Interface())
				case "llike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval42.String()))
				case "rlike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval42.String()))
				case "alike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval42.String()))
				case "nllike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval42.String()))
				case "nrlike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval42.String()))
				case "nalike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval42.String()))
				case "in":
					m.InsertDatetime.AndWhereIn(fval42.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop42)
				}
			}
		}
	}
	ftyp43, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval43 := val.FieldByName("UpdateDatetime")
		fop43, ok := ftyp43.Tag.Lookup("op")
		for fval43.Kind() == reflect.Ptr && !fval43.IsNil() {
			fval43 = fval43.Elem()
		}
		if fval43.Kind() != reflect.Ptr {
			if !ok {
				m.UpdateDatetime.AndWhereEq(fval43.Interface())
			} else {
				switch fop43 {
				case "=":
					m.UpdateDatetime.AndWhereEq(fval43.Interface())
				case "!=":
					m.UpdateDatetime.AndWhereNeq(fval43.Interface())
				case ">":
					m.UpdateDatetime.AndWhereGt(fval43.Interface())
				case ">=":
					m.UpdateDatetime.AndWhereGte(fval43.Interface())
				case "<":
					m.UpdateDatetime.AndWhereLt(fval43.Interface())
				case "<=":
					m.UpdateDatetime.AndWhereLte(fval43.Interface())
				case "llike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval43.String()))
				case "rlike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval43.String()))
				case "alike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval43.String()))
				case "nllike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval43.String()))
				case "nrlike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval43.String()))
				case "nalike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval43.String()))
				case "in":
					m.UpdateDatetime.AndWhereIn(fval43.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop43)
				}
			}
		}
	}
	ftyp44, exists := typ.FieldByName("Status")
	if exists {
		fval44 := val.FieldByName("Status")
		fop44, ok := ftyp44.Tag.Lookup("op")
		for fval44.Kind() == reflect.Ptr && !fval44.IsNil() {
			fval44 = fval44.Elem()
		}
		if fval44.Kind() != reflect.Ptr {
			if !ok {
				m.Status.AndWhereEq(fval44.Interface())
			} else {
				switch fop44 {
				case "=":
					m.Status.AndWhereEq(fval44.Interface())
				case "!=":
					m.Status.AndWhereNeq(fval44.Interface())
				case ">":
					m.Status.AndWhereGt(fval44.Interface())
				case ">=":
					m.Status.AndWhereGte(fval44.Interface())
				case "<":
					m.Status.AndWhereLt(fval44.Interface())
				case "<=":
					m.Status.AndWhereLte(fval44.Interface())
				case "llike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s", fval44.String()))
				case "rlike":
					m.Status.AndWhereLike(fmt.Sprintf("%s%%", fval44.String()))
				case "alike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval44.String()))
				case "nllike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval44.String()))
				case "nrlike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval44.String()))
				case "nalike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval44.String()))
				case "in":
					m.Status.AndWhereIn(fval44.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop44)
				}
			}
		}
	}
	ftyp45, exists := typ.FieldByName("StudentStatus")
	if exists {
		fval45 := val.FieldByName("StudentStatus")
		fop45, ok := ftyp45.Tag.Lookup("op")
		for fval45.Kind() == reflect.Ptr && !fval45.IsNil() {
			fval45 = fval45.Elem()
		}
		if fval45.Kind() != reflect.Ptr {
			if !ok {
				m.StudentStatus.AndWhereEq(fval45.Interface())
			} else {
				switch fop45 {
				case "=":
					m.StudentStatus.AndWhereEq(fval45.Interface())
				case "!=":
					m.StudentStatus.AndWhereNeq(fval45.Interface())
				case ">":
					m.StudentStatus.AndWhereGt(fval45.Interface())
				case ">=":
					m.StudentStatus.AndWhereGte(fval45.Interface())
				case "<":
					m.StudentStatus.AndWhereLt(fval45.Interface())
				case "<=":
					m.StudentStatus.AndWhereLte(fval45.Interface())
				case "llike":
					m.StudentStatus.AndWhereLike(fmt.Sprintf("%%%s", fval45.String()))
				case "rlike":
					m.StudentStatus.AndWhereLike(fmt.Sprintf("%s%%", fval45.String()))
				case "alike":
					m.StudentStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval45.String()))
				case "nllike":
					m.StudentStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval45.String()))
				case "nrlike":
					m.StudentStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval45.String()))
				case "nalike":
					m.StudentStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval45.String()))
				case "in":
					m.StudentStatus.AndWhereIn(fval45.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop45)
				}
			}
		}
	}
	ftyp46, exists := typ.FieldByName("IsAuth")
	if exists {
		fval46 := val.FieldByName("IsAuth")
		fop46, ok := ftyp46.Tag.Lookup("op")
		for fval46.Kind() == reflect.Ptr && !fval46.IsNil() {
			fval46 = fval46.Elem()
		}
		if fval46.Kind() != reflect.Ptr {
			if !ok {
				m.IsAuth.AndWhereEq(fval46.Interface())
			} else {
				switch fop46 {
				case "=":
					m.IsAuth.AndWhereEq(fval46.Interface())
				case "!=":
					m.IsAuth.AndWhereNeq(fval46.Interface())
				case ">":
					m.IsAuth.AndWhereGt(fval46.Interface())
				case ">=":
					m.IsAuth.AndWhereGte(fval46.Interface())
				case "<":
					m.IsAuth.AndWhereLt(fval46.Interface())
				case "<=":
					m.IsAuth.AndWhereLte(fval46.Interface())
				case "llike":
					m.IsAuth.AndWhereLike(fmt.Sprintf("%%%s", fval46.String()))
				case "rlike":
					m.IsAuth.AndWhereLike(fmt.Sprintf("%s%%", fval46.String()))
				case "alike":
					m.IsAuth.AndWhereLike(fmt.Sprintf("%%%s%%", fval46.String()))
				case "nllike":
					m.IsAuth.AndWhereNotLike(fmt.Sprintf("%%%s", fval46.String()))
				case "nrlike":
					m.IsAuth.AndWhereNotLike(fmt.Sprintf("%s%%", fval46.String()))
				case "nalike":
					m.IsAuth.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval46.String()))
				case "in":
					m.IsAuth.AndWhereIn(fval46.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop46)
				}
			}
		}
	}
	ftyp47, exists := typ.FieldByName("Campus")
	if exists {
		fval47 := val.FieldByName("Campus")
		fop47, ok := ftyp47.Tag.Lookup("op")
		for fval47.Kind() == reflect.Ptr && !fval47.IsNil() {
			fval47 = fval47.Elem()
		}
		if fval47.Kind() != reflect.Ptr {
			if !ok {
				m.Campus.AndWhereEq(fval47.Interface())
			} else {
				switch fop47 {
				case "=":
					m.Campus.AndWhereEq(fval47.Interface())
				case "!=":
					m.Campus.AndWhereNeq(fval47.Interface())
				case ">":
					m.Campus.AndWhereGt(fval47.Interface())
				case ">=":
					m.Campus.AndWhereGte(fval47.Interface())
				case "<":
					m.Campus.AndWhereLt(fval47.Interface())
				case "<=":
					m.Campus.AndWhereLte(fval47.Interface())
				case "llike":
					m.Campus.AndWhereLike(fmt.Sprintf("%%%s", fval47.String()))
				case "rlike":
					m.Campus.AndWhereLike(fmt.Sprintf("%s%%", fval47.String()))
				case "alike":
					m.Campus.AndWhereLike(fmt.Sprintf("%%%s%%", fval47.String()))
				case "nllike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%%%s", fval47.String()))
				case "nrlike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%s%%", fval47.String()))
				case "nalike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval47.String()))
				case "in":
					m.Campus.AndWhereIn(fval47.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop47)
				}
			}
		}
	}
	ftyp48, exists := typ.FieldByName("Zone")
	if exists {
		fval48 := val.FieldByName("Zone")
		fop48, ok := ftyp48.Tag.Lookup("op")
		for fval48.Kind() == reflect.Ptr && !fval48.IsNil() {
			fval48 = fval48.Elem()
		}
		if fval48.Kind() != reflect.Ptr {
			if !ok {
				m.Zone.AndWhereEq(fval48.Interface())
			} else {
				switch fop48 {
				case "=":
					m.Zone.AndWhereEq(fval48.Interface())
				case "!=":
					m.Zone.AndWhereNeq(fval48.Interface())
				case ">":
					m.Zone.AndWhereGt(fval48.Interface())
				case ">=":
					m.Zone.AndWhereGte(fval48.Interface())
				case "<":
					m.Zone.AndWhereLt(fval48.Interface())
				case "<=":
					m.Zone.AndWhereLte(fval48.Interface())
				case "llike":
					m.Zone.AndWhereLike(fmt.Sprintf("%%%s", fval48.String()))
				case "rlike":
					m.Zone.AndWhereLike(fmt.Sprintf("%s%%", fval48.String()))
				case "alike":
					m.Zone.AndWhereLike(fmt.Sprintf("%%%s%%", fval48.String()))
				case "nllike":
					m.Zone.AndWhereNotLike(fmt.Sprintf("%%%s", fval48.String()))
				case "nrlike":
					m.Zone.AndWhereNotLike(fmt.Sprintf("%s%%", fval48.String()))
				case "nalike":
					m.Zone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval48.String()))
				case "in":
					m.Zone.AndWhereIn(fval48.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop48)
				}
			}
		}
	}
	ftyp49, exists := typ.FieldByName("Building")
	if exists {
		fval49 := val.FieldByName("Building")
		fop49, ok := ftyp49.Tag.Lookup("op")
		for fval49.Kind() == reflect.Ptr && !fval49.IsNil() {
			fval49 = fval49.Elem()
		}
		if fval49.Kind() != reflect.Ptr {
			if !ok {
				m.Building.AndWhereEq(fval49.Interface())
			} else {
				switch fop49 {
				case "=":
					m.Building.AndWhereEq(fval49.Interface())
				case "!=":
					m.Building.AndWhereNeq(fval49.Interface())
				case ">":
					m.Building.AndWhereGt(fval49.Interface())
				case ">=":
					m.Building.AndWhereGte(fval49.Interface())
				case "<":
					m.Building.AndWhereLt(fval49.Interface())
				case "<=":
					m.Building.AndWhereLte(fval49.Interface())
				case "llike":
					m.Building.AndWhereLike(fmt.Sprintf("%%%s", fval49.String()))
				case "rlike":
					m.Building.AndWhereLike(fmt.Sprintf("%s%%", fval49.String()))
				case "alike":
					m.Building.AndWhereLike(fmt.Sprintf("%%%s%%", fval49.String()))
				case "nllike":
					m.Building.AndWhereNotLike(fmt.Sprintf("%%%s", fval49.String()))
				case "nrlike":
					m.Building.AndWhereNotLike(fmt.Sprintf("%s%%", fval49.String()))
				case "nalike":
					m.Building.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval49.String()))
				case "in":
					m.Building.AndWhereIn(fval49.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop49)
				}
			}
		}
	}
	ftyp50, exists := typ.FieldByName("Unit")
	if exists {
		fval50 := val.FieldByName("Unit")
		fop50, ok := ftyp50.Tag.Lookup("op")
		for fval50.Kind() == reflect.Ptr && !fval50.IsNil() {
			fval50 = fval50.Elem()
		}
		if fval50.Kind() != reflect.Ptr {
			if !ok {
				m.Unit.AndWhereEq(fval50.Interface())
			} else {
				switch fop50 {
				case "=":
					m.Unit.AndWhereEq(fval50.Interface())
				case "!=":
					m.Unit.AndWhereNeq(fval50.Interface())
				case ">":
					m.Unit.AndWhereGt(fval50.Interface())
				case ">=":
					m.Unit.AndWhereGte(fval50.Interface())
				case "<":
					m.Unit.AndWhereLt(fval50.Interface())
				case "<=":
					m.Unit.AndWhereLte(fval50.Interface())
				case "llike":
					m.Unit.AndWhereLike(fmt.Sprintf("%%%s", fval50.String()))
				case "rlike":
					m.Unit.AndWhereLike(fmt.Sprintf("%s%%", fval50.String()))
				case "alike":
					m.Unit.AndWhereLike(fmt.Sprintf("%%%s%%", fval50.String()))
				case "nllike":
					m.Unit.AndWhereNotLike(fmt.Sprintf("%%%s", fval50.String()))
				case "nrlike":
					m.Unit.AndWhereNotLike(fmt.Sprintf("%s%%", fval50.String()))
				case "nalike":
					m.Unit.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval50.String()))
				case "in":
					m.Unit.AndWhereIn(fval50.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop50)
				}
			}
		}
	}
	ftyp51, exists := typ.FieldByName("Room")
	if exists {
		fval51 := val.FieldByName("Room")
		fop51, ok := ftyp51.Tag.Lookup("op")
		for fval51.Kind() == reflect.Ptr && !fval51.IsNil() {
			fval51 = fval51.Elem()
		}
		if fval51.Kind() != reflect.Ptr {
			if !ok {
				m.Room.AndWhereEq(fval51.Interface())
			} else {
				switch fop51 {
				case "=":
					m.Room.AndWhereEq(fval51.Interface())
				case "!=":
					m.Room.AndWhereNeq(fval51.Interface())
				case ">":
					m.Room.AndWhereGt(fval51.Interface())
				case ">=":
					m.Room.AndWhereGte(fval51.Interface())
				case "<":
					m.Room.AndWhereLt(fval51.Interface())
				case "<=":
					m.Room.AndWhereLte(fval51.Interface())
				case "llike":
					m.Room.AndWhereLike(fmt.Sprintf("%%%s", fval51.String()))
				case "rlike":
					m.Room.AndWhereLike(fmt.Sprintf("%s%%", fval51.String()))
				case "alike":
					m.Room.AndWhereLike(fmt.Sprintf("%%%s%%", fval51.String()))
				case "nllike":
					m.Room.AndWhereNotLike(fmt.Sprintf("%%%s", fval51.String()))
				case "nrlike":
					m.Room.AndWhereNotLike(fmt.Sprintf("%s%%", fval51.String()))
				case "nalike":
					m.Room.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval51.String()))
				case "in":
					m.Room.AndWhereIn(fval51.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop51)
				}
			}
		}
	}
	ftyp52, exists := typ.FieldByName("Bed")
	if exists {
		fval52 := val.FieldByName("Bed")
		fop52, ok := ftyp52.Tag.Lookup("op")
		for fval52.Kind() == reflect.Ptr && !fval52.IsNil() {
			fval52 = fval52.Elem()
		}
		if fval52.Kind() != reflect.Ptr {
			if !ok {
				m.Bed.AndWhereEq(fval52.Interface())
			} else {
				switch fop52 {
				case "=":
					m.Bed.AndWhereEq(fval52.Interface())
				case "!=":
					m.Bed.AndWhereNeq(fval52.Interface())
				case ">":
					m.Bed.AndWhereGt(fval52.Interface())
				case ">=":
					m.Bed.AndWhereGte(fval52.Interface())
				case "<":
					m.Bed.AndWhereLt(fval52.Interface())
				case "<=":
					m.Bed.AndWhereLte(fval52.Interface())
				case "llike":
					m.Bed.AndWhereLike(fmt.Sprintf("%%%s", fval52.String()))
				case "rlike":
					m.Bed.AndWhereLike(fmt.Sprintf("%s%%", fval52.String()))
				case "alike":
					m.Bed.AndWhereLike(fmt.Sprintf("%%%s%%", fval52.String()))
				case "nllike":
					m.Bed.AndWhereNotLike(fmt.Sprintf("%%%s", fval52.String()))
				case "nrlike":
					m.Bed.AndWhereNotLike(fmt.Sprintf("%s%%", fval52.String()))
				case "nalike":
					m.Bed.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval52.String()))
				case "in":
					m.Bed.AndWhereIn(fval52.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop52)
				}
			}
		}
	}
	ftyp53, exists := typ.FieldByName("StatusSort")
	if exists {
		fval53 := val.FieldByName("StatusSort")
		fop53, ok := ftyp53.Tag.Lookup("op")
		for fval53.Kind() == reflect.Ptr && !fval53.IsNil() {
			fval53 = fval53.Elem()
		}
		if fval53.Kind() != reflect.Ptr {
			if !ok {
				m.StatusSort.AndWhereEq(fval53.Interface())
			} else {
				switch fop53 {
				case "=":
					m.StatusSort.AndWhereEq(fval53.Interface())
				case "!=":
					m.StatusSort.AndWhereNeq(fval53.Interface())
				case ">":
					m.StatusSort.AndWhereGt(fval53.Interface())
				case ">=":
					m.StatusSort.AndWhereGte(fval53.Interface())
				case "<":
					m.StatusSort.AndWhereLt(fval53.Interface())
				case "<=":
					m.StatusSort.AndWhereLte(fval53.Interface())
				case "llike":
					m.StatusSort.AndWhereLike(fmt.Sprintf("%%%s", fval53.String()))
				case "rlike":
					m.StatusSort.AndWhereLike(fmt.Sprintf("%s%%", fval53.String()))
				case "alike":
					m.StatusSort.AndWhereLike(fmt.Sprintf("%%%s%%", fval53.String()))
				case "nllike":
					m.StatusSort.AndWhereNotLike(fmt.Sprintf("%%%s", fval53.String()))
				case "nrlike":
					m.StatusSort.AndWhereNotLike(fmt.Sprintf("%s%%", fval53.String()))
				case "nalike":
					m.StatusSort.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval53.String()))
				case "in":
					m.StatusSort.AndWhereIn(fval53.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop53)
				}
			}
		}
	}
	ftyp54, exists := typ.FieldByName("Height")
	if exists {
		fval54 := val.FieldByName("Height")
		fop54, ok := ftyp54.Tag.Lookup("op")
		for fval54.Kind() == reflect.Ptr && !fval54.IsNil() {
			fval54 = fval54.Elem()
		}
		if fval54.Kind() != reflect.Ptr {
			if !ok {
				m.Height.AndWhereEq(fval54.Interface())
			} else {
				switch fop54 {
				case "=":
					m.Height.AndWhereEq(fval54.Interface())
				case "!=":
					m.Height.AndWhereNeq(fval54.Interface())
				case ">":
					m.Height.AndWhereGt(fval54.Interface())
				case ">=":
					m.Height.AndWhereGte(fval54.Interface())
				case "<":
					m.Height.AndWhereLt(fval54.Interface())
				case "<=":
					m.Height.AndWhereLte(fval54.Interface())
				case "llike":
					m.Height.AndWhereLike(fmt.Sprintf("%%%s", fval54.String()))
				case "rlike":
					m.Height.AndWhereLike(fmt.Sprintf("%s%%", fval54.String()))
				case "alike":
					m.Height.AndWhereLike(fmt.Sprintf("%%%s%%", fval54.String()))
				case "nllike":
					m.Height.AndWhereNotLike(fmt.Sprintf("%%%s", fval54.String()))
				case "nrlike":
					m.Height.AndWhereNotLike(fmt.Sprintf("%s%%", fval54.String()))
				case "nalike":
					m.Height.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval54.String()))
				case "in":
					m.Height.AndWhereIn(fval54.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop54)
				}
			}
		}
	}
	ftyp55, exists := typ.FieldByName("Weight")
	if exists {
		fval55 := val.FieldByName("Weight")
		fop55, ok := ftyp55.Tag.Lookup("op")
		for fval55.Kind() == reflect.Ptr && !fval55.IsNil() {
			fval55 = fval55.Elem()
		}
		if fval55.Kind() != reflect.Ptr {
			if !ok {
				m.Weight.AndWhereEq(fval55.Interface())
			} else {
				switch fop55 {
				case "=":
					m.Weight.AndWhereEq(fval55.Interface())
				case "!=":
					m.Weight.AndWhereNeq(fval55.Interface())
				case ">":
					m.Weight.AndWhereGt(fval55.Interface())
				case ">=":
					m.Weight.AndWhereGte(fval55.Interface())
				case "<":
					m.Weight.AndWhereLt(fval55.Interface())
				case "<=":
					m.Weight.AndWhereLte(fval55.Interface())
				case "llike":
					m.Weight.AndWhereLike(fmt.Sprintf("%%%s", fval55.String()))
				case "rlike":
					m.Weight.AndWhereLike(fmt.Sprintf("%s%%", fval55.String()))
				case "alike":
					m.Weight.AndWhereLike(fmt.Sprintf("%%%s%%", fval55.String()))
				case "nllike":
					m.Weight.AndWhereNotLike(fmt.Sprintf("%%%s", fval55.String()))
				case "nrlike":
					m.Weight.AndWhereNotLike(fmt.Sprintf("%s%%", fval55.String()))
				case "nalike":
					m.Weight.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval55.String()))
				case "in":
					m.Weight.AndWhereIn(fval55.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop55)
				}
			}
		}
	}
	ftyp56, exists := typ.FieldByName("FootSize")
	if exists {
		fval56 := val.FieldByName("FootSize")
		fop56, ok := ftyp56.Tag.Lookup("op")
		for fval56.Kind() == reflect.Ptr && !fval56.IsNil() {
			fval56 = fval56.Elem()
		}
		if fval56.Kind() != reflect.Ptr {
			if !ok {
				m.FootSize.AndWhereEq(fval56.Interface())
			} else {
				switch fop56 {
				case "=":
					m.FootSize.AndWhereEq(fval56.Interface())
				case "!=":
					m.FootSize.AndWhereNeq(fval56.Interface())
				case ">":
					m.FootSize.AndWhereGt(fval56.Interface())
				case ">=":
					m.FootSize.AndWhereGte(fval56.Interface())
				case "<":
					m.FootSize.AndWhereLt(fval56.Interface())
				case "<=":
					m.FootSize.AndWhereLte(fval56.Interface())
				case "llike":
					m.FootSize.AndWhereLike(fmt.Sprintf("%%%s", fval56.String()))
				case "rlike":
					m.FootSize.AndWhereLike(fmt.Sprintf("%s%%", fval56.String()))
				case "alike":
					m.FootSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval56.String()))
				case "nllike":
					m.FootSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval56.String()))
				case "nrlike":
					m.FootSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval56.String()))
				case "nalike":
					m.FootSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval56.String()))
				case "in":
					m.FootSize.AndWhereIn(fval56.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop56)
				}
			}
		}
	}
	ftyp57, exists := typ.FieldByName("ClothSize")
	if exists {
		fval57 := val.FieldByName("ClothSize")
		fop57, ok := ftyp57.Tag.Lookup("op")
		for fval57.Kind() == reflect.Ptr && !fval57.IsNil() {
			fval57 = fval57.Elem()
		}
		if fval57.Kind() != reflect.Ptr {
			if !ok {
				m.ClothSize.AndWhereEq(fval57.Interface())
			} else {
				switch fop57 {
				case "=":
					m.ClothSize.AndWhereEq(fval57.Interface())
				case "!=":
					m.ClothSize.AndWhereNeq(fval57.Interface())
				case ">":
					m.ClothSize.AndWhereGt(fval57.Interface())
				case ">=":
					m.ClothSize.AndWhereGte(fval57.Interface())
				case "<":
					m.ClothSize.AndWhereLt(fval57.Interface())
				case "<=":
					m.ClothSize.AndWhereLte(fval57.Interface())
				case "llike":
					m.ClothSize.AndWhereLike(fmt.Sprintf("%%%s", fval57.String()))
				case "rlike":
					m.ClothSize.AndWhereLike(fmt.Sprintf("%s%%", fval57.String()))
				case "alike":
					m.ClothSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval57.String()))
				case "nllike":
					m.ClothSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval57.String()))
				case "nrlike":
					m.ClothSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval57.String()))
				case "nalike":
					m.ClothSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval57.String()))
				case "in":
					m.ClothSize.AndWhereIn(fval57.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop57)
				}
			}
		}
	}
	ftyp58, exists := typ.FieldByName("HeadSize")
	if exists {
		fval58 := val.FieldByName("HeadSize")
		fop58, ok := ftyp58.Tag.Lookup("op")
		for fval58.Kind() == reflect.Ptr && !fval58.IsNil() {
			fval58 = fval58.Elem()
		}
		if fval58.Kind() != reflect.Ptr {
			if !ok {
				m.HeadSize.AndWhereEq(fval58.Interface())
			} else {
				switch fop58 {
				case "=":
					m.HeadSize.AndWhereEq(fval58.Interface())
				case "!=":
					m.HeadSize.AndWhereNeq(fval58.Interface())
				case ">":
					m.HeadSize.AndWhereGt(fval58.Interface())
				case ">=":
					m.HeadSize.AndWhereGte(fval58.Interface())
				case "<":
					m.HeadSize.AndWhereLt(fval58.Interface())
				case "<=":
					m.HeadSize.AndWhereLte(fval58.Interface())
				case "llike":
					m.HeadSize.AndWhereLike(fmt.Sprintf("%%%s", fval58.String()))
				case "rlike":
					m.HeadSize.AndWhereLike(fmt.Sprintf("%s%%", fval58.String()))
				case "alike":
					m.HeadSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval58.String()))
				case "nllike":
					m.HeadSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval58.String()))
				case "nrlike":
					m.HeadSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval58.String()))
				case "nalike":
					m.HeadSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval58.String()))
				case "in":
					m.HeadSize.AndWhereIn(fval58.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop58)
				}
			}
		}
	}
	ftyp59, exists := typ.FieldByName("Remark1")
	if exists {
		fval59 := val.FieldByName("Remark1")
		fop59, ok := ftyp59.Tag.Lookup("op")
		for fval59.Kind() == reflect.Ptr && !fval59.IsNil() {
			fval59 = fval59.Elem()
		}
		if fval59.Kind() != reflect.Ptr {
			if !ok {
				m.Remark1.AndWhereEq(fval59.Interface())
			} else {
				switch fop59 {
				case "=":
					m.Remark1.AndWhereEq(fval59.Interface())
				case "!=":
					m.Remark1.AndWhereNeq(fval59.Interface())
				case ">":
					m.Remark1.AndWhereGt(fval59.Interface())
				case ">=":
					m.Remark1.AndWhereGte(fval59.Interface())
				case "<":
					m.Remark1.AndWhereLt(fval59.Interface())
				case "<=":
					m.Remark1.AndWhereLte(fval59.Interface())
				case "llike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval59.String()))
				case "rlike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval59.String()))
				case "alike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval59.String()))
				case "nllike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval59.String()))
				case "nrlike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval59.String()))
				case "nalike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval59.String()))
				case "in":
					m.Remark1.AndWhereIn(fval59.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop59)
				}
			}
		}
	}
	ftyp60, exists := typ.FieldByName("Remark2")
	if exists {
		fval60 := val.FieldByName("Remark2")
		fop60, ok := ftyp60.Tag.Lookup("op")
		for fval60.Kind() == reflect.Ptr && !fval60.IsNil() {
			fval60 = fval60.Elem()
		}
		if fval60.Kind() != reflect.Ptr {
			if !ok {
				m.Remark2.AndWhereEq(fval60.Interface())
			} else {
				switch fop60 {
				case "=":
					m.Remark2.AndWhereEq(fval60.Interface())
				case "!=":
					m.Remark2.AndWhereNeq(fval60.Interface())
				case ">":
					m.Remark2.AndWhereGt(fval60.Interface())
				case ">=":
					m.Remark2.AndWhereGte(fval60.Interface())
				case "<":
					m.Remark2.AndWhereLt(fval60.Interface())
				case "<=":
					m.Remark2.AndWhereLte(fval60.Interface())
				case "llike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval60.String()))
				case "rlike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval60.String()))
				case "alike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval60.String()))
				case "nllike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval60.String()))
				case "nrlike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval60.String()))
				case "nalike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval60.String()))
				case "in":
					m.Remark2.AndWhereIn(fval60.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop60)
				}
			}
		}
	}
	ftyp61, exists := typ.FieldByName("Remark3")
	if exists {
		fval61 := val.FieldByName("Remark3")
		fop61, ok := ftyp61.Tag.Lookup("op")
		for fval61.Kind() == reflect.Ptr && !fval61.IsNil() {
			fval61 = fval61.Elem()
		}
		if fval61.Kind() != reflect.Ptr {
			if !ok {
				m.Remark3.AndWhereEq(fval61.Interface())
			} else {
				switch fop61 {
				case "=":
					m.Remark3.AndWhereEq(fval61.Interface())
				case "!=":
					m.Remark3.AndWhereNeq(fval61.Interface())
				case ">":
					m.Remark3.AndWhereGt(fval61.Interface())
				case ">=":
					m.Remark3.AndWhereGte(fval61.Interface())
				case "<":
					m.Remark3.AndWhereLt(fval61.Interface())
				case "<=":
					m.Remark3.AndWhereLte(fval61.Interface())
				case "llike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval61.String()))
				case "rlike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval61.String()))
				case "alike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval61.String()))
				case "nllike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval61.String()))
				case "nrlike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval61.String()))
				case "nalike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval61.String()))
				case "in":
					m.Remark3.AndWhereIn(fval61.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop61)
				}
			}
		}
	}
	ftyp62, exists := typ.FieldByName("Remark4")
	if exists {
		fval62 := val.FieldByName("Remark4")
		fop62, ok := ftyp62.Tag.Lookup("op")
		for fval62.Kind() == reflect.Ptr && !fval62.IsNil() {
			fval62 = fval62.Elem()
		}
		if fval62.Kind() != reflect.Ptr {
			if !ok {
				m.Remark4.AndWhereEq(fval62.Interface())
			} else {
				switch fop62 {
				case "=":
					m.Remark4.AndWhereEq(fval62.Interface())
				case "!=":
					m.Remark4.AndWhereNeq(fval62.Interface())
				case ">":
					m.Remark4.AndWhereGt(fval62.Interface())
				case ">=":
					m.Remark4.AndWhereGte(fval62.Interface())
				case "<":
					m.Remark4.AndWhereLt(fval62.Interface())
				case "<=":
					m.Remark4.AndWhereLte(fval62.Interface())
				case "llike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval62.String()))
				case "rlike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval62.String()))
				case "alike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval62.String()))
				case "nllike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval62.String()))
				case "nrlike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval62.String()))
				case "nalike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval62.String()))
				case "in":
					m.Remark4.AndWhereIn(fval62.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop62)
				}
			}
		}
	}
	ftyp63, exists := typ.FieldByName("IsPayment")
	if exists {
		fval63 := val.FieldByName("IsPayment")
		fop63, ok := ftyp63.Tag.Lookup("op")
		for fval63.Kind() == reflect.Ptr && !fval63.IsNil() {
			fval63 = fval63.Elem()
		}
		if fval63.Kind() != reflect.Ptr {
			if !ok {
				m.IsPayment.AndWhereEq(fval63.Interface())
			} else {
				switch fop63 {
				case "=":
					m.IsPayment.AndWhereEq(fval63.Interface())
				case "!=":
					m.IsPayment.AndWhereNeq(fval63.Interface())
				case ">":
					m.IsPayment.AndWhereGt(fval63.Interface())
				case ">=":
					m.IsPayment.AndWhereGte(fval63.Interface())
				case "<":
					m.IsPayment.AndWhereLt(fval63.Interface())
				case "<=":
					m.IsPayment.AndWhereLte(fval63.Interface())
				case "llike":
					m.IsPayment.AndWhereLike(fmt.Sprintf("%%%s", fval63.String()))
				case "rlike":
					m.IsPayment.AndWhereLike(fmt.Sprintf("%s%%", fval63.String()))
				case "alike":
					m.IsPayment.AndWhereLike(fmt.Sprintf("%%%s%%", fval63.String()))
				case "nllike":
					m.IsPayment.AndWhereNotLike(fmt.Sprintf("%%%s", fval63.String()))
				case "nrlike":
					m.IsPayment.AndWhereNotLike(fmt.Sprintf("%s%%", fval63.String()))
				case "nalike":
					m.IsPayment.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval63.String()))
				case "in":
					m.IsPayment.AndWhereIn(fval63.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop63)
				}
			}
		}
	}
	ftyp64, exists := typ.FieldByName("IsCheckIn")
	if exists {
		fval64 := val.FieldByName("IsCheckIn")
		fop64, ok := ftyp64.Tag.Lookup("op")
		for fval64.Kind() == reflect.Ptr && !fval64.IsNil() {
			fval64 = fval64.Elem()
		}
		if fval64.Kind() != reflect.Ptr {
			if !ok {
				m.IsCheckIn.AndWhereEq(fval64.Interface())
			} else {
				switch fop64 {
				case "=":
					m.IsCheckIn.AndWhereEq(fval64.Interface())
				case "!=":
					m.IsCheckIn.AndWhereNeq(fval64.Interface())
				case ">":
					m.IsCheckIn.AndWhereGt(fval64.Interface())
				case ">=":
					m.IsCheckIn.AndWhereGte(fval64.Interface())
				case "<":
					m.IsCheckIn.AndWhereLt(fval64.Interface())
				case "<=":
					m.IsCheckIn.AndWhereLte(fval64.Interface())
				case "llike":
					m.IsCheckIn.AndWhereLike(fmt.Sprintf("%%%s", fval64.String()))
				case "rlike":
					m.IsCheckIn.AndWhereLike(fmt.Sprintf("%s%%", fval64.String()))
				case "alike":
					m.IsCheckIn.AndWhereLike(fmt.Sprintf("%%%s%%", fval64.String()))
				case "nllike":
					m.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%%%s", fval64.String()))
				case "nrlike":
					m.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%s%%", fval64.String()))
				case "nalike":
					m.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval64.String()))
				case "in":
					m.IsCheckIn.AndWhereIn(fval64.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop64)
				}
			}
		}
	}
	ftyp65, exists := typ.FieldByName("GetMilitaryTC")
	if exists {
		fval65 := val.FieldByName("GetMilitaryTC")
		fop65, ok := ftyp65.Tag.Lookup("op")
		for fval65.Kind() == reflect.Ptr && !fval65.IsNil() {
			fval65 = fval65.Elem()
		}
		if fval65.Kind() != reflect.Ptr {
			if !ok {
				m.GetMilitaryTC.AndWhereEq(fval65.Interface())
			} else {
				switch fop65 {
				case "=":
					m.GetMilitaryTC.AndWhereEq(fval65.Interface())
				case "!=":
					m.GetMilitaryTC.AndWhereNeq(fval65.Interface())
				case ">":
					m.GetMilitaryTC.AndWhereGt(fval65.Interface())
				case ">=":
					m.GetMilitaryTC.AndWhereGte(fval65.Interface())
				case "<":
					m.GetMilitaryTC.AndWhereLt(fval65.Interface())
				case "<=":
					m.GetMilitaryTC.AndWhereLte(fval65.Interface())
				case "llike":
					m.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%%%s", fval65.String()))
				case "rlike":
					m.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%s%%", fval65.String()))
				case "alike":
					m.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%%%s%%", fval65.String()))
				case "nllike":
					m.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%%%s", fval65.String()))
				case "nrlike":
					m.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%s%%", fval65.String()))
				case "nalike":
					m.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval65.String()))
				case "in":
					m.GetMilitaryTC.AndWhereIn(fval65.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop65)
				}
			}
		}
	}
	ftyp66, exists := typ.FieldByName("OriginAreaName")
	if exists {
		fval66 := val.FieldByName("OriginAreaName")
		fop66, ok := ftyp66.Tag.Lookup("op")
		for fval66.Kind() == reflect.Ptr && !fval66.IsNil() {
			fval66 = fval66.Elem()
		}
		if fval66.Kind() != reflect.Ptr {
			if !ok {
				m.OriginAreaName.AndWhereEq(fval66.Interface())
			} else {
				switch fop66 {
				case "=":
					m.OriginAreaName.AndWhereEq(fval66.Interface())
				case "!=":
					m.OriginAreaName.AndWhereNeq(fval66.Interface())
				case ">":
					m.OriginAreaName.AndWhereGt(fval66.Interface())
				case ">=":
					m.OriginAreaName.AndWhereGte(fval66.Interface())
				case "<":
					m.OriginAreaName.AndWhereLt(fval66.Interface())
				case "<=":
					m.OriginAreaName.AndWhereLte(fval66.Interface())
				case "llike":
					m.OriginAreaName.AndWhereLike(fmt.Sprintf("%%%s", fval66.String()))
				case "rlike":
					m.OriginAreaName.AndWhereLike(fmt.Sprintf("%s%%", fval66.String()))
				case "alike":
					m.OriginAreaName.AndWhereLike(fmt.Sprintf("%%%s%%", fval66.String()))
				case "nllike":
					m.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%%%s", fval66.String()))
				case "nrlike":
					m.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%s%%", fval66.String()))
				case "nalike":
					m.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval66.String()))
				case "in":
					m.OriginAreaName.AndWhereIn(fval66.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop66)
				}
			}
		}
	}
	return m, nil
}
func (l *StudentbasicinfoList) FromQuery(query interface{}) (*StudentbasicinfoList, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				l.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					l.Id.AndWhereEq(fval0.Interface())
				case "!=":
					l.Id.AndWhereNeq(fval0.Interface())
				case ">":
					l.Id.AndWhereGt(fval0.Interface())
				case ">=":
					l.Id.AndWhereGte(fval0.Interface())
				case "<":
					l.Id.AndWhereLt(fval0.Interface())
				case "<=":
					l.Id.AndWhereLte(fval0.Interface())
				case "llike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					l.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					l.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("RecordId")
	if exists {
		fval1 := val.FieldByName("RecordId")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				l.RecordId.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					l.RecordId.AndWhereEq(fval1.Interface())
				case "!=":
					l.RecordId.AndWhereNeq(fval1.Interface())
				case ">":
					l.RecordId.AndWhereGt(fval1.Interface())
				case ">=":
					l.RecordId.AndWhereGte(fval1.Interface())
				case "<":
					l.RecordId.AndWhereLt(fval1.Interface())
				case "<=":
					l.RecordId.AndWhereLte(fval1.Interface())
				case "llike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					l.RecordId.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("IntelUserCode")
	if exists {
		fval2 := val.FieldByName("IntelUserCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				l.IntelUserCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					l.IntelUserCode.AndWhereEq(fval2.Interface())
				case "!=":
					l.IntelUserCode.AndWhereNeq(fval2.Interface())
				case ">":
					l.IntelUserCode.AndWhereGt(fval2.Interface())
				case ">=":
					l.IntelUserCode.AndWhereGte(fval2.Interface())
				case "<":
					l.IntelUserCode.AndWhereLt(fval2.Interface())
				case "<=":
					l.IntelUserCode.AndWhereLte(fval2.Interface())
				case "llike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					l.IntelUserCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					l.IntelUserCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					l.IntelUserCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("Class")
	if exists {
		fval3 := val.FieldByName("Class")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				l.Class.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					l.Class.AndWhereEq(fval3.Interface())
				case "!=":
					l.Class.AndWhereNeq(fval3.Interface())
				case ">":
					l.Class.AndWhereGt(fval3.Interface())
				case ">=":
					l.Class.AndWhereGte(fval3.Interface())
				case "<":
					l.Class.AndWhereLt(fval3.Interface())
				case "<=":
					l.Class.AndWhereLte(fval3.Interface())
				case "llike":
					l.Class.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					l.Class.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					l.Class.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					l.Class.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					l.Class.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					l.Class.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					l.Class.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("OtherName")
	if exists {
		fval4 := val.FieldByName("OtherName")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				l.OtherName.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					l.OtherName.AndWhereEq(fval4.Interface())
				case "!=":
					l.OtherName.AndWhereNeq(fval4.Interface())
				case ">":
					l.OtherName.AndWhereGt(fval4.Interface())
				case ">=":
					l.OtherName.AndWhereGte(fval4.Interface())
				case "<":
					l.OtherName.AndWhereLt(fval4.Interface())
				case "<=":
					l.OtherName.AndWhereLte(fval4.Interface())
				case "llike":
					l.OtherName.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					l.OtherName.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					l.OtherName.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					l.OtherName.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					l.OtherName.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					l.OtherName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					l.OtherName.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("NameInPinyin")
	if exists {
		fval5 := val.FieldByName("NameInPinyin")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				l.NameInPinyin.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					l.NameInPinyin.AndWhereEq(fval5.Interface())
				case "!=":
					l.NameInPinyin.AndWhereNeq(fval5.Interface())
				case ">":
					l.NameInPinyin.AndWhereGt(fval5.Interface())
				case ">=":
					l.NameInPinyin.AndWhereGte(fval5.Interface())
				case "<":
					l.NameInPinyin.AndWhereLt(fval5.Interface())
				case "<=":
					l.NameInPinyin.AndWhereLte(fval5.Interface())
				case "llike":
					l.NameInPinyin.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					l.NameInPinyin.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					l.NameInPinyin.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					l.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					l.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					l.NameInPinyin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					l.NameInPinyin.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("EnglishName")
	if exists {
		fval6 := val.FieldByName("EnglishName")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				l.EnglishName.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					l.EnglishName.AndWhereEq(fval6.Interface())
				case "!=":
					l.EnglishName.AndWhereNeq(fval6.Interface())
				case ">":
					l.EnglishName.AndWhereGt(fval6.Interface())
				case ">=":
					l.EnglishName.AndWhereGte(fval6.Interface())
				case "<":
					l.EnglishName.AndWhereLt(fval6.Interface())
				case "<=":
					l.EnglishName.AndWhereLte(fval6.Interface())
				case "llike":
					l.EnglishName.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					l.EnglishName.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					l.EnglishName.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					l.EnglishName.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					l.EnglishName.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					l.EnglishName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					l.EnglishName.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("CountryCode")
	if exists {
		fval7 := val.FieldByName("CountryCode")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				l.CountryCode.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					l.CountryCode.AndWhereEq(fval7.Interface())
				case "!=":
					l.CountryCode.AndWhereNeq(fval7.Interface())
				case ">":
					l.CountryCode.AndWhereGt(fval7.Interface())
				case ">=":
					l.CountryCode.AndWhereGte(fval7.Interface())
				case "<":
					l.CountryCode.AndWhereLt(fval7.Interface())
				case "<=":
					l.CountryCode.AndWhereLte(fval7.Interface())
				case "llike":
					l.CountryCode.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					l.CountryCode.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					l.CountryCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					l.CountryCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					l.CountryCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					l.CountryCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					l.CountryCode.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("NationalityCode")
	if exists {
		fval8 := val.FieldByName("NationalityCode")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				l.NationalityCode.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					l.NationalityCode.AndWhereEq(fval8.Interface())
				case "!=":
					l.NationalityCode.AndWhereNeq(fval8.Interface())
				case ">":
					l.NationalityCode.AndWhereGt(fval8.Interface())
				case ">=":
					l.NationalityCode.AndWhereGte(fval8.Interface())
				case "<":
					l.NationalityCode.AndWhereLt(fval8.Interface())
				case "<=":
					l.NationalityCode.AndWhereLte(fval8.Interface())
				case "llike":
					l.NationalityCode.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					l.NationalityCode.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					l.NationalityCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					l.NationalityCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					l.NationalityCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					l.NationalityCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					l.NationalityCode.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("Birthday")
	if exists {
		fval9 := val.FieldByName("Birthday")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				l.Birthday.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					l.Birthday.AndWhereEq(fval9.Interface())
				case "!=":
					l.Birthday.AndWhereNeq(fval9.Interface())
				case ">":
					l.Birthday.AndWhereGt(fval9.Interface())
				case ">=":
					l.Birthday.AndWhereGte(fval9.Interface())
				case "<":
					l.Birthday.AndWhereLt(fval9.Interface())
				case "<=":
					l.Birthday.AndWhereLte(fval9.Interface())
				case "llike":
					l.Birthday.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					l.Birthday.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					l.Birthday.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					l.Birthday.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					l.Birthday.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					l.Birthday.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					l.Birthday.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("PoliticalCode")
	if exists {
		fval10 := val.FieldByName("PoliticalCode")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				l.PoliticalCode.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					l.PoliticalCode.AndWhereEq(fval10.Interface())
				case "!=":
					l.PoliticalCode.AndWhereNeq(fval10.Interface())
				case ">":
					l.PoliticalCode.AndWhereGt(fval10.Interface())
				case ">=":
					l.PoliticalCode.AndWhereGte(fval10.Interface())
				case "<":
					l.PoliticalCode.AndWhereLt(fval10.Interface())
				case "<=":
					l.PoliticalCode.AndWhereLte(fval10.Interface())
				case "llike":
					l.PoliticalCode.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					l.PoliticalCode.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					l.PoliticalCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					l.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					l.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					l.PoliticalCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					l.PoliticalCode.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("QQAcct")
	if exists {
		fval11 := val.FieldByName("QQAcct")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				l.QQAcct.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					l.QQAcct.AndWhereEq(fval11.Interface())
				case "!=":
					l.QQAcct.AndWhereNeq(fval11.Interface())
				case ">":
					l.QQAcct.AndWhereGt(fval11.Interface())
				case ">=":
					l.QQAcct.AndWhereGte(fval11.Interface())
				case "<":
					l.QQAcct.AndWhereLt(fval11.Interface())
				case "<=":
					l.QQAcct.AndWhereLte(fval11.Interface())
				case "llike":
					l.QQAcct.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					l.QQAcct.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					l.QQAcct.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					l.QQAcct.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					l.QQAcct.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					l.QQAcct.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					l.QQAcct.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("WeChatAcct")
	if exists {
		fval12 := val.FieldByName("WeChatAcct")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				l.WeChatAcct.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					l.WeChatAcct.AndWhereEq(fval12.Interface())
				case "!=":
					l.WeChatAcct.AndWhereNeq(fval12.Interface())
				case ">":
					l.WeChatAcct.AndWhereGt(fval12.Interface())
				case ">=":
					l.WeChatAcct.AndWhereGte(fval12.Interface())
				case "<":
					l.WeChatAcct.AndWhereLt(fval12.Interface())
				case "<=":
					l.WeChatAcct.AndWhereLte(fval12.Interface())
				case "llike":
					l.WeChatAcct.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					l.WeChatAcct.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					l.WeChatAcct.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					l.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					l.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					l.WeChatAcct.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					l.WeChatAcct.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("BankCardNumber")
	if exists {
		fval13 := val.FieldByName("BankCardNumber")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				l.BankCardNumber.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					l.BankCardNumber.AndWhereEq(fval13.Interface())
				case "!=":
					l.BankCardNumber.AndWhereNeq(fval13.Interface())
				case ">":
					l.BankCardNumber.AndWhereGt(fval13.Interface())
				case ">=":
					l.BankCardNumber.AndWhereGte(fval13.Interface())
				case "<":
					l.BankCardNumber.AndWhereLt(fval13.Interface())
				case "<=":
					l.BankCardNumber.AndWhereLte(fval13.Interface())
				case "llike":
					l.BankCardNumber.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					l.BankCardNumber.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					l.BankCardNumber.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					l.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					l.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					l.BankCardNumber.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					l.BankCardNumber.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("AccountBankCode")
	if exists {
		fval14 := val.FieldByName("AccountBankCode")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				l.AccountBankCode.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					l.AccountBankCode.AndWhereEq(fval14.Interface())
				case "!=":
					l.AccountBankCode.AndWhereNeq(fval14.Interface())
				case ">":
					l.AccountBankCode.AndWhereGt(fval14.Interface())
				case ">=":
					l.AccountBankCode.AndWhereGte(fval14.Interface())
				case "<":
					l.AccountBankCode.AndWhereLt(fval14.Interface())
				case "<=":
					l.AccountBankCode.AndWhereLte(fval14.Interface())
				case "llike":
					l.AccountBankCode.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					l.AccountBankCode.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					l.AccountBankCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					l.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					l.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					l.AccountBankCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					l.AccountBankCode.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("AllPowerfulCardNum")
	if exists {
		fval15 := val.FieldByName("AllPowerfulCardNum")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				l.AllPowerfulCardNum.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					l.AllPowerfulCardNum.AndWhereEq(fval15.Interface())
				case "!=":
					l.AllPowerfulCardNum.AndWhereNeq(fval15.Interface())
				case ">":
					l.AllPowerfulCardNum.AndWhereGt(fval15.Interface())
				case ">=":
					l.AllPowerfulCardNum.AndWhereGte(fval15.Interface())
				case "<":
					l.AllPowerfulCardNum.AndWhereLt(fval15.Interface())
				case "<=":
					l.AllPowerfulCardNum.AndWhereLte(fval15.Interface())
				case "llike":
					l.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					l.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					l.AllPowerfulCardNum.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					l.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					l.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					l.AllPowerfulCardNum.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					l.AllPowerfulCardNum.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("MaritalCode")
	if exists {
		fval16 := val.FieldByName("MaritalCode")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				l.MaritalCode.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					l.MaritalCode.AndWhereEq(fval16.Interface())
				case "!=":
					l.MaritalCode.AndWhereNeq(fval16.Interface())
				case ">":
					l.MaritalCode.AndWhereGt(fval16.Interface())
				case ">=":
					l.MaritalCode.AndWhereGte(fval16.Interface())
				case "<":
					l.MaritalCode.AndWhereLt(fval16.Interface())
				case "<=":
					l.MaritalCode.AndWhereLte(fval16.Interface())
				case "llike":
					l.MaritalCode.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					l.MaritalCode.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					l.MaritalCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					l.MaritalCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					l.MaritalCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					l.MaritalCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					l.MaritalCode.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("OriginAreaCode")
	if exists {
		fval17 := val.FieldByName("OriginAreaCode")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				l.OriginAreaCode.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					l.OriginAreaCode.AndWhereEq(fval17.Interface())
				case "!=":
					l.OriginAreaCode.AndWhereNeq(fval17.Interface())
				case ">":
					l.OriginAreaCode.AndWhereGt(fval17.Interface())
				case ">=":
					l.OriginAreaCode.AndWhereGte(fval17.Interface())
				case "<":
					l.OriginAreaCode.AndWhereLt(fval17.Interface())
				case "<=":
					l.OriginAreaCode.AndWhereLte(fval17.Interface())
				case "llike":
					l.OriginAreaCode.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					l.OriginAreaCode.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					l.OriginAreaCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					l.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					l.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					l.OriginAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					l.OriginAreaCode.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("StudentAreaCode")
	if exists {
		fval18 := val.FieldByName("StudentAreaCode")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				l.StudentAreaCode.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					l.StudentAreaCode.AndWhereEq(fval18.Interface())
				case "!=":
					l.StudentAreaCode.AndWhereNeq(fval18.Interface())
				case ">":
					l.StudentAreaCode.AndWhereGt(fval18.Interface())
				case ">=":
					l.StudentAreaCode.AndWhereGte(fval18.Interface())
				case "<":
					l.StudentAreaCode.AndWhereLt(fval18.Interface())
				case "<=":
					l.StudentAreaCode.AndWhereLte(fval18.Interface())
				case "llike":
					l.StudentAreaCode.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					l.StudentAreaCode.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					l.StudentAreaCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					l.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					l.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					l.StudentAreaCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					l.StudentAreaCode.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Hobbies")
	if exists {
		fval19 := val.FieldByName("Hobbies")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				l.Hobbies.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					l.Hobbies.AndWhereEq(fval19.Interface())
				case "!=":
					l.Hobbies.AndWhereNeq(fval19.Interface())
				case ">":
					l.Hobbies.AndWhereGt(fval19.Interface())
				case ">=":
					l.Hobbies.AndWhereGte(fval19.Interface())
				case "<":
					l.Hobbies.AndWhereLt(fval19.Interface())
				case "<=":
					l.Hobbies.AndWhereLte(fval19.Interface())
				case "llike":
					l.Hobbies.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					l.Hobbies.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					l.Hobbies.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					l.Hobbies.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					l.Hobbies.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					l.Hobbies.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					l.Hobbies.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("Creed")
	if exists {
		fval20 := val.FieldByName("Creed")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				l.Creed.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					l.Creed.AndWhereEq(fval20.Interface())
				case "!=":
					l.Creed.AndWhereNeq(fval20.Interface())
				case ">":
					l.Creed.AndWhereGt(fval20.Interface())
				case ">=":
					l.Creed.AndWhereGte(fval20.Interface())
				case "<":
					l.Creed.AndWhereLt(fval20.Interface())
				case "<=":
					l.Creed.AndWhereLte(fval20.Interface())
				case "llike":
					l.Creed.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					l.Creed.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					l.Creed.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					l.Creed.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					l.Creed.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					l.Creed.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					l.Creed.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("TrainTicketinterval")
	if exists {
		fval21 := val.FieldByName("TrainTicketinterval")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				l.TrainTicketinterval.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					l.TrainTicketinterval.AndWhereEq(fval21.Interface())
				case "!=":
					l.TrainTicketinterval.AndWhereNeq(fval21.Interface())
				case ">":
					l.TrainTicketinterval.AndWhereGt(fval21.Interface())
				case ">=":
					l.TrainTicketinterval.AndWhereGte(fval21.Interface())
				case "<":
					l.TrainTicketinterval.AndWhereLt(fval21.Interface())
				case "<=":
					l.TrainTicketinterval.AndWhereLte(fval21.Interface())
				case "llike":
					l.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					l.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					l.TrainTicketinterval.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					l.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					l.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					l.TrainTicketinterval.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					l.TrainTicketinterval.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("FamilyAddress")
	if exists {
		fval22 := val.FieldByName("FamilyAddress")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				l.FamilyAddress.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					l.FamilyAddress.AndWhereEq(fval22.Interface())
				case "!=":
					l.FamilyAddress.AndWhereNeq(fval22.Interface())
				case ">":
					l.FamilyAddress.AndWhereGt(fval22.Interface())
				case ">=":
					l.FamilyAddress.AndWhereGte(fval22.Interface())
				case "<":
					l.FamilyAddress.AndWhereLt(fval22.Interface())
				case "<=":
					l.FamilyAddress.AndWhereLte(fval22.Interface())
				case "llike":
					l.FamilyAddress.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					l.FamilyAddress.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					l.FamilyAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					l.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					l.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					l.FamilyAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					l.FamilyAddress.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("DetailAddress")
	if exists {
		fval23 := val.FieldByName("DetailAddress")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				l.DetailAddress.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					l.DetailAddress.AndWhereEq(fval23.Interface())
				case "!=":
					l.DetailAddress.AndWhereNeq(fval23.Interface())
				case ">":
					l.DetailAddress.AndWhereGt(fval23.Interface())
				case ">=":
					l.DetailAddress.AndWhereGte(fval23.Interface())
				case "<":
					l.DetailAddress.AndWhereLt(fval23.Interface())
				case "<=":
					l.DetailAddress.AndWhereLte(fval23.Interface())
				case "llike":
					l.DetailAddress.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					l.DetailAddress.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					l.DetailAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					l.DetailAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					l.DetailAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					l.DetailAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					l.DetailAddress.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("PostCode")
	if exists {
		fval24 := val.FieldByName("PostCode")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				l.PostCode.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					l.PostCode.AndWhereEq(fval24.Interface())
				case "!=":
					l.PostCode.AndWhereNeq(fval24.Interface())
				case ">":
					l.PostCode.AndWhereGt(fval24.Interface())
				case ">=":
					l.PostCode.AndWhereGte(fval24.Interface())
				case "<":
					l.PostCode.AndWhereLt(fval24.Interface())
				case "<=":
					l.PostCode.AndWhereLte(fval24.Interface())
				case "llike":
					l.PostCode.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					l.PostCode.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					l.PostCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					l.PostCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					l.PostCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					l.PostCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					l.PostCode.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("HomePhone")
	if exists {
		fval25 := val.FieldByName("HomePhone")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				l.HomePhone.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					l.HomePhone.AndWhereEq(fval25.Interface())
				case "!=":
					l.HomePhone.AndWhereNeq(fval25.Interface())
				case ">":
					l.HomePhone.AndWhereGt(fval25.Interface())
				case ">=":
					l.HomePhone.AndWhereGte(fval25.Interface())
				case "<":
					l.HomePhone.AndWhereLt(fval25.Interface())
				case "<=":
					l.HomePhone.AndWhereLte(fval25.Interface())
				case "llike":
					l.HomePhone.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					l.HomePhone.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					l.HomePhone.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					l.HomePhone.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					l.HomePhone.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					l.HomePhone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					l.HomePhone.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("EnrollmentDate")
	if exists {
		fval26 := val.FieldByName("EnrollmentDate")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				l.EnrollmentDate.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					l.EnrollmentDate.AndWhereEq(fval26.Interface())
				case "!=":
					l.EnrollmentDate.AndWhereNeq(fval26.Interface())
				case ">":
					l.EnrollmentDate.AndWhereGt(fval26.Interface())
				case ">=":
					l.EnrollmentDate.AndWhereGte(fval26.Interface())
				case "<":
					l.EnrollmentDate.AndWhereLt(fval26.Interface())
				case "<=":
					l.EnrollmentDate.AndWhereLte(fval26.Interface())
				case "llike":
					l.EnrollmentDate.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					l.EnrollmentDate.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					l.EnrollmentDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					l.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					l.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					l.EnrollmentDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					l.EnrollmentDate.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("GraduationDate")
	if exists {
		fval27 := val.FieldByName("GraduationDate")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				l.GraduationDate.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					l.GraduationDate.AndWhereEq(fval27.Interface())
				case "!=":
					l.GraduationDate.AndWhereNeq(fval27.Interface())
				case ">":
					l.GraduationDate.AndWhereGt(fval27.Interface())
				case ">=":
					l.GraduationDate.AndWhereGte(fval27.Interface())
				case "<":
					l.GraduationDate.AndWhereLt(fval27.Interface())
				case "<=":
					l.GraduationDate.AndWhereLte(fval27.Interface())
				case "llike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					l.GraduationDate.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	ftyp28, exists := typ.FieldByName("MidSchoolAddress")
	if exists {
		fval28 := val.FieldByName("MidSchoolAddress")
		fop28, ok := ftyp28.Tag.Lookup("op")
		for fval28.Kind() == reflect.Ptr && !fval28.IsNil() {
			fval28 = fval28.Elem()
		}
		if fval28.Kind() != reflect.Ptr {
			if !ok {
				l.MidSchoolAddress.AndWhereEq(fval28.Interface())
			} else {
				switch fop28 {
				case "=":
					l.MidSchoolAddress.AndWhereEq(fval28.Interface())
				case "!=":
					l.MidSchoolAddress.AndWhereNeq(fval28.Interface())
				case ">":
					l.MidSchoolAddress.AndWhereGt(fval28.Interface())
				case ">=":
					l.MidSchoolAddress.AndWhereGte(fval28.Interface())
				case "<":
					l.MidSchoolAddress.AndWhereLt(fval28.Interface())
				case "<=":
					l.MidSchoolAddress.AndWhereLte(fval28.Interface())
				case "llike":
					l.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%%%s", fval28.String()))
				case "rlike":
					l.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%s%%", fval28.String()))
				case "alike":
					l.MidSchoolAddress.AndWhereLike(fmt.Sprintf("%%%s%%", fval28.String()))
				case "nllike":
					l.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%%%s", fval28.String()))
				case "nrlike":
					l.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%s%%", fval28.String()))
				case "nalike":
					l.MidSchoolAddress.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval28.String()))
				case "in":
					l.MidSchoolAddress.AndWhereIn(fval28.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop28)
				}
			}
		}
	}
	ftyp29, exists := typ.FieldByName("MidSchoolName")
	if exists {
		fval29 := val.FieldByName("MidSchoolName")
		fop29, ok := ftyp29.Tag.Lookup("op")
		for fval29.Kind() == reflect.Ptr && !fval29.IsNil() {
			fval29 = fval29.Elem()
		}
		if fval29.Kind() != reflect.Ptr {
			if !ok {
				l.MidSchoolName.AndWhereEq(fval29.Interface())
			} else {
				switch fop29 {
				case "=":
					l.MidSchoolName.AndWhereEq(fval29.Interface())
				case "!=":
					l.MidSchoolName.AndWhereNeq(fval29.Interface())
				case ">":
					l.MidSchoolName.AndWhereGt(fval29.Interface())
				case ">=":
					l.MidSchoolName.AndWhereGte(fval29.Interface())
				case "<":
					l.MidSchoolName.AndWhereLt(fval29.Interface())
				case "<=":
					l.MidSchoolName.AndWhereLte(fval29.Interface())
				case "llike":
					l.MidSchoolName.AndWhereLike(fmt.Sprintf("%%%s", fval29.String()))
				case "rlike":
					l.MidSchoolName.AndWhereLike(fmt.Sprintf("%s%%", fval29.String()))
				case "alike":
					l.MidSchoolName.AndWhereLike(fmt.Sprintf("%%%s%%", fval29.String()))
				case "nllike":
					l.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%%%s", fval29.String()))
				case "nrlike":
					l.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%s%%", fval29.String()))
				case "nalike":
					l.MidSchoolName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval29.String()))
				case "in":
					l.MidSchoolName.AndWhereIn(fval29.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop29)
				}
			}
		}
	}
	ftyp30, exists := typ.FieldByName("Referee")
	if exists {
		fval30 := val.FieldByName("Referee")
		fop30, ok := ftyp30.Tag.Lookup("op")
		for fval30.Kind() == reflect.Ptr && !fval30.IsNil() {
			fval30 = fval30.Elem()
		}
		if fval30.Kind() != reflect.Ptr {
			if !ok {
				l.Referee.AndWhereEq(fval30.Interface())
			} else {
				switch fop30 {
				case "=":
					l.Referee.AndWhereEq(fval30.Interface())
				case "!=":
					l.Referee.AndWhereNeq(fval30.Interface())
				case ">":
					l.Referee.AndWhereGt(fval30.Interface())
				case ">=":
					l.Referee.AndWhereGte(fval30.Interface())
				case "<":
					l.Referee.AndWhereLt(fval30.Interface())
				case "<=":
					l.Referee.AndWhereLte(fval30.Interface())
				case "llike":
					l.Referee.AndWhereLike(fmt.Sprintf("%%%s", fval30.String()))
				case "rlike":
					l.Referee.AndWhereLike(fmt.Sprintf("%s%%", fval30.String()))
				case "alike":
					l.Referee.AndWhereLike(fmt.Sprintf("%%%s%%", fval30.String()))
				case "nllike":
					l.Referee.AndWhereNotLike(fmt.Sprintf("%%%s", fval30.String()))
				case "nrlike":
					l.Referee.AndWhereNotLike(fmt.Sprintf("%s%%", fval30.String()))
				case "nalike":
					l.Referee.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval30.String()))
				case "in":
					l.Referee.AndWhereIn(fval30.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop30)
				}
			}
		}
	}
	ftyp31, exists := typ.FieldByName("RefereeDuty")
	if exists {
		fval31 := val.FieldByName("RefereeDuty")
		fop31, ok := ftyp31.Tag.Lookup("op")
		for fval31.Kind() == reflect.Ptr && !fval31.IsNil() {
			fval31 = fval31.Elem()
		}
		if fval31.Kind() != reflect.Ptr {
			if !ok {
				l.RefereeDuty.AndWhereEq(fval31.Interface())
			} else {
				switch fop31 {
				case "=":
					l.RefereeDuty.AndWhereEq(fval31.Interface())
				case "!=":
					l.RefereeDuty.AndWhereNeq(fval31.Interface())
				case ">":
					l.RefereeDuty.AndWhereGt(fval31.Interface())
				case ">=":
					l.RefereeDuty.AndWhereGte(fval31.Interface())
				case "<":
					l.RefereeDuty.AndWhereLt(fval31.Interface())
				case "<=":
					l.RefereeDuty.AndWhereLte(fval31.Interface())
				case "llike":
					l.RefereeDuty.AndWhereLike(fmt.Sprintf("%%%s", fval31.String()))
				case "rlike":
					l.RefereeDuty.AndWhereLike(fmt.Sprintf("%s%%", fval31.String()))
				case "alike":
					l.RefereeDuty.AndWhereLike(fmt.Sprintf("%%%s%%", fval31.String()))
				case "nllike":
					l.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%%%s", fval31.String()))
				case "nrlike":
					l.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%s%%", fval31.String()))
				case "nalike":
					l.RefereeDuty.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval31.String()))
				case "in":
					l.RefereeDuty.AndWhereIn(fval31.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop31)
				}
			}
		}
	}
	ftyp32, exists := typ.FieldByName("RefereePhone")
	if exists {
		fval32 := val.FieldByName("RefereePhone")
		fop32, ok := ftyp32.Tag.Lookup("op")
		for fval32.Kind() == reflect.Ptr && !fval32.IsNil() {
			fval32 = fval32.Elem()
		}
		if fval32.Kind() != reflect.Ptr {
			if !ok {
				l.RefereePhone.AndWhereEq(fval32.Interface())
			} else {
				switch fop32 {
				case "=":
					l.RefereePhone.AndWhereEq(fval32.Interface())
				case "!=":
					l.RefereePhone.AndWhereNeq(fval32.Interface())
				case ">":
					l.RefereePhone.AndWhereGt(fval32.Interface())
				case ">=":
					l.RefereePhone.AndWhereGte(fval32.Interface())
				case "<":
					l.RefereePhone.AndWhereLt(fval32.Interface())
				case "<=":
					l.RefereePhone.AndWhereLte(fval32.Interface())
				case "llike":
					l.RefereePhone.AndWhereLike(fmt.Sprintf("%%%s", fval32.String()))
				case "rlike":
					l.RefereePhone.AndWhereLike(fmt.Sprintf("%s%%", fval32.String()))
				case "alike":
					l.RefereePhone.AndWhereLike(fmt.Sprintf("%%%s%%", fval32.String()))
				case "nllike":
					l.RefereePhone.AndWhereNotLike(fmt.Sprintf("%%%s", fval32.String()))
				case "nrlike":
					l.RefereePhone.AndWhereNotLike(fmt.Sprintf("%s%%", fval32.String()))
				case "nalike":
					l.RefereePhone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval32.String()))
				case "in":
					l.RefereePhone.AndWhereIn(fval32.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop32)
				}
			}
		}
	}
	ftyp33, exists := typ.FieldByName("AdmissionTicketNo")
	if exists {
		fval33 := val.FieldByName("AdmissionTicketNo")
		fop33, ok := ftyp33.Tag.Lookup("op")
		for fval33.Kind() == reflect.Ptr && !fval33.IsNil() {
			fval33 = fval33.Elem()
		}
		if fval33.Kind() != reflect.Ptr {
			if !ok {
				l.AdmissionTicketNo.AndWhereEq(fval33.Interface())
			} else {
				switch fop33 {
				case "=":
					l.AdmissionTicketNo.AndWhereEq(fval33.Interface())
				case "!=":
					l.AdmissionTicketNo.AndWhereNeq(fval33.Interface())
				case ">":
					l.AdmissionTicketNo.AndWhereGt(fval33.Interface())
				case ">=":
					l.AdmissionTicketNo.AndWhereGte(fval33.Interface())
				case "<":
					l.AdmissionTicketNo.AndWhereLt(fval33.Interface())
				case "<=":
					l.AdmissionTicketNo.AndWhereLte(fval33.Interface())
				case "llike":
					l.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%%%s", fval33.String()))
				case "rlike":
					l.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%s%%", fval33.String()))
				case "alike":
					l.AdmissionTicketNo.AndWhereLike(fmt.Sprintf("%%%s%%", fval33.String()))
				case "nllike":
					l.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%%%s", fval33.String()))
				case "nrlike":
					l.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%s%%", fval33.String()))
				case "nalike":
					l.AdmissionTicketNo.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval33.String()))
				case "in":
					l.AdmissionTicketNo.AndWhereIn(fval33.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop33)
				}
			}
		}
	}
	ftyp34, exists := typ.FieldByName("CollegeEntranceExamScores")
	if exists {
		fval34 := val.FieldByName("CollegeEntranceExamScores")
		fop34, ok := ftyp34.Tag.Lookup("op")
		for fval34.Kind() == reflect.Ptr && !fval34.IsNil() {
			fval34 = fval34.Elem()
		}
		if fval34.Kind() != reflect.Ptr {
			if !ok {
				l.CollegeEntranceExamScores.AndWhereEq(fval34.Interface())
			} else {
				switch fop34 {
				case "=":
					l.CollegeEntranceExamScores.AndWhereEq(fval34.Interface())
				case "!=":
					l.CollegeEntranceExamScores.AndWhereNeq(fval34.Interface())
				case ">":
					l.CollegeEntranceExamScores.AndWhereGt(fval34.Interface())
				case ">=":
					l.CollegeEntranceExamScores.AndWhereGte(fval34.Interface())
				case "<":
					l.CollegeEntranceExamScores.AndWhereLt(fval34.Interface())
				case "<=":
					l.CollegeEntranceExamScores.AndWhereLte(fval34.Interface())
				case "llike":
					l.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%%%s", fval34.String()))
				case "rlike":
					l.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%s%%", fval34.String()))
				case "alike":
					l.CollegeEntranceExamScores.AndWhereLike(fmt.Sprintf("%%%s%%", fval34.String()))
				case "nllike":
					l.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%%%s", fval34.String()))
				case "nrlike":
					l.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%s%%", fval34.String()))
				case "nalike":
					l.CollegeEntranceExamScores.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval34.String()))
				case "in":
					l.CollegeEntranceExamScores.AndWhereIn(fval34.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop34)
				}
			}
		}
	}
	ftyp35, exists := typ.FieldByName("AdmissionYear")
	if exists {
		fval35 := val.FieldByName("AdmissionYear")
		fop35, ok := ftyp35.Tag.Lookup("op")
		for fval35.Kind() == reflect.Ptr && !fval35.IsNil() {
			fval35 = fval35.Elem()
		}
		if fval35.Kind() != reflect.Ptr {
			if !ok {
				l.AdmissionYear.AndWhereEq(fval35.Interface())
			} else {
				switch fop35 {
				case "=":
					l.AdmissionYear.AndWhereEq(fval35.Interface())
				case "!=":
					l.AdmissionYear.AndWhereNeq(fval35.Interface())
				case ">":
					l.AdmissionYear.AndWhereGt(fval35.Interface())
				case ">=":
					l.AdmissionYear.AndWhereGte(fval35.Interface())
				case "<":
					l.AdmissionYear.AndWhereLt(fval35.Interface())
				case "<=":
					l.AdmissionYear.AndWhereLte(fval35.Interface())
				case "llike":
					l.AdmissionYear.AndWhereLike(fmt.Sprintf("%%%s", fval35.String()))
				case "rlike":
					l.AdmissionYear.AndWhereLike(fmt.Sprintf("%s%%", fval35.String()))
				case "alike":
					l.AdmissionYear.AndWhereLike(fmt.Sprintf("%%%s%%", fval35.String()))
				case "nllike":
					l.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%%%s", fval35.String()))
				case "nrlike":
					l.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%s%%", fval35.String()))
				case "nalike":
					l.AdmissionYear.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval35.String()))
				case "in":
					l.AdmissionYear.AndWhereIn(fval35.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop35)
				}
			}
		}
	}
	ftyp36, exists := typ.FieldByName("ForeignLanguageCode")
	if exists {
		fval36 := val.FieldByName("ForeignLanguageCode")
		fop36, ok := ftyp36.Tag.Lookup("op")
		for fval36.Kind() == reflect.Ptr && !fval36.IsNil() {
			fval36 = fval36.Elem()
		}
		if fval36.Kind() != reflect.Ptr {
			if !ok {
				l.ForeignLanguageCode.AndWhereEq(fval36.Interface())
			} else {
				switch fop36 {
				case "=":
					l.ForeignLanguageCode.AndWhereEq(fval36.Interface())
				case "!=":
					l.ForeignLanguageCode.AndWhereNeq(fval36.Interface())
				case ">":
					l.ForeignLanguageCode.AndWhereGt(fval36.Interface())
				case ">=":
					l.ForeignLanguageCode.AndWhereGte(fval36.Interface())
				case "<":
					l.ForeignLanguageCode.AndWhereLt(fval36.Interface())
				case "<=":
					l.ForeignLanguageCode.AndWhereLte(fval36.Interface())
				case "llike":
					l.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%%%s", fval36.String()))
				case "rlike":
					l.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%s%%", fval36.String()))
				case "alike":
					l.ForeignLanguageCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval36.String()))
				case "nllike":
					l.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval36.String()))
				case "nrlike":
					l.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval36.String()))
				case "nalike":
					l.ForeignLanguageCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval36.String()))
				case "in":
					l.ForeignLanguageCode.AndWhereIn(fval36.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop36)
				}
			}
		}
	}
	ftyp37, exists := typ.FieldByName("StudentOrigin")
	if exists {
		fval37 := val.FieldByName("StudentOrigin")
		fop37, ok := ftyp37.Tag.Lookup("op")
		for fval37.Kind() == reflect.Ptr && !fval37.IsNil() {
			fval37 = fval37.Elem()
		}
		if fval37.Kind() != reflect.Ptr {
			if !ok {
				l.StudentOrigin.AndWhereEq(fval37.Interface())
			} else {
				switch fop37 {
				case "=":
					l.StudentOrigin.AndWhereEq(fval37.Interface())
				case "!=":
					l.StudentOrigin.AndWhereNeq(fval37.Interface())
				case ">":
					l.StudentOrigin.AndWhereGt(fval37.Interface())
				case ">=":
					l.StudentOrigin.AndWhereGte(fval37.Interface())
				case "<":
					l.StudentOrigin.AndWhereLt(fval37.Interface())
				case "<=":
					l.StudentOrigin.AndWhereLte(fval37.Interface())
				case "llike":
					l.StudentOrigin.AndWhereLike(fmt.Sprintf("%%%s", fval37.String()))
				case "rlike":
					l.StudentOrigin.AndWhereLike(fmt.Sprintf("%s%%", fval37.String()))
				case "alike":
					l.StudentOrigin.AndWhereLike(fmt.Sprintf("%%%s%%", fval37.String()))
				case "nllike":
					l.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%%%s", fval37.String()))
				case "nrlike":
					l.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%s%%", fval37.String()))
				case "nalike":
					l.StudentOrigin.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval37.String()))
				case "in":
					l.StudentOrigin.AndWhereIn(fval37.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop37)
				}
			}
		}
	}
	ftyp38, exists := typ.FieldByName("BizType")
	if exists {
		fval38 := val.FieldByName("BizType")
		fop38, ok := ftyp38.Tag.Lookup("op")
		for fval38.Kind() == reflect.Ptr && !fval38.IsNil() {
			fval38 = fval38.Elem()
		}
		if fval38.Kind() != reflect.Ptr {
			if !ok {
				l.BizType.AndWhereEq(fval38.Interface())
			} else {
				switch fop38 {
				case "=":
					l.BizType.AndWhereEq(fval38.Interface())
				case "!=":
					l.BizType.AndWhereNeq(fval38.Interface())
				case ">":
					l.BizType.AndWhereGt(fval38.Interface())
				case ">=":
					l.BizType.AndWhereGte(fval38.Interface())
				case "<":
					l.BizType.AndWhereLt(fval38.Interface())
				case "<=":
					l.BizType.AndWhereLte(fval38.Interface())
				case "llike":
					l.BizType.AndWhereLike(fmt.Sprintf("%%%s", fval38.String()))
				case "rlike":
					l.BizType.AndWhereLike(fmt.Sprintf("%s%%", fval38.String()))
				case "alike":
					l.BizType.AndWhereLike(fmt.Sprintf("%%%s%%", fval38.String()))
				case "nllike":
					l.BizType.AndWhereNotLike(fmt.Sprintf("%%%s", fval38.String()))
				case "nrlike":
					l.BizType.AndWhereNotLike(fmt.Sprintf("%s%%", fval38.String()))
				case "nalike":
					l.BizType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval38.String()))
				case "in":
					l.BizType.AndWhereIn(fval38.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop38)
				}
			}
		}
	}
	ftyp39, exists := typ.FieldByName("TaskCode")
	if exists {
		fval39 := val.FieldByName("TaskCode")
		fop39, ok := ftyp39.Tag.Lookup("op")
		for fval39.Kind() == reflect.Ptr && !fval39.IsNil() {
			fval39 = fval39.Elem()
		}
		if fval39.Kind() != reflect.Ptr {
			if !ok {
				l.TaskCode.AndWhereEq(fval39.Interface())
			} else {
				switch fop39 {
				case "=":
					l.TaskCode.AndWhereEq(fval39.Interface())
				case "!=":
					l.TaskCode.AndWhereNeq(fval39.Interface())
				case ">":
					l.TaskCode.AndWhereGt(fval39.Interface())
				case ">=":
					l.TaskCode.AndWhereGte(fval39.Interface())
				case "<":
					l.TaskCode.AndWhereLt(fval39.Interface())
				case "<=":
					l.TaskCode.AndWhereLte(fval39.Interface())
				case "llike":
					l.TaskCode.AndWhereLike(fmt.Sprintf("%%%s", fval39.String()))
				case "rlike":
					l.TaskCode.AndWhereLike(fmt.Sprintf("%s%%", fval39.String()))
				case "alike":
					l.TaskCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval39.String()))
				case "nllike":
					l.TaskCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval39.String()))
				case "nrlike":
					l.TaskCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval39.String()))
				case "nalike":
					l.TaskCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval39.String()))
				case "in":
					l.TaskCode.AndWhereIn(fval39.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop39)
				}
			}
		}
	}
	ftyp40, exists := typ.FieldByName("ApproveStatus")
	if exists {
		fval40 := val.FieldByName("ApproveStatus")
		fop40, ok := ftyp40.Tag.Lookup("op")
		for fval40.Kind() == reflect.Ptr && !fval40.IsNil() {
			fval40 = fval40.Elem()
		}
		if fval40.Kind() != reflect.Ptr {
			if !ok {
				l.ApproveStatus.AndWhereEq(fval40.Interface())
			} else {
				switch fop40 {
				case "=":
					l.ApproveStatus.AndWhereEq(fval40.Interface())
				case "!=":
					l.ApproveStatus.AndWhereNeq(fval40.Interface())
				case ">":
					l.ApproveStatus.AndWhereGt(fval40.Interface())
				case ">=":
					l.ApproveStatus.AndWhereGte(fval40.Interface())
				case "<":
					l.ApproveStatus.AndWhereLt(fval40.Interface())
				case "<=":
					l.ApproveStatus.AndWhereLte(fval40.Interface())
				case "llike":
					l.ApproveStatus.AndWhereLike(fmt.Sprintf("%%%s", fval40.String()))
				case "rlike":
					l.ApproveStatus.AndWhereLike(fmt.Sprintf("%s%%", fval40.String()))
				case "alike":
					l.ApproveStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval40.String()))
				case "nllike":
					l.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval40.String()))
				case "nrlike":
					l.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval40.String()))
				case "nalike":
					l.ApproveStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval40.String()))
				case "in":
					l.ApproveStatus.AndWhereIn(fval40.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop40)
				}
			}
		}
	}
	ftyp41, exists := typ.FieldByName("Operator")
	if exists {
		fval41 := val.FieldByName("Operator")
		fop41, ok := ftyp41.Tag.Lookup("op")
		for fval41.Kind() == reflect.Ptr && !fval41.IsNil() {
			fval41 = fval41.Elem()
		}
		if fval41.Kind() != reflect.Ptr {
			if !ok {
				l.Operator.AndWhereEq(fval41.Interface())
			} else {
				switch fop41 {
				case "=":
					l.Operator.AndWhereEq(fval41.Interface())
				case "!=":
					l.Operator.AndWhereNeq(fval41.Interface())
				case ">":
					l.Operator.AndWhereGt(fval41.Interface())
				case ">=":
					l.Operator.AndWhereGte(fval41.Interface())
				case "<":
					l.Operator.AndWhereLt(fval41.Interface())
				case "<=":
					l.Operator.AndWhereLte(fval41.Interface())
				case "llike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval41.String()))
				case "rlike":
					l.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval41.String()))
				case "alike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval41.String()))
				case "nllike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval41.String()))
				case "nrlike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval41.String()))
				case "nalike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval41.String()))
				case "in":
					l.Operator.AndWhereIn(fval41.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop41)
				}
			}
		}
	}
	ftyp42, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval42 := val.FieldByName("InsertDatetime")
		fop42, ok := ftyp42.Tag.Lookup("op")
		for fval42.Kind() == reflect.Ptr && !fval42.IsNil() {
			fval42 = fval42.Elem()
		}
		if fval42.Kind() != reflect.Ptr {
			if !ok {
				l.InsertDatetime.AndWhereEq(fval42.Interface())
			} else {
				switch fop42 {
				case "=":
					l.InsertDatetime.AndWhereEq(fval42.Interface())
				case "!=":
					l.InsertDatetime.AndWhereNeq(fval42.Interface())
				case ">":
					l.InsertDatetime.AndWhereGt(fval42.Interface())
				case ">=":
					l.InsertDatetime.AndWhereGte(fval42.Interface())
				case "<":
					l.InsertDatetime.AndWhereLt(fval42.Interface())
				case "<=":
					l.InsertDatetime.AndWhereLte(fval42.Interface())
				case "llike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval42.String()))
				case "rlike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval42.String()))
				case "alike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval42.String()))
				case "nllike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval42.String()))
				case "nrlike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval42.String()))
				case "nalike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval42.String()))
				case "in":
					l.InsertDatetime.AndWhereIn(fval42.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop42)
				}
			}
		}
	}
	ftyp43, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval43 := val.FieldByName("UpdateDatetime")
		fop43, ok := ftyp43.Tag.Lookup("op")
		for fval43.Kind() == reflect.Ptr && !fval43.IsNil() {
			fval43 = fval43.Elem()
		}
		if fval43.Kind() != reflect.Ptr {
			if !ok {
				l.UpdateDatetime.AndWhereEq(fval43.Interface())
			} else {
				switch fop43 {
				case "=":
					l.UpdateDatetime.AndWhereEq(fval43.Interface())
				case "!=":
					l.UpdateDatetime.AndWhereNeq(fval43.Interface())
				case ">":
					l.UpdateDatetime.AndWhereGt(fval43.Interface())
				case ">=":
					l.UpdateDatetime.AndWhereGte(fval43.Interface())
				case "<":
					l.UpdateDatetime.AndWhereLt(fval43.Interface())
				case "<=":
					l.UpdateDatetime.AndWhereLte(fval43.Interface())
				case "llike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval43.String()))
				case "rlike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval43.String()))
				case "alike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval43.String()))
				case "nllike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval43.String()))
				case "nrlike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval43.String()))
				case "nalike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval43.String()))
				case "in":
					l.UpdateDatetime.AndWhereIn(fval43.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop43)
				}
			}
		}
	}
	ftyp44, exists := typ.FieldByName("Status")
	if exists {
		fval44 := val.FieldByName("Status")
		fop44, ok := ftyp44.Tag.Lookup("op")
		for fval44.Kind() == reflect.Ptr && !fval44.IsNil() {
			fval44 = fval44.Elem()
		}
		if fval44.Kind() != reflect.Ptr {
			if !ok {
				l.Status.AndWhereEq(fval44.Interface())
			} else {
				switch fop44 {
				case "=":
					l.Status.AndWhereEq(fval44.Interface())
				case "!=":
					l.Status.AndWhereNeq(fval44.Interface())
				case ">":
					l.Status.AndWhereGt(fval44.Interface())
				case ">=":
					l.Status.AndWhereGte(fval44.Interface())
				case "<":
					l.Status.AndWhereLt(fval44.Interface())
				case "<=":
					l.Status.AndWhereLte(fval44.Interface())
				case "llike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s", fval44.String()))
				case "rlike":
					l.Status.AndWhereLike(fmt.Sprintf("%s%%", fval44.String()))
				case "alike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval44.String()))
				case "nllike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval44.String()))
				case "nrlike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval44.String()))
				case "nalike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval44.String()))
				case "in":
					l.Status.AndWhereIn(fval44.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop44)
				}
			}
		}
	}
	ftyp45, exists := typ.FieldByName("StudentStatus")
	if exists {
		fval45 := val.FieldByName("StudentStatus")
		fop45, ok := ftyp45.Tag.Lookup("op")
		for fval45.Kind() == reflect.Ptr && !fval45.IsNil() {
			fval45 = fval45.Elem()
		}
		if fval45.Kind() != reflect.Ptr {
			if !ok {
				l.StudentStatus.AndWhereEq(fval45.Interface())
			} else {
				switch fop45 {
				case "=":
					l.StudentStatus.AndWhereEq(fval45.Interface())
				case "!=":
					l.StudentStatus.AndWhereNeq(fval45.Interface())
				case ">":
					l.StudentStatus.AndWhereGt(fval45.Interface())
				case ">=":
					l.StudentStatus.AndWhereGte(fval45.Interface())
				case "<":
					l.StudentStatus.AndWhereLt(fval45.Interface())
				case "<=":
					l.StudentStatus.AndWhereLte(fval45.Interface())
				case "llike":
					l.StudentStatus.AndWhereLike(fmt.Sprintf("%%%s", fval45.String()))
				case "rlike":
					l.StudentStatus.AndWhereLike(fmt.Sprintf("%s%%", fval45.String()))
				case "alike":
					l.StudentStatus.AndWhereLike(fmt.Sprintf("%%%s%%", fval45.String()))
				case "nllike":
					l.StudentStatus.AndWhereNotLike(fmt.Sprintf("%%%s", fval45.String()))
				case "nrlike":
					l.StudentStatus.AndWhereNotLike(fmt.Sprintf("%s%%", fval45.String()))
				case "nalike":
					l.StudentStatus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval45.String()))
				case "in":
					l.StudentStatus.AndWhereIn(fval45.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop45)
				}
			}
		}
	}
	ftyp46, exists := typ.FieldByName("IsAuth")
	if exists {
		fval46 := val.FieldByName("IsAuth")
		fop46, ok := ftyp46.Tag.Lookup("op")
		for fval46.Kind() == reflect.Ptr && !fval46.IsNil() {
			fval46 = fval46.Elem()
		}
		if fval46.Kind() != reflect.Ptr {
			if !ok {
				l.IsAuth.AndWhereEq(fval46.Interface())
			} else {
				switch fop46 {
				case "=":
					l.IsAuth.AndWhereEq(fval46.Interface())
				case "!=":
					l.IsAuth.AndWhereNeq(fval46.Interface())
				case ">":
					l.IsAuth.AndWhereGt(fval46.Interface())
				case ">=":
					l.IsAuth.AndWhereGte(fval46.Interface())
				case "<":
					l.IsAuth.AndWhereLt(fval46.Interface())
				case "<=":
					l.IsAuth.AndWhereLte(fval46.Interface())
				case "llike":
					l.IsAuth.AndWhereLike(fmt.Sprintf("%%%s", fval46.String()))
				case "rlike":
					l.IsAuth.AndWhereLike(fmt.Sprintf("%s%%", fval46.String()))
				case "alike":
					l.IsAuth.AndWhereLike(fmt.Sprintf("%%%s%%", fval46.String()))
				case "nllike":
					l.IsAuth.AndWhereNotLike(fmt.Sprintf("%%%s", fval46.String()))
				case "nrlike":
					l.IsAuth.AndWhereNotLike(fmt.Sprintf("%s%%", fval46.String()))
				case "nalike":
					l.IsAuth.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval46.String()))
				case "in":
					l.IsAuth.AndWhereIn(fval46.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop46)
				}
			}
		}
	}
	ftyp47, exists := typ.FieldByName("Campus")
	if exists {
		fval47 := val.FieldByName("Campus")
		fop47, ok := ftyp47.Tag.Lookup("op")
		for fval47.Kind() == reflect.Ptr && !fval47.IsNil() {
			fval47 = fval47.Elem()
		}
		if fval47.Kind() != reflect.Ptr {
			if !ok {
				l.Campus.AndWhereEq(fval47.Interface())
			} else {
				switch fop47 {
				case "=":
					l.Campus.AndWhereEq(fval47.Interface())
				case "!=":
					l.Campus.AndWhereNeq(fval47.Interface())
				case ">":
					l.Campus.AndWhereGt(fval47.Interface())
				case ">=":
					l.Campus.AndWhereGte(fval47.Interface())
				case "<":
					l.Campus.AndWhereLt(fval47.Interface())
				case "<=":
					l.Campus.AndWhereLte(fval47.Interface())
				case "llike":
					l.Campus.AndWhereLike(fmt.Sprintf("%%%s", fval47.String()))
				case "rlike":
					l.Campus.AndWhereLike(fmt.Sprintf("%s%%", fval47.String()))
				case "alike":
					l.Campus.AndWhereLike(fmt.Sprintf("%%%s%%", fval47.String()))
				case "nllike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%%%s", fval47.String()))
				case "nrlike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%s%%", fval47.String()))
				case "nalike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval47.String()))
				case "in":
					l.Campus.AndWhereIn(fval47.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop47)
				}
			}
		}
	}
	ftyp48, exists := typ.FieldByName("Zone")
	if exists {
		fval48 := val.FieldByName("Zone")
		fop48, ok := ftyp48.Tag.Lookup("op")
		for fval48.Kind() == reflect.Ptr && !fval48.IsNil() {
			fval48 = fval48.Elem()
		}
		if fval48.Kind() != reflect.Ptr {
			if !ok {
				l.Zone.AndWhereEq(fval48.Interface())
			} else {
				switch fop48 {
				case "=":
					l.Zone.AndWhereEq(fval48.Interface())
				case "!=":
					l.Zone.AndWhereNeq(fval48.Interface())
				case ">":
					l.Zone.AndWhereGt(fval48.Interface())
				case ">=":
					l.Zone.AndWhereGte(fval48.Interface())
				case "<":
					l.Zone.AndWhereLt(fval48.Interface())
				case "<=":
					l.Zone.AndWhereLte(fval48.Interface())
				case "llike":
					l.Zone.AndWhereLike(fmt.Sprintf("%%%s", fval48.String()))
				case "rlike":
					l.Zone.AndWhereLike(fmt.Sprintf("%s%%", fval48.String()))
				case "alike":
					l.Zone.AndWhereLike(fmt.Sprintf("%%%s%%", fval48.String()))
				case "nllike":
					l.Zone.AndWhereNotLike(fmt.Sprintf("%%%s", fval48.String()))
				case "nrlike":
					l.Zone.AndWhereNotLike(fmt.Sprintf("%s%%", fval48.String()))
				case "nalike":
					l.Zone.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval48.String()))
				case "in":
					l.Zone.AndWhereIn(fval48.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop48)
				}
			}
		}
	}
	ftyp49, exists := typ.FieldByName("Building")
	if exists {
		fval49 := val.FieldByName("Building")
		fop49, ok := ftyp49.Tag.Lookup("op")
		for fval49.Kind() == reflect.Ptr && !fval49.IsNil() {
			fval49 = fval49.Elem()
		}
		if fval49.Kind() != reflect.Ptr {
			if !ok {
				l.Building.AndWhereEq(fval49.Interface())
			} else {
				switch fop49 {
				case "=":
					l.Building.AndWhereEq(fval49.Interface())
				case "!=":
					l.Building.AndWhereNeq(fval49.Interface())
				case ">":
					l.Building.AndWhereGt(fval49.Interface())
				case ">=":
					l.Building.AndWhereGte(fval49.Interface())
				case "<":
					l.Building.AndWhereLt(fval49.Interface())
				case "<=":
					l.Building.AndWhereLte(fval49.Interface())
				case "llike":
					l.Building.AndWhereLike(fmt.Sprintf("%%%s", fval49.String()))
				case "rlike":
					l.Building.AndWhereLike(fmt.Sprintf("%s%%", fval49.String()))
				case "alike":
					l.Building.AndWhereLike(fmt.Sprintf("%%%s%%", fval49.String()))
				case "nllike":
					l.Building.AndWhereNotLike(fmt.Sprintf("%%%s", fval49.String()))
				case "nrlike":
					l.Building.AndWhereNotLike(fmt.Sprintf("%s%%", fval49.String()))
				case "nalike":
					l.Building.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval49.String()))
				case "in":
					l.Building.AndWhereIn(fval49.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop49)
				}
			}
		}
	}
	ftyp50, exists := typ.FieldByName("Unit")
	if exists {
		fval50 := val.FieldByName("Unit")
		fop50, ok := ftyp50.Tag.Lookup("op")
		for fval50.Kind() == reflect.Ptr && !fval50.IsNil() {
			fval50 = fval50.Elem()
		}
		if fval50.Kind() != reflect.Ptr {
			if !ok {
				l.Unit.AndWhereEq(fval50.Interface())
			} else {
				switch fop50 {
				case "=":
					l.Unit.AndWhereEq(fval50.Interface())
				case "!=":
					l.Unit.AndWhereNeq(fval50.Interface())
				case ">":
					l.Unit.AndWhereGt(fval50.Interface())
				case ">=":
					l.Unit.AndWhereGte(fval50.Interface())
				case "<":
					l.Unit.AndWhereLt(fval50.Interface())
				case "<=":
					l.Unit.AndWhereLte(fval50.Interface())
				case "llike":
					l.Unit.AndWhereLike(fmt.Sprintf("%%%s", fval50.String()))
				case "rlike":
					l.Unit.AndWhereLike(fmt.Sprintf("%s%%", fval50.String()))
				case "alike":
					l.Unit.AndWhereLike(fmt.Sprintf("%%%s%%", fval50.String()))
				case "nllike":
					l.Unit.AndWhereNotLike(fmt.Sprintf("%%%s", fval50.String()))
				case "nrlike":
					l.Unit.AndWhereNotLike(fmt.Sprintf("%s%%", fval50.String()))
				case "nalike":
					l.Unit.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval50.String()))
				case "in":
					l.Unit.AndWhereIn(fval50.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop50)
				}
			}
		}
	}
	ftyp51, exists := typ.FieldByName("Room")
	if exists {
		fval51 := val.FieldByName("Room")
		fop51, ok := ftyp51.Tag.Lookup("op")
		for fval51.Kind() == reflect.Ptr && !fval51.IsNil() {
			fval51 = fval51.Elem()
		}
		if fval51.Kind() != reflect.Ptr {
			if !ok {
				l.Room.AndWhereEq(fval51.Interface())
			} else {
				switch fop51 {
				case "=":
					l.Room.AndWhereEq(fval51.Interface())
				case "!=":
					l.Room.AndWhereNeq(fval51.Interface())
				case ">":
					l.Room.AndWhereGt(fval51.Interface())
				case ">=":
					l.Room.AndWhereGte(fval51.Interface())
				case "<":
					l.Room.AndWhereLt(fval51.Interface())
				case "<=":
					l.Room.AndWhereLte(fval51.Interface())
				case "llike":
					l.Room.AndWhereLike(fmt.Sprintf("%%%s", fval51.String()))
				case "rlike":
					l.Room.AndWhereLike(fmt.Sprintf("%s%%", fval51.String()))
				case "alike":
					l.Room.AndWhereLike(fmt.Sprintf("%%%s%%", fval51.String()))
				case "nllike":
					l.Room.AndWhereNotLike(fmt.Sprintf("%%%s", fval51.String()))
				case "nrlike":
					l.Room.AndWhereNotLike(fmt.Sprintf("%s%%", fval51.String()))
				case "nalike":
					l.Room.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval51.String()))
				case "in":
					l.Room.AndWhereIn(fval51.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop51)
				}
			}
		}
	}
	ftyp52, exists := typ.FieldByName("Bed")
	if exists {
		fval52 := val.FieldByName("Bed")
		fop52, ok := ftyp52.Tag.Lookup("op")
		for fval52.Kind() == reflect.Ptr && !fval52.IsNil() {
			fval52 = fval52.Elem()
		}
		if fval52.Kind() != reflect.Ptr {
			if !ok {
				l.Bed.AndWhereEq(fval52.Interface())
			} else {
				switch fop52 {
				case "=":
					l.Bed.AndWhereEq(fval52.Interface())
				case "!=":
					l.Bed.AndWhereNeq(fval52.Interface())
				case ">":
					l.Bed.AndWhereGt(fval52.Interface())
				case ">=":
					l.Bed.AndWhereGte(fval52.Interface())
				case "<":
					l.Bed.AndWhereLt(fval52.Interface())
				case "<=":
					l.Bed.AndWhereLte(fval52.Interface())
				case "llike":
					l.Bed.AndWhereLike(fmt.Sprintf("%%%s", fval52.String()))
				case "rlike":
					l.Bed.AndWhereLike(fmt.Sprintf("%s%%", fval52.String()))
				case "alike":
					l.Bed.AndWhereLike(fmt.Sprintf("%%%s%%", fval52.String()))
				case "nllike":
					l.Bed.AndWhereNotLike(fmt.Sprintf("%%%s", fval52.String()))
				case "nrlike":
					l.Bed.AndWhereNotLike(fmt.Sprintf("%s%%", fval52.String()))
				case "nalike":
					l.Bed.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval52.String()))
				case "in":
					l.Bed.AndWhereIn(fval52.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop52)
				}
			}
		}
	}
	ftyp53, exists := typ.FieldByName("StatusSort")
	if exists {
		fval53 := val.FieldByName("StatusSort")
		fop53, ok := ftyp53.Tag.Lookup("op")
		for fval53.Kind() == reflect.Ptr && !fval53.IsNil() {
			fval53 = fval53.Elem()
		}
		if fval53.Kind() != reflect.Ptr {
			if !ok {
				l.StatusSort.AndWhereEq(fval53.Interface())
			} else {
				switch fop53 {
				case "=":
					l.StatusSort.AndWhereEq(fval53.Interface())
				case "!=":
					l.StatusSort.AndWhereNeq(fval53.Interface())
				case ">":
					l.StatusSort.AndWhereGt(fval53.Interface())
				case ">=":
					l.StatusSort.AndWhereGte(fval53.Interface())
				case "<":
					l.StatusSort.AndWhereLt(fval53.Interface())
				case "<=":
					l.StatusSort.AndWhereLte(fval53.Interface())
				case "llike":
					l.StatusSort.AndWhereLike(fmt.Sprintf("%%%s", fval53.String()))
				case "rlike":
					l.StatusSort.AndWhereLike(fmt.Sprintf("%s%%", fval53.String()))
				case "alike":
					l.StatusSort.AndWhereLike(fmt.Sprintf("%%%s%%", fval53.String()))
				case "nllike":
					l.StatusSort.AndWhereNotLike(fmt.Sprintf("%%%s", fval53.String()))
				case "nrlike":
					l.StatusSort.AndWhereNotLike(fmt.Sprintf("%s%%", fval53.String()))
				case "nalike":
					l.StatusSort.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval53.String()))
				case "in":
					l.StatusSort.AndWhereIn(fval53.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop53)
				}
			}
		}
	}
	ftyp54, exists := typ.FieldByName("Height")
	if exists {
		fval54 := val.FieldByName("Height")
		fop54, ok := ftyp54.Tag.Lookup("op")
		for fval54.Kind() == reflect.Ptr && !fval54.IsNil() {
			fval54 = fval54.Elem()
		}
		if fval54.Kind() != reflect.Ptr {
			if !ok {
				l.Height.AndWhereEq(fval54.Interface())
			} else {
				switch fop54 {
				case "=":
					l.Height.AndWhereEq(fval54.Interface())
				case "!=":
					l.Height.AndWhereNeq(fval54.Interface())
				case ">":
					l.Height.AndWhereGt(fval54.Interface())
				case ">=":
					l.Height.AndWhereGte(fval54.Interface())
				case "<":
					l.Height.AndWhereLt(fval54.Interface())
				case "<=":
					l.Height.AndWhereLte(fval54.Interface())
				case "llike":
					l.Height.AndWhereLike(fmt.Sprintf("%%%s", fval54.String()))
				case "rlike":
					l.Height.AndWhereLike(fmt.Sprintf("%s%%", fval54.String()))
				case "alike":
					l.Height.AndWhereLike(fmt.Sprintf("%%%s%%", fval54.String()))
				case "nllike":
					l.Height.AndWhereNotLike(fmt.Sprintf("%%%s", fval54.String()))
				case "nrlike":
					l.Height.AndWhereNotLike(fmt.Sprintf("%s%%", fval54.String()))
				case "nalike":
					l.Height.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval54.String()))
				case "in":
					l.Height.AndWhereIn(fval54.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop54)
				}
			}
		}
	}
	ftyp55, exists := typ.FieldByName("Weight")
	if exists {
		fval55 := val.FieldByName("Weight")
		fop55, ok := ftyp55.Tag.Lookup("op")
		for fval55.Kind() == reflect.Ptr && !fval55.IsNil() {
			fval55 = fval55.Elem()
		}
		if fval55.Kind() != reflect.Ptr {
			if !ok {
				l.Weight.AndWhereEq(fval55.Interface())
			} else {
				switch fop55 {
				case "=":
					l.Weight.AndWhereEq(fval55.Interface())
				case "!=":
					l.Weight.AndWhereNeq(fval55.Interface())
				case ">":
					l.Weight.AndWhereGt(fval55.Interface())
				case ">=":
					l.Weight.AndWhereGte(fval55.Interface())
				case "<":
					l.Weight.AndWhereLt(fval55.Interface())
				case "<=":
					l.Weight.AndWhereLte(fval55.Interface())
				case "llike":
					l.Weight.AndWhereLike(fmt.Sprintf("%%%s", fval55.String()))
				case "rlike":
					l.Weight.AndWhereLike(fmt.Sprintf("%s%%", fval55.String()))
				case "alike":
					l.Weight.AndWhereLike(fmt.Sprintf("%%%s%%", fval55.String()))
				case "nllike":
					l.Weight.AndWhereNotLike(fmt.Sprintf("%%%s", fval55.String()))
				case "nrlike":
					l.Weight.AndWhereNotLike(fmt.Sprintf("%s%%", fval55.String()))
				case "nalike":
					l.Weight.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval55.String()))
				case "in":
					l.Weight.AndWhereIn(fval55.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop55)
				}
			}
		}
	}
	ftyp56, exists := typ.FieldByName("FootSize")
	if exists {
		fval56 := val.FieldByName("FootSize")
		fop56, ok := ftyp56.Tag.Lookup("op")
		for fval56.Kind() == reflect.Ptr && !fval56.IsNil() {
			fval56 = fval56.Elem()
		}
		if fval56.Kind() != reflect.Ptr {
			if !ok {
				l.FootSize.AndWhereEq(fval56.Interface())
			} else {
				switch fop56 {
				case "=":
					l.FootSize.AndWhereEq(fval56.Interface())
				case "!=":
					l.FootSize.AndWhereNeq(fval56.Interface())
				case ">":
					l.FootSize.AndWhereGt(fval56.Interface())
				case ">=":
					l.FootSize.AndWhereGte(fval56.Interface())
				case "<":
					l.FootSize.AndWhereLt(fval56.Interface())
				case "<=":
					l.FootSize.AndWhereLte(fval56.Interface())
				case "llike":
					l.FootSize.AndWhereLike(fmt.Sprintf("%%%s", fval56.String()))
				case "rlike":
					l.FootSize.AndWhereLike(fmt.Sprintf("%s%%", fval56.String()))
				case "alike":
					l.FootSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval56.String()))
				case "nllike":
					l.FootSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval56.String()))
				case "nrlike":
					l.FootSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval56.String()))
				case "nalike":
					l.FootSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval56.String()))
				case "in":
					l.FootSize.AndWhereIn(fval56.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop56)
				}
			}
		}
	}
	ftyp57, exists := typ.FieldByName("ClothSize")
	if exists {
		fval57 := val.FieldByName("ClothSize")
		fop57, ok := ftyp57.Tag.Lookup("op")
		for fval57.Kind() == reflect.Ptr && !fval57.IsNil() {
			fval57 = fval57.Elem()
		}
		if fval57.Kind() != reflect.Ptr {
			if !ok {
				l.ClothSize.AndWhereEq(fval57.Interface())
			} else {
				switch fop57 {
				case "=":
					l.ClothSize.AndWhereEq(fval57.Interface())
				case "!=":
					l.ClothSize.AndWhereNeq(fval57.Interface())
				case ">":
					l.ClothSize.AndWhereGt(fval57.Interface())
				case ">=":
					l.ClothSize.AndWhereGte(fval57.Interface())
				case "<":
					l.ClothSize.AndWhereLt(fval57.Interface())
				case "<=":
					l.ClothSize.AndWhereLte(fval57.Interface())
				case "llike":
					l.ClothSize.AndWhereLike(fmt.Sprintf("%%%s", fval57.String()))
				case "rlike":
					l.ClothSize.AndWhereLike(fmt.Sprintf("%s%%", fval57.String()))
				case "alike":
					l.ClothSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval57.String()))
				case "nllike":
					l.ClothSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval57.String()))
				case "nrlike":
					l.ClothSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval57.String()))
				case "nalike":
					l.ClothSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval57.String()))
				case "in":
					l.ClothSize.AndWhereIn(fval57.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop57)
				}
			}
		}
	}
	ftyp58, exists := typ.FieldByName("HeadSize")
	if exists {
		fval58 := val.FieldByName("HeadSize")
		fop58, ok := ftyp58.Tag.Lookup("op")
		for fval58.Kind() == reflect.Ptr && !fval58.IsNil() {
			fval58 = fval58.Elem()
		}
		if fval58.Kind() != reflect.Ptr {
			if !ok {
				l.HeadSize.AndWhereEq(fval58.Interface())
			} else {
				switch fop58 {
				case "=":
					l.HeadSize.AndWhereEq(fval58.Interface())
				case "!=":
					l.HeadSize.AndWhereNeq(fval58.Interface())
				case ">":
					l.HeadSize.AndWhereGt(fval58.Interface())
				case ">=":
					l.HeadSize.AndWhereGte(fval58.Interface())
				case "<":
					l.HeadSize.AndWhereLt(fval58.Interface())
				case "<=":
					l.HeadSize.AndWhereLte(fval58.Interface())
				case "llike":
					l.HeadSize.AndWhereLike(fmt.Sprintf("%%%s", fval58.String()))
				case "rlike":
					l.HeadSize.AndWhereLike(fmt.Sprintf("%s%%", fval58.String()))
				case "alike":
					l.HeadSize.AndWhereLike(fmt.Sprintf("%%%s%%", fval58.String()))
				case "nllike":
					l.HeadSize.AndWhereNotLike(fmt.Sprintf("%%%s", fval58.String()))
				case "nrlike":
					l.HeadSize.AndWhereNotLike(fmt.Sprintf("%s%%", fval58.String()))
				case "nalike":
					l.HeadSize.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval58.String()))
				case "in":
					l.HeadSize.AndWhereIn(fval58.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop58)
				}
			}
		}
	}
	ftyp59, exists := typ.FieldByName("Remark1")
	if exists {
		fval59 := val.FieldByName("Remark1")
		fop59, ok := ftyp59.Tag.Lookup("op")
		for fval59.Kind() == reflect.Ptr && !fval59.IsNil() {
			fval59 = fval59.Elem()
		}
		if fval59.Kind() != reflect.Ptr {
			if !ok {
				l.Remark1.AndWhereEq(fval59.Interface())
			} else {
				switch fop59 {
				case "=":
					l.Remark1.AndWhereEq(fval59.Interface())
				case "!=":
					l.Remark1.AndWhereNeq(fval59.Interface())
				case ">":
					l.Remark1.AndWhereGt(fval59.Interface())
				case ">=":
					l.Remark1.AndWhereGte(fval59.Interface())
				case "<":
					l.Remark1.AndWhereLt(fval59.Interface())
				case "<=":
					l.Remark1.AndWhereLte(fval59.Interface())
				case "llike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval59.String()))
				case "rlike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval59.String()))
				case "alike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval59.String()))
				case "nllike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval59.String()))
				case "nrlike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval59.String()))
				case "nalike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval59.String()))
				case "in":
					l.Remark1.AndWhereIn(fval59.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop59)
				}
			}
		}
	}
	ftyp60, exists := typ.FieldByName("Remark2")
	if exists {
		fval60 := val.FieldByName("Remark2")
		fop60, ok := ftyp60.Tag.Lookup("op")
		for fval60.Kind() == reflect.Ptr && !fval60.IsNil() {
			fval60 = fval60.Elem()
		}
		if fval60.Kind() != reflect.Ptr {
			if !ok {
				l.Remark2.AndWhereEq(fval60.Interface())
			} else {
				switch fop60 {
				case "=":
					l.Remark2.AndWhereEq(fval60.Interface())
				case "!=":
					l.Remark2.AndWhereNeq(fval60.Interface())
				case ">":
					l.Remark2.AndWhereGt(fval60.Interface())
				case ">=":
					l.Remark2.AndWhereGte(fval60.Interface())
				case "<":
					l.Remark2.AndWhereLt(fval60.Interface())
				case "<=":
					l.Remark2.AndWhereLte(fval60.Interface())
				case "llike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval60.String()))
				case "rlike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval60.String()))
				case "alike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval60.String()))
				case "nllike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval60.String()))
				case "nrlike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval60.String()))
				case "nalike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval60.String()))
				case "in":
					l.Remark2.AndWhereIn(fval60.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop60)
				}
			}
		}
	}
	ftyp61, exists := typ.FieldByName("Remark3")
	if exists {
		fval61 := val.FieldByName("Remark3")
		fop61, ok := ftyp61.Tag.Lookup("op")
		for fval61.Kind() == reflect.Ptr && !fval61.IsNil() {
			fval61 = fval61.Elem()
		}
		if fval61.Kind() != reflect.Ptr {
			if !ok {
				l.Remark3.AndWhereEq(fval61.Interface())
			} else {
				switch fop61 {
				case "=":
					l.Remark3.AndWhereEq(fval61.Interface())
				case "!=":
					l.Remark3.AndWhereNeq(fval61.Interface())
				case ">":
					l.Remark3.AndWhereGt(fval61.Interface())
				case ">=":
					l.Remark3.AndWhereGte(fval61.Interface())
				case "<":
					l.Remark3.AndWhereLt(fval61.Interface())
				case "<=":
					l.Remark3.AndWhereLte(fval61.Interface())
				case "llike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval61.String()))
				case "rlike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval61.String()))
				case "alike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval61.String()))
				case "nllike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval61.String()))
				case "nrlike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval61.String()))
				case "nalike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval61.String()))
				case "in":
					l.Remark3.AndWhereIn(fval61.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop61)
				}
			}
		}
	}
	ftyp62, exists := typ.FieldByName("Remark4")
	if exists {
		fval62 := val.FieldByName("Remark4")
		fop62, ok := ftyp62.Tag.Lookup("op")
		for fval62.Kind() == reflect.Ptr && !fval62.IsNil() {
			fval62 = fval62.Elem()
		}
		if fval62.Kind() != reflect.Ptr {
			if !ok {
				l.Remark4.AndWhereEq(fval62.Interface())
			} else {
				switch fop62 {
				case "=":
					l.Remark4.AndWhereEq(fval62.Interface())
				case "!=":
					l.Remark4.AndWhereNeq(fval62.Interface())
				case ">":
					l.Remark4.AndWhereGt(fval62.Interface())
				case ">=":
					l.Remark4.AndWhereGte(fval62.Interface())
				case "<":
					l.Remark4.AndWhereLt(fval62.Interface())
				case "<=":
					l.Remark4.AndWhereLte(fval62.Interface())
				case "llike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval62.String()))
				case "rlike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval62.String()))
				case "alike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval62.String()))
				case "nllike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval62.String()))
				case "nrlike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval62.String()))
				case "nalike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval62.String()))
				case "in":
					l.Remark4.AndWhereIn(fval62.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop62)
				}
			}
		}
	}
	ftyp63, exists := typ.FieldByName("IsPayment")
	if exists {
		fval63 := val.FieldByName("IsPayment")
		fop63, ok := ftyp63.Tag.Lookup("op")
		for fval63.Kind() == reflect.Ptr && !fval63.IsNil() {
			fval63 = fval63.Elem()
		}
		if fval63.Kind() != reflect.Ptr {
			if !ok {
				l.IsPayment.AndWhereEq(fval63.Interface())
			} else {
				switch fop63 {
				case "=":
					l.IsPayment.AndWhereEq(fval63.Interface())
				case "!=":
					l.IsPayment.AndWhereNeq(fval63.Interface())
				case ">":
					l.IsPayment.AndWhereGt(fval63.Interface())
				case ">=":
					l.IsPayment.AndWhereGte(fval63.Interface())
				case "<":
					l.IsPayment.AndWhereLt(fval63.Interface())
				case "<=":
					l.IsPayment.AndWhereLte(fval63.Interface())
				case "llike":
					l.IsPayment.AndWhereLike(fmt.Sprintf("%%%s", fval63.String()))
				case "rlike":
					l.IsPayment.AndWhereLike(fmt.Sprintf("%s%%", fval63.String()))
				case "alike":
					l.IsPayment.AndWhereLike(fmt.Sprintf("%%%s%%", fval63.String()))
				case "nllike":
					l.IsPayment.AndWhereNotLike(fmt.Sprintf("%%%s", fval63.String()))
				case "nrlike":
					l.IsPayment.AndWhereNotLike(fmt.Sprintf("%s%%", fval63.String()))
				case "nalike":
					l.IsPayment.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval63.String()))
				case "in":
					l.IsPayment.AndWhereIn(fval63.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop63)
				}
			}
		}
	}
	ftyp64, exists := typ.FieldByName("IsCheckIn")
	if exists {
		fval64 := val.FieldByName("IsCheckIn")
		fop64, ok := ftyp64.Tag.Lookup("op")
		for fval64.Kind() == reflect.Ptr && !fval64.IsNil() {
			fval64 = fval64.Elem()
		}
		if fval64.Kind() != reflect.Ptr {
			if !ok {
				l.IsCheckIn.AndWhereEq(fval64.Interface())
			} else {
				switch fop64 {
				case "=":
					l.IsCheckIn.AndWhereEq(fval64.Interface())
				case "!=":
					l.IsCheckIn.AndWhereNeq(fval64.Interface())
				case ">":
					l.IsCheckIn.AndWhereGt(fval64.Interface())
				case ">=":
					l.IsCheckIn.AndWhereGte(fval64.Interface())
				case "<":
					l.IsCheckIn.AndWhereLt(fval64.Interface())
				case "<=":
					l.IsCheckIn.AndWhereLte(fval64.Interface())
				case "llike":
					l.IsCheckIn.AndWhereLike(fmt.Sprintf("%%%s", fval64.String()))
				case "rlike":
					l.IsCheckIn.AndWhereLike(fmt.Sprintf("%s%%", fval64.String()))
				case "alike":
					l.IsCheckIn.AndWhereLike(fmt.Sprintf("%%%s%%", fval64.String()))
				case "nllike":
					l.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%%%s", fval64.String()))
				case "nrlike":
					l.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%s%%", fval64.String()))
				case "nalike":
					l.IsCheckIn.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval64.String()))
				case "in":
					l.IsCheckIn.AndWhereIn(fval64.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop64)
				}
			}
		}
	}
	ftyp65, exists := typ.FieldByName("GetMilitaryTC")
	if exists {
		fval65 := val.FieldByName("GetMilitaryTC")
		fop65, ok := ftyp65.Tag.Lookup("op")
		for fval65.Kind() == reflect.Ptr && !fval65.IsNil() {
			fval65 = fval65.Elem()
		}
		if fval65.Kind() != reflect.Ptr {
			if !ok {
				l.GetMilitaryTC.AndWhereEq(fval65.Interface())
			} else {
				switch fop65 {
				case "=":
					l.GetMilitaryTC.AndWhereEq(fval65.Interface())
				case "!=":
					l.GetMilitaryTC.AndWhereNeq(fval65.Interface())
				case ">":
					l.GetMilitaryTC.AndWhereGt(fval65.Interface())
				case ">=":
					l.GetMilitaryTC.AndWhereGte(fval65.Interface())
				case "<":
					l.GetMilitaryTC.AndWhereLt(fval65.Interface())
				case "<=":
					l.GetMilitaryTC.AndWhereLte(fval65.Interface())
				case "llike":
					l.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%%%s", fval65.String()))
				case "rlike":
					l.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%s%%", fval65.String()))
				case "alike":
					l.GetMilitaryTC.AndWhereLike(fmt.Sprintf("%%%s%%", fval65.String()))
				case "nllike":
					l.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%%%s", fval65.String()))
				case "nrlike":
					l.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%s%%", fval65.String()))
				case "nalike":
					l.GetMilitaryTC.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval65.String()))
				case "in":
					l.GetMilitaryTC.AndWhereIn(fval65.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop65)
				}
			}
		}
	}
	ftyp66, exists := typ.FieldByName("OriginAreaName")
	if exists {
		fval66 := val.FieldByName("OriginAreaName")
		fop66, ok := ftyp66.Tag.Lookup("op")
		for fval66.Kind() == reflect.Ptr && !fval66.IsNil() {
			fval66 = fval66.Elem()
		}
		if fval66.Kind() != reflect.Ptr {
			if !ok {
				l.OriginAreaName.AndWhereEq(fval66.Interface())
			} else {
				switch fop66 {
				case "=":
					l.OriginAreaName.AndWhereEq(fval66.Interface())
				case "!=":
					l.OriginAreaName.AndWhereNeq(fval66.Interface())
				case ">":
					l.OriginAreaName.AndWhereGt(fval66.Interface())
				case ">=":
					l.OriginAreaName.AndWhereGte(fval66.Interface())
				case "<":
					l.OriginAreaName.AndWhereLt(fval66.Interface())
				case "<=":
					l.OriginAreaName.AndWhereLte(fval66.Interface())
				case "llike":
					l.OriginAreaName.AndWhereLike(fmt.Sprintf("%%%s", fval66.String()))
				case "rlike":
					l.OriginAreaName.AndWhereLike(fmt.Sprintf("%s%%", fval66.String()))
				case "alike":
					l.OriginAreaName.AndWhereLike(fmt.Sprintf("%%%s%%", fval66.String()))
				case "nllike":
					l.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%%%s", fval66.String()))
				case "nrlike":
					l.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%s%%", fval66.String()))
				case "nalike":
					l.OriginAreaName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval66.String()))
				case "in":
					l.OriginAreaName.AndWhereIn(fval66.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop66)
				}
			}
		}
	}
	return l, nil
}
func NewClass() *Class {
	m := &Class{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", "ClassCode", "ClassCode", 2)
	m.ClassName.Init(m, "ClassName", "ClassName", "ClassName", "ClassName", 3)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 4)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", "ResearchArea", "ResearchArea", 5)
	m.Grade.Init(m, "Grade", "Grade", "Grade", "Grade", 6)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", "TrainingMode", "TrainingMode", 7)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", "EntranceDate", "EntranceDate", 8)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 9)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", "ProgramLength", "ProgramLength", 10)
	m.StudentType.Init(m, "StudentType", "StudentType", "StudentType", "StudentType", 11)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", "CredentialsType", "CredentialsType", 12)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", "DegreeType", "DegreeType", 13)
	m.Counselor.Init(m, "Counselor", "Counselor", "Counselor", "Counselor", 14)
	m.Adviser.Init(m, "Adviser", "Adviser", "Adviser", "Adviser", 15)
	m.Leadership.Init(m, "Leadership", "Leadership", "Leadership", "Leadership", 16)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", "Supervisor", "Supervisor", 17)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", "Assistant1", "Assistant1", 18)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", "Assistant2", "Assistant2", 19)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 20)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 21)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 22)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 23)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 24)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 25)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 26)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 27)
	m.InitRel()
	return m
}

func newSubClass(parent nborm.Model) *Class {
	m := &Class{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", "ClassCode", "ClassCode", 2)
	m.ClassName.Init(m, "ClassName", "ClassName", "ClassName", "ClassName", 3)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 4)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", "ResearchArea", "ResearchArea", 5)
	m.Grade.Init(m, "Grade", "Grade", "Grade", "Grade", 6)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", "TrainingMode", "TrainingMode", 7)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", "EntranceDate", "EntranceDate", 8)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 9)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", "ProgramLength", "ProgramLength", 10)
	m.StudentType.Init(m, "StudentType", "StudentType", "StudentType", "StudentType", 11)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", "CredentialsType", "CredentialsType", 12)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", "DegreeType", "DegreeType", 13)
	m.Counselor.Init(m, "Counselor", "Counselor", "Counselor", "Counselor", 14)
	m.Adviser.Init(m, "Adviser", "Adviser", "Adviser", "Adviser", 15)
	m.Leadership.Init(m, "Leadership", "Leadership", "Leadership", "Leadership", 16)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", "Supervisor", "Supervisor", 17)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", "Assistant1", "Assistant1", 18)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", "Assistant2", "Assistant2", 19)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 20)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 21)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 22)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 23)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 24)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 25)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 26)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 27)
	return m
}

func (m *Class) InitRel() {
	m.ClassGrade = newSubGrade(m)
	var relInfo0 *nborm.RelationInfo
	relInfo0 = relInfo0.Append("ClassGrade", m.ClassGrade, nborm.NewExpr("@=@", &m.Grade, &m.ClassGrade.GradeCode))
	m.AppendRelation(relInfo0)
	m.Students = newSubStudentbasicinfoList(m)
	m.Students.dupMap = make(map[string]int)
	var relInfo1 *nborm.RelationInfo
	relInfo1 = relInfo1.Append("Students", m.Students, nborm.NewExpr("@=@", &m.ClassCode, &m.Students.Class))
	m.AppendRelation(relInfo1)
	m.AddRelInited()
}

func (m *Class) DB() string {
	return "*"
}

func (m *Class) Tab() string {
	return "class"
}

func (m *Class) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"Id", "Id", &m.Id, 0},
		{"RecordId", "RecordId", &m.RecordId, 1},
		{"ClassCode", "ClassCode", &m.ClassCode, 2},
		{"ClassName", "ClassName", &m.ClassName, 3},
		{"Campus", "Campus", &m.Campus, 4},
		{"ResearchArea", "ResearchArea", &m.ResearchArea, 5},
		{"Grade", "Grade", &m.Grade, 6},
		{"TrainingMode", "TrainingMode", &m.TrainingMode, 7},
		{"EntranceDate", "EntranceDate", &m.EntranceDate, 8},
		{"GraduationDate", "GraduationDate", &m.GraduationDate, 9},
		{"ProgramLength", "ProgramLength", &m.ProgramLength, 10},
		{"StudentType", "StudentType", &m.StudentType, 11},
		{"CredentialsType", "CredentialsType", &m.CredentialsType, 12},
		{"DegreeType", "DegreeType", &m.DegreeType, 13},
		{"Counselor", "Counselor", &m.Counselor, 14},
		{"Adviser", "Adviser", &m.Adviser, 15},
		{"Leadership", "Leadership", &m.Leadership, 16},
		{"Supervisor", "Supervisor", &m.Supervisor, 17},
		{"Assistant1", "Assistant1", &m.Assistant1, 18},
		{"Assistant2", "Assistant2", &m.Assistant2, 19},
		{"Operator", "Operator", &m.Operator, 20},
		{"InsertDatetime", "InsertDatetime", &m.InsertDatetime, 21},
		{"UpdateDatetime", "UpdateDatetime", &m.UpdateDatetime, 22},
		{"Status", "Status", &m.Status, 23},
		{"Remark1", "Remark1", &m.Remark1, 24},
		{"Remark2", "Remark2", &m.Remark2, 25},
		{"Remark3", "Remark3", &m.Remark3, 26},
		{"Remark4", "Remark4", &m.Remark4, 27},
	}
}

func (m *Class) AutoIncField() nborm.Field {
	return &m.Id
}

func (m *Class) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.Id,
	}
}

func (m *Class) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.ClassCode,
			&m.Status,
		},
	}
}
func (m Class) MarshalJSON() ([]byte, error) {
	if !m.IsSynced() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	buffer.WriteString("{\n\"Aggs\": ")
	metaB, err := json.MarshalIndent(m.Meta, "", "\t")
	if err != nil {
		return nil, err
	}
	buffer.Write(metaB)
	if m.Id.IsValid() {
		buffer.WriteString(",\n\"Id\": ")
		IdB, err := json.MarshalIndent(m.Id, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdB)
	}
	if m.RecordId.IsValid() {
		buffer.WriteString(",\n\"RecordId\": ")
		RecordIdB, err := json.MarshalIndent(m.RecordId, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RecordIdB)
	}
	if m.ClassCode.IsValid() {
		buffer.WriteString(",\n\"ClassCode\": ")
		ClassCodeB, err := json.MarshalIndent(m.ClassCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClassCodeB)
	}
	if m.ClassName.IsValid() {
		buffer.WriteString(",\n\"ClassName\": ")
		ClassNameB, err := json.MarshalIndent(m.ClassName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClassNameB)
	}
	if m.Campus.IsValid() {
		buffer.WriteString(",\n\"Campus\": ")
		CampusB, err := json.MarshalIndent(m.Campus, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CampusB)
	}
	if m.ResearchArea.IsValid() {
		buffer.WriteString(",\n\"ResearchArea\": ")
		ResearchAreaB, err := json.MarshalIndent(m.ResearchArea, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ResearchAreaB)
	}
	if m.Grade.IsValid() {
		buffer.WriteString(",\n\"Grade\": ")
		GradeB, err := json.MarshalIndent(m.Grade, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GradeB)
	}
	if m.TrainingMode.IsValid() {
		buffer.WriteString(",\n\"TrainingMode\": ")
		TrainingModeB, err := json.MarshalIndent(m.TrainingMode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(TrainingModeB)
	}
	if m.EntranceDate.IsValid() {
		buffer.WriteString(",\n\"EntranceDate\": ")
		EntranceDateB, err := json.MarshalIndent(m.EntranceDate, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EntranceDateB)
	}
	if m.GraduationDate.IsValid() {
		buffer.WriteString(",\n\"GraduationDate\": ")
		GraduationDateB, err := json.MarshalIndent(m.GraduationDate, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GraduationDateB)
	}
	if m.ProgramLength.IsValid() {
		buffer.WriteString(",\n\"ProgramLength\": ")
		ProgramLengthB, err := json.MarshalIndent(m.ProgramLength, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ProgramLengthB)
	}
	if m.StudentType.IsValid() {
		buffer.WriteString(",\n\"StudentType\": ")
		StudentTypeB, err := json.MarshalIndent(m.StudentType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentTypeB)
	}
	if m.CredentialsType.IsValid() {
		buffer.WriteString(",\n\"CredentialsType\": ")
		CredentialsTypeB, err := json.MarshalIndent(m.CredentialsType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CredentialsTypeB)
	}
	if m.DegreeType.IsValid() {
		buffer.WriteString(",\n\"DegreeType\": ")
		DegreeTypeB, err := json.MarshalIndent(m.DegreeType, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(DegreeTypeB)
	}
	if m.Counselor.IsValid() {
		buffer.WriteString(",\n\"Counselor\": ")
		CounselorB, err := json.MarshalIndent(m.Counselor, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CounselorB)
	}
	if m.Adviser.IsValid() {
		buffer.WriteString(",\n\"Adviser\": ")
		AdviserB, err := json.MarshalIndent(m.Adviser, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AdviserB)
	}
	if m.Leadership.IsValid() {
		buffer.WriteString(",\n\"Leadership\": ")
		LeadershipB, err := json.MarshalIndent(m.Leadership, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(LeadershipB)
	}
	if m.Supervisor.IsValid() {
		buffer.WriteString(",\n\"Supervisor\": ")
		SupervisorB, err := json.MarshalIndent(m.Supervisor, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(SupervisorB)
	}
	if m.Assistant1.IsValid() {
		buffer.WriteString(",\n\"Assistant1\": ")
		Assistant1B, err := json.MarshalIndent(m.Assistant1, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Assistant1B)
	}
	if m.Assistant2.IsValid() {
		buffer.WriteString(",\n\"Assistant2\": ")
		Assistant2B, err := json.MarshalIndent(m.Assistant2, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Assistant2B)
	}
	if m.Operator.IsValid() {
		buffer.WriteString(",\n\"Operator\": ")
		OperatorB, err := json.MarshalIndent(m.Operator, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OperatorB)
	}
	if m.InsertDatetime.IsValid() {
		buffer.WriteString(",\n\"InsertDatetime\": ")
		InsertDatetimeB, err := json.MarshalIndent(m.InsertDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(InsertDatetimeB)
	}
	if m.UpdateDatetime.IsValid() {
		buffer.WriteString(",\n\"UpdateDatetime\": ")
		UpdateDatetimeB, err := json.MarshalIndent(m.UpdateDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateDatetimeB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.Remark1.IsValid() {
		buffer.WriteString(",\n\"Remark1\": ")
		Remark1B, err := json.MarshalIndent(m.Remark1, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark1B)
	}
	if m.Remark2.IsValid() {
		buffer.WriteString(",\n\"Remark2\": ")
		Remark2B, err := json.MarshalIndent(m.Remark2, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark2B)
	}
	if m.Remark3.IsValid() {
		buffer.WriteString(",\n\"Remark3\": ")
		Remark3B, err := json.MarshalIndent(m.Remark3, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark3B)
	}
	if m.Remark4.IsValid() {
		buffer.WriteString(",\n\"Remark4\": ")
		Remark4B, err := json.MarshalIndent(m.Remark4, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark4B)
	}
	if m.ClassGrade != nil && m.ClassGrade.IsSynced() {
		buffer.WriteString(",\n\"ClassGrade\": ")
		ClassGradeB, err := json.MarshalIndent(m.ClassGrade, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClassGradeB)
	}
	if m.Students != nil && m.Students.Len() > 0 {
		buffer.WriteString(",\n\"Students\": ")
		StudentsB, err := json.MarshalIndent(m.Students, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StudentsB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type ClassList struct {
	Class  `json:"-"`
	dupMap map[string]int
	List   []*Class
	Total  int
}

func (m *Class) Collapse() {
	if m.ClassGrade != nil && m.ClassGrade.IsSynced() {
		m.ClassGrade.Collapse()
	}
	if m.Students != nil && m.Students.IsSynced() {
		m.Students.Collapse()
	}
}

func NewClassList() *ClassList {
	l := &ClassList{
		Class{},
		make(map[string]int),
		make([]*Class, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.ClassCode.Init(l, "ClassCode", "ClassCode", "ClassCode", "ClassCode", 2)
	l.ClassName.Init(l, "ClassName", "ClassName", "ClassName", "ClassName", 3)
	l.Campus.Init(l, "Campus", "Campus", "Campus", "Campus", 4)
	l.ResearchArea.Init(l, "ResearchArea", "ResearchArea", "ResearchArea", "ResearchArea", 5)
	l.Grade.Init(l, "Grade", "Grade", "Grade", "Grade", 6)
	l.TrainingMode.Init(l, "TrainingMode", "TrainingMode", "TrainingMode", "TrainingMode", 7)
	l.EntranceDate.Init(l, "EntranceDate", "EntranceDate", "EntranceDate", "EntranceDate", 8)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 9)
	l.ProgramLength.Init(l, "ProgramLength", "ProgramLength", "ProgramLength", "ProgramLength", 10)
	l.StudentType.Init(l, "StudentType", "StudentType", "StudentType", "StudentType", 11)
	l.CredentialsType.Init(l, "CredentialsType", "CredentialsType", "CredentialsType", "CredentialsType", 12)
	l.DegreeType.Init(l, "DegreeType", "DegreeType", "DegreeType", "DegreeType", 13)
	l.Counselor.Init(l, "Counselor", "Counselor", "Counselor", "Counselor", 14)
	l.Adviser.Init(l, "Adviser", "Adviser", "Adviser", "Adviser", 15)
	l.Leadership.Init(l, "Leadership", "Leadership", "Leadership", "Leadership", 16)
	l.Supervisor.Init(l, "Supervisor", "Supervisor", "Supervisor", "Supervisor", 17)
	l.Assistant1.Init(l, "Assistant1", "Assistant1", "Assistant1", "Assistant1", 18)
	l.Assistant2.Init(l, "Assistant2", "Assistant2", "Assistant2", "Assistant2", 19)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 20)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 21)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 22)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 23)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 24)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 25)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 26)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 27)
	l.InitRel()
	return l
}

func newSubClassList(parent nborm.Model) *ClassList {
	l := &ClassList{
		Class{},
		make(map[string]int),
		make([]*Class, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.ClassCode.Init(l, "ClassCode", "ClassCode", "ClassCode", "ClassCode", 2)
	l.ClassName.Init(l, "ClassName", "ClassName", "ClassName", "ClassName", 3)
	l.Campus.Init(l, "Campus", "Campus", "Campus", "Campus", 4)
	l.ResearchArea.Init(l, "ResearchArea", "ResearchArea", "ResearchArea", "ResearchArea", 5)
	l.Grade.Init(l, "Grade", "Grade", "Grade", "Grade", 6)
	l.TrainingMode.Init(l, "TrainingMode", "TrainingMode", "TrainingMode", "TrainingMode", 7)
	l.EntranceDate.Init(l, "EntranceDate", "EntranceDate", "EntranceDate", "EntranceDate", 8)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 9)
	l.ProgramLength.Init(l, "ProgramLength", "ProgramLength", "ProgramLength", "ProgramLength", 10)
	l.StudentType.Init(l, "StudentType", "StudentType", "StudentType", "StudentType", 11)
	l.CredentialsType.Init(l, "CredentialsType", "CredentialsType", "CredentialsType", "CredentialsType", 12)
	l.DegreeType.Init(l, "DegreeType", "DegreeType", "DegreeType", "DegreeType", 13)
	l.Counselor.Init(l, "Counselor", "Counselor", "Counselor", "Counselor", 14)
	l.Adviser.Init(l, "Adviser", "Adviser", "Adviser", "Adviser", 15)
	l.Leadership.Init(l, "Leadership", "Leadership", "Leadership", "Leadership", 16)
	l.Supervisor.Init(l, "Supervisor", "Supervisor", "Supervisor", "Supervisor", 17)
	l.Assistant1.Init(l, "Assistant1", "Assistant1", "Assistant1", "Assistant1", 18)
	l.Assistant2.Init(l, "Assistant2", "Assistant2", "Assistant2", "Assistant2", 19)
	l.Operator.Init(l, "Operator", "Operator", "Operator", "Operator", 20)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 21)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 22)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 23)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 24)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 25)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 26)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 27)
	return l
}

func (l *ClassList) NewModel() nborm.Model {
	m := &Class{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.RecordId.Init(m, "RecordId", "RecordId", "RecordId", "RecordId", 1)
	l.RecordId.CopyStatus(&m.RecordId)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", "ClassCode", "ClassCode", 2)
	l.ClassCode.CopyStatus(&m.ClassCode)
	m.ClassName.Init(m, "ClassName", "ClassName", "ClassName", "ClassName", 3)
	l.ClassName.CopyStatus(&m.ClassName)
	m.Campus.Init(m, "Campus", "Campus", "Campus", "Campus", 4)
	l.Campus.CopyStatus(&m.Campus)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", "ResearchArea", "ResearchArea", 5)
	l.ResearchArea.CopyStatus(&m.ResearchArea)
	m.Grade.Init(m, "Grade", "Grade", "Grade", "Grade", 6)
	l.Grade.CopyStatus(&m.Grade)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", "TrainingMode", "TrainingMode", 7)
	l.TrainingMode.CopyStatus(&m.TrainingMode)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", "EntranceDate", "EntranceDate", 8)
	l.EntranceDate.CopyStatus(&m.EntranceDate)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", "GraduationDate", "GraduationDate", 9)
	l.GraduationDate.CopyStatus(&m.GraduationDate)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", "ProgramLength", "ProgramLength", 10)
	l.ProgramLength.CopyStatus(&m.ProgramLength)
	m.StudentType.Init(m, "StudentType", "StudentType", "StudentType", "StudentType", 11)
	l.StudentType.CopyStatus(&m.StudentType)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", "CredentialsType", "CredentialsType", 12)
	l.CredentialsType.CopyStatus(&m.CredentialsType)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", "DegreeType", "DegreeType", 13)
	l.DegreeType.CopyStatus(&m.DegreeType)
	m.Counselor.Init(m, "Counselor", "Counselor", "Counselor", "Counselor", 14)
	l.Counselor.CopyStatus(&m.Counselor)
	m.Adviser.Init(m, "Adviser", "Adviser", "Adviser", "Adviser", 15)
	l.Adviser.CopyStatus(&m.Adviser)
	m.Leadership.Init(m, "Leadership", "Leadership", "Leadership", "Leadership", 16)
	l.Leadership.CopyStatus(&m.Leadership)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", "Supervisor", "Supervisor", 17)
	l.Supervisor.CopyStatus(&m.Supervisor)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", "Assistant1", "Assistant1", 18)
	l.Assistant1.CopyStatus(&m.Assistant1)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", "Assistant2", "Assistant2", 19)
	l.Assistant2.CopyStatus(&m.Assistant2)
	m.Operator.Init(m, "Operator", "Operator", "Operator", "Operator", 20)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 21)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 22)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 23)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 24)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 25)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 26)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 27)
	l.Remark4.CopyStatus(&m.Remark4)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *ClassList) SetTotal(total int) {
	l.Total = total
}

func (l *ClassList) GetTotal() int {
	return l.Total
}

func (l *ClassList) Len() int {
	return len(l.List)
}

func (l *ClassList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l ClassList) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, 1024)
	bs = append(bs, []byte("{")...)
	ListB, err := json.MarshalIndent(l.List, "", "\t")
	if err != nil {
		return nil, err
	}
	ListB = append([]byte("\"List\": "), ListB...)
	bs = append(bs, ListB...)
	bs = append(bs, []byte(", ")...)
	TotalB, err := json.MarshalIndent(l.Total, "", "\t")
	if err != nil {
		return nil, err
	}
	TotalB = append([]byte("\"Total\": "), TotalB...)
	bs = append(bs, TotalB...)
	bs = append(bs, []byte("}")...)
	return bs, nil
}

func (l *ClassList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.Class)
}

func (l *ClassList) UnmarshalMeta(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &l.Class)
}

func (l *ClassList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List[idx].ClassGrade = l.List[l.Len()-1].ClassGrade
		l.List[idx].Students.checkDup()
		l.List[idx].Students.List = append(l.List[idx].Students.List, l.List[l.Len()-1].Students.List...)
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *ClassList) Filter(f func(m *Class) bool) []*Class {
	ll := make([]*Class, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *ClassList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.Id.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Id.AnyValue()))
	}
	if lastModel.RecordId.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RecordId.AnyValue()))
	}
	if lastModel.ClassCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ClassCode.AnyValue()))
	}
	if lastModel.ClassName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ClassName.AnyValue()))
	}
	if lastModel.Campus.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Campus.AnyValue()))
	}
	if lastModel.ResearchArea.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ResearchArea.AnyValue()))
	}
	if lastModel.Grade.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Grade.AnyValue()))
	}
	if lastModel.TrainingMode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.TrainingMode.AnyValue()))
	}
	if lastModel.EntranceDate.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.EntranceDate.AnyValue()))
	}
	if lastModel.GraduationDate.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.GraduationDate.AnyValue()))
	}
	if lastModel.ProgramLength.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ProgramLength.AnyValue()))
	}
	if lastModel.StudentType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.StudentType.AnyValue()))
	}
	if lastModel.CredentialsType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.CredentialsType.AnyValue()))
	}
	if lastModel.DegreeType.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.DegreeType.AnyValue()))
	}
	if lastModel.Counselor.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Counselor.AnyValue()))
	}
	if lastModel.Adviser.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Adviser.AnyValue()))
	}
	if lastModel.Leadership.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Leadership.AnyValue()))
	}
	if lastModel.Supervisor.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Supervisor.AnyValue()))
	}
	if lastModel.Assistant1.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Assistant1.AnyValue()))
	}
	if lastModel.Assistant2.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Assistant2.AnyValue()))
	}
	if lastModel.Operator.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Operator.AnyValue()))
	}
	if lastModel.InsertDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.InsertDatetime.AnyValue()))
	}
	if lastModel.UpdateDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateDatetime.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.Remark1.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark1.AnyValue()))
	}
	if lastModel.Remark2.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark2.AnyValue()))
	}
	if lastModel.Remark3.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark3.AnyValue()))
	}
	if lastModel.Remark4.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark4.AnyValue()))
	}
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *ClassList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

func (m *Class) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (l *ClassList) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

type ClassCacheElem struct {
	hashValue  string
	model      *Class
	modifyTime time.Time
}

type ClassListCacheElem struct {
	hashValue  string
	list       *ClassList
	modifyTime time.Time
}

type ClassCacheManager struct {
	container map[string]*ClassCacheElem
	query     chan string
	in        chan *ClassCacheElem
	out       chan *ClassCacheElem
}

type ClassListCacheManager struct {
	container map[string]*ClassListCacheElem
	query     chan string
	in        chan *ClassListCacheElem
	out       chan *ClassListCacheElem
}

func newClassCacheManager() *ClassCacheManager {
	return &ClassCacheManager{
		make(map[string]*ClassCacheElem),
		make(chan string),
		make(chan *ClassCacheElem),
		make(chan *ClassCacheElem),
	}
}

func newClassListCacheManager() *ClassListCacheManager {
	return &ClassListCacheManager{
		make(map[string]*ClassListCacheElem),
		make(chan string),
		make(chan *ClassListCacheElem),
		make(chan *ClassListCacheElem),
	}
}

func (mgr *ClassCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *ClassListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var ClassCache = newClassCacheManager()

var ClassListCache = newClassListCacheManager()

func (m *Class) GetCache(hashVal string, timeout time.Duration) bool {
	ClassCache.query <- hashVal
	elem := <-ClassCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *Class) SetCache(hashValue string) {
	ClassCache.in <- &ClassCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *ClassList) GetListCache(hashValue string, timeout time.Duration) bool {
	ClassListCache.query <- hashValue
	elem := <-ClassListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *ClassList) SetListCache(hashValue string) {
	ClassListCache.in <- &ClassListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func (m *Class) FromQuery(query interface{}) (*Class, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				m.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					m.Id.AndWhereEq(fval0.Interface())
				case "!=":
					m.Id.AndWhereNeq(fval0.Interface())
				case ">":
					m.Id.AndWhereGt(fval0.Interface())
				case ">=":
					m.Id.AndWhereGte(fval0.Interface())
				case "<":
					m.Id.AndWhereLt(fval0.Interface())
				case "<=":
					m.Id.AndWhereLte(fval0.Interface())
				case "llike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					m.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					m.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("RecordId")
	if exists {
		fval1 := val.FieldByName("RecordId")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				m.RecordId.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					m.RecordId.AndWhereEq(fval1.Interface())
				case "!=":
					m.RecordId.AndWhereNeq(fval1.Interface())
				case ">":
					m.RecordId.AndWhereGt(fval1.Interface())
				case ">=":
					m.RecordId.AndWhereGte(fval1.Interface())
				case "<":
					m.RecordId.AndWhereLt(fval1.Interface())
				case "<=":
					m.RecordId.AndWhereLte(fval1.Interface())
				case "llike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					m.RecordId.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					m.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					m.RecordId.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("ClassCode")
	if exists {
		fval2 := val.FieldByName("ClassCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				m.ClassCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					m.ClassCode.AndWhereEq(fval2.Interface())
				case "!=":
					m.ClassCode.AndWhereNeq(fval2.Interface())
				case ">":
					m.ClassCode.AndWhereGt(fval2.Interface())
				case ">=":
					m.ClassCode.AndWhereGte(fval2.Interface())
				case "<":
					m.ClassCode.AndWhereLt(fval2.Interface())
				case "<=":
					m.ClassCode.AndWhereLte(fval2.Interface())
				case "llike":
					m.ClassCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					m.ClassCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					m.ClassCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					m.ClassCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					m.ClassCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					m.ClassCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					m.ClassCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("ClassName")
	if exists {
		fval3 := val.FieldByName("ClassName")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				m.ClassName.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					m.ClassName.AndWhereEq(fval3.Interface())
				case "!=":
					m.ClassName.AndWhereNeq(fval3.Interface())
				case ">":
					m.ClassName.AndWhereGt(fval3.Interface())
				case ">=":
					m.ClassName.AndWhereGte(fval3.Interface())
				case "<":
					m.ClassName.AndWhereLt(fval3.Interface())
				case "<=":
					m.ClassName.AndWhereLte(fval3.Interface())
				case "llike":
					m.ClassName.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					m.ClassName.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					m.ClassName.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					m.ClassName.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					m.ClassName.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					m.ClassName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					m.ClassName.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("Campus")
	if exists {
		fval4 := val.FieldByName("Campus")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				m.Campus.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					m.Campus.AndWhereEq(fval4.Interface())
				case "!=":
					m.Campus.AndWhereNeq(fval4.Interface())
				case ">":
					m.Campus.AndWhereGt(fval4.Interface())
				case ">=":
					m.Campus.AndWhereGte(fval4.Interface())
				case "<":
					m.Campus.AndWhereLt(fval4.Interface())
				case "<=":
					m.Campus.AndWhereLte(fval4.Interface())
				case "llike":
					m.Campus.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					m.Campus.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					m.Campus.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					m.Campus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					m.Campus.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("ResearchArea")
	if exists {
		fval5 := val.FieldByName("ResearchArea")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				m.ResearchArea.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					m.ResearchArea.AndWhereEq(fval5.Interface())
				case "!=":
					m.ResearchArea.AndWhereNeq(fval5.Interface())
				case ">":
					m.ResearchArea.AndWhereGt(fval5.Interface())
				case ">=":
					m.ResearchArea.AndWhereGte(fval5.Interface())
				case "<":
					m.ResearchArea.AndWhereLt(fval5.Interface())
				case "<=":
					m.ResearchArea.AndWhereLte(fval5.Interface())
				case "llike":
					m.ResearchArea.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					m.ResearchArea.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					m.ResearchArea.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					m.ResearchArea.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					m.ResearchArea.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					m.ResearchArea.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					m.ResearchArea.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("Grade")
	if exists {
		fval6 := val.FieldByName("Grade")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				m.Grade.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					m.Grade.AndWhereEq(fval6.Interface())
				case "!=":
					m.Grade.AndWhereNeq(fval6.Interface())
				case ">":
					m.Grade.AndWhereGt(fval6.Interface())
				case ">=":
					m.Grade.AndWhereGte(fval6.Interface())
				case "<":
					m.Grade.AndWhereLt(fval6.Interface())
				case "<=":
					m.Grade.AndWhereLte(fval6.Interface())
				case "llike":
					m.Grade.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					m.Grade.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					m.Grade.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					m.Grade.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					m.Grade.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					m.Grade.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					m.Grade.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("TrainingMode")
	if exists {
		fval7 := val.FieldByName("TrainingMode")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				m.TrainingMode.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					m.TrainingMode.AndWhereEq(fval7.Interface())
				case "!=":
					m.TrainingMode.AndWhereNeq(fval7.Interface())
				case ">":
					m.TrainingMode.AndWhereGt(fval7.Interface())
				case ">=":
					m.TrainingMode.AndWhereGte(fval7.Interface())
				case "<":
					m.TrainingMode.AndWhereLt(fval7.Interface())
				case "<=":
					m.TrainingMode.AndWhereLte(fval7.Interface())
				case "llike":
					m.TrainingMode.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					m.TrainingMode.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					m.TrainingMode.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					m.TrainingMode.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					m.TrainingMode.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					m.TrainingMode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					m.TrainingMode.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("EntranceDate")
	if exists {
		fval8 := val.FieldByName("EntranceDate")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				m.EntranceDate.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					m.EntranceDate.AndWhereEq(fval8.Interface())
				case "!=":
					m.EntranceDate.AndWhereNeq(fval8.Interface())
				case ">":
					m.EntranceDate.AndWhereGt(fval8.Interface())
				case ">=":
					m.EntranceDate.AndWhereGte(fval8.Interface())
				case "<":
					m.EntranceDate.AndWhereLt(fval8.Interface())
				case "<=":
					m.EntranceDate.AndWhereLte(fval8.Interface())
				case "llike":
					m.EntranceDate.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					m.EntranceDate.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					m.EntranceDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					m.EntranceDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					m.EntranceDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					m.EntranceDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					m.EntranceDate.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("GraduationDate")
	if exists {
		fval9 := val.FieldByName("GraduationDate")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				m.GraduationDate.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					m.GraduationDate.AndWhereEq(fval9.Interface())
				case "!=":
					m.GraduationDate.AndWhereNeq(fval9.Interface())
				case ">":
					m.GraduationDate.AndWhereGt(fval9.Interface())
				case ">=":
					m.GraduationDate.AndWhereGte(fval9.Interface())
				case "<":
					m.GraduationDate.AndWhereLt(fval9.Interface())
				case "<=":
					m.GraduationDate.AndWhereLte(fval9.Interface())
				case "llike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					m.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					m.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					m.GraduationDate.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("ProgramLength")
	if exists {
		fval10 := val.FieldByName("ProgramLength")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				m.ProgramLength.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					m.ProgramLength.AndWhereEq(fval10.Interface())
				case "!=":
					m.ProgramLength.AndWhereNeq(fval10.Interface())
				case ">":
					m.ProgramLength.AndWhereGt(fval10.Interface())
				case ">=":
					m.ProgramLength.AndWhereGte(fval10.Interface())
				case "<":
					m.ProgramLength.AndWhereLt(fval10.Interface())
				case "<=":
					m.ProgramLength.AndWhereLte(fval10.Interface())
				case "llike":
					m.ProgramLength.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					m.ProgramLength.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					m.ProgramLength.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					m.ProgramLength.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					m.ProgramLength.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					m.ProgramLength.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					m.ProgramLength.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("StudentType")
	if exists {
		fval11 := val.FieldByName("StudentType")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				m.StudentType.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					m.StudentType.AndWhereEq(fval11.Interface())
				case "!=":
					m.StudentType.AndWhereNeq(fval11.Interface())
				case ">":
					m.StudentType.AndWhereGt(fval11.Interface())
				case ">=":
					m.StudentType.AndWhereGte(fval11.Interface())
				case "<":
					m.StudentType.AndWhereLt(fval11.Interface())
				case "<=":
					m.StudentType.AndWhereLte(fval11.Interface())
				case "llike":
					m.StudentType.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					m.StudentType.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					m.StudentType.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					m.StudentType.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					m.StudentType.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					m.StudentType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					m.StudentType.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("CredentialsType")
	if exists {
		fval12 := val.FieldByName("CredentialsType")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				m.CredentialsType.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					m.CredentialsType.AndWhereEq(fval12.Interface())
				case "!=":
					m.CredentialsType.AndWhereNeq(fval12.Interface())
				case ">":
					m.CredentialsType.AndWhereGt(fval12.Interface())
				case ">=":
					m.CredentialsType.AndWhereGte(fval12.Interface())
				case "<":
					m.CredentialsType.AndWhereLt(fval12.Interface())
				case "<=":
					m.CredentialsType.AndWhereLte(fval12.Interface())
				case "llike":
					m.CredentialsType.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					m.CredentialsType.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					m.CredentialsType.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					m.CredentialsType.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					m.CredentialsType.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					m.CredentialsType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					m.CredentialsType.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("DegreeType")
	if exists {
		fval13 := val.FieldByName("DegreeType")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				m.DegreeType.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					m.DegreeType.AndWhereEq(fval13.Interface())
				case "!=":
					m.DegreeType.AndWhereNeq(fval13.Interface())
				case ">":
					m.DegreeType.AndWhereGt(fval13.Interface())
				case ">=":
					m.DegreeType.AndWhereGte(fval13.Interface())
				case "<":
					m.DegreeType.AndWhereLt(fval13.Interface())
				case "<=":
					m.DegreeType.AndWhereLte(fval13.Interface())
				case "llike":
					m.DegreeType.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					m.DegreeType.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					m.DegreeType.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					m.DegreeType.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					m.DegreeType.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					m.DegreeType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					m.DegreeType.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("Counselor")
	if exists {
		fval14 := val.FieldByName("Counselor")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				m.Counselor.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					m.Counselor.AndWhereEq(fval14.Interface())
				case "!=":
					m.Counselor.AndWhereNeq(fval14.Interface())
				case ">":
					m.Counselor.AndWhereGt(fval14.Interface())
				case ">=":
					m.Counselor.AndWhereGte(fval14.Interface())
				case "<":
					m.Counselor.AndWhereLt(fval14.Interface())
				case "<=":
					m.Counselor.AndWhereLte(fval14.Interface())
				case "llike":
					m.Counselor.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					m.Counselor.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					m.Counselor.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					m.Counselor.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					m.Counselor.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					m.Counselor.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					m.Counselor.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("Adviser")
	if exists {
		fval15 := val.FieldByName("Adviser")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				m.Adviser.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					m.Adviser.AndWhereEq(fval15.Interface())
				case "!=":
					m.Adviser.AndWhereNeq(fval15.Interface())
				case ">":
					m.Adviser.AndWhereGt(fval15.Interface())
				case ">=":
					m.Adviser.AndWhereGte(fval15.Interface())
				case "<":
					m.Adviser.AndWhereLt(fval15.Interface())
				case "<=":
					m.Adviser.AndWhereLte(fval15.Interface())
				case "llike":
					m.Adviser.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					m.Adviser.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					m.Adviser.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					m.Adviser.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					m.Adviser.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					m.Adviser.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					m.Adviser.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("Leadership")
	if exists {
		fval16 := val.FieldByName("Leadership")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				m.Leadership.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					m.Leadership.AndWhereEq(fval16.Interface())
				case "!=":
					m.Leadership.AndWhereNeq(fval16.Interface())
				case ">":
					m.Leadership.AndWhereGt(fval16.Interface())
				case ">=":
					m.Leadership.AndWhereGte(fval16.Interface())
				case "<":
					m.Leadership.AndWhereLt(fval16.Interface())
				case "<=":
					m.Leadership.AndWhereLte(fval16.Interface())
				case "llike":
					m.Leadership.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					m.Leadership.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					m.Leadership.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					m.Leadership.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					m.Leadership.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					m.Leadership.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					m.Leadership.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("Supervisor")
	if exists {
		fval17 := val.FieldByName("Supervisor")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				m.Supervisor.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					m.Supervisor.AndWhereEq(fval17.Interface())
				case "!=":
					m.Supervisor.AndWhereNeq(fval17.Interface())
				case ">":
					m.Supervisor.AndWhereGt(fval17.Interface())
				case ">=":
					m.Supervisor.AndWhereGte(fval17.Interface())
				case "<":
					m.Supervisor.AndWhereLt(fval17.Interface())
				case "<=":
					m.Supervisor.AndWhereLte(fval17.Interface())
				case "llike":
					m.Supervisor.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					m.Supervisor.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					m.Supervisor.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					m.Supervisor.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					m.Supervisor.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					m.Supervisor.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					m.Supervisor.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("Assistant1")
	if exists {
		fval18 := val.FieldByName("Assistant1")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				m.Assistant1.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					m.Assistant1.AndWhereEq(fval18.Interface())
				case "!=":
					m.Assistant1.AndWhereNeq(fval18.Interface())
				case ">":
					m.Assistant1.AndWhereGt(fval18.Interface())
				case ">=":
					m.Assistant1.AndWhereGte(fval18.Interface())
				case "<":
					m.Assistant1.AndWhereLt(fval18.Interface())
				case "<=":
					m.Assistant1.AndWhereLte(fval18.Interface())
				case "llike":
					m.Assistant1.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					m.Assistant1.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					m.Assistant1.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					m.Assistant1.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					m.Assistant1.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					m.Assistant1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					m.Assistant1.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Assistant2")
	if exists {
		fval19 := val.FieldByName("Assistant2")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				m.Assistant2.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					m.Assistant2.AndWhereEq(fval19.Interface())
				case "!=":
					m.Assistant2.AndWhereNeq(fval19.Interface())
				case ">":
					m.Assistant2.AndWhereGt(fval19.Interface())
				case ">=":
					m.Assistant2.AndWhereGte(fval19.Interface())
				case "<":
					m.Assistant2.AndWhereLt(fval19.Interface())
				case "<=":
					m.Assistant2.AndWhereLte(fval19.Interface())
				case "llike":
					m.Assistant2.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					m.Assistant2.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					m.Assistant2.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					m.Assistant2.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					m.Assistant2.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					m.Assistant2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					m.Assistant2.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("Operator")
	if exists {
		fval20 := val.FieldByName("Operator")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				m.Operator.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					m.Operator.AndWhereEq(fval20.Interface())
				case "!=":
					m.Operator.AndWhereNeq(fval20.Interface())
				case ">":
					m.Operator.AndWhereGt(fval20.Interface())
				case ">=":
					m.Operator.AndWhereGte(fval20.Interface())
				case "<":
					m.Operator.AndWhereLt(fval20.Interface())
				case "<=":
					m.Operator.AndWhereLte(fval20.Interface())
				case "llike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					m.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					m.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					m.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					m.Operator.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval21 := val.FieldByName("InsertDatetime")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				m.InsertDatetime.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					m.InsertDatetime.AndWhereEq(fval21.Interface())
				case "!=":
					m.InsertDatetime.AndWhereNeq(fval21.Interface())
				case ">":
					m.InsertDatetime.AndWhereGt(fval21.Interface())
				case ">=":
					m.InsertDatetime.AndWhereGte(fval21.Interface())
				case "<":
					m.InsertDatetime.AndWhereLt(fval21.Interface())
				case "<=":
					m.InsertDatetime.AndWhereLte(fval21.Interface())
				case "llike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					m.InsertDatetime.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval22 := val.FieldByName("UpdateDatetime")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				m.UpdateDatetime.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					m.UpdateDatetime.AndWhereEq(fval22.Interface())
				case "!=":
					m.UpdateDatetime.AndWhereNeq(fval22.Interface())
				case ">":
					m.UpdateDatetime.AndWhereGt(fval22.Interface())
				case ">=":
					m.UpdateDatetime.AndWhereGte(fval22.Interface())
				case "<":
					m.UpdateDatetime.AndWhereLt(fval22.Interface())
				case "<=":
					m.UpdateDatetime.AndWhereLte(fval22.Interface())
				case "llike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					m.UpdateDatetime.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("Status")
	if exists {
		fval23 := val.FieldByName("Status")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				m.Status.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					m.Status.AndWhereEq(fval23.Interface())
				case "!=":
					m.Status.AndWhereNeq(fval23.Interface())
				case ">":
					m.Status.AndWhereGt(fval23.Interface())
				case ">=":
					m.Status.AndWhereGte(fval23.Interface())
				case "<":
					m.Status.AndWhereLt(fval23.Interface())
				case "<=":
					m.Status.AndWhereLte(fval23.Interface())
				case "llike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					m.Status.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					m.Status.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("Remark1")
	if exists {
		fval24 := val.FieldByName("Remark1")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				m.Remark1.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					m.Remark1.AndWhereEq(fval24.Interface())
				case "!=":
					m.Remark1.AndWhereNeq(fval24.Interface())
				case ">":
					m.Remark1.AndWhereGt(fval24.Interface())
				case ">=":
					m.Remark1.AndWhereGte(fval24.Interface())
				case "<":
					m.Remark1.AndWhereLt(fval24.Interface())
				case "<=":
					m.Remark1.AndWhereLte(fval24.Interface())
				case "llike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					m.Remark1.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("Remark2")
	if exists {
		fval25 := val.FieldByName("Remark2")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				m.Remark2.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					m.Remark2.AndWhereEq(fval25.Interface())
				case "!=":
					m.Remark2.AndWhereNeq(fval25.Interface())
				case ">":
					m.Remark2.AndWhereGt(fval25.Interface())
				case ">=":
					m.Remark2.AndWhereGte(fval25.Interface())
				case "<":
					m.Remark2.AndWhereLt(fval25.Interface())
				case "<=":
					m.Remark2.AndWhereLte(fval25.Interface())
				case "llike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					m.Remark2.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("Remark3")
	if exists {
		fval26 := val.FieldByName("Remark3")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				m.Remark3.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					m.Remark3.AndWhereEq(fval26.Interface())
				case "!=":
					m.Remark3.AndWhereNeq(fval26.Interface())
				case ">":
					m.Remark3.AndWhereGt(fval26.Interface())
				case ">=":
					m.Remark3.AndWhereGte(fval26.Interface())
				case "<":
					m.Remark3.AndWhereLt(fval26.Interface())
				case "<=":
					m.Remark3.AndWhereLte(fval26.Interface())
				case "llike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					m.Remark3.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("Remark4")
	if exists {
		fval27 := val.FieldByName("Remark4")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				m.Remark4.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					m.Remark4.AndWhereEq(fval27.Interface())
				case "!=":
					m.Remark4.AndWhereNeq(fval27.Interface())
				case ">":
					m.Remark4.AndWhereGt(fval27.Interface())
				case ">=":
					m.Remark4.AndWhereGte(fval27.Interface())
				case "<":
					m.Remark4.AndWhereLt(fval27.Interface())
				case "<=":
					m.Remark4.AndWhereLte(fval27.Interface())
				case "llike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					m.Remark4.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	return m, nil
}
func (l *ClassList) FromQuery(query interface{}) (*ClassList, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				l.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					l.Id.AndWhereEq(fval0.Interface())
				case "!=":
					l.Id.AndWhereNeq(fval0.Interface())
				case ">":
					l.Id.AndWhereGt(fval0.Interface())
				case ">=":
					l.Id.AndWhereGte(fval0.Interface())
				case "<":
					l.Id.AndWhereLt(fval0.Interface())
				case "<=":
					l.Id.AndWhereLte(fval0.Interface())
				case "llike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					l.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					l.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("RecordId")
	if exists {
		fval1 := val.FieldByName("RecordId")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				l.RecordId.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					l.RecordId.AndWhereEq(fval1.Interface())
				case "!=":
					l.RecordId.AndWhereNeq(fval1.Interface())
				case ">":
					l.RecordId.AndWhereGt(fval1.Interface())
				case ">=":
					l.RecordId.AndWhereGte(fval1.Interface())
				case "<":
					l.RecordId.AndWhereLt(fval1.Interface())
				case "<=":
					l.RecordId.AndWhereLte(fval1.Interface())
				case "llike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					l.RecordId.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					l.RecordId.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					l.RecordId.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("ClassCode")
	if exists {
		fval2 := val.FieldByName("ClassCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				l.ClassCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					l.ClassCode.AndWhereEq(fval2.Interface())
				case "!=":
					l.ClassCode.AndWhereNeq(fval2.Interface())
				case ">":
					l.ClassCode.AndWhereGt(fval2.Interface())
				case ">=":
					l.ClassCode.AndWhereGte(fval2.Interface())
				case "<":
					l.ClassCode.AndWhereLt(fval2.Interface())
				case "<=":
					l.ClassCode.AndWhereLte(fval2.Interface())
				case "llike":
					l.ClassCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					l.ClassCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					l.ClassCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					l.ClassCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					l.ClassCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					l.ClassCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					l.ClassCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("ClassName")
	if exists {
		fval3 := val.FieldByName("ClassName")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				l.ClassName.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					l.ClassName.AndWhereEq(fval3.Interface())
				case "!=":
					l.ClassName.AndWhereNeq(fval3.Interface())
				case ">":
					l.ClassName.AndWhereGt(fval3.Interface())
				case ">=":
					l.ClassName.AndWhereGte(fval3.Interface())
				case "<":
					l.ClassName.AndWhereLt(fval3.Interface())
				case "<=":
					l.ClassName.AndWhereLte(fval3.Interface())
				case "llike":
					l.ClassName.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					l.ClassName.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					l.ClassName.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					l.ClassName.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					l.ClassName.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					l.ClassName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					l.ClassName.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("Campus")
	if exists {
		fval4 := val.FieldByName("Campus")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				l.Campus.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					l.Campus.AndWhereEq(fval4.Interface())
				case "!=":
					l.Campus.AndWhereNeq(fval4.Interface())
				case ">":
					l.Campus.AndWhereGt(fval4.Interface())
				case ">=":
					l.Campus.AndWhereGte(fval4.Interface())
				case "<":
					l.Campus.AndWhereLt(fval4.Interface())
				case "<=":
					l.Campus.AndWhereLte(fval4.Interface())
				case "llike":
					l.Campus.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					l.Campus.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					l.Campus.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					l.Campus.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					l.Campus.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("ResearchArea")
	if exists {
		fval5 := val.FieldByName("ResearchArea")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				l.ResearchArea.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					l.ResearchArea.AndWhereEq(fval5.Interface())
				case "!=":
					l.ResearchArea.AndWhereNeq(fval5.Interface())
				case ">":
					l.ResearchArea.AndWhereGt(fval5.Interface())
				case ">=":
					l.ResearchArea.AndWhereGte(fval5.Interface())
				case "<":
					l.ResearchArea.AndWhereLt(fval5.Interface())
				case "<=":
					l.ResearchArea.AndWhereLte(fval5.Interface())
				case "llike":
					l.ResearchArea.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					l.ResearchArea.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					l.ResearchArea.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					l.ResearchArea.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					l.ResearchArea.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					l.ResearchArea.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					l.ResearchArea.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("Grade")
	if exists {
		fval6 := val.FieldByName("Grade")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				l.Grade.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					l.Grade.AndWhereEq(fval6.Interface())
				case "!=":
					l.Grade.AndWhereNeq(fval6.Interface())
				case ">":
					l.Grade.AndWhereGt(fval6.Interface())
				case ">=":
					l.Grade.AndWhereGte(fval6.Interface())
				case "<":
					l.Grade.AndWhereLt(fval6.Interface())
				case "<=":
					l.Grade.AndWhereLte(fval6.Interface())
				case "llike":
					l.Grade.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					l.Grade.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					l.Grade.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					l.Grade.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					l.Grade.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					l.Grade.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					l.Grade.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("TrainingMode")
	if exists {
		fval7 := val.FieldByName("TrainingMode")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				l.TrainingMode.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					l.TrainingMode.AndWhereEq(fval7.Interface())
				case "!=":
					l.TrainingMode.AndWhereNeq(fval7.Interface())
				case ">":
					l.TrainingMode.AndWhereGt(fval7.Interface())
				case ">=":
					l.TrainingMode.AndWhereGte(fval7.Interface())
				case "<":
					l.TrainingMode.AndWhereLt(fval7.Interface())
				case "<=":
					l.TrainingMode.AndWhereLte(fval7.Interface())
				case "llike":
					l.TrainingMode.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					l.TrainingMode.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					l.TrainingMode.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					l.TrainingMode.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					l.TrainingMode.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					l.TrainingMode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					l.TrainingMode.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("EntranceDate")
	if exists {
		fval8 := val.FieldByName("EntranceDate")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				l.EntranceDate.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					l.EntranceDate.AndWhereEq(fval8.Interface())
				case "!=":
					l.EntranceDate.AndWhereNeq(fval8.Interface())
				case ">":
					l.EntranceDate.AndWhereGt(fval8.Interface())
				case ">=":
					l.EntranceDate.AndWhereGte(fval8.Interface())
				case "<":
					l.EntranceDate.AndWhereLt(fval8.Interface())
				case "<=":
					l.EntranceDate.AndWhereLte(fval8.Interface())
				case "llike":
					l.EntranceDate.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					l.EntranceDate.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					l.EntranceDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					l.EntranceDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					l.EntranceDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					l.EntranceDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					l.EntranceDate.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("GraduationDate")
	if exists {
		fval9 := val.FieldByName("GraduationDate")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				l.GraduationDate.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					l.GraduationDate.AndWhereEq(fval9.Interface())
				case "!=":
					l.GraduationDate.AndWhereNeq(fval9.Interface())
				case ">":
					l.GraduationDate.AndWhereGt(fval9.Interface())
				case ">=":
					l.GraduationDate.AndWhereGte(fval9.Interface())
				case "<":
					l.GraduationDate.AndWhereLt(fval9.Interface())
				case "<=":
					l.GraduationDate.AndWhereLte(fval9.Interface())
				case "llike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					l.GraduationDate.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					l.GraduationDate.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					l.GraduationDate.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	ftyp10, exists := typ.FieldByName("ProgramLength")
	if exists {
		fval10 := val.FieldByName("ProgramLength")
		fop10, ok := ftyp10.Tag.Lookup("op")
		for fval10.Kind() == reflect.Ptr && !fval10.IsNil() {
			fval10 = fval10.Elem()
		}
		if fval10.Kind() != reflect.Ptr {
			if !ok {
				l.ProgramLength.AndWhereEq(fval10.Interface())
			} else {
				switch fop10 {
				case "=":
					l.ProgramLength.AndWhereEq(fval10.Interface())
				case "!=":
					l.ProgramLength.AndWhereNeq(fval10.Interface())
				case ">":
					l.ProgramLength.AndWhereGt(fval10.Interface())
				case ">=":
					l.ProgramLength.AndWhereGte(fval10.Interface())
				case "<":
					l.ProgramLength.AndWhereLt(fval10.Interface())
				case "<=":
					l.ProgramLength.AndWhereLte(fval10.Interface())
				case "llike":
					l.ProgramLength.AndWhereLike(fmt.Sprintf("%%%s", fval10.String()))
				case "rlike":
					l.ProgramLength.AndWhereLike(fmt.Sprintf("%s%%", fval10.String()))
				case "alike":
					l.ProgramLength.AndWhereLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "nllike":
					l.ProgramLength.AndWhereNotLike(fmt.Sprintf("%%%s", fval10.String()))
				case "nrlike":
					l.ProgramLength.AndWhereNotLike(fmt.Sprintf("%s%%", fval10.String()))
				case "nalike":
					l.ProgramLength.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval10.String()))
				case "in":
					l.ProgramLength.AndWhereIn(fval10.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop10)
				}
			}
		}
	}
	ftyp11, exists := typ.FieldByName("StudentType")
	if exists {
		fval11 := val.FieldByName("StudentType")
		fop11, ok := ftyp11.Tag.Lookup("op")
		for fval11.Kind() == reflect.Ptr && !fval11.IsNil() {
			fval11 = fval11.Elem()
		}
		if fval11.Kind() != reflect.Ptr {
			if !ok {
				l.StudentType.AndWhereEq(fval11.Interface())
			} else {
				switch fop11 {
				case "=":
					l.StudentType.AndWhereEq(fval11.Interface())
				case "!=":
					l.StudentType.AndWhereNeq(fval11.Interface())
				case ">":
					l.StudentType.AndWhereGt(fval11.Interface())
				case ">=":
					l.StudentType.AndWhereGte(fval11.Interface())
				case "<":
					l.StudentType.AndWhereLt(fval11.Interface())
				case "<=":
					l.StudentType.AndWhereLte(fval11.Interface())
				case "llike":
					l.StudentType.AndWhereLike(fmt.Sprintf("%%%s", fval11.String()))
				case "rlike":
					l.StudentType.AndWhereLike(fmt.Sprintf("%s%%", fval11.String()))
				case "alike":
					l.StudentType.AndWhereLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "nllike":
					l.StudentType.AndWhereNotLike(fmt.Sprintf("%%%s", fval11.String()))
				case "nrlike":
					l.StudentType.AndWhereNotLike(fmt.Sprintf("%s%%", fval11.String()))
				case "nalike":
					l.StudentType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval11.String()))
				case "in":
					l.StudentType.AndWhereIn(fval11.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop11)
				}
			}
		}
	}
	ftyp12, exists := typ.FieldByName("CredentialsType")
	if exists {
		fval12 := val.FieldByName("CredentialsType")
		fop12, ok := ftyp12.Tag.Lookup("op")
		for fval12.Kind() == reflect.Ptr && !fval12.IsNil() {
			fval12 = fval12.Elem()
		}
		if fval12.Kind() != reflect.Ptr {
			if !ok {
				l.CredentialsType.AndWhereEq(fval12.Interface())
			} else {
				switch fop12 {
				case "=":
					l.CredentialsType.AndWhereEq(fval12.Interface())
				case "!=":
					l.CredentialsType.AndWhereNeq(fval12.Interface())
				case ">":
					l.CredentialsType.AndWhereGt(fval12.Interface())
				case ">=":
					l.CredentialsType.AndWhereGte(fval12.Interface())
				case "<":
					l.CredentialsType.AndWhereLt(fval12.Interface())
				case "<=":
					l.CredentialsType.AndWhereLte(fval12.Interface())
				case "llike":
					l.CredentialsType.AndWhereLike(fmt.Sprintf("%%%s", fval12.String()))
				case "rlike":
					l.CredentialsType.AndWhereLike(fmt.Sprintf("%s%%", fval12.String()))
				case "alike":
					l.CredentialsType.AndWhereLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "nllike":
					l.CredentialsType.AndWhereNotLike(fmt.Sprintf("%%%s", fval12.String()))
				case "nrlike":
					l.CredentialsType.AndWhereNotLike(fmt.Sprintf("%s%%", fval12.String()))
				case "nalike":
					l.CredentialsType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval12.String()))
				case "in":
					l.CredentialsType.AndWhereIn(fval12.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop12)
				}
			}
		}
	}
	ftyp13, exists := typ.FieldByName("DegreeType")
	if exists {
		fval13 := val.FieldByName("DegreeType")
		fop13, ok := ftyp13.Tag.Lookup("op")
		for fval13.Kind() == reflect.Ptr && !fval13.IsNil() {
			fval13 = fval13.Elem()
		}
		if fval13.Kind() != reflect.Ptr {
			if !ok {
				l.DegreeType.AndWhereEq(fval13.Interface())
			} else {
				switch fop13 {
				case "=":
					l.DegreeType.AndWhereEq(fval13.Interface())
				case "!=":
					l.DegreeType.AndWhereNeq(fval13.Interface())
				case ">":
					l.DegreeType.AndWhereGt(fval13.Interface())
				case ">=":
					l.DegreeType.AndWhereGte(fval13.Interface())
				case "<":
					l.DegreeType.AndWhereLt(fval13.Interface())
				case "<=":
					l.DegreeType.AndWhereLte(fval13.Interface())
				case "llike":
					l.DegreeType.AndWhereLike(fmt.Sprintf("%%%s", fval13.String()))
				case "rlike":
					l.DegreeType.AndWhereLike(fmt.Sprintf("%s%%", fval13.String()))
				case "alike":
					l.DegreeType.AndWhereLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "nllike":
					l.DegreeType.AndWhereNotLike(fmt.Sprintf("%%%s", fval13.String()))
				case "nrlike":
					l.DegreeType.AndWhereNotLike(fmt.Sprintf("%s%%", fval13.String()))
				case "nalike":
					l.DegreeType.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval13.String()))
				case "in":
					l.DegreeType.AndWhereIn(fval13.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop13)
				}
			}
		}
	}
	ftyp14, exists := typ.FieldByName("Counselor")
	if exists {
		fval14 := val.FieldByName("Counselor")
		fop14, ok := ftyp14.Tag.Lookup("op")
		for fval14.Kind() == reflect.Ptr && !fval14.IsNil() {
			fval14 = fval14.Elem()
		}
		if fval14.Kind() != reflect.Ptr {
			if !ok {
				l.Counselor.AndWhereEq(fval14.Interface())
			} else {
				switch fop14 {
				case "=":
					l.Counselor.AndWhereEq(fval14.Interface())
				case "!=":
					l.Counselor.AndWhereNeq(fval14.Interface())
				case ">":
					l.Counselor.AndWhereGt(fval14.Interface())
				case ">=":
					l.Counselor.AndWhereGte(fval14.Interface())
				case "<":
					l.Counselor.AndWhereLt(fval14.Interface())
				case "<=":
					l.Counselor.AndWhereLte(fval14.Interface())
				case "llike":
					l.Counselor.AndWhereLike(fmt.Sprintf("%%%s", fval14.String()))
				case "rlike":
					l.Counselor.AndWhereLike(fmt.Sprintf("%s%%", fval14.String()))
				case "alike":
					l.Counselor.AndWhereLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "nllike":
					l.Counselor.AndWhereNotLike(fmt.Sprintf("%%%s", fval14.String()))
				case "nrlike":
					l.Counselor.AndWhereNotLike(fmt.Sprintf("%s%%", fval14.String()))
				case "nalike":
					l.Counselor.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval14.String()))
				case "in":
					l.Counselor.AndWhereIn(fval14.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop14)
				}
			}
		}
	}
	ftyp15, exists := typ.FieldByName("Adviser")
	if exists {
		fval15 := val.FieldByName("Adviser")
		fop15, ok := ftyp15.Tag.Lookup("op")
		for fval15.Kind() == reflect.Ptr && !fval15.IsNil() {
			fval15 = fval15.Elem()
		}
		if fval15.Kind() != reflect.Ptr {
			if !ok {
				l.Adviser.AndWhereEq(fval15.Interface())
			} else {
				switch fop15 {
				case "=":
					l.Adviser.AndWhereEq(fval15.Interface())
				case "!=":
					l.Adviser.AndWhereNeq(fval15.Interface())
				case ">":
					l.Adviser.AndWhereGt(fval15.Interface())
				case ">=":
					l.Adviser.AndWhereGte(fval15.Interface())
				case "<":
					l.Adviser.AndWhereLt(fval15.Interface())
				case "<=":
					l.Adviser.AndWhereLte(fval15.Interface())
				case "llike":
					l.Adviser.AndWhereLike(fmt.Sprintf("%%%s", fval15.String()))
				case "rlike":
					l.Adviser.AndWhereLike(fmt.Sprintf("%s%%", fval15.String()))
				case "alike":
					l.Adviser.AndWhereLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "nllike":
					l.Adviser.AndWhereNotLike(fmt.Sprintf("%%%s", fval15.String()))
				case "nrlike":
					l.Adviser.AndWhereNotLike(fmt.Sprintf("%s%%", fval15.String()))
				case "nalike":
					l.Adviser.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval15.String()))
				case "in":
					l.Adviser.AndWhereIn(fval15.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop15)
				}
			}
		}
	}
	ftyp16, exists := typ.FieldByName("Leadership")
	if exists {
		fval16 := val.FieldByName("Leadership")
		fop16, ok := ftyp16.Tag.Lookup("op")
		for fval16.Kind() == reflect.Ptr && !fval16.IsNil() {
			fval16 = fval16.Elem()
		}
		if fval16.Kind() != reflect.Ptr {
			if !ok {
				l.Leadership.AndWhereEq(fval16.Interface())
			} else {
				switch fop16 {
				case "=":
					l.Leadership.AndWhereEq(fval16.Interface())
				case "!=":
					l.Leadership.AndWhereNeq(fval16.Interface())
				case ">":
					l.Leadership.AndWhereGt(fval16.Interface())
				case ">=":
					l.Leadership.AndWhereGte(fval16.Interface())
				case "<":
					l.Leadership.AndWhereLt(fval16.Interface())
				case "<=":
					l.Leadership.AndWhereLte(fval16.Interface())
				case "llike":
					l.Leadership.AndWhereLike(fmt.Sprintf("%%%s", fval16.String()))
				case "rlike":
					l.Leadership.AndWhereLike(fmt.Sprintf("%s%%", fval16.String()))
				case "alike":
					l.Leadership.AndWhereLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "nllike":
					l.Leadership.AndWhereNotLike(fmt.Sprintf("%%%s", fval16.String()))
				case "nrlike":
					l.Leadership.AndWhereNotLike(fmt.Sprintf("%s%%", fval16.String()))
				case "nalike":
					l.Leadership.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval16.String()))
				case "in":
					l.Leadership.AndWhereIn(fval16.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop16)
				}
			}
		}
	}
	ftyp17, exists := typ.FieldByName("Supervisor")
	if exists {
		fval17 := val.FieldByName("Supervisor")
		fop17, ok := ftyp17.Tag.Lookup("op")
		for fval17.Kind() == reflect.Ptr && !fval17.IsNil() {
			fval17 = fval17.Elem()
		}
		if fval17.Kind() != reflect.Ptr {
			if !ok {
				l.Supervisor.AndWhereEq(fval17.Interface())
			} else {
				switch fop17 {
				case "=":
					l.Supervisor.AndWhereEq(fval17.Interface())
				case "!=":
					l.Supervisor.AndWhereNeq(fval17.Interface())
				case ">":
					l.Supervisor.AndWhereGt(fval17.Interface())
				case ">=":
					l.Supervisor.AndWhereGte(fval17.Interface())
				case "<":
					l.Supervisor.AndWhereLt(fval17.Interface())
				case "<=":
					l.Supervisor.AndWhereLte(fval17.Interface())
				case "llike":
					l.Supervisor.AndWhereLike(fmt.Sprintf("%%%s", fval17.String()))
				case "rlike":
					l.Supervisor.AndWhereLike(fmt.Sprintf("%s%%", fval17.String()))
				case "alike":
					l.Supervisor.AndWhereLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "nllike":
					l.Supervisor.AndWhereNotLike(fmt.Sprintf("%%%s", fval17.String()))
				case "nrlike":
					l.Supervisor.AndWhereNotLike(fmt.Sprintf("%s%%", fval17.String()))
				case "nalike":
					l.Supervisor.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval17.String()))
				case "in":
					l.Supervisor.AndWhereIn(fval17.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop17)
				}
			}
		}
	}
	ftyp18, exists := typ.FieldByName("Assistant1")
	if exists {
		fval18 := val.FieldByName("Assistant1")
		fop18, ok := ftyp18.Tag.Lookup("op")
		for fval18.Kind() == reflect.Ptr && !fval18.IsNil() {
			fval18 = fval18.Elem()
		}
		if fval18.Kind() != reflect.Ptr {
			if !ok {
				l.Assistant1.AndWhereEq(fval18.Interface())
			} else {
				switch fop18 {
				case "=":
					l.Assistant1.AndWhereEq(fval18.Interface())
				case "!=":
					l.Assistant1.AndWhereNeq(fval18.Interface())
				case ">":
					l.Assistant1.AndWhereGt(fval18.Interface())
				case ">=":
					l.Assistant1.AndWhereGte(fval18.Interface())
				case "<":
					l.Assistant1.AndWhereLt(fval18.Interface())
				case "<=":
					l.Assistant1.AndWhereLte(fval18.Interface())
				case "llike":
					l.Assistant1.AndWhereLike(fmt.Sprintf("%%%s", fval18.String()))
				case "rlike":
					l.Assistant1.AndWhereLike(fmt.Sprintf("%s%%", fval18.String()))
				case "alike":
					l.Assistant1.AndWhereLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "nllike":
					l.Assistant1.AndWhereNotLike(fmt.Sprintf("%%%s", fval18.String()))
				case "nrlike":
					l.Assistant1.AndWhereNotLike(fmt.Sprintf("%s%%", fval18.String()))
				case "nalike":
					l.Assistant1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval18.String()))
				case "in":
					l.Assistant1.AndWhereIn(fval18.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop18)
				}
			}
		}
	}
	ftyp19, exists := typ.FieldByName("Assistant2")
	if exists {
		fval19 := val.FieldByName("Assistant2")
		fop19, ok := ftyp19.Tag.Lookup("op")
		for fval19.Kind() == reflect.Ptr && !fval19.IsNil() {
			fval19 = fval19.Elem()
		}
		if fval19.Kind() != reflect.Ptr {
			if !ok {
				l.Assistant2.AndWhereEq(fval19.Interface())
			} else {
				switch fop19 {
				case "=":
					l.Assistant2.AndWhereEq(fval19.Interface())
				case "!=":
					l.Assistant2.AndWhereNeq(fval19.Interface())
				case ">":
					l.Assistant2.AndWhereGt(fval19.Interface())
				case ">=":
					l.Assistant2.AndWhereGte(fval19.Interface())
				case "<":
					l.Assistant2.AndWhereLt(fval19.Interface())
				case "<=":
					l.Assistant2.AndWhereLte(fval19.Interface())
				case "llike":
					l.Assistant2.AndWhereLike(fmt.Sprintf("%%%s", fval19.String()))
				case "rlike":
					l.Assistant2.AndWhereLike(fmt.Sprintf("%s%%", fval19.String()))
				case "alike":
					l.Assistant2.AndWhereLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "nllike":
					l.Assistant2.AndWhereNotLike(fmt.Sprintf("%%%s", fval19.String()))
				case "nrlike":
					l.Assistant2.AndWhereNotLike(fmt.Sprintf("%s%%", fval19.String()))
				case "nalike":
					l.Assistant2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval19.String()))
				case "in":
					l.Assistant2.AndWhereIn(fval19.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop19)
				}
			}
		}
	}
	ftyp20, exists := typ.FieldByName("Operator")
	if exists {
		fval20 := val.FieldByName("Operator")
		fop20, ok := ftyp20.Tag.Lookup("op")
		for fval20.Kind() == reflect.Ptr && !fval20.IsNil() {
			fval20 = fval20.Elem()
		}
		if fval20.Kind() != reflect.Ptr {
			if !ok {
				l.Operator.AndWhereEq(fval20.Interface())
			} else {
				switch fop20 {
				case "=":
					l.Operator.AndWhereEq(fval20.Interface())
				case "!=":
					l.Operator.AndWhereNeq(fval20.Interface())
				case ">":
					l.Operator.AndWhereGt(fval20.Interface())
				case ">=":
					l.Operator.AndWhereGte(fval20.Interface())
				case "<":
					l.Operator.AndWhereLt(fval20.Interface())
				case "<=":
					l.Operator.AndWhereLte(fval20.Interface())
				case "llike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s", fval20.String()))
				case "rlike":
					l.Operator.AndWhereLike(fmt.Sprintf("%s%%", fval20.String()))
				case "alike":
					l.Operator.AndWhereLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "nllike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s", fval20.String()))
				case "nrlike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%s%%", fval20.String()))
				case "nalike":
					l.Operator.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval20.String()))
				case "in":
					l.Operator.AndWhereIn(fval20.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop20)
				}
			}
		}
	}
	ftyp21, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval21 := val.FieldByName("InsertDatetime")
		fop21, ok := ftyp21.Tag.Lookup("op")
		for fval21.Kind() == reflect.Ptr && !fval21.IsNil() {
			fval21 = fval21.Elem()
		}
		if fval21.Kind() != reflect.Ptr {
			if !ok {
				l.InsertDatetime.AndWhereEq(fval21.Interface())
			} else {
				switch fop21 {
				case "=":
					l.InsertDatetime.AndWhereEq(fval21.Interface())
				case "!=":
					l.InsertDatetime.AndWhereNeq(fval21.Interface())
				case ">":
					l.InsertDatetime.AndWhereGt(fval21.Interface())
				case ">=":
					l.InsertDatetime.AndWhereGte(fval21.Interface())
				case "<":
					l.InsertDatetime.AndWhereLt(fval21.Interface())
				case "<=":
					l.InsertDatetime.AndWhereLte(fval21.Interface())
				case "llike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval21.String()))
				case "rlike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval21.String()))
				case "alike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "nllike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval21.String()))
				case "nrlike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval21.String()))
				case "nalike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval21.String()))
				case "in":
					l.InsertDatetime.AndWhereIn(fval21.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop21)
				}
			}
		}
	}
	ftyp22, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval22 := val.FieldByName("UpdateDatetime")
		fop22, ok := ftyp22.Tag.Lookup("op")
		for fval22.Kind() == reflect.Ptr && !fval22.IsNil() {
			fval22 = fval22.Elem()
		}
		if fval22.Kind() != reflect.Ptr {
			if !ok {
				l.UpdateDatetime.AndWhereEq(fval22.Interface())
			} else {
				switch fop22 {
				case "=":
					l.UpdateDatetime.AndWhereEq(fval22.Interface())
				case "!=":
					l.UpdateDatetime.AndWhereNeq(fval22.Interface())
				case ">":
					l.UpdateDatetime.AndWhereGt(fval22.Interface())
				case ">=":
					l.UpdateDatetime.AndWhereGte(fval22.Interface())
				case "<":
					l.UpdateDatetime.AndWhereLt(fval22.Interface())
				case "<=":
					l.UpdateDatetime.AndWhereLte(fval22.Interface())
				case "llike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval22.String()))
				case "rlike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval22.String()))
				case "alike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "nllike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval22.String()))
				case "nrlike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval22.String()))
				case "nalike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval22.String()))
				case "in":
					l.UpdateDatetime.AndWhereIn(fval22.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop22)
				}
			}
		}
	}
	ftyp23, exists := typ.FieldByName("Status")
	if exists {
		fval23 := val.FieldByName("Status")
		fop23, ok := ftyp23.Tag.Lookup("op")
		for fval23.Kind() == reflect.Ptr && !fval23.IsNil() {
			fval23 = fval23.Elem()
		}
		if fval23.Kind() != reflect.Ptr {
			if !ok {
				l.Status.AndWhereEq(fval23.Interface())
			} else {
				switch fop23 {
				case "=":
					l.Status.AndWhereEq(fval23.Interface())
				case "!=":
					l.Status.AndWhereNeq(fval23.Interface())
				case ">":
					l.Status.AndWhereGt(fval23.Interface())
				case ">=":
					l.Status.AndWhereGte(fval23.Interface())
				case "<":
					l.Status.AndWhereLt(fval23.Interface())
				case "<=":
					l.Status.AndWhereLte(fval23.Interface())
				case "llike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s", fval23.String()))
				case "rlike":
					l.Status.AndWhereLike(fmt.Sprintf("%s%%", fval23.String()))
				case "alike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "nllike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval23.String()))
				case "nrlike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval23.String()))
				case "nalike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval23.String()))
				case "in":
					l.Status.AndWhereIn(fval23.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop23)
				}
			}
		}
	}
	ftyp24, exists := typ.FieldByName("Remark1")
	if exists {
		fval24 := val.FieldByName("Remark1")
		fop24, ok := ftyp24.Tag.Lookup("op")
		for fval24.Kind() == reflect.Ptr && !fval24.IsNil() {
			fval24 = fval24.Elem()
		}
		if fval24.Kind() != reflect.Ptr {
			if !ok {
				l.Remark1.AndWhereEq(fval24.Interface())
			} else {
				switch fop24 {
				case "=":
					l.Remark1.AndWhereEq(fval24.Interface())
				case "!=":
					l.Remark1.AndWhereNeq(fval24.Interface())
				case ">":
					l.Remark1.AndWhereGt(fval24.Interface())
				case ">=":
					l.Remark1.AndWhereGte(fval24.Interface())
				case "<":
					l.Remark1.AndWhereLt(fval24.Interface())
				case "<=":
					l.Remark1.AndWhereLte(fval24.Interface())
				case "llike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval24.String()))
				case "rlike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval24.String()))
				case "alike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "nllike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval24.String()))
				case "nrlike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval24.String()))
				case "nalike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval24.String()))
				case "in":
					l.Remark1.AndWhereIn(fval24.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop24)
				}
			}
		}
	}
	ftyp25, exists := typ.FieldByName("Remark2")
	if exists {
		fval25 := val.FieldByName("Remark2")
		fop25, ok := ftyp25.Tag.Lookup("op")
		for fval25.Kind() == reflect.Ptr && !fval25.IsNil() {
			fval25 = fval25.Elem()
		}
		if fval25.Kind() != reflect.Ptr {
			if !ok {
				l.Remark2.AndWhereEq(fval25.Interface())
			} else {
				switch fop25 {
				case "=":
					l.Remark2.AndWhereEq(fval25.Interface())
				case "!=":
					l.Remark2.AndWhereNeq(fval25.Interface())
				case ">":
					l.Remark2.AndWhereGt(fval25.Interface())
				case ">=":
					l.Remark2.AndWhereGte(fval25.Interface())
				case "<":
					l.Remark2.AndWhereLt(fval25.Interface())
				case "<=":
					l.Remark2.AndWhereLte(fval25.Interface())
				case "llike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval25.String()))
				case "rlike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval25.String()))
				case "alike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "nllike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval25.String()))
				case "nrlike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval25.String()))
				case "nalike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval25.String()))
				case "in":
					l.Remark2.AndWhereIn(fval25.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop25)
				}
			}
		}
	}
	ftyp26, exists := typ.FieldByName("Remark3")
	if exists {
		fval26 := val.FieldByName("Remark3")
		fop26, ok := ftyp26.Tag.Lookup("op")
		for fval26.Kind() == reflect.Ptr && !fval26.IsNil() {
			fval26 = fval26.Elem()
		}
		if fval26.Kind() != reflect.Ptr {
			if !ok {
				l.Remark3.AndWhereEq(fval26.Interface())
			} else {
				switch fop26 {
				case "=":
					l.Remark3.AndWhereEq(fval26.Interface())
				case "!=":
					l.Remark3.AndWhereNeq(fval26.Interface())
				case ">":
					l.Remark3.AndWhereGt(fval26.Interface())
				case ">=":
					l.Remark3.AndWhereGte(fval26.Interface())
				case "<":
					l.Remark3.AndWhereLt(fval26.Interface())
				case "<=":
					l.Remark3.AndWhereLte(fval26.Interface())
				case "llike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval26.String()))
				case "rlike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval26.String()))
				case "alike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "nllike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval26.String()))
				case "nrlike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval26.String()))
				case "nalike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval26.String()))
				case "in":
					l.Remark3.AndWhereIn(fval26.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop26)
				}
			}
		}
	}
	ftyp27, exists := typ.FieldByName("Remark4")
	if exists {
		fval27 := val.FieldByName("Remark4")
		fop27, ok := ftyp27.Tag.Lookup("op")
		for fval27.Kind() == reflect.Ptr && !fval27.IsNil() {
			fval27 = fval27.Elem()
		}
		if fval27.Kind() != reflect.Ptr {
			if !ok {
				l.Remark4.AndWhereEq(fval27.Interface())
			} else {
				switch fop27 {
				case "=":
					l.Remark4.AndWhereEq(fval27.Interface())
				case "!=":
					l.Remark4.AndWhereNeq(fval27.Interface())
				case ">":
					l.Remark4.AndWhereGt(fval27.Interface())
				case ">=":
					l.Remark4.AndWhereGte(fval27.Interface())
				case "<":
					l.Remark4.AndWhereLt(fval27.Interface())
				case "<=":
					l.Remark4.AndWhereLte(fval27.Interface())
				case "llike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval27.String()))
				case "rlike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval27.String()))
				case "alike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "nllike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval27.String()))
				case "nrlike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval27.String()))
				case "nalike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval27.String()))
				case "in":
					l.Remark4.AndWhereIn(fval27.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop27)
				}
			}
		}
	}
	return l, nil
}
func NewGrade() *Grade {
	m := &Grade{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.GradeName.Init(m, "GradeName", "GradeName", "GradeName", "GradeName", 1)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", "GradeCode", "GradeCode", 2)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 3)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 4)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 5)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 6)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 7)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 8)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 9)
	m.InitRel()
	return m
}

func newSubGrade(parent nborm.Model) *Grade {
	m := &Grade{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	m.GradeName.Init(m, "GradeName", "GradeName", "GradeName", "GradeName", 1)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", "GradeCode", "GradeCode", 2)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 3)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 4)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 5)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 6)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 7)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 8)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 9)
	return m
}

func (m *Grade) InitRel() {
	m.Classes = newSubClassList(m)
	m.Classes.dupMap = make(map[string]int)
	var relInfo0 *nborm.RelationInfo
	relInfo0 = relInfo0.Append("Classes", m.Classes, nborm.NewExpr("@=@", &m.GradeCode, &m.Classes.Grade))
	m.AppendRelation(relInfo0)
	m.AddRelInited()
}

func (m *Grade) DB() string {
	return "*"
}

func (m *Grade) Tab() string {
	return "grade"
}

func (m *Grade) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"Id", "Id", &m.Id, 0},
		{"GradeName", "GradeName", &m.GradeName, 1},
		{"GradeCode", "GradeCode", &m.GradeCode, 2},
		{"InsertDatetime", "InsertDatetime", &m.InsertDatetime, 3},
		{"UpdateDatetime", "UpdateDatetime", &m.UpdateDatetime, 4},
		{"Status", "Status", &m.Status, 5},
		{"Remark1", "Remark1", &m.Remark1, 6},
		{"Remark2", "Remark2", &m.Remark2, 7},
		{"Remark3", "Remark3", &m.Remark3, 8},
		{"Remark4", "Remark4", &m.Remark4, 9},
	}
}

func (m *Grade) AutoIncField() nborm.Field {
	return &m.Id
}

func (m *Grade) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.Id,
	}
}

func (m *Grade) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.GradeCode,
			&m.Status,
		},
	}
}
func (m Grade) MarshalJSON() ([]byte, error) {
	if !m.IsSynced() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	buffer.WriteString("{\n\"Aggs\": ")
	metaB, err := json.MarshalIndent(m.Meta, "", "\t")
	if err != nil {
		return nil, err
	}
	buffer.Write(metaB)
	if m.Id.IsValid() {
		buffer.WriteString(",\n\"Id\": ")
		IdB, err := json.MarshalIndent(m.Id, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IdB)
	}
	if m.GradeName.IsValid() {
		buffer.WriteString(",\n\"GradeName\": ")
		GradeNameB, err := json.MarshalIndent(m.GradeName, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GradeNameB)
	}
	if m.GradeCode.IsValid() {
		buffer.WriteString(",\n\"GradeCode\": ")
		GradeCodeB, err := json.MarshalIndent(m.GradeCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(GradeCodeB)
	}
	if m.InsertDatetime.IsValid() {
		buffer.WriteString(",\n\"InsertDatetime\": ")
		InsertDatetimeB, err := json.MarshalIndent(m.InsertDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(InsertDatetimeB)
	}
	if m.UpdateDatetime.IsValid() {
		buffer.WriteString(",\n\"UpdateDatetime\": ")
		UpdateDatetimeB, err := json.MarshalIndent(m.UpdateDatetime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateDatetimeB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.Remark1.IsValid() {
		buffer.WriteString(",\n\"Remark1\": ")
		Remark1B, err := json.MarshalIndent(m.Remark1, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark1B)
	}
	if m.Remark2.IsValid() {
		buffer.WriteString(",\n\"Remark2\": ")
		Remark2B, err := json.MarshalIndent(m.Remark2, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark2B)
	}
	if m.Remark3.IsValid() {
		buffer.WriteString(",\n\"Remark3\": ")
		Remark3B, err := json.MarshalIndent(m.Remark3, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark3B)
	}
	if m.Remark4.IsValid() {
		buffer.WriteString(",\n\"Remark4\": ")
		Remark4B, err := json.MarshalIndent(m.Remark4, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(Remark4B)
	}
	if m.Classes != nil && m.Classes.Len() > 0 {
		buffer.WriteString(",\n\"Classes\": ")
		ClassesB, err := json.MarshalIndent(m.Classes, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ClassesB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type GradeList struct {
	Grade  `json:"-"`
	dupMap map[string]int
	List   []*Grade
	Total  int
}

func (m *Grade) Collapse() {
	if m.Classes != nil && m.Classes.IsSynced() {
		m.Classes.Collapse()
	}
}

func NewGradeList() *GradeList {
	l := &GradeList{
		Grade{},
		make(map[string]int),
		make([]*Grade, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.GradeName.Init(l, "GradeName", "GradeName", "GradeName", "GradeName", 1)
	l.GradeCode.Init(l, "GradeCode", "GradeCode", "GradeCode", "GradeCode", 2)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 3)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 4)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 5)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 6)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 7)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 8)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 9)
	l.InitRel()
	return l
}

func newSubGradeList(parent nborm.Model) *GradeList {
	l := &GradeList{
		Grade{},
		make(map[string]int),
		make([]*Grade, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.Id.Init(l, "Id", "Id", "Id", "Id", 0)
	l.GradeName.Init(l, "GradeName", "GradeName", "GradeName", "GradeName", 1)
	l.GradeCode.Init(l, "GradeCode", "GradeCode", "GradeCode", "GradeCode", 2)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 3)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 4)
	l.Status.Init(l, "Status", "Status", "Status", "Status", 5)
	l.Remark1.Init(l, "Remark1", "Remark1", "Remark1", "Remark1", 6)
	l.Remark2.Init(l, "Remark2", "Remark2", "Remark2", "Remark2", 7)
	l.Remark3.Init(l, "Remark3", "Remark3", "Remark3", "Remark3", 8)
	l.Remark4.Init(l, "Remark4", "Remark4", "Remark4", "Remark4", 9)
	return l
}

func (l *GradeList) NewModel() nborm.Model {
	m := &Grade{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.GradeName.Init(m, "GradeName", "GradeName", "GradeName", "GradeName", 1)
	l.GradeName.CopyStatus(&m.GradeName)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", "GradeCode", "GradeCode", 2)
	l.GradeCode.CopyStatus(&m.GradeCode)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", "InsertDatetime", "InsertDatetime", 3)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", "UpdateDatetime", 4)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", "Status", "Status", 5)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", "Remark1", "Remark1", 6)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", "Remark2", "Remark2", 7)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", "Remark3", "Remark3", 8)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", "Remark4", "Remark4", 9)
	l.Remark4.CopyStatus(&m.Remark4)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *GradeList) SetTotal(total int) {
	l.Total = total
}

func (l *GradeList) GetTotal() int {
	return l.Total
}

func (l *GradeList) Len() int {
	return len(l.List)
}

func (l *GradeList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l GradeList) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, 1024)
	bs = append(bs, []byte("{")...)
	ListB, err := json.MarshalIndent(l.List, "", "\t")
	if err != nil {
		return nil, err
	}
	ListB = append([]byte("\"List\": "), ListB...)
	bs = append(bs, ListB...)
	bs = append(bs, []byte(", ")...)
	TotalB, err := json.MarshalIndent(l.Total, "", "\t")
	if err != nil {
		return nil, err
	}
	TotalB = append([]byte("\"Total\": "), TotalB...)
	bs = append(bs, TotalB...)
	bs = append(bs, []byte("}")...)
	return bs, nil
}

func (l *GradeList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &l.Grade)
}

func (l *GradeList) UnmarshalMeta(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &l.Grade)
}

func (l *GradeList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List[idx].Classes.checkDup()
		l.List[idx].Classes.List = append(l.List[idx].Classes.List, l.List[l.Len()-1].Classes.List...)
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *GradeList) Filter(f func(m *Grade) bool) []*Grade {
	ll := make([]*Grade, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *GradeList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.Id.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Id.AnyValue()))
	}
	if lastModel.GradeName.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.GradeName.AnyValue()))
	}
	if lastModel.GradeCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.GradeCode.AnyValue()))
	}
	if lastModel.InsertDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.InsertDatetime.AnyValue()))
	}
	if lastModel.UpdateDatetime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateDatetime.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.Remark1.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark1.AnyValue()))
	}
	if lastModel.Remark2.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark2.AnyValue()))
	}
	if lastModel.Remark3.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark3.AnyValue()))
	}
	if lastModel.Remark4.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Remark4.AnyValue()))
	}
	builder.WriteString(lastModel.AggCheckDup())
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *GradeList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

func (m *Grade) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (l *GradeList) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

type GradeCacheElem struct {
	hashValue  string
	model      *Grade
	modifyTime time.Time
}

type GradeListCacheElem struct {
	hashValue  string
	list       *GradeList
	modifyTime time.Time
}

type GradeCacheManager struct {
	container map[string]*GradeCacheElem
	query     chan string
	in        chan *GradeCacheElem
	out       chan *GradeCacheElem
}

type GradeListCacheManager struct {
	container map[string]*GradeListCacheElem
	query     chan string
	in        chan *GradeListCacheElem
	out       chan *GradeListCacheElem
}

func newGradeCacheManager() *GradeCacheManager {
	return &GradeCacheManager{
		make(map[string]*GradeCacheElem),
		make(chan string),
		make(chan *GradeCacheElem),
		make(chan *GradeCacheElem),
	}
}

func newGradeListCacheManager() *GradeListCacheManager {
	return &GradeListCacheManager{
		make(map[string]*GradeListCacheElem),
		make(chan string),
		make(chan *GradeListCacheElem),
		make(chan *GradeListCacheElem),
	}
}

func (mgr *GradeCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *GradeListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var GradeCache = newGradeCacheManager()

var GradeListCache = newGradeListCacheManager()

func (m *Grade) GetCache(hashVal string, timeout time.Duration) bool {
	GradeCache.query <- hashVal
	elem := <-GradeCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *Grade) SetCache(hashValue string) {
	GradeCache.in <- &GradeCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *GradeList) GetListCache(hashValue string, timeout time.Duration) bool {
	GradeListCache.query <- hashValue
	elem := <-GradeListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *GradeList) SetListCache(hashValue string) {
	GradeListCache.in <- &GradeListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func (m *Grade) FromQuery(query interface{}) (*Grade, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				m.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					m.Id.AndWhereEq(fval0.Interface())
				case "!=":
					m.Id.AndWhereNeq(fval0.Interface())
				case ">":
					m.Id.AndWhereGt(fval0.Interface())
				case ">=":
					m.Id.AndWhereGte(fval0.Interface())
				case "<":
					m.Id.AndWhereLt(fval0.Interface())
				case "<=":
					m.Id.AndWhereLte(fval0.Interface())
				case "llike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					m.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					m.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					m.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					m.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("GradeName")
	if exists {
		fval1 := val.FieldByName("GradeName")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				m.GradeName.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					m.GradeName.AndWhereEq(fval1.Interface())
				case "!=":
					m.GradeName.AndWhereNeq(fval1.Interface())
				case ">":
					m.GradeName.AndWhereGt(fval1.Interface())
				case ">=":
					m.GradeName.AndWhereGte(fval1.Interface())
				case "<":
					m.GradeName.AndWhereLt(fval1.Interface())
				case "<=":
					m.GradeName.AndWhereLte(fval1.Interface())
				case "llike":
					m.GradeName.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					m.GradeName.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					m.GradeName.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					m.GradeName.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					m.GradeName.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					m.GradeName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					m.GradeName.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("GradeCode")
	if exists {
		fval2 := val.FieldByName("GradeCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				m.GradeCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					m.GradeCode.AndWhereEq(fval2.Interface())
				case "!=":
					m.GradeCode.AndWhereNeq(fval2.Interface())
				case ">":
					m.GradeCode.AndWhereGt(fval2.Interface())
				case ">=":
					m.GradeCode.AndWhereGte(fval2.Interface())
				case "<":
					m.GradeCode.AndWhereLt(fval2.Interface())
				case "<=":
					m.GradeCode.AndWhereLte(fval2.Interface())
				case "llike":
					m.GradeCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					m.GradeCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					m.GradeCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					m.GradeCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					m.GradeCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					m.GradeCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					m.GradeCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval3 := val.FieldByName("InsertDatetime")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				m.InsertDatetime.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					m.InsertDatetime.AndWhereEq(fval3.Interface())
				case "!=":
					m.InsertDatetime.AndWhereNeq(fval3.Interface())
				case ">":
					m.InsertDatetime.AndWhereGt(fval3.Interface())
				case ">=":
					m.InsertDatetime.AndWhereGte(fval3.Interface())
				case "<":
					m.InsertDatetime.AndWhereLt(fval3.Interface())
				case "<=":
					m.InsertDatetime.AndWhereLte(fval3.Interface())
				case "llike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					m.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					m.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					m.InsertDatetime.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval4 := val.FieldByName("UpdateDatetime")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				m.UpdateDatetime.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					m.UpdateDatetime.AndWhereEq(fval4.Interface())
				case "!=":
					m.UpdateDatetime.AndWhereNeq(fval4.Interface())
				case ">":
					m.UpdateDatetime.AndWhereGt(fval4.Interface())
				case ">=":
					m.UpdateDatetime.AndWhereGte(fval4.Interface())
				case "<":
					m.UpdateDatetime.AndWhereLt(fval4.Interface())
				case "<=":
					m.UpdateDatetime.AndWhereLte(fval4.Interface())
				case "llike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					m.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					m.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					m.UpdateDatetime.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("Status")
	if exists {
		fval5 := val.FieldByName("Status")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				m.Status.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					m.Status.AndWhereEq(fval5.Interface())
				case "!=":
					m.Status.AndWhereNeq(fval5.Interface())
				case ">":
					m.Status.AndWhereGt(fval5.Interface())
				case ">=":
					m.Status.AndWhereGte(fval5.Interface())
				case "<":
					m.Status.AndWhereLt(fval5.Interface())
				case "<=":
					m.Status.AndWhereLte(fval5.Interface())
				case "llike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					m.Status.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					m.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					m.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					m.Status.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("Remark1")
	if exists {
		fval6 := val.FieldByName("Remark1")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				m.Remark1.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					m.Remark1.AndWhereEq(fval6.Interface())
				case "!=":
					m.Remark1.AndWhereNeq(fval6.Interface())
				case ">":
					m.Remark1.AndWhereGt(fval6.Interface())
				case ">=":
					m.Remark1.AndWhereGte(fval6.Interface())
				case "<":
					m.Remark1.AndWhereLt(fval6.Interface())
				case "<=":
					m.Remark1.AndWhereLte(fval6.Interface())
				case "llike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					m.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					m.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					m.Remark1.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("Remark2")
	if exists {
		fval7 := val.FieldByName("Remark2")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				m.Remark2.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					m.Remark2.AndWhereEq(fval7.Interface())
				case "!=":
					m.Remark2.AndWhereNeq(fval7.Interface())
				case ">":
					m.Remark2.AndWhereGt(fval7.Interface())
				case ">=":
					m.Remark2.AndWhereGte(fval7.Interface())
				case "<":
					m.Remark2.AndWhereLt(fval7.Interface())
				case "<=":
					m.Remark2.AndWhereLte(fval7.Interface())
				case "llike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					m.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					m.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					m.Remark2.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("Remark3")
	if exists {
		fval8 := val.FieldByName("Remark3")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				m.Remark3.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					m.Remark3.AndWhereEq(fval8.Interface())
				case "!=":
					m.Remark3.AndWhereNeq(fval8.Interface())
				case ">":
					m.Remark3.AndWhereGt(fval8.Interface())
				case ">=":
					m.Remark3.AndWhereGte(fval8.Interface())
				case "<":
					m.Remark3.AndWhereLt(fval8.Interface())
				case "<=":
					m.Remark3.AndWhereLte(fval8.Interface())
				case "llike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					m.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					m.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					m.Remark3.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("Remark4")
	if exists {
		fval9 := val.FieldByName("Remark4")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				m.Remark4.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					m.Remark4.AndWhereEq(fval9.Interface())
				case "!=":
					m.Remark4.AndWhereNeq(fval9.Interface())
				case ">":
					m.Remark4.AndWhereGt(fval9.Interface())
				case ">=":
					m.Remark4.AndWhereGte(fval9.Interface())
				case "<":
					m.Remark4.AndWhereLt(fval9.Interface())
				case "<=":
					m.Remark4.AndWhereLte(fval9.Interface())
				case "llike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					m.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					m.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					m.Remark4.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	return m, nil
}
func (l *GradeList) FromQuery(query interface{}) (*GradeList, error) {
	val, typ := reflect.ValueOf(query), reflect.TypeOf(query)
	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("FromQuery() only support struct: %s(%s)", typ.Name(), typ.Kind())
	}
	ftyp0, exists := typ.FieldByName("Id")
	if exists {
		fval0 := val.FieldByName("Id")
		fop0, ok := ftyp0.Tag.Lookup("op")
		for fval0.Kind() == reflect.Ptr && !fval0.IsNil() {
			fval0 = fval0.Elem()
		}
		if fval0.Kind() != reflect.Ptr {
			if !ok {
				l.Id.AndWhereEq(fval0.Interface())
			} else {
				switch fop0 {
				case "=":
					l.Id.AndWhereEq(fval0.Interface())
				case "!=":
					l.Id.AndWhereNeq(fval0.Interface())
				case ">":
					l.Id.AndWhereGt(fval0.Interface())
				case ">=":
					l.Id.AndWhereGte(fval0.Interface())
				case "<":
					l.Id.AndWhereLt(fval0.Interface())
				case "<=":
					l.Id.AndWhereLte(fval0.Interface())
				case "llike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s", fval0.String()))
				case "rlike":
					l.Id.AndWhereLike(fmt.Sprintf("%s%%", fval0.String()))
				case "alike":
					l.Id.AndWhereLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "nllike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s", fval0.String()))
				case "nrlike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%s%%", fval0.String()))
				case "nalike":
					l.Id.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval0.String()))
				case "in":
					l.Id.AndWhereIn(fval0.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop0)
				}
			}
		}
	}
	ftyp1, exists := typ.FieldByName("GradeName")
	if exists {
		fval1 := val.FieldByName("GradeName")
		fop1, ok := ftyp1.Tag.Lookup("op")
		for fval1.Kind() == reflect.Ptr && !fval1.IsNil() {
			fval1 = fval1.Elem()
		}
		if fval1.Kind() != reflect.Ptr {
			if !ok {
				l.GradeName.AndWhereEq(fval1.Interface())
			} else {
				switch fop1 {
				case "=":
					l.GradeName.AndWhereEq(fval1.Interface())
				case "!=":
					l.GradeName.AndWhereNeq(fval1.Interface())
				case ">":
					l.GradeName.AndWhereGt(fval1.Interface())
				case ">=":
					l.GradeName.AndWhereGte(fval1.Interface())
				case "<":
					l.GradeName.AndWhereLt(fval1.Interface())
				case "<=":
					l.GradeName.AndWhereLte(fval1.Interface())
				case "llike":
					l.GradeName.AndWhereLike(fmt.Sprintf("%%%s", fval1.String()))
				case "rlike":
					l.GradeName.AndWhereLike(fmt.Sprintf("%s%%", fval1.String()))
				case "alike":
					l.GradeName.AndWhereLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "nllike":
					l.GradeName.AndWhereNotLike(fmt.Sprintf("%%%s", fval1.String()))
				case "nrlike":
					l.GradeName.AndWhereNotLike(fmt.Sprintf("%s%%", fval1.String()))
				case "nalike":
					l.GradeName.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval1.String()))
				case "in":
					l.GradeName.AndWhereIn(fval1.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop1)
				}
			}
		}
	}
	ftyp2, exists := typ.FieldByName("GradeCode")
	if exists {
		fval2 := val.FieldByName("GradeCode")
		fop2, ok := ftyp2.Tag.Lookup("op")
		for fval2.Kind() == reflect.Ptr && !fval2.IsNil() {
			fval2 = fval2.Elem()
		}
		if fval2.Kind() != reflect.Ptr {
			if !ok {
				l.GradeCode.AndWhereEq(fval2.Interface())
			} else {
				switch fop2 {
				case "=":
					l.GradeCode.AndWhereEq(fval2.Interface())
				case "!=":
					l.GradeCode.AndWhereNeq(fval2.Interface())
				case ">":
					l.GradeCode.AndWhereGt(fval2.Interface())
				case ">=":
					l.GradeCode.AndWhereGte(fval2.Interface())
				case "<":
					l.GradeCode.AndWhereLt(fval2.Interface())
				case "<=":
					l.GradeCode.AndWhereLte(fval2.Interface())
				case "llike":
					l.GradeCode.AndWhereLike(fmt.Sprintf("%%%s", fval2.String()))
				case "rlike":
					l.GradeCode.AndWhereLike(fmt.Sprintf("%s%%", fval2.String()))
				case "alike":
					l.GradeCode.AndWhereLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "nllike":
					l.GradeCode.AndWhereNotLike(fmt.Sprintf("%%%s", fval2.String()))
				case "nrlike":
					l.GradeCode.AndWhereNotLike(fmt.Sprintf("%s%%", fval2.String()))
				case "nalike":
					l.GradeCode.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval2.String()))
				case "in":
					l.GradeCode.AndWhereIn(fval2.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop2)
				}
			}
		}
	}
	ftyp3, exists := typ.FieldByName("InsertDatetime")
	if exists {
		fval3 := val.FieldByName("InsertDatetime")
		fop3, ok := ftyp3.Tag.Lookup("op")
		for fval3.Kind() == reflect.Ptr && !fval3.IsNil() {
			fval3 = fval3.Elem()
		}
		if fval3.Kind() != reflect.Ptr {
			if !ok {
				l.InsertDatetime.AndWhereEq(fval3.Interface())
			} else {
				switch fop3 {
				case "=":
					l.InsertDatetime.AndWhereEq(fval3.Interface())
				case "!=":
					l.InsertDatetime.AndWhereNeq(fval3.Interface())
				case ">":
					l.InsertDatetime.AndWhereGt(fval3.Interface())
				case ">=":
					l.InsertDatetime.AndWhereGte(fval3.Interface())
				case "<":
					l.InsertDatetime.AndWhereLt(fval3.Interface())
				case "<=":
					l.InsertDatetime.AndWhereLte(fval3.Interface())
				case "llike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval3.String()))
				case "rlike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval3.String()))
				case "alike":
					l.InsertDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "nllike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval3.String()))
				case "nrlike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval3.String()))
				case "nalike":
					l.InsertDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval3.String()))
				case "in":
					l.InsertDatetime.AndWhereIn(fval3.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop3)
				}
			}
		}
	}
	ftyp4, exists := typ.FieldByName("UpdateDatetime")
	if exists {
		fval4 := val.FieldByName("UpdateDatetime")
		fop4, ok := ftyp4.Tag.Lookup("op")
		for fval4.Kind() == reflect.Ptr && !fval4.IsNil() {
			fval4 = fval4.Elem()
		}
		if fval4.Kind() != reflect.Ptr {
			if !ok {
				l.UpdateDatetime.AndWhereEq(fval4.Interface())
			} else {
				switch fop4 {
				case "=":
					l.UpdateDatetime.AndWhereEq(fval4.Interface())
				case "!=":
					l.UpdateDatetime.AndWhereNeq(fval4.Interface())
				case ">":
					l.UpdateDatetime.AndWhereGt(fval4.Interface())
				case ">=":
					l.UpdateDatetime.AndWhereGte(fval4.Interface())
				case "<":
					l.UpdateDatetime.AndWhereLt(fval4.Interface())
				case "<=":
					l.UpdateDatetime.AndWhereLte(fval4.Interface())
				case "llike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s", fval4.String()))
				case "rlike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%s%%", fval4.String()))
				case "alike":
					l.UpdateDatetime.AndWhereLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "nllike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s", fval4.String()))
				case "nrlike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%s%%", fval4.String()))
				case "nalike":
					l.UpdateDatetime.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval4.String()))
				case "in":
					l.UpdateDatetime.AndWhereIn(fval4.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop4)
				}
			}
		}
	}
	ftyp5, exists := typ.FieldByName("Status")
	if exists {
		fval5 := val.FieldByName("Status")
		fop5, ok := ftyp5.Tag.Lookup("op")
		for fval5.Kind() == reflect.Ptr && !fval5.IsNil() {
			fval5 = fval5.Elem()
		}
		if fval5.Kind() != reflect.Ptr {
			if !ok {
				l.Status.AndWhereEq(fval5.Interface())
			} else {
				switch fop5 {
				case "=":
					l.Status.AndWhereEq(fval5.Interface())
				case "!=":
					l.Status.AndWhereNeq(fval5.Interface())
				case ">":
					l.Status.AndWhereGt(fval5.Interface())
				case ">=":
					l.Status.AndWhereGte(fval5.Interface())
				case "<":
					l.Status.AndWhereLt(fval5.Interface())
				case "<=":
					l.Status.AndWhereLte(fval5.Interface())
				case "llike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s", fval5.String()))
				case "rlike":
					l.Status.AndWhereLike(fmt.Sprintf("%s%%", fval5.String()))
				case "alike":
					l.Status.AndWhereLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "nllike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s", fval5.String()))
				case "nrlike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%s%%", fval5.String()))
				case "nalike":
					l.Status.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval5.String()))
				case "in":
					l.Status.AndWhereIn(fval5.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop5)
				}
			}
		}
	}
	ftyp6, exists := typ.FieldByName("Remark1")
	if exists {
		fval6 := val.FieldByName("Remark1")
		fop6, ok := ftyp6.Tag.Lookup("op")
		for fval6.Kind() == reflect.Ptr && !fval6.IsNil() {
			fval6 = fval6.Elem()
		}
		if fval6.Kind() != reflect.Ptr {
			if !ok {
				l.Remark1.AndWhereEq(fval6.Interface())
			} else {
				switch fop6 {
				case "=":
					l.Remark1.AndWhereEq(fval6.Interface())
				case "!=":
					l.Remark1.AndWhereNeq(fval6.Interface())
				case ">":
					l.Remark1.AndWhereGt(fval6.Interface())
				case ">=":
					l.Remark1.AndWhereGte(fval6.Interface())
				case "<":
					l.Remark1.AndWhereLt(fval6.Interface())
				case "<=":
					l.Remark1.AndWhereLte(fval6.Interface())
				case "llike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s", fval6.String()))
				case "rlike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%s%%", fval6.String()))
				case "alike":
					l.Remark1.AndWhereLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "nllike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s", fval6.String()))
				case "nrlike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%s%%", fval6.String()))
				case "nalike":
					l.Remark1.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval6.String()))
				case "in":
					l.Remark1.AndWhereIn(fval6.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop6)
				}
			}
		}
	}
	ftyp7, exists := typ.FieldByName("Remark2")
	if exists {
		fval7 := val.FieldByName("Remark2")
		fop7, ok := ftyp7.Tag.Lookup("op")
		for fval7.Kind() == reflect.Ptr && !fval7.IsNil() {
			fval7 = fval7.Elem()
		}
		if fval7.Kind() != reflect.Ptr {
			if !ok {
				l.Remark2.AndWhereEq(fval7.Interface())
			} else {
				switch fop7 {
				case "=":
					l.Remark2.AndWhereEq(fval7.Interface())
				case "!=":
					l.Remark2.AndWhereNeq(fval7.Interface())
				case ">":
					l.Remark2.AndWhereGt(fval7.Interface())
				case ">=":
					l.Remark2.AndWhereGte(fval7.Interface())
				case "<":
					l.Remark2.AndWhereLt(fval7.Interface())
				case "<=":
					l.Remark2.AndWhereLte(fval7.Interface())
				case "llike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s", fval7.String()))
				case "rlike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%s%%", fval7.String()))
				case "alike":
					l.Remark2.AndWhereLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "nllike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s", fval7.String()))
				case "nrlike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%s%%", fval7.String()))
				case "nalike":
					l.Remark2.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval7.String()))
				case "in":
					l.Remark2.AndWhereIn(fval7.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop7)
				}
			}
		}
	}
	ftyp8, exists := typ.FieldByName("Remark3")
	if exists {
		fval8 := val.FieldByName("Remark3")
		fop8, ok := ftyp8.Tag.Lookup("op")
		for fval8.Kind() == reflect.Ptr && !fval8.IsNil() {
			fval8 = fval8.Elem()
		}
		if fval8.Kind() != reflect.Ptr {
			if !ok {
				l.Remark3.AndWhereEq(fval8.Interface())
			} else {
				switch fop8 {
				case "=":
					l.Remark3.AndWhereEq(fval8.Interface())
				case "!=":
					l.Remark3.AndWhereNeq(fval8.Interface())
				case ">":
					l.Remark3.AndWhereGt(fval8.Interface())
				case ">=":
					l.Remark3.AndWhereGte(fval8.Interface())
				case "<":
					l.Remark3.AndWhereLt(fval8.Interface())
				case "<=":
					l.Remark3.AndWhereLte(fval8.Interface())
				case "llike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s", fval8.String()))
				case "rlike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%s%%", fval8.String()))
				case "alike":
					l.Remark3.AndWhereLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "nllike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s", fval8.String()))
				case "nrlike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%s%%", fval8.String()))
				case "nalike":
					l.Remark3.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval8.String()))
				case "in":
					l.Remark3.AndWhereIn(fval8.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop8)
				}
			}
		}
	}
	ftyp9, exists := typ.FieldByName("Remark4")
	if exists {
		fval9 := val.FieldByName("Remark4")
		fop9, ok := ftyp9.Tag.Lookup("op")
		for fval9.Kind() == reflect.Ptr && !fval9.IsNil() {
			fval9 = fval9.Elem()
		}
		if fval9.Kind() != reflect.Ptr {
			if !ok {
				l.Remark4.AndWhereEq(fval9.Interface())
			} else {
				switch fop9 {
				case "=":
					l.Remark4.AndWhereEq(fval9.Interface())
				case "!=":
					l.Remark4.AndWhereNeq(fval9.Interface())
				case ">":
					l.Remark4.AndWhereGt(fval9.Interface())
				case ">=":
					l.Remark4.AndWhereGte(fval9.Interface())
				case "<":
					l.Remark4.AndWhereLt(fval9.Interface())
				case "<=":
					l.Remark4.AndWhereLte(fval9.Interface())
				case "llike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s", fval9.String()))
				case "rlike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%s%%", fval9.String()))
				case "alike":
					l.Remark4.AndWhereLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "nllike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s", fval9.String()))
				case "nrlike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%s%%", fval9.String()))
				case "nalike":
					l.Remark4.AndWhereNotLike(fmt.Sprintf("%%%s%%", fval9.String()))
				case "in":
					l.Remark4.AndWhereIn(fval9.Interface())
				default:
					return nil, fmt.Errorf("unknown op tag: %s", fop9)
				}
			}
		}
	}
	return l, nil
}
func init() {
	go UserCache.run()
	go UserListCache.run()
	go StudentbasicinfoCache.run()
	go StudentbasicinfoListCache.run()
	go ClassCache.run()
	go ClassListCache.run()
	go GradeCache.run()
	go GradeListCache.run()
}
