package model

import (
	"github.com/wangjun861205/nborm"
)

//db:*
//tab:employ_account
//pk:SID
//uk:Phone
//uk:ID
type EmployAccount struct {
	nborm.Meta `json:"Aggs"`

	SID           nborm.Int      `col:"SID" auto_increment:"true"`
	ID            nborm.String   `col:"ID"`
	Phone         nborm.String   `col:"Phone"`
	Password      nborm.String   `col:"Password"`
	Status        nborm.Int      `col:"Status"`
	ErrorCount    nborm.Int      `col:"ErrorCount"`
	LastErrorTime nborm.Datetime `col:"LastErrorTime"`
	LastLogin     nborm.Datetime `col:"LastLogin"`
	CreateTime    nborm.Datetime `col:"CreateTime"`
	UpdateTime    nborm.Datetime `col:"UpdateTime"`
}

//db:*
//tab:employ_enterprise
//pk:SID
//uk:AccountID
//uk:Email
//uk:UniformCode
//uk:ID
type EmployEnterprise struct {
	nborm.Meta `json:"Aggs"`

	SID            nborm.Int      `col:"SID" auto_increment:"true"`
	ID             nborm.String   `col:"ID"`
	AccountID      nborm.String   `col:"AccountID"`
	Email          nborm.String   `col:"Email"`
	UniformCode    nborm.String   `col:"UniformCode"`
	Name           nborm.String   `col:"Name"`
	RegisterCityID nborm.String   `col:"RegisterCityID"`
	SectorID       nborm.String   `col:"SectorID"`
	NatureID       nborm.String   `col:"NatureID"`
	ScopeID        nborm.String   `col:"ScopeID"`
	OfficeCityID   nborm.String   `col:"OfficeCityID"`
	OfficeAddress  nborm.String   `col:"OfficeAddress"`
	Website        nborm.String   `col:"Website"`
	Contact        nborm.String   `col:"Contact"`
	ContactPhone   nborm.String   `col:"ContactPhone"`
	EmployFromThis nborm.Int      `col:"EmployFromThis"`
	Introduction   nborm.String   `col:"Introduction"`
	Zipcode        nborm.String   `col:"Zipcode"`
	Status         nborm.Int      `col:"Status"`
	UpdateHash     nborm.String   `col:"UpdateHash"`
	RejectReason   nborm.String   `col:"RejectReason"`
	LicenseImageID nborm.String   `col:"LicenseImageID"`
	CreateTime     nborm.Datetime `col:"CreateTime"`
	UpdateTime     nborm.Datetime `col:"UpdateTime"`
	Account        *EmployAccount `rel:"[@@.AccountID=@$.ID]"`
}
