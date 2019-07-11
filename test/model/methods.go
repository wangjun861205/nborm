package model

import (
	"github.com/wangjun861205/nborm"
)

func NewEnterpriseAccount() *EnterpriseAccount {
	m := &EnterpriseAccount{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseAccount) InitRel() {
	m.Enterprise = &EnterpriseList{}
	m.Enterprise.SetParent(m)
	nborm.InitModel(m.Enterprise)
	m.AddRelInited()
}
func (m *EnterpriseAccount) DB() string {
	return "qdxg"
}

func (m *EnterpriseAccount) Tab() string {
	return "enterprise_account"
}

func (m *EnterpriseAccount) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"Phone", "Phone", &m.Phone},
		{"Email", "Email", &m.Email},
		{"Password", "Password", &m.Password},
	}
}

func (m *EnterpriseAccount) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseAccount) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseAccount) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseAccount) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&m.Enterprise.AccountID,
			},
			m.Enterprise,
			"Enterprise",
		},
	}
}

func (m EnterpriseAccount) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseAccount) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseAccountList struct {
	EnterpriseAccount
	List  []*EnterpriseAccount
	Total int
}

func NewEnterpriseAccountList() *EnterpriseAccountList {
	l := &EnterpriseAccountList{
		EnterpriseAccount{},
		make([]*EnterpriseAccount, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseAccountList) NewModel() nborm.Model {
	m := &EnterpriseAccount{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseAccountList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseAccountList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseAccountList) Len() int {
	return len(l.List)
}

func (l *EnterpriseAccountList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseAccountList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseAccountList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseAccountList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.Enterprise.List = append(lm.Enterprise.List, rm.Enterprise.List...)
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *EnterpriseAccountList) Filter(f func(m *EnterpriseAccount) bool) []*EnterpriseAccount {
	ll := make([]*EnterpriseAccount, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseReviewStatus() *EnterpriseReviewStatus {
	m := &EnterpriseReviewStatus{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseReviewStatus) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseReviewStatus) DB() string {
	return "qdxg"
}

func (m *EnterpriseReviewStatus) Tab() string {
	return "enterprise_revice_status"
}

func (m *EnterpriseReviewStatus) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"ReviewStatus", "ReviewStatus", &m.ReviewStatus},
		{"Operator", "Operator", &m.Operator},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *EnterpriseReviewStatus) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseReviewStatus) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseReviewStatus) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseReviewStatus) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseReviewStatus) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseReviewStatus) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseReviewStatusList struct {
	EnterpriseReviewStatus
	List  []*EnterpriseReviewStatus
	Total int
}

func NewEnterpriseReviewStatusList() *EnterpriseReviewStatusList {
	l := &EnterpriseReviewStatusList{
		EnterpriseReviewStatus{},
		make([]*EnterpriseReviewStatus, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseReviewStatusList) NewModel() nborm.Model {
	m := &EnterpriseReviewStatus{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseReviewStatusList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseReviewStatusList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseReviewStatusList) Len() int {
	return len(l.List)
}

func (l *EnterpriseReviewStatusList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseReviewStatusList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseReviewStatusList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseReviewStatusList) Collapse() {
	return
}

func (l *EnterpriseReviewStatusList) Filter(f func(m *EnterpriseReviewStatus) bool) []*EnterpriseReviewStatus {
	ll := make([]*EnterpriseReviewStatus, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseStatistic() *EnterpriseStatistic {
	m := &EnterpriseStatistic{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseStatistic) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseStatistic) DB() string {
	return "qdxg"
}

func (m *EnterpriseStatistic) Tab() string {
	return "enterprise_statistic"
}

func (m *EnterpriseStatistic) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"SubmitCount", "SubmitCount", &m.SubmitCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *EnterpriseStatistic) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseStatistic) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseStatistic) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseStatistic) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseStatistic) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseStatistic) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseStatisticList struct {
	EnterpriseStatistic
	List  []*EnterpriseStatistic
	Total int
}

func NewEnterpriseStatisticList() *EnterpriseStatisticList {
	l := &EnterpriseStatisticList{
		EnterpriseStatistic{},
		make([]*EnterpriseStatistic, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseStatisticList) NewModel() nborm.Model {
	m := &EnterpriseStatistic{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseStatisticList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseStatisticList) Len() int {
	return len(l.List)
}

func (l *EnterpriseStatisticList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseStatisticList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseStatisticList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseStatisticList) Collapse() {
	return
}

func (l *EnterpriseStatisticList) Filter(f func(m *EnterpriseStatistic) bool) []*EnterpriseStatistic {
	ll := make([]*EnterpriseStatistic, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterprise() *Enterprise {
	m := &Enterprise{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *Enterprise) InitRel() {
	m.Account = &EnterpriseAccount{}
	m.Account.SetParent(m)
	nborm.InitModel(m.Account)
	m.Sector = &JobFlag{}
	m.Sector.SetParent(m)
	nborm.InitModel(m.Sector)
	m.AddRelInited()
}
func (m *Enterprise) DB() string {
	return "qdxg"
}

func (m *Enterprise) Tab() string {
	return "enterprise"
}

func (m *Enterprise) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"AccountID", "AccountID", &m.AccountID},
		{"UniformCode", "UniformCode", &m.UniformCode},
		{"Name", "Name", &m.Name},
		{"RegisterCityID", "RegisterCityID", &m.RegisterCityID},
		{"RegisterAddress", "RegisterAddress", &m.RegisterAddress},
		{"SectorID", "SectorID", &m.SectorID},
		{"NatureID", "NatureID", &m.NatureID},
		{"ScopeID", "ScopeID", &m.ScopeID},
		{"Website", "Website", &m.Website},
		{"Contact", "Contact", &m.Contact},
		{"EmployeeFromThis", "EmployeeFromThis", &m.EmployeeFromThis},
		{"Introduction", "Introduction", &m.Introduction},
		{"ZipCode", "ZipCode", &m.ZipCode},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *Enterprise) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *Enterprise) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *Enterprise) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *Enterprise) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.AccountID,
				&m.Account.ID,
			},
			m.Account,
			"Account",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.SectorID,
				&m.Sector.ID,
			},
			m.Sector,
			"Sector",
		},
	}
}

func (m Enterprise) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *Enterprise) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseList struct {
	Enterprise
	List  []*Enterprise
	Total int
}

func NewEnterpriseList() *EnterpriseList {
	l := &EnterpriseList{
		Enterprise{},
		make([]*Enterprise, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseList) NewModel() nborm.Model {
	m := &Enterprise{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseList) Len() int {
	return len(l.List)
}

func (l *EnterpriseList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.Account = rm.Account
		lm.Sector = rm.Sector
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *EnterpriseList) Filter(f func(m *Enterprise) bool) []*Enterprise {
	ll := make([]*Enterprise, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseAttachment() *EnterpriseAttachment {
	m := &EnterpriseAttachment{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseAttachment) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseAttachment) DB() string {
	return "qdxg"
}

func (m *EnterpriseAttachment) Tab() string {
	return "enterprise_attachment"
}

func (m *EnterpriseAttachment) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Type", "Type", &m.Type},
		{"URL", "URL", &m.URL},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *EnterpriseAttachment) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseAttachment) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseAttachment) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseAttachment) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseAttachment) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseAttachment) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseAttachmentList struct {
	EnterpriseAttachment
	List  []*EnterpriseAttachment
	Total int
}

func NewEnterpriseAttachmentList() *EnterpriseAttachmentList {
	l := &EnterpriseAttachmentList{
		EnterpriseAttachment{},
		make([]*EnterpriseAttachment, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseAttachmentList) NewModel() nborm.Model {
	m := &EnterpriseAttachment{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseAttachmentList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseAttachmentList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseAttachmentList) Len() int {
	return len(l.List)
}

func (l *EnterpriseAttachmentList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseAttachmentList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseAttachmentList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseAttachmentList) Collapse() {
	return
}

func (l *EnterpriseAttachmentList) Filter(f func(m *EnterpriseAttachment) bool) []*EnterpriseAttachment {
	ll := make([]*EnterpriseAttachment, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidEnterpriseTag() *MidEnterpriseTag {
	m := &MidEnterpriseTag{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidEnterpriseTag) InitRel() {
	m.AddRelInited()
}
func (m *MidEnterpriseTag) DB() string {
	return "qdxg"
}

func (m *MidEnterpriseTag) Tab() string {
	return "mid_enterprise__tag"
}

func (m *MidEnterpriseTag) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"TagID", "TagID", &m.TagID},
	}
}

func (m *MidEnterpriseTag) AutoIncField() nborm.Field {
	return nil
}

func (m *MidEnterpriseTag) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.EnterpriseID,
		&m.TagID,
	}
}

func (m *MidEnterpriseTag) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidEnterpriseTag) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidEnterpriseTag) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidEnterpriseTag) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidEnterpriseTagList struct {
	MidEnterpriseTag
	List  []*MidEnterpriseTag
	Total int
}

func NewMidEnterpriseTagList() *MidEnterpriseTagList {
	l := &MidEnterpriseTagList{
		MidEnterpriseTag{},
		make([]*MidEnterpriseTag, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidEnterpriseTagList) NewModel() nborm.Model {
	m := &MidEnterpriseTag{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidEnterpriseTagList) SetTotal(total int) {
	l.Total = total
}

func (l *MidEnterpriseTagList) GetTotal() int {
	return l.Total
}

func (l *MidEnterpriseTagList) Len() int {
	return len(l.List)
}

func (l *MidEnterpriseTagList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidEnterpriseTagList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidEnterpriseTagList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidEnterpriseTagList) Collapse() {
	return
}

func (l *MidEnterpriseTagList) Filter(f func(m *MidEnterpriseTag) bool) []*MidEnterpriseTag {
	ll := make([]*MidEnterpriseTag, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseJobStatistic() *EnterpriseJobStatistic {
	m := &EnterpriseJobStatistic{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseJobStatistic) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseJobStatistic) DB() string {
	return "qdxg"
}

func (m *EnterpriseJobStatistic) Tab() string {
	return "enterprise_job_static"
}

func (m *EnterpriseJobStatistic) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobID", "JobID", &m.JobID},
		{"SubmitCount", "SubmitCount", &m.SubmitCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *EnterpriseJobStatistic) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseJobStatistic) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJobStatistic) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseJobStatistic) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseJobStatistic) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseJobStatistic) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseJobStatisticList struct {
	EnterpriseJobStatistic
	List  []*EnterpriseJobStatistic
	Total int
}

func NewEnterpriseJobStatisticList() *EnterpriseJobStatisticList {
	l := &EnterpriseJobStatisticList{
		EnterpriseJobStatistic{},
		make([]*EnterpriseJobStatistic, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseJobStatisticList) NewModel() nborm.Model {
	m := &EnterpriseJobStatistic{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobStatisticList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseJobStatisticList) Len() int {
	return len(l.List)
}

func (l *EnterpriseJobStatisticList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseJobStatisticList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseJobStatisticList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseJobStatisticList) Collapse() {
	return
}

func (l *EnterpriseJobStatisticList) Filter(f func(m *EnterpriseJobStatistic) bool) []*EnterpriseJobStatistic {
	ll := make([]*EnterpriseJobStatistic, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeEnterpriseJob() *MidStudentResumeEnterpriseJob {
	m := &MidStudentResumeEnterpriseJob{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeEnterpriseJob) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeEnterpriseJob) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeEnterpriseJob) Tab() string {
	return "mid_student_resume__enterprise_job"
}

func (m *MidStudentResumeEnterpriseJob) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"JobID", "JobID", &m.JobID},
		{"ReviewStatus", "ReviewStatus", &m.ReviewStatus},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentResumeEnterpriseJob) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeEnterpriseJob) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.JobID,
	}
}

func (m *MidStudentResumeEnterpriseJob) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeEnterpriseJob) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeEnterpriseJob) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeEnterpriseJob) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeEnterpriseJobList struct {
	MidStudentResumeEnterpriseJob
	List  []*MidStudentResumeEnterpriseJob
	Total int
}

func NewMidStudentResumeEnterpriseJobList() *MidStudentResumeEnterpriseJobList {
	l := &MidStudentResumeEnterpriseJobList{
		MidStudentResumeEnterpriseJob{},
		make([]*MidStudentResumeEnterpriseJob, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeEnterpriseJobList) NewModel() nborm.Model {
	m := &MidStudentResumeEnterpriseJob{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeEnterpriseJobList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeEnterpriseJobList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeEnterpriseJobList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeEnterpriseJobList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeEnterpriseJobList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeEnterpriseJobList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeEnterpriseJobList) Collapse() {
	return
}

func (l *MidStudentResumeEnterpriseJobList) Filter(f func(m *MidStudentResumeEnterpriseJob) bool) []*MidStudentResumeEnterpriseJob {
	ll := make([]*MidStudentResumeEnterpriseJob, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseJob() *EnterpriseJob {
	m := &EnterpriseJob{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseJob) InitRel() {
	m.Enterprise = &EnterpriseList{}
	m.Enterprise.SetParent(m)
	nborm.InitModel(m.Enterprise)
	m.StudentResumes = &StudentResumeList{}
	m.StudentResumes.SetParent(m)
	nborm.InitModel(m.StudentResumes)
	var mm0 *MidStudentResumeEnterpriseJob
	mm0 = &MidStudentResumeEnterpriseJob{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	mm0.AndExprJoinWhere(nborm.NewExpr("@=1 OR @=2", &mm0.ReviewStatus, &mm0.ReviewStatus))
	m.AppendMidTab(mm0)
	m.AddRelInited()
}
func (m *EnterpriseJob) DB() string {
	return "qdxg"
}

func (m *EnterpriseJob) Tab() string {
	return "enterprise_job"
}

func (m *EnterpriseJob) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Name", "Name", &m.Name},
		{"CityID", "CityID", &m.CityID},
		{"Address", "Address", &m.Address},
		{"TypeID", "TypeID", &m.TypeID},
		{"Gender", "Gender", &m.Gender},
		{"MajorCode", "MajorCode", &m.MajorCode},
		{"DegreeID", "DegreeID", &m.DegreeID},
		{"LanguageSkillID", "LanguageSkillID", &m.LanguageSkillID},
		{"Description", "Description", &m.Description},
		{"SalaryRangeID", "SalaryRangeID", &m.SalaryRangeID},
		{"Welfare", "Welfare", &m.Welfare},
		{"Vacancies", "Vacancies", &m.Vacancies},
		{"ExpiredAt", "ExpiredAt", &m.ExpiredAt},
		{"Status", "Status", &m.Status},
		{"Comment", "Comment", &m.Comment},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *EnterpriseJob) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseJob) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJob) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseJob) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeEnterpriseJob)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.EnterpriseID,
				&m.Enterprise.ID,
			},
			m.Enterprise,
			"Enterprise",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.JobID,
				&mm0.ResumeID,
				&m.StudentResumes.ID,
			},
			m.StudentResumes,
			"StudentResumes",
		},
	}
}

func (m EnterpriseJob) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseJob) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseJobList struct {
	EnterpriseJob
	List  []*EnterpriseJob
	Total int
}

func NewEnterpriseJobList() *EnterpriseJobList {
	l := &EnterpriseJobList{
		EnterpriseJob{},
		make([]*EnterpriseJob, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseJobList) NewModel() nborm.Model {
	m := &EnterpriseJob{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseJobList) Len() int {
	return len(l.List)
}

func (l *EnterpriseJobList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseJobList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseJobList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseJobList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.Enterprise.List = append(lm.Enterprise.List, rm.Enterprise.List...)
		lm.StudentResumes.List = append(lm.StudentResumes.List, rm.StudentResumes.List...)
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *EnterpriseJobList) Filter(f func(m *EnterpriseJob) bool) []*EnterpriseJob {
	ll := make([]*EnterpriseJob, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentJobFairRead() *MidStudentJobFairRead {
	m := &MidStudentJobFairRead{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentJobFairRead) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentJobFairRead) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairRead) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairRead) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairRead) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentJobFairRead) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairRead) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentJobFairRead) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentJobFairRead) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentJobFairRead) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentJobFairReadList struct {
	MidStudentJobFairRead
	List  []*MidStudentJobFairRead
	Total int
}

func NewMidStudentJobFairReadList() *MidStudentJobFairReadList {
	l := &MidStudentJobFairReadList{
		MidStudentJobFairRead{},
		make([]*MidStudentJobFairRead, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentJobFairReadList) NewModel() nborm.Model {
	m := &MidStudentJobFairRead{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairReadList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairReadList) GetTotal() int {
	return l.Total
}

func (l *MidStudentJobFairReadList) Len() int {
	return len(l.List)
}

func (l *MidStudentJobFairReadList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentJobFairReadList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentJobFairReadList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentJobFairReadList) Collapse() {
	return
}

func (l *MidStudentJobFairReadList) Filter(f func(m *MidStudentJobFairRead) bool) []*MidStudentJobFairRead {
	ll := make([]*MidStudentJobFairRead, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentJobFairEnroll() *MidStudentJobFairEnroll {
	m := &MidStudentJobFairEnroll{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentJobFairEnroll) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentJobFairEnroll) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairEnroll) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairEnroll) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairEnroll) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentJobFairEnroll) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairEnroll) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentJobFairEnroll) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentJobFairEnroll) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentJobFairEnroll) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentJobFairEnrollList struct {
	MidStudentJobFairEnroll
	List  []*MidStudentJobFairEnroll
	Total int
}

func NewMidStudentJobFairEnrollList() *MidStudentJobFairEnrollList {
	l := &MidStudentJobFairEnrollList{
		MidStudentJobFairEnroll{},
		make([]*MidStudentJobFairEnroll, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentJobFairEnrollList) NewModel() nborm.Model {
	m := &MidStudentJobFairEnroll{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairEnrollList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairEnrollList) GetTotal() int {
	return l.Total
}

func (l *MidStudentJobFairEnrollList) Len() int {
	return len(l.List)
}

func (l *MidStudentJobFairEnrollList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentJobFairEnrollList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentJobFairEnrollList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentJobFairEnrollList) Collapse() {
	return
}

func (l *MidStudentJobFairEnrollList) Filter(f func(m *MidStudentJobFairEnroll) bool) []*MidStudentJobFairEnroll {
	ll := make([]*MidStudentJobFairEnroll, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentJobFairShare() *MidStudentJobFairShare {
	m := &MidStudentJobFairShare{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentJobFairShare) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentJobFairShare) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairShare) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairShare) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairShare) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentJobFairShare) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairShare) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentJobFairShare) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentJobFairShare) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentJobFairShare) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentJobFairShareList struct {
	MidStudentJobFairShare
	List  []*MidStudentJobFairShare
	Total int
}

func NewMidStudentJobFairShareList() *MidStudentJobFairShareList {
	l := &MidStudentJobFairShareList{
		MidStudentJobFairShare{},
		make([]*MidStudentJobFairShare, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentJobFairShareList) NewModel() nborm.Model {
	m := &MidStudentJobFairShare{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairShareList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairShareList) GetTotal() int {
	return l.Total
}

func (l *MidStudentJobFairShareList) Len() int {
	return len(l.List)
}

func (l *MidStudentJobFairShareList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentJobFairShareList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentJobFairShareList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentJobFairShareList) Collapse() {
	return
}

func (l *MidStudentJobFairShareList) Filter(f func(m *MidStudentJobFairShare) bool) []*MidStudentJobFairShare {
	ll := make([]*MidStudentJobFairShare, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewJobFairStatistic() *JobFairStatistic {
	m := &JobFairStatistic{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *JobFairStatistic) InitRel() {
	m.AddRelInited()
}
func (m *JobFairStatistic) DB() string {
	return "qdxg"
}

func (m *JobFairStatistic) Tab() string {
	return "job_fair_statistic"
}

func (m *JobFairStatistic) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"ReadCount", "ReadCount", &m.ReadCount},
		{"EnrollCount", "EnrollCount", &m.EnrollCount},
		{"ShareCount", "ShareCount", &m.ShareCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *JobFairStatistic) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *JobFairStatistic) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *JobFairStatistic) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *JobFairStatistic) Relations() nborm.RelationInfoList {
	return nil
}

func (m JobFairStatistic) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *JobFairStatistic) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type JobFairStatisticList struct {
	JobFairStatistic
	List  []*JobFairStatistic
	Total int
}

func NewJobFairStatisticList() *JobFairStatisticList {
	l := &JobFairStatisticList{
		JobFairStatistic{},
		make([]*JobFairStatistic, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *JobFairStatisticList) NewModel() nborm.Model {
	m := &JobFairStatistic{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *JobFairStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFairStatisticList) GetTotal() int {
	return l.Total
}

func (l *JobFairStatisticList) Len() int {
	return len(l.List)
}

func (l *JobFairStatisticList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l JobFairStatisticList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *JobFairStatisticList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *JobFairStatisticList) Collapse() {
	return
}

func (l *JobFairStatisticList) Filter(f func(m *JobFairStatistic) bool) []*JobFairStatistic {
	ll := make([]*JobFairStatistic, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewJobFair() *JobFair {
	m := &JobFair{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *JobFair) InitRel() {
	m.AddRelInited()
}
func (m *JobFair) DB() string {
	return "qdxg"
}

func (m *JobFair) Tab() string {
	return "job_fair"
}

func (m *JobFair) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"Name", "Name", &m.Name},
		{"StartTime", "StartTime", &m.StartTime},
		{"EndTime", "EndTime", &m.EndTime},
		{"Description", "Description", &m.Description},
		{"Status", "Status", &m.Status},
		{"Comment", "Comment", &m.Comment},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *JobFair) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *JobFair) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *JobFair) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *JobFair) Relations() nborm.RelationInfoList {
	return nil
}

func (m JobFair) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *JobFair) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type JobFairList struct {
	JobFair
	List  []*JobFair
	Total int
}

func NewJobFairList() *JobFairList {
	l := &JobFairList{
		JobFair{},
		make([]*JobFair, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *JobFairList) NewModel() nborm.Model {
	m := &JobFair{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *JobFairList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFairList) GetTotal() int {
	return l.Total
}

func (l *JobFairList) Len() int {
	return len(l.List)
}

func (l *JobFairList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l JobFairList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *JobFairList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *JobFairList) Collapse() {
	return
}

func (l *JobFairList) Filter(f func(m *JobFair) bool) []*JobFair {
	ll := make([]*JobFair, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewJobFlag() *JobFlag {
	m := &JobFlag{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *JobFlag) InitRel() {
	m.AddRelInited()
}
func (m *JobFlag) DB() string {
	return "qdxg"
}

func (m *JobFlag) Tab() string {
	return "job_flag"
}

func (m *JobFlag) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"Name", "Name", &m.Name},
		{"Type", "Type", &m.Type},
		{"Value", "Value", &m.Value},
		{"Order", "Order", &m.Order},
		{"ParentID", "ParentID", &m.ParentID},
		{"Status", "Status", &m.Status},
		{"Operator", "Operator", &m.Operator},
		{"Comment", "Comment", &m.Comment},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *JobFlag) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *JobFlag) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *JobFlag) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *JobFlag) Relations() nborm.RelationInfoList {
	return nil
}

func (m JobFlag) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *JobFlag) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type JobFlagList struct {
	JobFlag
	List  []*JobFlag
	Total int
}

func NewJobFlagList() *JobFlagList {
	l := &JobFlagList{
		JobFlag{},
		make([]*JobFlag, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *JobFlagList) NewModel() nborm.Model {
	m := &JobFlag{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *JobFlagList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFlagList) GetTotal() int {
	return l.Total
}

func (l *JobFlagList) Len() int {
	return len(l.List)
}

func (l *JobFlagList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l JobFlagList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *JobFlagList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *JobFlagList) Collapse() {
	return
}

func (l *JobFlagList) Filter(f func(m *JobFlag) bool) []*JobFlag {
	ll := make([]*JobFlag, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeLanguageSkill() *MidStudentResumeLanguageSkill {
	m := &MidStudentResumeLanguageSkill{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeLanguageSkill) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeLanguageSkill) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeLanguageSkill) Tab() string {
	return "mid_student_resume__language_skill"
}

func (m *MidStudentResumeLanguageSkill) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"LanguageSkillID", "LanguageSkillID", &m.LanguageSkillID},
	}
}

func (m *MidStudentResumeLanguageSkill) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeLanguageSkill) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.LanguageSkillID,
	}
}

func (m *MidStudentResumeLanguageSkill) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeLanguageSkill) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeLanguageSkill) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeLanguageSkill) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeLanguageSkillList struct {
	MidStudentResumeLanguageSkill
	List  []*MidStudentResumeLanguageSkill
	Total int
}

func NewMidStudentResumeLanguageSkillList() *MidStudentResumeLanguageSkillList {
	l := &MidStudentResumeLanguageSkillList{
		MidStudentResumeLanguageSkill{},
		make([]*MidStudentResumeLanguageSkill, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeLanguageSkillList) NewModel() nborm.Model {
	m := &MidStudentResumeLanguageSkill{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeLanguageSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeLanguageSkillList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeLanguageSkillList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeLanguageSkillList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeLanguageSkillList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeLanguageSkillList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeLanguageSkillList) Collapse() {
	return
}

func (l *MidStudentResumeLanguageSkillList) Filter(f func(m *MidStudentResumeLanguageSkill) bool) []*MidStudentResumeLanguageSkill {
	ll := make([]*MidStudentResumeLanguageSkill, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeStudentTrain() *MidStudentResumeStudentTrain {
	m := &MidStudentResumeStudentTrain{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeStudentTrain) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeStudentTrain) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentTrain) Tab() string {
	return "mid_student_resume__student_train"
}

func (m *MidStudentResumeStudentTrain) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"TrainID", "TrainID", &m.TrainID},
	}
}

func (m *MidStudentResumeStudentTrain) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeStudentTrain) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.TrainID,
	}
}

func (m *MidStudentResumeStudentTrain) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeStudentTrain) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeStudentTrain) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeStudentTrain) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeStudentTrainList struct {
	MidStudentResumeStudentTrain
	List  []*MidStudentResumeStudentTrain
	Total int
}

func NewMidStudentResumeStudentTrainList() *MidStudentResumeStudentTrainList {
	l := &MidStudentResumeStudentTrainList{
		MidStudentResumeStudentTrain{},
		make([]*MidStudentResumeStudentTrain, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeStudentTrainList) NewModel() nborm.Model {
	m := &MidStudentResumeStudentTrain{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentTrainList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentTrainList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeStudentTrainList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeStudentTrainList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeStudentTrainList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeStudentTrainList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeStudentTrainList) Collapse() {
	return
}

func (l *MidStudentResumeStudentTrainList) Filter(f func(m *MidStudentResumeStudentTrain) bool) []*MidStudentResumeStudentTrain {
	ll := make([]*MidStudentResumeStudentTrain, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeStudentHonor() *MidStudentResumeStudentHonor {
	m := &MidStudentResumeStudentHonor{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeStudentHonor) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeStudentHonor) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentHonor) Tab() string {
	return "mid_student_resume__student_honor"
}

func (m *MidStudentResumeStudentHonor) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"HonorID", "HonorID", &m.HonorID},
	}
}

func (m *MidStudentResumeStudentHonor) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeStudentHonor) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.HonorID,
	}
}

func (m *MidStudentResumeStudentHonor) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeStudentHonor) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeStudentHonor) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeStudentHonor) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeStudentHonorList struct {
	MidStudentResumeStudentHonor
	List  []*MidStudentResumeStudentHonor
	Total int
}

func NewMidStudentResumeStudentHonorList() *MidStudentResumeStudentHonorList {
	l := &MidStudentResumeStudentHonorList{
		MidStudentResumeStudentHonor{},
		make([]*MidStudentResumeStudentHonor, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeStudentHonorList) NewModel() nborm.Model {
	m := &MidStudentResumeStudentHonor{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentHonorList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentHonorList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeStudentHonorList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeStudentHonorList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeStudentHonorList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeStudentHonorList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeStudentHonorList) Collapse() {
	return
}

func (l *MidStudentResumeStudentHonorList) Filter(f func(m *MidStudentResumeStudentHonor) bool) []*MidStudentResumeStudentHonor {
	ll := make([]*MidStudentResumeStudentHonor, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeStudentExperience() *MidStudentResumeStudentExperience {
	m := &MidStudentResumeStudentExperience{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeStudentExperience) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeStudentExperience) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentExperience) Tab() string {
	return "mid_student_resume__student_experience"
}

func (m *MidStudentResumeStudentExperience) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"ExperienceID", "ExperienceID", &m.ExperienceID},
	}
}

func (m *MidStudentResumeStudentExperience) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeStudentExperience) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.ExperienceID,
	}
}

func (m *MidStudentResumeStudentExperience) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeStudentExperience) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeStudentExperience) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeStudentExperience) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeStudentExperienceList struct {
	MidStudentResumeStudentExperience
	List  []*MidStudentResumeStudentExperience
	Total int
}

func NewMidStudentResumeStudentExperienceList() *MidStudentResumeStudentExperienceList {
	l := &MidStudentResumeStudentExperienceList{
		MidStudentResumeStudentExperience{},
		make([]*MidStudentResumeStudentExperience, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeStudentExperienceList) NewModel() nborm.Model {
	m := &MidStudentResumeStudentExperience{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentExperienceList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentExperienceList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeStudentExperienceList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeStudentExperienceList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeStudentExperienceList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeStudentExperienceList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeStudentExperienceList) Collapse() {
	return
}

func (l *MidStudentResumeStudentExperienceList) Filter(f func(m *MidStudentResumeStudentExperience) bool) []*MidStudentResumeStudentExperience {
	ll := make([]*MidStudentResumeStudentExperience, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewMidStudentResumeStudentSkill() *MidStudentResumeStudentSkill {
	m := &MidStudentResumeStudentSkill{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *MidStudentResumeStudentSkill) InitRel() {
	m.AddRelInited()
}
func (m *MidStudentResumeStudentSkill) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentSkill) Tab() string {
	return "mid_student_resume__student_skill"
}

func (m *MidStudentResumeStudentSkill) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"SkillID", "SkillID", &m.SkillID},
	}
}

func (m *MidStudentResumeStudentSkill) AutoIncField() nborm.Field {
	return nil
}

func (m *MidStudentResumeStudentSkill) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ResumeID,
		&m.SkillID,
	}
}

func (m *MidStudentResumeStudentSkill) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *MidStudentResumeStudentSkill) Relations() nborm.RelationInfoList {
	return nil
}

func (m MidStudentResumeStudentSkill) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *MidStudentResumeStudentSkill) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type MidStudentResumeStudentSkillList struct {
	MidStudentResumeStudentSkill
	List  []*MidStudentResumeStudentSkill
	Total int
}

func NewMidStudentResumeStudentSkillList() *MidStudentResumeStudentSkillList {
	l := &MidStudentResumeStudentSkillList{
		MidStudentResumeStudentSkill{},
		make([]*MidStudentResumeStudentSkill, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *MidStudentResumeStudentSkillList) NewModel() nborm.Model {
	m := &MidStudentResumeStudentSkill{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentSkillList) GetTotal() int {
	return l.Total
}

func (l *MidStudentResumeStudentSkillList) Len() int {
	return len(l.List)
}

func (l *MidStudentResumeStudentSkillList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l MidStudentResumeStudentSkillList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *MidStudentResumeStudentSkillList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *MidStudentResumeStudentSkillList) Collapse() {
	return
}

func (l *MidStudentResumeStudentSkillList) Filter(f func(m *MidStudentResumeStudentSkill) bool) []*MidStudentResumeStudentSkill {
	ll := make([]*MidStudentResumeStudentSkill, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentTrain() *StudentTrain {
	m := &StudentTrain{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentTrain) InitRel() {
	m.StudentResume = &StudentResume{}
	m.StudentResume.SetParent(m)
	nborm.InitModel(m.StudentResume)
	var mm0 *MidStudentResumeStudentTrain
	mm0 = &MidStudentResumeStudentTrain{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	m.AppendMidTab(mm0)
	m.AddRelInited()
}
func (m *StudentTrain) DB() string {
	return "qdxg"
}

func (m *StudentTrain) Tab() string {
	return "student_train"
}

func (m *StudentTrain) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"InstitutionName", "InstitutionName", &m.InstitutionName},
		{"StartDate", "StartDate", &m.StartDate},
		{"EndDate", "EndDate", &m.EndDate},
		{"Degree", "Degree", &m.Degree},
		{"Major", "Major", &m.Major},
		{"Description", "Description", &m.Description},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentTrain) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentTrain) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentTrain) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentTrain) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeStudentTrain)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.TrainID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
			"StudentResume",
		},
	}
}

func (m StudentTrain) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentTrain) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentTrainList struct {
	StudentTrain
	List  []*StudentTrain
	Total int
}

func NewStudentTrainList() *StudentTrainList {
	l := &StudentTrainList{
		StudentTrain{},
		make([]*StudentTrain, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentTrainList) NewModel() nborm.Model {
	m := &StudentTrain{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentTrainList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentTrainList) GetTotal() int {
	return l.Total
}

func (l *StudentTrainList) Len() int {
	return len(l.List)
}

func (l *StudentTrainList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentTrainList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentTrainList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentTrainList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.StudentResume = rm.StudentResume
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *StudentTrainList) Filter(f func(m *StudentTrain) bool) []*StudentTrain {
	ll := make([]*StudentTrain, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentHonor() *StudentHonor {
	m := &StudentHonor{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentHonor) InitRel() {
	m.StudentResume = &StudentResume{}
	m.StudentResume.SetParent(m)
	nborm.InitModel(m.StudentResume)
	var mm0 *MidStudentResumeStudentHonor
	mm0 = &MidStudentResumeStudentHonor{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	m.AppendMidTab(mm0)
	m.AddRelInited()
}
func (m *StudentHonor) DB() string {
	return "qdxg"
}

func (m *StudentHonor) Tab() string {
	return "student_honor"
}

func (m *StudentHonor) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"Name", "Name", &m.Name},
		{"Description", "Description", &m.Description},
		{"GrantDate", "GrantDate", &m.GrantDate},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentHonor) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentHonor) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentHonor) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentHonor) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeStudentHonor)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.HonorID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
			"StudentResume",
		},
	}
}

func (m StudentHonor) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentHonor) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentHonorList struct {
	StudentHonor
	List  []*StudentHonor
	Total int
}

func NewStudentHonorList() *StudentHonorList {
	l := &StudentHonorList{
		StudentHonor{},
		make([]*StudentHonor, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentHonorList) NewModel() nborm.Model {
	m := &StudentHonor{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentHonorList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentHonorList) GetTotal() int {
	return l.Total
}

func (l *StudentHonorList) Len() int {
	return len(l.List)
}

func (l *StudentHonorList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentHonorList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentHonorList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentHonorList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.StudentResume = rm.StudentResume
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *StudentHonorList) Filter(f func(m *StudentHonor) bool) []*StudentHonor {
	ll := make([]*StudentHonor, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentExperience() *StudentExperience {
	m := &StudentExperience{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentExperience) InitRel() {
	m.StudentResume = &StudentResume{}
	m.StudentResume.SetParent(m)
	nborm.InitModel(m.StudentResume)
	var mm0 *MidStudentResumeStudentExperience
	mm0 = &MidStudentResumeStudentExperience{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	m.AppendMidTab(mm0)
	m.AddRelInited()
}
func (m *StudentExperience) DB() string {
	return "qdxg"
}

func (m *StudentExperience) Tab() string {
	return "student_experience"
}

func (m *StudentExperience) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"CompanyName", "CompanyName", &m.CompanyName},
		{"StartDate", "StartDate", &m.StartDate},
		{"EndDate", "EndDate", &m.EndDate},
		{"SectorID", "SectorID", &m.SectorID},
		{"Position", "Position", &m.Position},
		{"Description", "Description", &m.Description},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentExperience) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentExperience) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentExperience) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentExperience) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeStudentExperience)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.ExperienceID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
			"StudentResume",
		},
	}
}

func (m StudentExperience) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentExperience) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentExperienceList struct {
	StudentExperience
	List  []*StudentExperience
	Total int
}

func NewStudentExperienceList() *StudentExperienceList {
	l := &StudentExperienceList{
		StudentExperience{},
		make([]*StudentExperience, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentExperienceList) NewModel() nborm.Model {
	m := &StudentExperience{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentExperienceList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentExperienceList) GetTotal() int {
	return l.Total
}

func (l *StudentExperienceList) Len() int {
	return len(l.List)
}

func (l *StudentExperienceList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentExperienceList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentExperienceList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentExperienceList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.StudentResume = rm.StudentResume
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *StudentExperienceList) Filter(f func(m *StudentExperience) bool) []*StudentExperience {
	ll := make([]*StudentExperience, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentSkill() *StudentSkill {
	m := &StudentSkill{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentSkill) InitRel() {
	m.StudentResume = &StudentResume{}
	m.StudentResume.SetParent(m)
	nborm.InitModel(m.StudentResume)
	var mm0 *MidStudentResumeStudentSkill
	mm0 = &MidStudentResumeStudentSkill{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	m.AppendMidTab(mm0)
	m.AddRelInited()
}
func (m *StudentSkill) DB() string {
	return "qdxg"
}

func (m *StudentSkill) Tab() string {
	return "student_skill"
}

func (m *StudentSkill) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"Name", "Name", &m.Name},
		{"Description", "Description", &m.Description},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentSkill) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentSkill) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentSkill) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentSkill) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeStudentSkill)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.SkillID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
			"StudentResume",
		},
	}
}

func (m StudentSkill) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentSkill) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentSkillList struct {
	StudentSkill
	List  []*StudentSkill
	Total int
}

func NewStudentSkillList() *StudentSkillList {
	l := &StudentSkillList{
		StudentSkill{},
		make([]*StudentSkill, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentSkillList) NewModel() nborm.Model {
	m := &StudentSkill{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentSkillList) GetTotal() int {
	return l.Total
}

func (l *StudentSkillList) Len() int {
	return len(l.List)
}

func (l *StudentSkillList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentSkillList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentSkillList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentSkillList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.StudentResume = rm.StudentResume
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *StudentSkillList) Filter(f func(m *StudentSkill) bool) []*StudentSkill {
	ll := make([]*StudentSkill, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentResume() *StudentResume {
	m := &StudentResume{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentResume) InitRel() {
	m.StudentTrain = &StudentTrainList{}
	m.StudentTrain.SetParent(m)
	nborm.InitModel(m.StudentTrain)
	m.StudentHonor = &StudentHonorList{}
	m.StudentHonor.SetParent(m)
	nborm.InitModel(m.StudentHonor)
	m.StudentExperience = &StudentExperienceList{}
	m.StudentExperience.SetParent(m)
	nborm.InitModel(m.StudentExperience)
	m.StudentSkill = &StudentSkillList{}
	m.StudentSkill.SetParent(m)
	nborm.InitModel(m.StudentSkill)
	m.StudentLanguageType = &JobFlagList{}
	m.StudentLanguageType.SetParent(m)
	nborm.InitModel(m.StudentLanguageType)
	var mm0 *MidStudentResumeStudentTrain
	var mm1 *MidStudentResumeStudentHonor
	var mm2 *MidStudentResumeStudentExperience
	var mm3 *MidStudentResumeStudentSkill
	mm0 = &MidStudentResumeStudentTrain{}
	mm0.SetParent(m)
	nborm.InitModel(mm0)
	m.AppendMidTab(mm0)
	mm1 = &MidStudentResumeStudentHonor{}
	mm1.SetParent(m)
	nborm.InitModel(mm1)
	m.AppendMidTab(mm1)
	mm2 = &MidStudentResumeStudentExperience{}
	mm2.SetParent(m)
	nborm.InitModel(mm2)
	m.AppendMidTab(mm2)
	mm3 = &MidStudentResumeStudentSkill{}
	mm3.SetParent(m)
	nborm.InitModel(mm3)
	m.AppendMidTab(mm3)
	m.AddRelInited()
}
func (m *StudentResume) DB() string {
	return "qdxg"
}

func (m *StudentResume) Tab() string {
	return "student_resume"
}

func (m *StudentResume) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"Introduction", "Introduction", &m.Introduction},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentResume) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentResume) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentResume) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentResume) Relations() nborm.RelationInfoList {
	if !m.IsRelInited() {
		m.InitRel()
	}
	mm0 := m.GetMidTabs()[0].(*MidStudentResumeStudentTrain)
	mm1 := m.GetMidTabs()[1].(*MidStudentResumeStudentHonor)
	mm2 := m.GetMidTabs()[2].(*MidStudentResumeStudentExperience)
	mm3 := m.GetMidTabs()[3].(*MidStudentResumeStudentSkill)
	return nborm.RelationInfoList{
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm0.ResumeID,
				&mm0.TrainID,
				&m.StudentTrain.ID,
			},
			m.StudentTrain,
			"StudentTrain",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm1.ResumeID,
				&mm1.HonorID,
				&m.StudentHonor.ID,
			},
			m.StudentHonor,
			"StudentHonor",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm2.ResumeID,
				&mm2.ExperienceID,
				&m.StudentExperience.ID,
			},
			m.StudentExperience,
			"StudentExperience",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&mm3.ResumeID,
				&mm3.SkillID,
				&m.StudentSkill.ID,
			},
			m.StudentSkill,
			"StudentSkill",
		},
		nborm.RelationInfo{
			nborm.FieldList{
				&m.ID,
				&m.StudentLanguageType.ID,
			},
			m.StudentLanguageType,
			"StudentLanguageType",
		},
	}
}

func (m StudentResume) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentResume) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentResumeList struct {
	StudentResume
	List  []*StudentResume
	Total int
}

func NewStudentResumeList() *StudentResumeList {
	l := &StudentResumeList{
		StudentResume{},
		make([]*StudentResume, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentResumeList) NewModel() nborm.Model {
	m := &StudentResume{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentResumeList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentResumeList) GetTotal() int {
	return l.Total
}

func (l *StudentResumeList) Len() int {
	return len(l.List)
}

func (l *StudentResumeList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentResumeList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentResumeList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentResumeList) Collapse() {
	if len(l.List) < 2 {
		return
	}
	lm := l.List[len(l.List)-2]
	rm := l.List[len(l.List)-1]
	if nborm.IsPrimaryKeyEqual(lm, rm) {
		lm.StudentTrain.List = append(lm.StudentTrain.List, rm.StudentTrain.List...)
		lm.StudentHonor.List = append(lm.StudentHonor.List, rm.StudentHonor.List...)
		lm.StudentExperience.List = append(lm.StudentExperience.List, rm.StudentExperience.List...)
		lm.StudentSkill.List = append(lm.StudentSkill.List, rm.StudentSkill.List...)
		lm.StudentLanguageType.List = append(lm.StudentLanguageType.List, rm.StudentLanguageType.List...)
		l.List = l.List[:len(l.List)-1]
	}
}

func (l *StudentResumeList) Filter(f func(m *StudentResume) bool) []*StudentResume {
	ll := make([]*StudentResume, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseSnapshot() *EnterpriseSnapshot {
	m := &EnterpriseSnapshot{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseSnapshot) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseSnapshot) DB() string {
	return "qdxg"
}

func (m *EnterpriseSnapshot) Tab() string {
	return "enterprise_snapshot"
}

func (m *EnterpriseSnapshot) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Info", "Info", &m.Info},
	}
}

func (m *EnterpriseSnapshot) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseSnapshot) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseSnapshot) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseSnapshot) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseSnapshot) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseSnapshot) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseSnapshotList struct {
	EnterpriseSnapshot
	List  []*EnterpriseSnapshot
	Total int
}

func NewEnterpriseSnapshotList() *EnterpriseSnapshotList {
	l := &EnterpriseSnapshotList{
		EnterpriseSnapshot{},
		make([]*EnterpriseSnapshot, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseSnapshotList) NewModel() nborm.Model {
	m := &EnterpriseSnapshot{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseSnapshotList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseSnapshotList) Len() int {
	return len(l.List)
}

func (l *EnterpriseSnapshotList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseSnapshotList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseSnapshotList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseSnapshotList) Collapse() {
	return
}

func (l *EnterpriseSnapshotList) Filter(f func(m *EnterpriseSnapshot) bool) []*EnterpriseSnapshot {
	ll := make([]*EnterpriseSnapshot, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewEnterpriseJobSnapshot() *EnterpriseJobSnapshot {
	m := &EnterpriseJobSnapshot{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *EnterpriseJobSnapshot) InitRel() {
	m.AddRelInited()
}
func (m *EnterpriseJobSnapshot) DB() string {
	return "qdxg"
}

func (m *EnterpriseJobSnapshot) Tab() string {
	return "enterprise_job_snapshot"
}

func (m *EnterpriseJobSnapshot) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobID", "JobID", &m.JobID},
		{"Into", "Into", &m.Into},
	}
}

func (m *EnterpriseJobSnapshot) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *EnterpriseJobSnapshot) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJobSnapshot) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *EnterpriseJobSnapshot) Relations() nborm.RelationInfoList {
	return nil
}

func (m EnterpriseJobSnapshot) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *EnterpriseJobSnapshot) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type EnterpriseJobSnapshotList struct {
	EnterpriseJobSnapshot
	List  []*EnterpriseJobSnapshot
	Total int
}

func NewEnterpriseJobSnapshotList() *EnterpriseJobSnapshotList {
	l := &EnterpriseJobSnapshotList{
		EnterpriseJobSnapshot{},
		make([]*EnterpriseJobSnapshot, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *EnterpriseJobSnapshotList) NewModel() nborm.Model {
	m := &EnterpriseJobSnapshot{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobSnapshotList) GetTotal() int {
	return l.Total
}

func (l *EnterpriseJobSnapshotList) Len() int {
	return len(l.List)
}

func (l *EnterpriseJobSnapshotList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l EnterpriseJobSnapshotList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *EnterpriseJobSnapshotList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *EnterpriseJobSnapshotList) Collapse() {
	return
}

func (l *EnterpriseJobSnapshotList) Filter(f func(m *EnterpriseJobSnapshot) bool) []*EnterpriseJobSnapshot {
	ll := make([]*EnterpriseJobSnapshot, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}

func NewStudentResumeSnapshot() *StudentResumeSnapshot {
	m := &StudentResumeSnapshot{}
	nborm.InitModel(m)
	m.InitRel()
	return m
}

func (m *StudentResumeSnapshot) InitRel() {
	m.AddRelInited()
}
func (m *StudentResumeSnapshot) DB() string {
	return "qdxg"
}

func (m *StudentResumeSnapshot) Tab() string {
	return "student_resume_snapshot"
}

func (m *StudentResumeSnapshot) FieldInfos() nborm.FieldInfoList {
	return nborm.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"Info", "Info", &m.Info},
	}
}

func (m *StudentResumeSnapshot) AutoIncField() nborm.Field {
	return &m.ID
}

func (m *StudentResumeSnapshot) PrimaryKey() nborm.FieldList {
	return nborm.FieldList{
		&m.ID,
	}
}

func (m *StudentResumeSnapshot) UniqueKeys() []nborm.FieldList {
	return nil
}
func (m *StudentResumeSnapshot) Relations() nborm.RelationInfoList {
	return nil
}

func (m StudentResumeSnapshot) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&m), nil
}

func (m *StudentResumeSnapshot) UnmarshalJSON(data []byte) error {
	return nborm.UnmarshalModel(data, m)
}

type StudentResumeSnapshotList struct {
	StudentResumeSnapshot
	List  []*StudentResumeSnapshot
	Total int
}

func NewStudentResumeSnapshotList() *StudentResumeSnapshotList {
	l := &StudentResumeSnapshotList{
		StudentResumeSnapshot{},
		make([]*StudentResumeSnapshot, 0, 32),
		0,
	}
	nborm.InitModel(l)
	l.InitRel()
	return l
}

func (l *StudentResumeSnapshotList) NewModel() nborm.Model {
	m := &StudentResumeSnapshot{}
	m.SetParent(l.GetParent())
	nborm.InitModel(m)
	m.InitRel()
	l.List = append(l.List, m)
	return m
}

func (l *StudentResumeSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentResumeSnapshotList) GetTotal() int {
	return l.Total
}

func (l *StudentResumeSnapshotList) Len() int {
	return len(l.List)
}

func (l *StudentResumeSnapshotList) GetList() []nborm.Model {
	modelList := make([]nborm.Model, 0, l.Len())
	for _, m := range l.List {
		modelList = append(modelList, m)
	}
	return modelList
}

func (l StudentResumeSnapshotList) MarshalJSON() ([]byte, error) {
	return nborm.MarshalModel(&l), nil
}

func (l *StudentResumeSnapshotList) UnmarshalJSON(b []byte) error {
	return nborm.UnmarshalModel(b, l)
}

func (l *StudentResumeSnapshotList) Collapse() {
	return
}

func (l *StudentResumeSnapshotList) Filter(f func(m *StudentResumeSnapshot) bool) []*StudentResumeSnapshot {
	ll := make([]*StudentResumeSnapshot, 0, l.Len())
	for _, m := range l.List {
		if f(m) {
			ll = append(ll, m)
		}
	}
	return ll
}
