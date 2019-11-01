package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wangjun861205/nborm"
	"strings"
	"time"
)

func NewUser() *User {
	m := &User{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 1)
	m.UserCode.Init(m, "UserCode", "UserCode", 2)
	m.Name.Init(m, "Name", "Name", 3)
	m.Sex.Init(m, "Sex", "Sex", 4)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", 5)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", 6)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", 7)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", 8)
	m.UserType.Init(m, "UserType", "UserType", 9)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", 10)
	m.Type.Init(m, "Type", "Type", 11)
	m.Password.Init(m, "Password", "Password", 12)
	m.Phone.Init(m, "Phone", "Phone", 13)
	m.Email.Init(m, "Email", "Email", 14)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", 15)
	m.Question.Init(m, "Question", "Question", 16)
	m.Answer.Init(m, "Answer", "Answer", 17)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", 18)
	m.Operator.Init(m, "Operator", "Operator", 19)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 20)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 21)
	m.Status.Init(m, "Status", "Status", 22)
	m.Remark1.Init(m, "Remark1", "Remark1", 23)
	m.Remark2.Init(m, "Remark2", "Remark2", 24)
	m.Remark3.Init(m, "Remark3", "Remark3", 25)
	m.Remark4.Init(m, "Remark4", "Remark4", 26)
	m.Nonego.Init(m, "Nonego", "Nonego", 27)
	m.InitRel()
	return m
}

func newSubUser(parent nborm.Model) *User {
	m := &User{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 1)
	m.UserCode.Init(m, "UserCode", "UserCode", 2)
	m.Name.Init(m, "Name", "Name", 3)
	m.Sex.Init(m, "Sex", "Sex", 4)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", 5)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", 6)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", 7)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", 8)
	m.UserType.Init(m, "UserType", "UserType", 9)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", 10)
	m.Type.Init(m, "Type", "Type", 11)
	m.Password.Init(m, "Password", "Password", 12)
	m.Phone.Init(m, "Phone", "Phone", 13)
	m.Email.Init(m, "Email", "Email", 14)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", 15)
	m.Question.Init(m, "Question", "Question", 16)
	m.Answer.Init(m, "Answer", "Answer", 17)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", 18)
	m.Operator.Init(m, "Operator", "Operator", 19)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 20)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 21)
	m.Status.Init(m, "Status", "Status", 22)
	m.Remark1.Init(m, "Remark1", "Remark1", 23)
	m.Remark2.Init(m, "Remark2", "Remark2", 24)
	m.Remark3.Init(m, "Remark3", "Remark3", 25)
	m.Remark4.Init(m, "Remark4", "Remark4", 26)
	m.Nonego.Init(m, "Nonego", "Nonego", 27)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", 1)
	l.UserCode.Init(l, "UserCode", "UserCode", 2)
	l.Name.Init(l, "Name", "Name", 3)
	l.Sex.Init(l, "Sex", "Sex", 4)
	l.IdentityType.Init(l, "IdentityType", "IdentityType", 5)
	l.IdentityNum.Init(l, "IdentityNum", "IdentityNum", 6)
	l.ExpirationDate.Init(l, "ExpirationDate", "ExpirationDate", 7)
	l.UniversityCode.Init(l, "UniversityCode", "UniversityCode", 8)
	l.UserType.Init(l, "UserType", "UserType", 9)
	l.EnrollmentStatus.Init(l, "EnrollmentStatus", "EnrollmentStatus", 10)
	l.Type.Init(l, "Type", "Type", 11)
	l.Password.Init(l, "Password", "Password", 12)
	l.Phone.Init(l, "Phone", "Phone", 13)
	l.Email.Init(l, "Email", "Email", 14)
	l.PictureURL.Init(l, "PictureURL", "PictureURL", 15)
	l.Question.Init(l, "Question", "Question", 16)
	l.Answer.Init(l, "Answer", "Answer", 17)
	l.AvailableLogin.Init(l, "AvailableLogin", "AvailableLogin", 18)
	l.Operator.Init(l, "Operator", "Operator", 19)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 20)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 21)
	l.Status.Init(l, "Status", "Status", 22)
	l.Remark1.Init(l, "Remark1", "Remark1", 23)
	l.Remark2.Init(l, "Remark2", "Remark2", 24)
	l.Remark3.Init(l, "Remark3", "Remark3", 25)
	l.Remark4.Init(l, "Remark4", "Remark4", 26)
	l.Nonego.Init(l, "Nonego", "Nonego", 27)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", 1)
	l.UserCode.Init(l, "UserCode", "UserCode", 2)
	l.Name.Init(l, "Name", "Name", 3)
	l.Sex.Init(l, "Sex", "Sex", 4)
	l.IdentityType.Init(l, "IdentityType", "IdentityType", 5)
	l.IdentityNum.Init(l, "IdentityNum", "IdentityNum", 6)
	l.ExpirationDate.Init(l, "ExpirationDate", "ExpirationDate", 7)
	l.UniversityCode.Init(l, "UniversityCode", "UniversityCode", 8)
	l.UserType.Init(l, "UserType", "UserType", 9)
	l.EnrollmentStatus.Init(l, "EnrollmentStatus", "EnrollmentStatus", 10)
	l.Type.Init(l, "Type", "Type", 11)
	l.Password.Init(l, "Password", "Password", 12)
	l.Phone.Init(l, "Phone", "Phone", 13)
	l.Email.Init(l, "Email", "Email", 14)
	l.PictureURL.Init(l, "PictureURL", "PictureURL", 15)
	l.Question.Init(l, "Question", "Question", 16)
	l.Answer.Init(l, "Answer", "Answer", 17)
	l.AvailableLogin.Init(l, "AvailableLogin", "AvailableLogin", 18)
	l.Operator.Init(l, "Operator", "Operator", 19)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 20)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 21)
	l.Status.Init(l, "Status", "Status", 22)
	l.Remark1.Init(l, "Remark1", "Remark1", 23)
	l.Remark2.Init(l, "Remark2", "Remark2", 24)
	l.Remark3.Init(l, "Remark3", "Remark3", 25)
	l.Remark4.Init(l, "Remark4", "Remark4", 26)
	l.Nonego.Init(l, "Nonego", "Nonego", 27)
	return l
}

func (l *UserList) NewModel() nborm.Model {
	m := &User{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 1)
	l.IntelUserCode.CopyStatus(&m.IntelUserCode)
	m.UserCode.Init(m, "UserCode", "UserCode", 2)
	l.UserCode.CopyStatus(&m.UserCode)
	m.Name.Init(m, "Name", "Name", 3)
	l.Name.CopyStatus(&m.Name)
	m.Sex.Init(m, "Sex", "Sex", 4)
	l.Sex.CopyStatus(&m.Sex)
	m.IdentityType.Init(m, "IdentityType", "IdentityType", 5)
	l.IdentityType.CopyStatus(&m.IdentityType)
	m.IdentityNum.Init(m, "IdentityNum", "IdentityNum", 6)
	l.IdentityNum.CopyStatus(&m.IdentityNum)
	m.ExpirationDate.Init(m, "ExpirationDate", "ExpirationDate", 7)
	l.ExpirationDate.CopyStatus(&m.ExpirationDate)
	m.UniversityCode.Init(m, "UniversityCode", "UniversityCode", 8)
	l.UniversityCode.CopyStatus(&m.UniversityCode)
	m.UserType.Init(m, "UserType", "UserType", 9)
	l.UserType.CopyStatus(&m.UserType)
	m.EnrollmentStatus.Init(m, "EnrollmentStatus", "EnrollmentStatus", 10)
	l.EnrollmentStatus.CopyStatus(&m.EnrollmentStatus)
	m.Type.Init(m, "Type", "Type", 11)
	l.Type.CopyStatus(&m.Type)
	m.Password.Init(m, "Password", "Password", 12)
	l.Password.CopyStatus(&m.Password)
	m.Phone.Init(m, "Phone", "Phone", 13)
	l.Phone.CopyStatus(&m.Phone)
	m.Email.Init(m, "Email", "Email", 14)
	l.Email.CopyStatus(&m.Email)
	m.PictureURL.Init(m, "PictureURL", "PictureURL", 15)
	l.PictureURL.CopyStatus(&m.PictureURL)
	m.Question.Init(m, "Question", "Question", 16)
	l.Question.CopyStatus(&m.Question)
	m.Answer.Init(m, "Answer", "Answer", 17)
	l.Answer.CopyStatus(&m.Answer)
	m.AvailableLogin.Init(m, "AvailableLogin", "AvailableLogin", 18)
	l.AvailableLogin.CopyStatus(&m.AvailableLogin)
	m.Operator.Init(m, "Operator", "Operator", 19)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 20)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 21)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", 22)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", 23)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", 24)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", 25)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", 26)
	l.Remark4.CopyStatus(&m.Remark4)
	m.Nonego.Init(m, "Nonego", "Nonego", 27)
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
	if string(b) == "[]" || string(b) == "null" {
		return nil
	}
	jl := struct {
		List  *[]*User
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
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

func NewStudentbasicinfo() *Studentbasicinfo {
	m := &Studentbasicinfo{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 2)
	m.Class.Init(m, "Class", "Class", 3)
	m.OtherName.Init(m, "OtherName", "OtherName", 4)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", 5)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", 6)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", 7)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", 8)
	m.Birthday.Init(m, "Birthday", "Birthday", 9)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", 10)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", 11)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", 12)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", 13)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", 14)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", 16)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", 17)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", 18)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", 19)
	m.Creed.Init(m, "Creed", "Creed", 20)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", 21)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", 22)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", 23)
	m.PostCode.Init(m, "PostCode", "PostCode", 24)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", 25)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", 26)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 27)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", 28)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", 29)
	m.Referee.Init(m, "Referee", "Referee", 30)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", 31)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", 32)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", 33)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", 35)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", 36)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", 37)
	m.BizType.Init(m, "BizType", "BizType", 38)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", 39)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", 40)
	m.Operator.Init(m, "Operator", "Operator", 41)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 42)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 43)
	m.Status.Init(m, "Status", "Status", 44)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", 45)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", 46)
	m.Campus.Init(m, "Campus", "Campus", 47)
	m.Zone.Init(m, "Zone", "Zone", 48)
	m.Building.Init(m, "Building", "Building", 49)
	m.Unit.Init(m, "Unit", "Unit", 50)
	m.Room.Init(m, "Room", "Room", 51)
	m.Bed.Init(m, "Bed", "Bed", 52)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", 53)
	m.Height.Init(m, "Height", "Height", 54)
	m.Weight.Init(m, "Weight", "Weight", 55)
	m.FootSize.Init(m, "FootSize", "FootSize", 56)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", 57)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", 58)
	m.Remark1.Init(m, "Remark1", "Remark1", 59)
	m.Remark2.Init(m, "Remark2", "Remark2", 60)
	m.Remark3.Init(m, "Remark3", "Remark3", 61)
	m.Remark4.Init(m, "Remark4", "Remark4", 62)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", 63)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", 64)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", 65)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", 66)
	m.InitRel()
	return m
}

func newSubStudentbasicinfo(parent nborm.Model) *Studentbasicinfo {
	m := &Studentbasicinfo{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 2)
	m.Class.Init(m, "Class", "Class", 3)
	m.OtherName.Init(m, "OtherName", "OtherName", 4)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", 5)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", 6)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", 7)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", 8)
	m.Birthday.Init(m, "Birthday", "Birthday", 9)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", 10)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", 11)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", 12)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", 13)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", 14)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", 16)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", 17)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", 18)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", 19)
	m.Creed.Init(m, "Creed", "Creed", 20)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", 21)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", 22)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", 23)
	m.PostCode.Init(m, "PostCode", "PostCode", 24)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", 25)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", 26)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 27)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", 28)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", 29)
	m.Referee.Init(m, "Referee", "Referee", 30)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", 31)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", 32)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", 33)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", 35)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", 36)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", 37)
	m.BizType.Init(m, "BizType", "BizType", 38)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", 39)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", 40)
	m.Operator.Init(m, "Operator", "Operator", 41)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 42)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 43)
	m.Status.Init(m, "Status", "Status", 44)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", 45)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", 46)
	m.Campus.Init(m, "Campus", "Campus", 47)
	m.Zone.Init(m, "Zone", "Zone", 48)
	m.Building.Init(m, "Building", "Building", 49)
	m.Unit.Init(m, "Unit", "Unit", 50)
	m.Room.Init(m, "Room", "Room", 51)
	m.Bed.Init(m, "Bed", "Bed", 52)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", 53)
	m.Height.Init(m, "Height", "Height", 54)
	m.Weight.Init(m, "Weight", "Weight", 55)
	m.FootSize.Init(m, "FootSize", "FootSize", 56)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", 57)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", 58)
	m.Remark1.Init(m, "Remark1", "Remark1", 59)
	m.Remark2.Init(m, "Remark2", "Remark2", 60)
	m.Remark3.Init(m, "Remark3", "Remark3", 61)
	m.Remark4.Init(m, "Remark4", "Remark4", 62)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", 63)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", 64)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", 65)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", 66)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", 1)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", 2)
	l.Class.Init(l, "Class", "Class", 3)
	l.OtherName.Init(l, "OtherName", "OtherName", 4)
	l.NameInPinyin.Init(l, "NameInPinyin", "NameInPinyin", 5)
	l.EnglishName.Init(l, "EnglishName", "EnglishName", 6)
	l.CountryCode.Init(l, "CountryCode", "CountryCode", 7)
	l.NationalityCode.Init(l, "NationalityCode", "NationalityCode", 8)
	l.Birthday.Init(l, "Birthday", "Birthday", 9)
	l.PoliticalCode.Init(l, "PoliticalCode", "PoliticalCode", 10)
	l.QQAcct.Init(l, "QQAcct", "QQAcct", 11)
	l.WeChatAcct.Init(l, "WeChatAcct", "WeChatAcct", 12)
	l.BankCardNumber.Init(l, "BankCardNumber", "BankCardNumber", 13)
	l.AccountBankCode.Init(l, "AccountBankCode", "AccountBankCode", 14)
	l.AllPowerfulCardNum.Init(l, "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.MaritalCode.Init(l, "MaritalCode", "MaritalCode", 16)
	l.OriginAreaCode.Init(l, "OriginAreaCode", "OriginAreaCode", 17)
	l.StudentAreaCode.Init(l, "StudentAreaCode", "StudentAreaCode", 18)
	l.Hobbies.Init(l, "Hobbies", "Hobbies", 19)
	l.Creed.Init(l, "Creed", "Creed", 20)
	l.TrainTicketinterval.Init(l, "TrainTicketinterval", "TrainTicketinterval", 21)
	l.FamilyAddress.Init(l, "FamilyAddress", "FamilyAddress", 22)
	l.DetailAddress.Init(l, "DetailAddress", "DetailAddress", 23)
	l.PostCode.Init(l, "PostCode", "PostCode", 24)
	l.HomePhone.Init(l, "HomePhone", "HomePhone", 25)
	l.EnrollmentDate.Init(l, "EnrollmentDate", "EnrollmentDate", 26)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", 27)
	l.MidSchoolAddress.Init(l, "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolName.Init(l, "MidSchoolName", "MidSchoolName", 29)
	l.Referee.Init(l, "Referee", "Referee", 30)
	l.RefereeDuty.Init(l, "RefereeDuty", "RefereeDuty", 31)
	l.RefereePhone.Init(l, "RefereePhone", "RefereePhone", 32)
	l.AdmissionTicketNo.Init(l, "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.CollegeEntranceExamScores.Init(l, "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.AdmissionYear.Init(l, "AdmissionYear", "AdmissionYear", 35)
	l.ForeignLanguageCode.Init(l, "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.StudentOrigin.Init(l, "StudentOrigin", "StudentOrigin", 37)
	l.BizType.Init(l, "BizType", "BizType", 38)
	l.TaskCode.Init(l, "TaskCode", "TaskCode", 39)
	l.ApproveStatus.Init(l, "ApproveStatus", "ApproveStatus", 40)
	l.Operator.Init(l, "Operator", "Operator", 41)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 42)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 43)
	l.Status.Init(l, "Status", "Status", 44)
	l.StudentStatus.Init(l, "StudentStatus", "StudentStatus", 45)
	l.IsAuth.Init(l, "IsAuth", "IsAuth", 46)
	l.Campus.Init(l, "Campus", "Campus", 47)
	l.Zone.Init(l, "Zone", "Zone", 48)
	l.Building.Init(l, "Building", "Building", 49)
	l.Unit.Init(l, "Unit", "Unit", 50)
	l.Room.Init(l, "Room", "Room", 51)
	l.Bed.Init(l, "Bed", "Bed", 52)
	l.StatusSort.Init(l, "StatusSort", "StatusSort", 53)
	l.Height.Init(l, "Height", "Height", 54)
	l.Weight.Init(l, "Weight", "Weight", 55)
	l.FootSize.Init(l, "FootSize", "FootSize", 56)
	l.ClothSize.Init(l, "ClothSize", "ClothSize", 57)
	l.HeadSize.Init(l, "HeadSize", "HeadSize", 58)
	l.Remark1.Init(l, "Remark1", "Remark1", 59)
	l.Remark2.Init(l, "Remark2", "Remark2", 60)
	l.Remark3.Init(l, "Remark3", "Remark3", 61)
	l.Remark4.Init(l, "Remark4", "Remark4", 62)
	l.IsPayment.Init(l, "IsPayment", "IsPayment", 63)
	l.IsCheckIn.Init(l, "isCheckIn", "IsCheckIn", 64)
	l.GetMilitaryTC.Init(l, "GetMilitaryTC", "GetMilitaryTC", 65)
	l.OriginAreaName.Init(l, "OriginAreaName", "OriginAreaName", 66)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", 1)
	l.IntelUserCode.Init(l, "IntelUserCode", "IntelUserCode", 2)
	l.Class.Init(l, "Class", "Class", 3)
	l.OtherName.Init(l, "OtherName", "OtherName", 4)
	l.NameInPinyin.Init(l, "NameInPinyin", "NameInPinyin", 5)
	l.EnglishName.Init(l, "EnglishName", "EnglishName", 6)
	l.CountryCode.Init(l, "CountryCode", "CountryCode", 7)
	l.NationalityCode.Init(l, "NationalityCode", "NationalityCode", 8)
	l.Birthday.Init(l, "Birthday", "Birthday", 9)
	l.PoliticalCode.Init(l, "PoliticalCode", "PoliticalCode", 10)
	l.QQAcct.Init(l, "QQAcct", "QQAcct", 11)
	l.WeChatAcct.Init(l, "WeChatAcct", "WeChatAcct", 12)
	l.BankCardNumber.Init(l, "BankCardNumber", "BankCardNumber", 13)
	l.AccountBankCode.Init(l, "AccountBankCode", "AccountBankCode", 14)
	l.AllPowerfulCardNum.Init(l, "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.MaritalCode.Init(l, "MaritalCode", "MaritalCode", 16)
	l.OriginAreaCode.Init(l, "OriginAreaCode", "OriginAreaCode", 17)
	l.StudentAreaCode.Init(l, "StudentAreaCode", "StudentAreaCode", 18)
	l.Hobbies.Init(l, "Hobbies", "Hobbies", 19)
	l.Creed.Init(l, "Creed", "Creed", 20)
	l.TrainTicketinterval.Init(l, "TrainTicketinterval", "TrainTicketinterval", 21)
	l.FamilyAddress.Init(l, "FamilyAddress", "FamilyAddress", 22)
	l.DetailAddress.Init(l, "DetailAddress", "DetailAddress", 23)
	l.PostCode.Init(l, "PostCode", "PostCode", 24)
	l.HomePhone.Init(l, "HomePhone", "HomePhone", 25)
	l.EnrollmentDate.Init(l, "EnrollmentDate", "EnrollmentDate", 26)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", 27)
	l.MidSchoolAddress.Init(l, "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolName.Init(l, "MidSchoolName", "MidSchoolName", 29)
	l.Referee.Init(l, "Referee", "Referee", 30)
	l.RefereeDuty.Init(l, "RefereeDuty", "RefereeDuty", 31)
	l.RefereePhone.Init(l, "RefereePhone", "RefereePhone", 32)
	l.AdmissionTicketNo.Init(l, "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.CollegeEntranceExamScores.Init(l, "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.AdmissionYear.Init(l, "AdmissionYear", "AdmissionYear", 35)
	l.ForeignLanguageCode.Init(l, "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.StudentOrigin.Init(l, "StudentOrigin", "StudentOrigin", 37)
	l.BizType.Init(l, "BizType", "BizType", 38)
	l.TaskCode.Init(l, "TaskCode", "TaskCode", 39)
	l.ApproveStatus.Init(l, "ApproveStatus", "ApproveStatus", 40)
	l.Operator.Init(l, "Operator", "Operator", 41)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 42)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 43)
	l.Status.Init(l, "Status", "Status", 44)
	l.StudentStatus.Init(l, "StudentStatus", "StudentStatus", 45)
	l.IsAuth.Init(l, "IsAuth", "IsAuth", 46)
	l.Campus.Init(l, "Campus", "Campus", 47)
	l.Zone.Init(l, "Zone", "Zone", 48)
	l.Building.Init(l, "Building", "Building", 49)
	l.Unit.Init(l, "Unit", "Unit", 50)
	l.Room.Init(l, "Room", "Room", 51)
	l.Bed.Init(l, "Bed", "Bed", 52)
	l.StatusSort.Init(l, "StatusSort", "StatusSort", 53)
	l.Height.Init(l, "Height", "Height", 54)
	l.Weight.Init(l, "Weight", "Weight", 55)
	l.FootSize.Init(l, "FootSize", "FootSize", 56)
	l.ClothSize.Init(l, "ClothSize", "ClothSize", 57)
	l.HeadSize.Init(l, "HeadSize", "HeadSize", 58)
	l.Remark1.Init(l, "Remark1", "Remark1", 59)
	l.Remark2.Init(l, "Remark2", "Remark2", 60)
	l.Remark3.Init(l, "Remark3", "Remark3", 61)
	l.Remark4.Init(l, "Remark4", "Remark4", 62)
	l.IsPayment.Init(l, "IsPayment", "IsPayment", 63)
	l.IsCheckIn.Init(l, "isCheckIn", "IsCheckIn", 64)
	l.GetMilitaryTC.Init(l, "GetMilitaryTC", "GetMilitaryTC", 65)
	l.OriginAreaName.Init(l, "OriginAreaName", "OriginAreaName", 66)
	return l
}

func (l *StudentbasicinfoList) NewModel() nborm.Model {
	m := &Studentbasicinfo{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	l.RecordId.CopyStatus(&m.RecordId)
	m.IntelUserCode.Init(m, "IntelUserCode", "IntelUserCode", 2)
	l.IntelUserCode.CopyStatus(&m.IntelUserCode)
	m.Class.Init(m, "Class", "Class", 3)
	l.Class.CopyStatus(&m.Class)
	m.OtherName.Init(m, "OtherName", "OtherName", 4)
	l.OtherName.CopyStatus(&m.OtherName)
	m.NameInPinyin.Init(m, "NameInPinyin", "NameInPinyin", 5)
	l.NameInPinyin.CopyStatus(&m.NameInPinyin)
	m.EnglishName.Init(m, "EnglishName", "EnglishName", 6)
	l.EnglishName.CopyStatus(&m.EnglishName)
	m.CountryCode.Init(m, "CountryCode", "CountryCode", 7)
	l.CountryCode.CopyStatus(&m.CountryCode)
	m.NationalityCode.Init(m, "NationalityCode", "NationalityCode", 8)
	l.NationalityCode.CopyStatus(&m.NationalityCode)
	m.Birthday.Init(m, "Birthday", "Birthday", 9)
	l.Birthday.CopyStatus(&m.Birthday)
	m.PoliticalCode.Init(m, "PoliticalCode", "PoliticalCode", 10)
	l.PoliticalCode.CopyStatus(&m.PoliticalCode)
	m.QQAcct.Init(m, "QQAcct", "QQAcct", 11)
	l.QQAcct.CopyStatus(&m.QQAcct)
	m.WeChatAcct.Init(m, "WeChatAcct", "WeChatAcct", 12)
	l.WeChatAcct.CopyStatus(&m.WeChatAcct)
	m.BankCardNumber.Init(m, "BankCardNumber", "BankCardNumber", 13)
	l.BankCardNumber.CopyStatus(&m.BankCardNumber)
	m.AccountBankCode.Init(m, "AccountBankCode", "AccountBankCode", 14)
	l.AccountBankCode.CopyStatus(&m.AccountBankCode)
	m.AllPowerfulCardNum.Init(m, "AllPowerfulCardNum", "AllPowerfulCardNum", 15)
	l.AllPowerfulCardNum.CopyStatus(&m.AllPowerfulCardNum)
	m.MaritalCode.Init(m, "MaritalCode", "MaritalCode", 16)
	l.MaritalCode.CopyStatus(&m.MaritalCode)
	m.OriginAreaCode.Init(m, "OriginAreaCode", "OriginAreaCode", 17)
	l.OriginAreaCode.CopyStatus(&m.OriginAreaCode)
	m.StudentAreaCode.Init(m, "StudentAreaCode", "StudentAreaCode", 18)
	l.StudentAreaCode.CopyStatus(&m.StudentAreaCode)
	m.Hobbies.Init(m, "Hobbies", "Hobbies", 19)
	l.Hobbies.CopyStatus(&m.Hobbies)
	m.Creed.Init(m, "Creed", "Creed", 20)
	l.Creed.CopyStatus(&m.Creed)
	m.TrainTicketinterval.Init(m, "TrainTicketinterval", "TrainTicketinterval", 21)
	l.TrainTicketinterval.CopyStatus(&m.TrainTicketinterval)
	m.FamilyAddress.Init(m, "FamilyAddress", "FamilyAddress", 22)
	l.FamilyAddress.CopyStatus(&m.FamilyAddress)
	m.DetailAddress.Init(m, "DetailAddress", "DetailAddress", 23)
	l.DetailAddress.CopyStatus(&m.DetailAddress)
	m.PostCode.Init(m, "PostCode", "PostCode", 24)
	l.PostCode.CopyStatus(&m.PostCode)
	m.HomePhone.Init(m, "HomePhone", "HomePhone", 25)
	l.HomePhone.CopyStatus(&m.HomePhone)
	m.EnrollmentDate.Init(m, "EnrollmentDate", "EnrollmentDate", 26)
	l.EnrollmentDate.CopyStatus(&m.EnrollmentDate)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 27)
	l.GraduationDate.CopyStatus(&m.GraduationDate)
	m.MidSchoolAddress.Init(m, "MidSchoolAddress", "MidSchoolAddress", 28)
	l.MidSchoolAddress.CopyStatus(&m.MidSchoolAddress)
	m.MidSchoolName.Init(m, "MidSchoolName", "MidSchoolName", 29)
	l.MidSchoolName.CopyStatus(&m.MidSchoolName)
	m.Referee.Init(m, "Referee", "Referee", 30)
	l.Referee.CopyStatus(&m.Referee)
	m.RefereeDuty.Init(m, "RefereeDuty", "RefereeDuty", 31)
	l.RefereeDuty.CopyStatus(&m.RefereeDuty)
	m.RefereePhone.Init(m, "RefereePhone", "RefereePhone", 32)
	l.RefereePhone.CopyStatus(&m.RefereePhone)
	m.AdmissionTicketNo.Init(m, "AdmissionTicketNo", "AdmissionTicketNo", 33)
	l.AdmissionTicketNo.CopyStatus(&m.AdmissionTicketNo)
	m.CollegeEntranceExamScores.Init(m, "CollegeEntranceExamScores", "CollegeEntranceExamScores", 34)
	l.CollegeEntranceExamScores.CopyStatus(&m.CollegeEntranceExamScores)
	m.AdmissionYear.Init(m, "AdmissionYear", "AdmissionYear", 35)
	l.AdmissionYear.CopyStatus(&m.AdmissionYear)
	m.ForeignLanguageCode.Init(m, "ForeignLanguageCode", "ForeignLanguageCode", 36)
	l.ForeignLanguageCode.CopyStatus(&m.ForeignLanguageCode)
	m.StudentOrigin.Init(m, "StudentOrigin", "StudentOrigin", 37)
	l.StudentOrigin.CopyStatus(&m.StudentOrigin)
	m.BizType.Init(m, "BizType", "BizType", 38)
	l.BizType.CopyStatus(&m.BizType)
	m.TaskCode.Init(m, "TaskCode", "TaskCode", 39)
	l.TaskCode.CopyStatus(&m.TaskCode)
	m.ApproveStatus.Init(m, "ApproveStatus", "ApproveStatus", 40)
	l.ApproveStatus.CopyStatus(&m.ApproveStatus)
	m.Operator.Init(m, "Operator", "Operator", 41)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 42)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 43)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", 44)
	l.Status.CopyStatus(&m.Status)
	m.StudentStatus.Init(m, "StudentStatus", "StudentStatus", 45)
	l.StudentStatus.CopyStatus(&m.StudentStatus)
	m.IsAuth.Init(m, "IsAuth", "IsAuth", 46)
	l.IsAuth.CopyStatus(&m.IsAuth)
	m.Campus.Init(m, "Campus", "Campus", 47)
	l.Campus.CopyStatus(&m.Campus)
	m.Zone.Init(m, "Zone", "Zone", 48)
	l.Zone.CopyStatus(&m.Zone)
	m.Building.Init(m, "Building", "Building", 49)
	l.Building.CopyStatus(&m.Building)
	m.Unit.Init(m, "Unit", "Unit", 50)
	l.Unit.CopyStatus(&m.Unit)
	m.Room.Init(m, "Room", "Room", 51)
	l.Room.CopyStatus(&m.Room)
	m.Bed.Init(m, "Bed", "Bed", 52)
	l.Bed.CopyStatus(&m.Bed)
	m.StatusSort.Init(m, "StatusSort", "StatusSort", 53)
	l.StatusSort.CopyStatus(&m.StatusSort)
	m.Height.Init(m, "Height", "Height", 54)
	l.Height.CopyStatus(&m.Height)
	m.Weight.Init(m, "Weight", "Weight", 55)
	l.Weight.CopyStatus(&m.Weight)
	m.FootSize.Init(m, "FootSize", "FootSize", 56)
	l.FootSize.CopyStatus(&m.FootSize)
	m.ClothSize.Init(m, "ClothSize", "ClothSize", 57)
	l.ClothSize.CopyStatus(&m.ClothSize)
	m.HeadSize.Init(m, "HeadSize", "HeadSize", 58)
	l.HeadSize.CopyStatus(&m.HeadSize)
	m.Remark1.Init(m, "Remark1", "Remark1", 59)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", 60)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", 61)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", 62)
	l.Remark4.CopyStatus(&m.Remark4)
	m.IsPayment.Init(m, "IsPayment", "IsPayment", 63)
	l.IsPayment.CopyStatus(&m.IsPayment)
	m.IsCheckIn.Init(m, "isCheckIn", "IsCheckIn", 64)
	l.IsCheckIn.CopyStatus(&m.IsCheckIn)
	m.GetMilitaryTC.Init(m, "GetMilitaryTC", "GetMilitaryTC", 65)
	l.GetMilitaryTC.CopyStatus(&m.GetMilitaryTC)
	m.OriginAreaName.Init(m, "OriginAreaName", "OriginAreaName", 66)
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
	if string(b) == "[]" || string(b) == "null" {
		return nil
	}
	jl := struct {
		List  *[]*Studentbasicinfo
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
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

func NewClass() *Class {
	m := &Class{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", 2)
	m.ClassName.Init(m, "ClassName", "ClassName", 3)
	m.Campus.Init(m, "Campus", "Campus", 4)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", 5)
	m.Grade.Init(m, "Grade", "Grade", 6)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", 7)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", 8)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 9)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", 10)
	m.StudentType.Init(m, "StudentType", "StudentType", 11)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", 12)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", 13)
	m.Counselor.Init(m, "Counselor", "Counselor", 14)
	m.Adviser.Init(m, "Adviser", "Adviser", 15)
	m.Leadership.Init(m, "Leadership", "Leadership", 16)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", 17)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", 18)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", 19)
	m.Operator.Init(m, "Operator", "Operator", 20)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 21)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 22)
	m.Status.Init(m, "Status", "Status", 23)
	m.Remark1.Init(m, "Remark1", "Remark1", 24)
	m.Remark2.Init(m, "Remark2", "Remark2", 25)
	m.Remark3.Init(m, "Remark3", "Remark3", 26)
	m.Remark4.Init(m, "Remark4", "Remark4", 27)
	m.InitRel()
	return m
}

func newSubClass(parent nborm.Model) *Class {
	m := &Class{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", 2)
	m.ClassName.Init(m, "ClassName", "ClassName", 3)
	m.Campus.Init(m, "Campus", "Campus", 4)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", 5)
	m.Grade.Init(m, "Grade", "Grade", 6)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", 7)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", 8)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 9)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", 10)
	m.StudentType.Init(m, "StudentType", "StudentType", 11)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", 12)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", 13)
	m.Counselor.Init(m, "Counselor", "Counselor", 14)
	m.Adviser.Init(m, "Adviser", "Adviser", 15)
	m.Leadership.Init(m, "Leadership", "Leadership", 16)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", 17)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", 18)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", 19)
	m.Operator.Init(m, "Operator", "Operator", 20)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 21)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 22)
	m.Status.Init(m, "Status", "Status", 23)
	m.Remark1.Init(m, "Remark1", "Remark1", 24)
	m.Remark2.Init(m, "Remark2", "Remark2", 25)
	m.Remark3.Init(m, "Remark3", "Remark3", 26)
	m.Remark4.Init(m, "Remark4", "Remark4", 27)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", 1)
	l.ClassCode.Init(l, "ClassCode", "ClassCode", 2)
	l.ClassName.Init(l, "ClassName", "ClassName", 3)
	l.Campus.Init(l, "Campus", "Campus", 4)
	l.ResearchArea.Init(l, "ResearchArea", "ResearchArea", 5)
	l.Grade.Init(l, "Grade", "Grade", 6)
	l.TrainingMode.Init(l, "TrainingMode", "TrainingMode", 7)
	l.EntranceDate.Init(l, "EntranceDate", "EntranceDate", 8)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", 9)
	l.ProgramLength.Init(l, "ProgramLength", "ProgramLength", 10)
	l.StudentType.Init(l, "StudentType", "StudentType", 11)
	l.CredentialsType.Init(l, "CredentialsType", "CredentialsType", 12)
	l.DegreeType.Init(l, "DegreeType", "DegreeType", 13)
	l.Counselor.Init(l, "Counselor", "Counselor", 14)
	l.Adviser.Init(l, "Adviser", "Adviser", 15)
	l.Leadership.Init(l, "Leadership", "Leadership", 16)
	l.Supervisor.Init(l, "Supervisor", "Supervisor", 17)
	l.Assistant1.Init(l, "Assistant1", "Assistant1", 18)
	l.Assistant2.Init(l, "Assistant2", "Assistant2", 19)
	l.Operator.Init(l, "Operator", "Operator", 20)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 21)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 22)
	l.Status.Init(l, "Status", "Status", 23)
	l.Remark1.Init(l, "Remark1", "Remark1", 24)
	l.Remark2.Init(l, "Remark2", "Remark2", 25)
	l.Remark3.Init(l, "Remark3", "Remark3", 26)
	l.Remark4.Init(l, "Remark4", "Remark4", 27)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.RecordId.Init(l, "RecordId", "RecordId", 1)
	l.ClassCode.Init(l, "ClassCode", "ClassCode", 2)
	l.ClassName.Init(l, "ClassName", "ClassName", 3)
	l.Campus.Init(l, "Campus", "Campus", 4)
	l.ResearchArea.Init(l, "ResearchArea", "ResearchArea", 5)
	l.Grade.Init(l, "Grade", "Grade", 6)
	l.TrainingMode.Init(l, "TrainingMode", "TrainingMode", 7)
	l.EntranceDate.Init(l, "EntranceDate", "EntranceDate", 8)
	l.GraduationDate.Init(l, "GraduationDate", "GraduationDate", 9)
	l.ProgramLength.Init(l, "ProgramLength", "ProgramLength", 10)
	l.StudentType.Init(l, "StudentType", "StudentType", 11)
	l.CredentialsType.Init(l, "CredentialsType", "CredentialsType", 12)
	l.DegreeType.Init(l, "DegreeType", "DegreeType", 13)
	l.Counselor.Init(l, "Counselor", "Counselor", 14)
	l.Adviser.Init(l, "Adviser", "Adviser", 15)
	l.Leadership.Init(l, "Leadership", "Leadership", 16)
	l.Supervisor.Init(l, "Supervisor", "Supervisor", 17)
	l.Assistant1.Init(l, "Assistant1", "Assistant1", 18)
	l.Assistant2.Init(l, "Assistant2", "Assistant2", 19)
	l.Operator.Init(l, "Operator", "Operator", 20)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 21)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 22)
	l.Status.Init(l, "Status", "Status", 23)
	l.Remark1.Init(l, "Remark1", "Remark1", 24)
	l.Remark2.Init(l, "Remark2", "Remark2", 25)
	l.Remark3.Init(l, "Remark3", "Remark3", 26)
	l.Remark4.Init(l, "Remark4", "Remark4", 27)
	return l
}

func (l *ClassList) NewModel() nborm.Model {
	m := &Class{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.RecordId.Init(m, "RecordId", "RecordId", 1)
	l.RecordId.CopyStatus(&m.RecordId)
	m.ClassCode.Init(m, "ClassCode", "ClassCode", 2)
	l.ClassCode.CopyStatus(&m.ClassCode)
	m.ClassName.Init(m, "ClassName", "ClassName", 3)
	l.ClassName.CopyStatus(&m.ClassName)
	m.Campus.Init(m, "Campus", "Campus", 4)
	l.Campus.CopyStatus(&m.Campus)
	m.ResearchArea.Init(m, "ResearchArea", "ResearchArea", 5)
	l.ResearchArea.CopyStatus(&m.ResearchArea)
	m.Grade.Init(m, "Grade", "Grade", 6)
	l.Grade.CopyStatus(&m.Grade)
	m.TrainingMode.Init(m, "TrainingMode", "TrainingMode", 7)
	l.TrainingMode.CopyStatus(&m.TrainingMode)
	m.EntranceDate.Init(m, "EntranceDate", "EntranceDate", 8)
	l.EntranceDate.CopyStatus(&m.EntranceDate)
	m.GraduationDate.Init(m, "GraduationDate", "GraduationDate", 9)
	l.GraduationDate.CopyStatus(&m.GraduationDate)
	m.ProgramLength.Init(m, "ProgramLength", "ProgramLength", 10)
	l.ProgramLength.CopyStatus(&m.ProgramLength)
	m.StudentType.Init(m, "StudentType", "StudentType", 11)
	l.StudentType.CopyStatus(&m.StudentType)
	m.CredentialsType.Init(m, "CredentialsType", "CredentialsType", 12)
	l.CredentialsType.CopyStatus(&m.CredentialsType)
	m.DegreeType.Init(m, "DegreeType", "DegreeType", 13)
	l.DegreeType.CopyStatus(&m.DegreeType)
	m.Counselor.Init(m, "Counselor", "Counselor", 14)
	l.Counselor.CopyStatus(&m.Counselor)
	m.Adviser.Init(m, "Adviser", "Adviser", 15)
	l.Adviser.CopyStatus(&m.Adviser)
	m.Leadership.Init(m, "Leadership", "Leadership", 16)
	l.Leadership.CopyStatus(&m.Leadership)
	m.Supervisor.Init(m, "Supervisor", "Supervisor", 17)
	l.Supervisor.CopyStatus(&m.Supervisor)
	m.Assistant1.Init(m, "Assistant1", "Assistant1", 18)
	l.Assistant1.CopyStatus(&m.Assistant1)
	m.Assistant2.Init(m, "Assistant2", "Assistant2", 19)
	l.Assistant2.CopyStatus(&m.Assistant2)
	m.Operator.Init(m, "Operator", "Operator", 20)
	l.Operator.CopyStatus(&m.Operator)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 21)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 22)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", 23)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", 24)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", 25)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", 26)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", 27)
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
	if string(b) == "[]" || string(b) == "null" {
		return nil
	}
	jl := struct {
		List  *[]*Class
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
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

func NewGrade() *Grade {
	m := &Grade{}
	m.Init(m, nil, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.GradeName.Init(m, "GradeName", "GradeName", 1)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", 2)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 3)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 4)
	m.Status.Init(m, "Status", "Status", 5)
	m.Remark1.Init(m, "Remark1", "Remark1", 6)
	m.Remark2.Init(m, "Remark2", "Remark2", 7)
	m.Remark3.Init(m, "Remark3", "Remark3", 8)
	m.Remark4.Init(m, "Remark4", "Remark4", 9)
	m.InitRel()
	return m
}

func newSubGrade(parent nborm.Model) *Grade {
	m := &Grade{}
	m.Init(m, parent, nil)
	m.Id.Init(m, "Id", "Id", 0)
	m.GradeName.Init(m, "GradeName", "GradeName", 1)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", 2)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 3)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 4)
	m.Status.Init(m, "Status", "Status", 5)
	m.Remark1.Init(m, "Remark1", "Remark1", 6)
	m.Remark2.Init(m, "Remark2", "Remark2", 7)
	m.Remark3.Init(m, "Remark3", "Remark3", 8)
	m.Remark4.Init(m, "Remark4", "Remark4", 9)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.GradeName.Init(l, "GradeName", "GradeName", 1)
	l.GradeCode.Init(l, "GradeCode", "GradeCode", 2)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 3)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 4)
	l.Status.Init(l, "Status", "Status", 5)
	l.Remark1.Init(l, "Remark1", "Remark1", 6)
	l.Remark2.Init(l, "Remark2", "Remark2", 7)
	l.Remark3.Init(l, "Remark3", "Remark3", 8)
	l.Remark4.Init(l, "Remark4", "Remark4", 9)
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
	l.Id.Init(l, "Id", "Id", 0)
	l.GradeName.Init(l, "GradeName", "GradeName", 1)
	l.GradeCode.Init(l, "GradeCode", "GradeCode", 2)
	l.InsertDatetime.Init(l, "InsertDatetime", "InsertDatetime", 3)
	l.UpdateDatetime.Init(l, "UpdateDatetime", "UpdateDatetime", 4)
	l.Status.Init(l, "Status", "Status", 5)
	l.Remark1.Init(l, "Remark1", "Remark1", 6)
	l.Remark2.Init(l, "Remark2", "Remark2", 7)
	l.Remark3.Init(l, "Remark3", "Remark3", 8)
	l.Remark4.Init(l, "Remark4", "Remark4", 9)
	return l
}

func (l *GradeList) NewModel() nborm.Model {
	m := &Grade{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.Id.Init(m, "Id", "Id", 0)
	l.Id.CopyStatus(&m.Id)
	m.GradeName.Init(m, "GradeName", "GradeName", 1)
	l.GradeName.CopyStatus(&m.GradeName)
	m.GradeCode.Init(m, "GradeCode", "GradeCode", 2)
	l.GradeCode.CopyStatus(&m.GradeCode)
	m.InsertDatetime.Init(m, "InsertDatetime", "InsertDatetime", 3)
	l.InsertDatetime.CopyStatus(&m.InsertDatetime)
	m.UpdateDatetime.Init(m, "UpdateDatetime", "UpdateDatetime", 4)
	l.UpdateDatetime.CopyStatus(&m.UpdateDatetime)
	m.Status.Init(m, "Status", "Status", 5)
	l.Status.CopyStatus(&m.Status)
	m.Remark1.Init(m, "Remark1", "Remark1", 6)
	l.Remark1.CopyStatus(&m.Remark1)
	m.Remark2.Init(m, "Remark2", "Remark2", 7)
	l.Remark2.CopyStatus(&m.Remark2)
	m.Remark3.Init(m, "Remark3", "Remark3", 8)
	l.Remark3.CopyStatus(&m.Remark3)
	m.Remark4.Init(m, "Remark4", "Remark4", 9)
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
	if string(b) == "[]" || string(b) == "null" {
		return nil
	}
	jl := struct {
		List  *[]*Grade
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
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
