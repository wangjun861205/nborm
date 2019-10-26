package model

import (
	"github.com/wangjun861205/nborm"
)

//db:*
//tab:user
//pk:Id
//uk:IntelUserCode
//uk:IntelUserCode,Status
type User struct {
	nborm.Meta `json:"Aggs"`

	Id               nborm.Int         `col:"Id" auto_increment:"true"`
	IntelUserCode    nborm.String      `col:"IntelUserCode"`
	UserCode         nborm.String      `col:"UserCode"`
	Name             nborm.String      `col:"Name"`
	Sex              nborm.String      `col:"Sex"`
	IdentityType     nborm.String      `col:"IdentityType"`
	IdentityNum      nborm.String      `col:"IdentityNum"`
	ExpirationDate   nborm.String      `col:"ExpirationDate"`
	UniversityCode   nborm.String      `col:"UniversityCode"`
	UserType         nborm.String      `col:"UserType"`
	EnrollmentStatus nborm.String      `col:"EnrollmentStatus"`
	Type             nborm.String      `col:"Type"`
	Password         nborm.String      `col:"Password"`
	Phone            nborm.String      `col:"Phone"`
	Email            nborm.String      `col:"Email"`
	PictureURL       nborm.String      `col:"PictureURL"`
	Question         nborm.String      `col:"Question"`
	Answer           nborm.String      `col:"Answer"`
	AvailableLogin   nborm.String      `col:"AvailableLogin"`
	Operator         nborm.String      `col:"Operator"`
	InsertDatetime   nborm.Datetime    `col:"InsertDatetime"`
	UpdateDatetime   nborm.Datetime    `col:"UpdateDatetime"`
	Status           nborm.Int         `col:"Status"`
	Remark1          nborm.String      `col:"Remark1"`
	Remark2          nborm.String      `col:"Remark2"`
	Remark3          nborm.String      `col:"Remark3"`
	Remark4          nborm.String      `col:"Remark4"`
	Nonego           nborm.Int         `col:"Nonego"`
	BasicInfo        *Studentbasicinfo `rel:"[@@.IntelUserCode=@$.IntelUserCode]"`
}

//db:*
//tab:studentbasicinfo
//pk:Id
//uk:IntelUserCode
type Studentbasicinfo struct {
	nborm.Meta `json:"Aggs"`

	Id                        nborm.Int      `col:"Id" auto_increment:"true"`
	RecordId                  nborm.String   `col:"RecordId"`
	IntelUserCode             nborm.String   `col:"IntelUserCode"`
	Class                     nborm.String   `col:"Class"`
	OtherName                 nborm.String   `col:"OtherName"`
	NameInPinyin              nborm.String   `col:"NameInPinyin"`
	EnglishName               nborm.String   `col:"EnglishName"`
	CountryCode               nborm.String   `col:"CountryCode"`
	NationalityCode           nborm.String   `col:"NationalityCode"`
	Birthday                  nborm.String   `col:"Birthday"`
	PoliticalCode             nborm.String   `col:"PoliticalCode"`
	QQAcct                    nborm.String   `col:"QQAcct"`
	WeChatAcct                nborm.String   `col:"WeChatAcct"`
	BankCardNumber            nborm.String   `col:"BankCardNumber"`
	AccountBankCode           nborm.String   `col:"AccountBankCode"`
	AllPowerfulCardNum        nborm.String   `col:"AllPowerfulCardNum"`
	MaritalCode               nborm.String   `col:"MaritalCode"`
	OriginAreaCode            nborm.String   `col:"OriginAreaCode"`
	StudentAreaCode           nborm.String   `col:"StudentAreaCode"`
	Hobbies                   nborm.String   `col:"Hobbies"`
	Creed                     nborm.String   `col:"Creed"`
	TrainTicketinterval       nborm.String   `col:"TrainTicketinterval"`
	FamilyAddress             nborm.String   `col:"FamilyAddress"`
	DetailAddress             nborm.String   `col:"DetailAddress"`
	PostCode                  nborm.String   `col:"PostCode"`
	HomePhone                 nborm.String   `col:"HomePhone"`
	EnrollmentDate            nborm.String   `col:"EnrollmentDate"`
	GraduationDate            nborm.String   `col:"GraduationDate"`
	MidSchoolAddress          nborm.String   `col:"MidSchoolAddress"`
	MidSchoolName             nborm.String   `col:"MidSchoolName"`
	Referee                   nborm.String   `col:"Referee"`
	RefereeDuty               nborm.String   `col:"RefereeDuty"`
	RefereePhone              nborm.String   `col:"RefereePhone"`
	AdmissionTicketNo         nborm.String   `col:"AdmissionTicketNo"`
	CollegeEntranceExamScores nborm.String   `col:"CollegeEntranceExamScores"`
	AdmissionYear             nborm.String   `col:"AdmissionYear"`
	ForeignLanguageCode       nborm.String   `col:"ForeignLanguageCode"`
	StudentOrigin             nborm.String   `col:"StudentOrigin"`
	BizType                   nborm.String   `col:"BizType"`
	TaskCode                  nborm.String   `col:"TaskCode"`
	ApproveStatus             nborm.String   `col:"ApproveStatus"`
	Operator                  nborm.String   `col:"Operator"`
	InsertDatetime            nborm.Datetime `col:"InsertDatetime"`
	UpdateDatetime            nborm.Datetime `col:"UpdateDatetime"`
	Status                    nborm.Int      `col:"Status"`
	StudentStatus             nborm.String   `col:"StudentStatus"`
	IsAuth                    nborm.Int      `col:"IsAuth"`
	Campus                    nborm.String   `col:"Campus"`
	Zone                      nborm.String   `col:"Zone"`
	Building                  nborm.String   `col:"Building"`
	Unit                      nborm.String   `col:"Unit"`
	Room                      nborm.String   `col:"Room"`
	Bed                       nborm.String   `col:"Bed"`
	StatusSort                nborm.String   `col:"StatusSort"`
	Height                    nborm.Int      `col:"Height"`
	Weight                    nborm.String   `col:"Weight"`
	FootSize                  nborm.Decimal  `col:"FootSize"`
	ClothSize                 nborm.String   `col:"ClothSize"`
	HeadSize                  nborm.Int      `col:"HeadSize"`
	Remark1                   nborm.String   `col:"Remark1"`
	Remark2                   nborm.String   `col:"Remark2"`
	Remark3                   nborm.String   `col:"Remark3"`
	Remark4                   nborm.String   `col:"Remark4"`
	IsPayment                 nborm.Int      `col:"IsPayment"`
	IsCheckIn                 nborm.Int      `col:"isCheckIn"`
	GetMilitaryTC             nborm.Int      `col:"GetMilitaryTC"`
	OriginAreaName            nborm.String   `col:"OriginAreaName"`
	User                      *User          `rel:"[@@.IntelUserCode=@$.IntelUserCode]"`
}
