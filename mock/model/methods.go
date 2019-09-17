package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wangjun861205/nborm"
	"strings"
	"time"
)

func NewEmployAccount() *EmployAccount {
	m := &EmployAccount{}
	m.Init(m, nil, nil)
	m.SID.Init(m, "SID", "SID", 0)
	m.ID.Init(m, "ID", "ID", 1)
	m.Phone.Init(m, "Phone", "Phone", 2)
	m.Password.Init(m, "Password", "Password", 3)
	m.Status.Init(m, "Status", "Status", 4)
	m.ErrorCount.Init(m, "ErrorCount", "ErrorCount", 5)
	m.LastErrorTime.Init(m, "LastErrorTime", "LastErrorTime", 6)
	m.LastLogin.Init(m, "LastLogin", "LastLogin", 7)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 8)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 9)
	m.InitRel()
	return m
}

func newSubEmployAccount(parent nborm.Model) *EmployAccount {
	m := &EmployAccount{}
	m.Init(m, parent, nil)
	m.SID.Init(m, "SID", "SID", 0)
	m.ID.Init(m, "ID", "ID", 1)
	m.Phone.Init(m, "Phone", "Phone", 2)
	m.Password.Init(m, "Password", "Password", 3)
	m.Status.Init(m, "Status", "Status", 4)
	m.ErrorCount.Init(m, "ErrorCount", "ErrorCount", 5)
	m.LastErrorTime.Init(m, "LastErrorTime", "LastErrorTime", 6)
	m.LastLogin.Init(m, "LastLogin", "LastLogin", 7)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 8)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 9)
	return m
}

func (m *EmployAccount) InitRel() {
	m.AddRelInited()
}

func (m *EmployAccount) DB() string {
	return "*"
}

func (m *EmployAccount) Tab() string {
	return "employ_account"
}

func (m *EmployAccount) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"SID", "SID", &m.SID, 0},
		{"ID", "ID", &m.ID, 1},
		{"Phone", "Phone", &m.Phone, 2},
		{"Password", "Password", &m.Password, 3},
		{"Status", "Status", &m.Status, 4},
		{"ErrorCount", "ErrorCount", &m.ErrorCount, 5},
		{"LastErrorTime", "LastErrorTime", &m.LastErrorTime, 6},
		{"LastLogin", "LastLogin", &m.LastLogin, 7},
		{"CreateTime", "CreateTime", &m.CreateTime, 8},
		{"UpdateTime", "UpdateTime", &m.UpdateTime, 9},
	}
}

func (m *EmployAccount) AutoIncField() nborm.Field {
	return &m.SID
}

func (m *EmployAccount) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.SID,
	}
}

func (m *EmployAccount) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.Phone,
		},
		{
			&m.ID,
		},
	}
}
func (m EmployAccount) MarshalJSON() ([]byte, error) {
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
	if m.SID.IsValid() {
		buffer.WriteString(",\n\"SID\": ")
		SIDB, err := json.MarshalIndent(m.SID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(SIDB)
	}
	if m.ID.IsValid() {
		buffer.WriteString(",\n\"ID\": ")
		IDB, err := json.MarshalIndent(m.ID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IDB)
	}
	if m.Phone.IsValid() {
		buffer.WriteString(",\n\"Phone\": ")
		PhoneB, err := json.MarshalIndent(m.Phone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PhoneB)
	}
	if m.Password.IsValid() {
		buffer.WriteString(",\n\"Password\": ")
		PasswordB, err := json.MarshalIndent(m.Password, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(PasswordB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.ErrorCount.IsValid() {
		buffer.WriteString(",\n\"ErrorCount\": ")
		ErrorCountB, err := json.MarshalIndent(m.ErrorCount, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ErrorCountB)
	}
	if m.LastErrorTime.IsValid() {
		buffer.WriteString(",\n\"LastErrorTime\": ")
		LastErrorTimeB, err := json.MarshalIndent(m.LastErrorTime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(LastErrorTimeB)
	}
	if m.LastLogin.IsValid() {
		buffer.WriteString(",\n\"LastLogin\": ")
		LastLoginB, err := json.MarshalIndent(m.LastLogin, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(LastLoginB)
	}
	if m.CreateTime.IsValid() {
		buffer.WriteString(",\n\"CreateTime\": ")
		CreateTimeB, err := json.MarshalIndent(m.CreateTime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CreateTimeB)
	}
	if m.UpdateTime.IsValid() {
		buffer.WriteString(",\n\"UpdateTime\": ")
		UpdateTimeB, err := json.MarshalIndent(m.UpdateTime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateTimeB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type EmployAccountList struct {
	EmployAccount `json:"-"`
	dupMap        map[string]int
	List          []*EmployAccount
	Total         int
}

func (m *EmployAccount) Collapse() {
}

func NewEmployAccountList() *EmployAccountList {
	l := &EmployAccountList{
		EmployAccount{},
		make(map[string]int),
		make([]*EmployAccount, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.SID.Init(l, "SID", "SID", 0)
	l.ID.Init(l, "ID", "ID", 1)
	l.Phone.Init(l, "Phone", "Phone", 2)
	l.Password.Init(l, "Password", "Password", 3)
	l.Status.Init(l, "Status", "Status", 4)
	l.ErrorCount.Init(l, "ErrorCount", "ErrorCount", 5)
	l.LastErrorTime.Init(l, "LastErrorTime", "LastErrorTime", 6)
	l.LastLogin.Init(l, "LastLogin", "LastLogin", 7)
	l.CreateTime.Init(l, "CreateTime", "CreateTime", 8)
	l.UpdateTime.Init(l, "UpdateTime", "UpdateTime", 9)
	l.InitRel()
	return l
}

func newSubEmployAccountList(parent nborm.Model) *EmployAccountList {
	l := &EmployAccountList{
		EmployAccount{},
		make(map[string]int),
		make([]*EmployAccount, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.SID.Init(l, "SID", "SID", 0)
	l.ID.Init(l, "ID", "ID", 1)
	l.Phone.Init(l, "Phone", "Phone", 2)
	l.Password.Init(l, "Password", "Password", 3)
	l.Status.Init(l, "Status", "Status", 4)
	l.ErrorCount.Init(l, "ErrorCount", "ErrorCount", 5)
	l.LastErrorTime.Init(l, "LastErrorTime", "LastErrorTime", 6)
	l.LastLogin.Init(l, "LastLogin", "LastLogin", 7)
	l.CreateTime.Init(l, "CreateTime", "CreateTime", 8)
	l.UpdateTime.Init(l, "UpdateTime", "UpdateTime", 9)
	return l
}

func (l *EmployAccountList) NewModel() nborm.Model {
	m := &EmployAccount{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.SID.Init(m, "SID", "SID", 0)
	l.SID.CopyStatus(&m.SID)
	m.ID.Init(m, "ID", "ID", 1)
	l.ID.CopyStatus(&m.ID)
	m.Phone.Init(m, "Phone", "Phone", 2)
	l.Phone.CopyStatus(&m.Phone)
	m.Password.Init(m, "Password", "Password", 3)
	l.Password.CopyStatus(&m.Password)
	m.Status.Init(m, "Status", "Status", 4)
	l.Status.CopyStatus(&m.Status)
	m.ErrorCount.Init(m, "ErrorCount", "ErrorCount", 5)
	l.ErrorCount.CopyStatus(&m.ErrorCount)
	m.LastErrorTime.Init(m, "LastErrorTime", "LastErrorTime", 6)
	l.LastErrorTime.CopyStatus(&m.LastErrorTime)
	m.LastLogin.Init(m, "LastLogin", "LastLogin", 7)
	l.LastLogin.CopyStatus(&m.LastLogin)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 8)
	l.CreateTime.CopyStatus(&m.CreateTime)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 9)
	l.UpdateTime.CopyStatus(&m.UpdateTime)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EmployAccountList) SetTotal(total int) {
	l.Total = total
}

func (l *EmployAccountList) GetTotal() int {
	return l.Total
}

func (l *EmployAccountList) Len() int {
	return len(l.List)
}

func (l *EmployAccountList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EmployAccountList) MarshalJSON() ([]byte, error) {
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

func (l *EmployAccountList) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	jl := struct {
		List  *[]*EmployAccount
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
}

func (l *EmployAccountList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *EmployAccountList) Filter(f func(m *EmployAccount) bool) []*EmployAccount {
	ll := make([]*EmployAccount, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *EmployAccountList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.SID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.SID.AnyValue()))
	}
	if lastModel.ID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ID.AnyValue()))
	}
	if lastModel.Phone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Phone.AnyValue()))
	}
	if lastModel.Password.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Password.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.ErrorCount.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ErrorCount.AnyValue()))
	}
	if lastModel.LastErrorTime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.LastErrorTime.AnyValue()))
	}
	if lastModel.LastLogin.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.LastLogin.AnyValue()))
	}
	if lastModel.CreateTime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.CreateTime.AnyValue()))
	}
	if lastModel.UpdateTime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateTime.AnyValue()))
	}
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *EmployAccountList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

type EmployAccountCacheElem struct {
	hashValue  string
	model      *EmployAccount
	modifyTime time.Time
}

type EmployAccountListCacheElem struct {
	hashValue  string
	list       *EmployAccountList
	modifyTime time.Time
}

type EmployAccountCacheManager struct {
	container map[string]*EmployAccountCacheElem
	query     chan string
	in        chan *EmployAccountCacheElem
	out       chan *EmployAccountCacheElem
}

type EmployAccountListCacheManager struct {
	container map[string]*EmployAccountListCacheElem
	query     chan string
	in        chan *EmployAccountListCacheElem
	out       chan *EmployAccountListCacheElem
}

func newEmployAccountCacheManager() *EmployAccountCacheManager {
	return &EmployAccountCacheManager{
		make(map[string]*EmployAccountCacheElem),
		make(chan string),
		make(chan *EmployAccountCacheElem),
		make(chan *EmployAccountCacheElem),
	}
}

func newEmployAccountListCacheManager() *EmployAccountListCacheManager {
	return &EmployAccountListCacheManager{
		make(map[string]*EmployAccountListCacheElem),
		make(chan string),
		make(chan *EmployAccountListCacheElem),
		make(chan *EmployAccountListCacheElem),
	}
}

func (mgr *EmployAccountCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *EmployAccountListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var EmployAccountCache = newEmployAccountCacheManager()

var EmployAccountListCache = newEmployAccountListCacheManager()

func (m *EmployAccount) GetCache(hashVal string, timeout time.Duration) bool {
	EmployAccountCache.query <- hashVal
	elem := <-EmployAccountCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *EmployAccount) SetCache(hashValue string) {
	EmployAccountCache.in <- &EmployAccountCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *EmployAccountList) GetListCache(hashValue string, timeout time.Duration) bool {
	EmployAccountListCache.query <- hashValue
	elem := <-EmployAccountListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *EmployAccountList) SetListCache(hashValue string) {
	EmployAccountListCache.in <- &EmployAccountListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func NewEmployEnterprise() *EmployEnterprise {
	m := &EmployEnterprise{}
	m.Init(m, nil, nil)
	m.SID.Init(m, "SID", "SID", 0)
	m.ID.Init(m, "ID", "ID", 1)
	m.AccountID.Init(m, "AccountID", "AccountID", 2)
	m.Email.Init(m, "Email", "Email", 3)
	m.UniformCode.Init(m, "UniformCode", "UniformCode", 4)
	m.Name.Init(m, "Name", "Name", 5)
	m.RegisterCityID.Init(m, "RegisterCityID", "RegisterCityID", 6)
	m.SectorID.Init(m, "SectorID", "SectorID", 7)
	m.NatureID.Init(m, "NatureID", "NatureID", 8)
	m.ScopeID.Init(m, "ScopeID", "ScopeID", 9)
	m.OfficeCityID.Init(m, "OfficeCityID", "OfficeCityID", 10)
	m.OfficeAddress.Init(m, "OfficeAddress", "OfficeAddress", 11)
	m.Website.Init(m, "Website", "Website", 12)
	m.Contact.Init(m, "Contact", "Contact", 13)
	m.ContactPhone.Init(m, "ContactPhone", "ContactPhone", 14)
	m.EmployFromThis.Init(m, "EmployFromThis", "EmployFromThis", 15)
	m.Introduction.Init(m, "Introduction", "Introduction", 16)
	m.Zipcode.Init(m, "Zipcode", "Zipcode", 17)
	m.Status.Init(m, "Status", "Status", 18)
	m.UpdateHash.Init(m, "UpdateHash", "UpdateHash", 19)
	m.RejectReason.Init(m, "RejectReason", "RejectReason", 20)
	m.LicenseImageID.Init(m, "LicenseImageID", "LicenseImageID", 21)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 22)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 23)
	m.InitRel()
	return m
}

func newSubEmployEnterprise(parent nborm.Model) *EmployEnterprise {
	m := &EmployEnterprise{}
	m.Init(m, parent, nil)
	m.SID.Init(m, "SID", "SID", 0)
	m.ID.Init(m, "ID", "ID", 1)
	m.AccountID.Init(m, "AccountID", "AccountID", 2)
	m.Email.Init(m, "Email", "Email", 3)
	m.UniformCode.Init(m, "UniformCode", "UniformCode", 4)
	m.Name.Init(m, "Name", "Name", 5)
	m.RegisterCityID.Init(m, "RegisterCityID", "RegisterCityID", 6)
	m.SectorID.Init(m, "SectorID", "SectorID", 7)
	m.NatureID.Init(m, "NatureID", "NatureID", 8)
	m.ScopeID.Init(m, "ScopeID", "ScopeID", 9)
	m.OfficeCityID.Init(m, "OfficeCityID", "OfficeCityID", 10)
	m.OfficeAddress.Init(m, "OfficeAddress", "OfficeAddress", 11)
	m.Website.Init(m, "Website", "Website", 12)
	m.Contact.Init(m, "Contact", "Contact", 13)
	m.ContactPhone.Init(m, "ContactPhone", "ContactPhone", 14)
	m.EmployFromThis.Init(m, "EmployFromThis", "EmployFromThis", 15)
	m.Introduction.Init(m, "Introduction", "Introduction", 16)
	m.Zipcode.Init(m, "Zipcode", "Zipcode", 17)
	m.Status.Init(m, "Status", "Status", 18)
	m.UpdateHash.Init(m, "UpdateHash", "UpdateHash", 19)
	m.RejectReason.Init(m, "RejectReason", "RejectReason", 20)
	m.LicenseImageID.Init(m, "LicenseImageID", "LicenseImageID", 21)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 22)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 23)
	return m
}

func (m *EmployEnterprise) InitRel() {
	m.AddRelInited()
}

func (m *EmployEnterprise) DB() string {
	return "*"
}

func (m *EmployEnterprise) Tab() string {
	return "employ_enterprise"
}

func (m *EmployEnterprise) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"SID", "SID", &m.SID, 0},
		{"ID", "ID", &m.ID, 1},
		{"AccountID", "AccountID", &m.AccountID, 2},
		{"Email", "Email", &m.Email, 3},
		{"UniformCode", "UniformCode", &m.UniformCode, 4},
		{"Name", "Name", &m.Name, 5},
		{"RegisterCityID", "RegisterCityID", &m.RegisterCityID, 6},
		{"SectorID", "SectorID", &m.SectorID, 7},
		{"NatureID", "NatureID", &m.NatureID, 8},
		{"ScopeID", "ScopeID", &m.ScopeID, 9},
		{"OfficeCityID", "OfficeCityID", &m.OfficeCityID, 10},
		{"OfficeAddress", "OfficeAddress", &m.OfficeAddress, 11},
		{"Website", "Website", &m.Website, 12},
		{"Contact", "Contact", &m.Contact, 13},
		{"ContactPhone", "ContactPhone", &m.ContactPhone, 14},
		{"EmployFromThis", "EmployFromThis", &m.EmployFromThis, 15},
		{"Introduction", "Introduction", &m.Introduction, 16},
		{"Zipcode", "Zipcode", &m.Zipcode, 17},
		{"Status", "Status", &m.Status, 18},
		{"UpdateHash", "UpdateHash", &m.UpdateHash, 19},
		{"RejectReason", "RejectReason", &m.RejectReason, 20},
		{"LicenseImageID", "LicenseImageID", &m.LicenseImageID, 21},
		{"CreateTime", "CreateTime", &m.CreateTime, 22},
		{"UpdateTime", "UpdateTime", &m.UpdateTime, 23},
	}
}

func (m *EmployEnterprise) AutoIncField() nborm.Field {
	return &m.SID
}

func (m *EmployEnterprise) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.SID,
	}
}

func (m *EmployEnterprise) UniqueKeys() []nborm.FieldList {
	return []nborm.FieldList{
		{
			&m.AccountID,
		},
		{
			&m.Email,
		},
		{
			&m.UniformCode,
		},
		{
			&m.ID,
		},
	}
}
func (m EmployEnterprise) MarshalJSON() ([]byte, error) {
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
	if m.SID.IsValid() {
		buffer.WriteString(",\n\"SID\": ")
		SIDB, err := json.MarshalIndent(m.SID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(SIDB)
	}
	if m.ID.IsValid() {
		buffer.WriteString(",\n\"ID\": ")
		IDB, err := json.MarshalIndent(m.ID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IDB)
	}
	if m.AccountID.IsValid() {
		buffer.WriteString(",\n\"AccountID\": ")
		AccountIDB, err := json.MarshalIndent(m.AccountID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(AccountIDB)
	}
	if m.Email.IsValid() {
		buffer.WriteString(",\n\"Email\": ")
		EmailB, err := json.MarshalIndent(m.Email, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EmailB)
	}
	if m.UniformCode.IsValid() {
		buffer.WriteString(",\n\"UniformCode\": ")
		UniformCodeB, err := json.MarshalIndent(m.UniformCode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UniformCodeB)
	}
	if m.Name.IsValid() {
		buffer.WriteString(",\n\"Name\": ")
		NameB, err := json.MarshalIndent(m.Name, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NameB)
	}
	if m.RegisterCityID.IsValid() {
		buffer.WriteString(",\n\"RegisterCityID\": ")
		RegisterCityIDB, err := json.MarshalIndent(m.RegisterCityID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RegisterCityIDB)
	}
	if m.SectorID.IsValid() {
		buffer.WriteString(",\n\"SectorID\": ")
		SectorIDB, err := json.MarshalIndent(m.SectorID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(SectorIDB)
	}
	if m.NatureID.IsValid() {
		buffer.WriteString(",\n\"NatureID\": ")
		NatureIDB, err := json.MarshalIndent(m.NatureID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(NatureIDB)
	}
	if m.ScopeID.IsValid() {
		buffer.WriteString(",\n\"ScopeID\": ")
		ScopeIDB, err := json.MarshalIndent(m.ScopeID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ScopeIDB)
	}
	if m.OfficeCityID.IsValid() {
		buffer.WriteString(",\n\"OfficeCityID\": ")
		OfficeCityIDB, err := json.MarshalIndent(m.OfficeCityID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OfficeCityIDB)
	}
	if m.OfficeAddress.IsValid() {
		buffer.WriteString(",\n\"OfficeAddress\": ")
		OfficeAddressB, err := json.MarshalIndent(m.OfficeAddress, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(OfficeAddressB)
	}
	if m.Website.IsValid() {
		buffer.WriteString(",\n\"Website\": ")
		WebsiteB, err := json.MarshalIndent(m.Website, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(WebsiteB)
	}
	if m.Contact.IsValid() {
		buffer.WriteString(",\n\"Contact\": ")
		ContactB, err := json.MarshalIndent(m.Contact, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ContactB)
	}
	if m.ContactPhone.IsValid() {
		buffer.WriteString(",\n\"ContactPhone\": ")
		ContactPhoneB, err := json.MarshalIndent(m.ContactPhone, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ContactPhoneB)
	}
	if m.EmployFromThis.IsValid() {
		buffer.WriteString(",\n\"EmployFromThis\": ")
		EmployFromThisB, err := json.MarshalIndent(m.EmployFromThis, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(EmployFromThisB)
	}
	if m.Introduction.IsValid() {
		buffer.WriteString(",\n\"Introduction\": ")
		IntroductionB, err := json.MarshalIndent(m.Introduction, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(IntroductionB)
	}
	if m.Zipcode.IsValid() {
		buffer.WriteString(",\n\"Zipcode\": ")
		ZipcodeB, err := json.MarshalIndent(m.Zipcode, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(ZipcodeB)
	}
	if m.Status.IsValid() {
		buffer.WriteString(",\n\"Status\": ")
		StatusB, err := json.MarshalIndent(m.Status, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(StatusB)
	}
	if m.UpdateHash.IsValid() {
		buffer.WriteString(",\n\"UpdateHash\": ")
		UpdateHashB, err := json.MarshalIndent(m.UpdateHash, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateHashB)
	}
	if m.RejectReason.IsValid() {
		buffer.WriteString(",\n\"RejectReason\": ")
		RejectReasonB, err := json.MarshalIndent(m.RejectReason, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(RejectReasonB)
	}
	if m.LicenseImageID.IsValid() {
		buffer.WriteString(",\n\"LicenseImageID\": ")
		LicenseImageIDB, err := json.MarshalIndent(m.LicenseImageID, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(LicenseImageIDB)
	}
	if m.CreateTime.IsValid() {
		buffer.WriteString(",\n\"CreateTime\": ")
		CreateTimeB, err := json.MarshalIndent(m.CreateTime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(CreateTimeB)
	}
	if m.UpdateTime.IsValid() {
		buffer.WriteString(",\n\"UpdateTime\": ")
		UpdateTimeB, err := json.MarshalIndent(m.UpdateTime, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(UpdateTimeB)
	}
	buffer.WriteString("\n}")
	return buffer.Bytes(), nil
}

type EmployEnterpriseList struct {
	EmployEnterprise `json:"-"`
	dupMap           map[string]int
	List             []*EmployEnterprise
	Total            int
}

func (m *EmployEnterprise) Collapse() {
}

func NewEmployEnterpriseList() *EmployEnterpriseList {
	l := &EmployEnterpriseList{
		EmployEnterprise{},
		make(map[string]int),
		make([]*EmployEnterprise, 0, 32),
		0,
	}
	l.Init(l, nil, nil)
	l.SID.Init(l, "SID", "SID", 0)
	l.ID.Init(l, "ID", "ID", 1)
	l.AccountID.Init(l, "AccountID", "AccountID", 2)
	l.Email.Init(l, "Email", "Email", 3)
	l.UniformCode.Init(l, "UniformCode", "UniformCode", 4)
	l.Name.Init(l, "Name", "Name", 5)
	l.RegisterCityID.Init(l, "RegisterCityID", "RegisterCityID", 6)
	l.SectorID.Init(l, "SectorID", "SectorID", 7)
	l.NatureID.Init(l, "NatureID", "NatureID", 8)
	l.ScopeID.Init(l, "ScopeID", "ScopeID", 9)
	l.OfficeCityID.Init(l, "OfficeCityID", "OfficeCityID", 10)
	l.OfficeAddress.Init(l, "OfficeAddress", "OfficeAddress", 11)
	l.Website.Init(l, "Website", "Website", 12)
	l.Contact.Init(l, "Contact", "Contact", 13)
	l.ContactPhone.Init(l, "ContactPhone", "ContactPhone", 14)
	l.EmployFromThis.Init(l, "EmployFromThis", "EmployFromThis", 15)
	l.Introduction.Init(l, "Introduction", "Introduction", 16)
	l.Zipcode.Init(l, "Zipcode", "Zipcode", 17)
	l.Status.Init(l, "Status", "Status", 18)
	l.UpdateHash.Init(l, "UpdateHash", "UpdateHash", 19)
	l.RejectReason.Init(l, "RejectReason", "RejectReason", 20)
	l.LicenseImageID.Init(l, "LicenseImageID", "LicenseImageID", 21)
	l.CreateTime.Init(l, "CreateTime", "CreateTime", 22)
	l.UpdateTime.Init(l, "UpdateTime", "UpdateTime", 23)
	l.InitRel()
	return l
}

func newSubEmployEnterpriseList(parent nborm.Model) *EmployEnterpriseList {
	l := &EmployEnterpriseList{
		EmployEnterprise{},
		make(map[string]int),
		make([]*EmployEnterprise, 0, 32),
		0,
	}
	l.Init(l, parent, nil)
	l.SID.Init(l, "SID", "SID", 0)
	l.ID.Init(l, "ID", "ID", 1)
	l.AccountID.Init(l, "AccountID", "AccountID", 2)
	l.Email.Init(l, "Email", "Email", 3)
	l.UniformCode.Init(l, "UniformCode", "UniformCode", 4)
	l.Name.Init(l, "Name", "Name", 5)
	l.RegisterCityID.Init(l, "RegisterCityID", "RegisterCityID", 6)
	l.SectorID.Init(l, "SectorID", "SectorID", 7)
	l.NatureID.Init(l, "NatureID", "NatureID", 8)
	l.ScopeID.Init(l, "ScopeID", "ScopeID", 9)
	l.OfficeCityID.Init(l, "OfficeCityID", "OfficeCityID", 10)
	l.OfficeAddress.Init(l, "OfficeAddress", "OfficeAddress", 11)
	l.Website.Init(l, "Website", "Website", 12)
	l.Contact.Init(l, "Contact", "Contact", 13)
	l.ContactPhone.Init(l, "ContactPhone", "ContactPhone", 14)
	l.EmployFromThis.Init(l, "EmployFromThis", "EmployFromThis", 15)
	l.Introduction.Init(l, "Introduction", "Introduction", 16)
	l.Zipcode.Init(l, "Zipcode", "Zipcode", 17)
	l.Status.Init(l, "Status", "Status", 18)
	l.UpdateHash.Init(l, "UpdateHash", "UpdateHash", 19)
	l.RejectReason.Init(l, "RejectReason", "RejectReason", 20)
	l.LicenseImageID.Init(l, "LicenseImageID", "LicenseImageID", 21)
	l.CreateTime.Init(l, "CreateTime", "CreateTime", 22)
	l.UpdateTime.Init(l, "UpdateTime", "UpdateTime", 23)
	return l
}

func (l *EmployEnterpriseList) NewModel() nborm.Model {
	m := &EmployEnterprise{}
	m.Init(m, nil, l)
	l.CopyAggs(m)
	m.SID.Init(m, "SID", "SID", 0)
	l.SID.CopyStatus(&m.SID)
	m.ID.Init(m, "ID", "ID", 1)
	l.ID.CopyStatus(&m.ID)
	m.AccountID.Init(m, "AccountID", "AccountID", 2)
	l.AccountID.CopyStatus(&m.AccountID)
	m.Email.Init(m, "Email", "Email", 3)
	l.Email.CopyStatus(&m.Email)
	m.UniformCode.Init(m, "UniformCode", "UniformCode", 4)
	l.UniformCode.CopyStatus(&m.UniformCode)
	m.Name.Init(m, "Name", "Name", 5)
	l.Name.CopyStatus(&m.Name)
	m.RegisterCityID.Init(m, "RegisterCityID", "RegisterCityID", 6)
	l.RegisterCityID.CopyStatus(&m.RegisterCityID)
	m.SectorID.Init(m, "SectorID", "SectorID", 7)
	l.SectorID.CopyStatus(&m.SectorID)
	m.NatureID.Init(m, "NatureID", "NatureID", 8)
	l.NatureID.CopyStatus(&m.NatureID)
	m.ScopeID.Init(m, "ScopeID", "ScopeID", 9)
	l.ScopeID.CopyStatus(&m.ScopeID)
	m.OfficeCityID.Init(m, "OfficeCityID", "OfficeCityID", 10)
	l.OfficeCityID.CopyStatus(&m.OfficeCityID)
	m.OfficeAddress.Init(m, "OfficeAddress", "OfficeAddress", 11)
	l.OfficeAddress.CopyStatus(&m.OfficeAddress)
	m.Website.Init(m, "Website", "Website", 12)
	l.Website.CopyStatus(&m.Website)
	m.Contact.Init(m, "Contact", "Contact", 13)
	l.Contact.CopyStatus(&m.Contact)
	m.ContactPhone.Init(m, "ContactPhone", "ContactPhone", 14)
	l.ContactPhone.CopyStatus(&m.ContactPhone)
	m.EmployFromThis.Init(m, "EmployFromThis", "EmployFromThis", 15)
	l.EmployFromThis.CopyStatus(&m.EmployFromThis)
	m.Introduction.Init(m, "Introduction", "Introduction", 16)
	l.Introduction.CopyStatus(&m.Introduction)
	m.Zipcode.Init(m, "Zipcode", "Zipcode", 17)
	l.Zipcode.CopyStatus(&m.Zipcode)
	m.Status.Init(m, "Status", "Status", 18)
	l.Status.CopyStatus(&m.Status)
	m.UpdateHash.Init(m, "UpdateHash", "UpdateHash", 19)
	l.UpdateHash.CopyStatus(&m.UpdateHash)
	m.RejectReason.Init(m, "RejectReason", "RejectReason", 20)
	l.RejectReason.CopyStatus(&m.RejectReason)
	m.LicenseImageID.Init(m, "LicenseImageID", "LicenseImageID", 21)
	l.LicenseImageID.CopyStatus(&m.LicenseImageID)
	m.CreateTime.Init(m, "CreateTime", "CreateTime", 22)
	l.CreateTime.CopyStatus(&m.CreateTime)
	m.UpdateTime.Init(m, "UpdateTime", "UpdateTime", 23)
	l.UpdateTime.CopyStatus(&m.UpdateTime)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EmployEnterpriseList) SetTotal(total int) {
	l.Total = total
}

func (l *EmployEnterpriseList) GetTotal() int {
	return l.Total
}

func (l *EmployEnterpriseList) Len() int {
	return len(l.List)
}

func (l *EmployEnterpriseList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EmployEnterpriseList) MarshalJSON() ([]byte, error) {
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

func (l *EmployEnterpriseList) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	jl := struct {
		List  *[]*EmployEnterprise
		Total *int
	}{
		&l.List,
		&l.Total,
	}
	return json.Unmarshal(b, &jl)
}

func (l *EmployEnterpriseList) Collapse() {
	idx := l.checkDup()
	if idx >= 0 {
		l.List = l.List[:len(l.List)-1]
		l.List[idx].Collapse()
	}
}

func (l *EmployEnterpriseList) Filter(f func(m *EmployEnterprise) bool) []*EmployEnterprise {
	ll := make([]*EmployEnterprise, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func (l *EmployEnterpriseList) checkDup() int {
	if l.Len() < 1 {
		return -1
	}
	var builder strings.Builder
	lastModel := l.List[l.Len()-1]
	if lastModel.SID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.SID.AnyValue()))
	}
	if lastModel.ID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ID.AnyValue()))
	}
	if lastModel.AccountID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.AccountID.AnyValue()))
	}
	if lastModel.Email.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Email.AnyValue()))
	}
	if lastModel.UniformCode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UniformCode.AnyValue()))
	}
	if lastModel.Name.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Name.AnyValue()))
	}
	if lastModel.RegisterCityID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RegisterCityID.AnyValue()))
	}
	if lastModel.SectorID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.SectorID.AnyValue()))
	}
	if lastModel.NatureID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.NatureID.AnyValue()))
	}
	if lastModel.ScopeID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ScopeID.AnyValue()))
	}
	if lastModel.OfficeCityID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.OfficeCityID.AnyValue()))
	}
	if lastModel.OfficeAddress.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.OfficeAddress.AnyValue()))
	}
	if lastModel.Website.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Website.AnyValue()))
	}
	if lastModel.Contact.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Contact.AnyValue()))
	}
	if lastModel.ContactPhone.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.ContactPhone.AnyValue()))
	}
	if lastModel.EmployFromThis.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.EmployFromThis.AnyValue()))
	}
	if lastModel.Introduction.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Introduction.AnyValue()))
	}
	if lastModel.Zipcode.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Zipcode.AnyValue()))
	}
	if lastModel.Status.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.Status.AnyValue()))
	}
	if lastModel.UpdateHash.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateHash.AnyValue()))
	}
	if lastModel.RejectReason.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.RejectReason.AnyValue()))
	}
	if lastModel.LicenseImageID.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.LicenseImageID.AnyValue()))
	}
	if lastModel.CreateTime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.CreateTime.AnyValue()))
	}
	if lastModel.UpdateTime.IsValid() {
		builder.WriteString(fmt.Sprintf("%v", lastModel.UpdateTime.AnyValue()))
	}
	if idx, ok := l.dupMap[builder.String()]; ok {
		return idx
	}
	l.dupMap[builder.String()] = l.Len() - 1
	return -1
}

func (l *EmployEnterpriseList) Slice(low, high int) {
	switch {
	case high <= l.Len():
		l.List = l.List[low:high]
	case low <= l.Len() && high > l.Len():
		l.List = l.List[low:]
	default:
		l.List = l.List[:0]
	}
}

type EmployEnterpriseCacheElem struct {
	hashValue  string
	model      *EmployEnterprise
	modifyTime time.Time
}

type EmployEnterpriseListCacheElem struct {
	hashValue  string
	list       *EmployEnterpriseList
	modifyTime time.Time
}

type EmployEnterpriseCacheManager struct {
	container map[string]*EmployEnterpriseCacheElem
	query     chan string
	in        chan *EmployEnterpriseCacheElem
	out       chan *EmployEnterpriseCacheElem
}

type EmployEnterpriseListCacheManager struct {
	container map[string]*EmployEnterpriseListCacheElem
	query     chan string
	in        chan *EmployEnterpriseListCacheElem
	out       chan *EmployEnterpriseListCacheElem
}

func newEmployEnterpriseCacheManager() *EmployEnterpriseCacheManager {
	return &EmployEnterpriseCacheManager{
		make(map[string]*EmployEnterpriseCacheElem),
		make(chan string),
		make(chan *EmployEnterpriseCacheElem),
		make(chan *EmployEnterpriseCacheElem),
	}
}

func newEmployEnterpriseListCacheManager() *EmployEnterpriseListCacheManager {
	return &EmployEnterpriseListCacheManager{
		make(map[string]*EmployEnterpriseListCacheElem),
		make(chan string),
		make(chan *EmployEnterpriseListCacheElem),
		make(chan *EmployEnterpriseListCacheElem),
	}
}

func (mgr *EmployEnterpriseCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

func (mgr *EmployEnterpriseListCacheManager) run() {
	for {
		select {
		case h := <-mgr.query:
			mgr.out <- mgr.container[h]
		case elem := <-mgr.in:
			mgr.container[elem.hashValue] = elem
		}
	}
}

var EmployEnterpriseCache = newEmployEnterpriseCacheManager()

var EmployEnterpriseListCache = newEmployEnterpriseListCacheManager()

func (m *EmployEnterprise) GetCache(hashVal string, timeout time.Duration) bool {
	EmployEnterpriseCache.query <- hashVal
	elem := <-EmployEnterpriseCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*m = *elem.model
	return true
}

func (m *EmployEnterprise) SetCache(hashValue string) {
	EmployEnterpriseCache.in <- &EmployEnterpriseCacheElem{
		hashValue,
		m,
		time.Now(),
	}
}

func (l *EmployEnterpriseList) GetListCache(hashValue string, timeout time.Duration) bool {
	EmployEnterpriseListCache.query <- hashValue
	elem := <-EmployEnterpriseListCache.out
	if elem == nil || time.Since(elem.modifyTime) > timeout {
		return false
	}
	*l = *elem.list
	return true
}

func (l *EmployEnterpriseList) SetListCache(hashValue string) {
	EmployEnterpriseListCache.in <- &EmployEnterpriseListCacheElem{
		hashValue,
		l,
		time.Now(),
	}
}

func init() {
	go EmployAccountCache.run()
	go EmployAccountListCache.run()
	go EmployEnterpriseCache.run()
	go EmployEnterpriseListCache.run()
}
