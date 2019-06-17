package job

import (
	"AntLinkCampus/CampusServer/handles/job/model"
)

func NewEnterpriseAccount() *EnterpriseAccount {
	m := &EnterpriseAccount{
		Enterprise: &EnterpriseList{},
	}
	model.InitModel(m)
	return m
}

func (m *EnterpriseAccount) DB() string {
	return "qdxg"
}

func (m *EnterpriseAccount) Tab() string {
	return "enterprise_account"
}

func (m *EnterpriseAccount) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Phone", "Phone", &m.Phone},
		{"Email", "Email", &m.Email},
		{"Password", "Password", &m.Password},
	}
}

func (m *EnterpriseAccount) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseAccount) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseAccount) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseAccount) Relations() model.RelationInfoList {
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.EnterpriseID,
				&m.Enterprise.ID,
			},
			m.Enterprise,
		},
	}

}

type EnterpriseAccountList struct {
	*EnterpriseAccount
	List  []*EnterpriseAccount
	Total int
}

func NewEnterpriseAccountList(cap int) *EnterpriseAccountList {
	return &EnterpriseAccountList{
		NewEnterpriseAccount(),
		make([]*EnterpriseAccount, 0, cap),
		0,
	}
}

func (l *EnterpriseAccountList) NewModel() model.Model {
	m := NewEnterpriseAccount()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseAccountList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseAccountList) Len() int {
	return len(l.List)
}

func NewEnterpriseReviewStatus() *EnterpriseReviewStatus {
	m := &EnterpriseReviewStatus{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseReviewStatus) DB() string {
	return "qdxg"
}

func (m *EnterpriseReviewStatus) Tab() string {
	return "enterprise_revice_status"
}

func (m *EnterpriseReviewStatus) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"ReviewStatus", "ReviewStatus", &m.ReviewStatus},
		{"Operator", "Operator", &m.Operator},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *EnterpriseReviewStatus) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseReviewStatus) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseReviewStatus) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseReviewStatus) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseReviewStatusList struct {
	*EnterpriseReviewStatus
	List  []*EnterpriseReviewStatus
	Total int
}

func NewEnterpriseReviewStatusList(cap int) *EnterpriseReviewStatusList {
	return &EnterpriseReviewStatusList{
		NewEnterpriseReviewStatus(),
		make([]*EnterpriseReviewStatus, 0, cap),
		0,
	}
}

func (l *EnterpriseReviewStatusList) NewModel() model.Model {
	m := NewEnterpriseReviewStatus()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseReviewStatusList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseReviewStatusList) Len() int {
	return len(l.List)
}

func NewEnterpriseStatistic() *EnterpriseStatistic {
	m := &EnterpriseStatistic{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseStatistic) DB() string {
	return "qdxg"
}

func (m *EnterpriseStatistic) Tab() string {
	return "enterprise_statistic"
}

func (m *EnterpriseStatistic) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"SubmitCount", "SubmitCount", &m.SubmitCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *EnterpriseStatistic) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseStatistic) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseStatistic) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseStatistic) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseStatisticList struct {
	*EnterpriseStatistic
	List  []*EnterpriseStatistic
	Total int
}

func NewEnterpriseStatisticList(cap int) *EnterpriseStatisticList {
	return &EnterpriseStatisticList{
		NewEnterpriseStatistic(),
		make([]*EnterpriseStatistic, 0, cap),
		0,
	}
}

func (l *EnterpriseStatisticList) NewModel() model.Model {
	m := NewEnterpriseStatistic()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseStatisticList) Len() int {
	return len(l.List)
}

func NewEnterprise() *Enterprise {
	m := &Enterprise{
		Account: &EnterpriseAccountList{},
	}
	model.InitModel(m)
	return m
}

func (m *Enterprise) DB() string {
	return "qdxg"
}

func (m *Enterprise) Tab() string {
	return "enterprise"
}

func (m *Enterprise) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
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

func (m *Enterprise) AutoIncField() model.Field {
	return &m.ID
}

func (m *Enterprise) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *Enterprise) UniqueKeys() []model.FieldList {
	return nil
}
func (m *Enterprise) Relations() model.RelationInfoList {
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&m.Account.EnterpriseID,
			},
			m.Account,
		},
	}

}

type EnterpriseList struct {
	*Enterprise
	List  []*Enterprise
	Total int
}

func NewEnterpriseList(cap int) *EnterpriseList {
	return &EnterpriseList{
		NewEnterprise(),
		make([]*Enterprise, 0, cap),
		0,
	}
}

func (l *EnterpriseList) NewModel() model.Model {
	m := NewEnterprise()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseList) Len() int {
	return len(l.List)
}

func NewEnterpriseAttachment() *EnterpriseAttachment {
	m := &EnterpriseAttachment{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseAttachment) DB() string {
	return "qdxg"
}

func (m *EnterpriseAttachment) Tab() string {
	return "enterprise_attachment"
}

func (m *EnterpriseAttachment) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Type", "Type", &m.Type},
		{"URL", "URL", &m.URL},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *EnterpriseAttachment) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseAttachment) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseAttachment) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseAttachment) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseAttachmentList struct {
	*EnterpriseAttachment
	List  []*EnterpriseAttachment
	Total int
}

func NewEnterpriseAttachmentList(cap int) *EnterpriseAttachmentList {
	return &EnterpriseAttachmentList{
		NewEnterpriseAttachment(),
		make([]*EnterpriseAttachment, 0, cap),
		0,
	}
}

func (l *EnterpriseAttachmentList) NewModel() model.Model {
	m := NewEnterpriseAttachment()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseAttachmentList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseAttachmentList) Len() int {
	return len(l.List)
}

func NewMidEnterpriseTag() *MidEnterpriseTag {
	m := &MidEnterpriseTag{}
	model.InitModel(m)
	return m
}

func (m *MidEnterpriseTag) DB() string {
	return "qdxg"
}

func (m *MidEnterpriseTag) Tab() string {
	return "mid_enterprise__tag"
}

func (m *MidEnterpriseTag) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"TagID", "TagID", &m.TagID},
	}
}

func (m *MidEnterpriseTag) AutoIncField() model.Field {
	return nil
}

func (m *MidEnterpriseTag) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.EnterpriseID,
		&m.TagID,
	}
}

func (m *MidEnterpriseTag) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidEnterpriseTag) Relations() model.RelationInfoList {
	return nil

}

type MidEnterpriseTagList struct {
	*MidEnterpriseTag
	List  []*MidEnterpriseTag
	Total int
}

func NewMidEnterpriseTagList(cap int) *MidEnterpriseTagList {
	return &MidEnterpriseTagList{
		NewMidEnterpriseTag(),
		make([]*MidEnterpriseTag, 0, cap),
		0,
	}
}

func (l *MidEnterpriseTagList) NewModel() model.Model {
	m := NewMidEnterpriseTag()
	l.List = append(l.List, m)
	return m
}

func (l *MidEnterpriseTagList) SetTotal(total int) {
	l.Total = total
}

func (l *MidEnterpriseTagList) Len() int {
	return len(l.List)
}

func NewEnterpriseJobStatistic() *EnterpriseJobStatistic {
	m := &EnterpriseJobStatistic{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseJobStatistic) DB() string {
	return "qdxg"
}

func (m *EnterpriseJobStatistic) Tab() string {
	return "enterprise_job_static"
}

func (m *EnterpriseJobStatistic) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobID", "JobID", &m.JobID},
		{"SubmitCount", "SubmitCount", &m.SubmitCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *EnterpriseJobStatistic) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseJobStatistic) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJobStatistic) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseJobStatistic) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseJobStatisticList struct {
	*EnterpriseJobStatistic
	List  []*EnterpriseJobStatistic
	Total int
}

func NewEnterpriseJobStatisticList(cap int) *EnterpriseJobStatisticList {
	return &EnterpriseJobStatisticList{
		NewEnterpriseJobStatistic(),
		make([]*EnterpriseJobStatistic, 0, cap),
		0,
	}
}

func (l *EnterpriseJobStatisticList) NewModel() model.Model {
	m := NewEnterpriseJobStatistic()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobStatisticList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeEnterpriseJob() *MidStudentResumeEnterpriseJob {
	m := &MidStudentResumeEnterpriseJob{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeEnterpriseJob) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeEnterpriseJob) Tab() string {
	return "mid_student_resume__enterprise_job"
}

func (m *MidStudentResumeEnterpriseJob) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"JobID", "JobID", &m.JobID},
		{"ReviewStatus", "ReviewStatus", &m.ReviewStatus},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentResumeEnterpriseJob) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeEnterpriseJob) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.JobID,
	}
}

func (m *MidStudentResumeEnterpriseJob) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeEnterpriseJob) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeEnterpriseJobList struct {
	*MidStudentResumeEnterpriseJob
	List  []*MidStudentResumeEnterpriseJob
	Total int
}

func NewMidStudentResumeEnterpriseJobList(cap int) *MidStudentResumeEnterpriseJobList {
	return &MidStudentResumeEnterpriseJobList{
		NewMidStudentResumeEnterpriseJob(),
		make([]*MidStudentResumeEnterpriseJob, 0, cap),
		0,
	}
}

func (l *MidStudentResumeEnterpriseJobList) NewModel() model.Model {
	m := NewMidStudentResumeEnterpriseJob()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeEnterpriseJobList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeEnterpriseJobList) Len() int {
	return len(l.List)
}

func NewEnterpriseJob() *EnterpriseJob {
	m := &EnterpriseJob{
		Enterprise:     &EnterpriseList{},
		StudentResumes: &StudentResumeList{},
	}
	model.InitModel(m)
	return m
}

func (m *EnterpriseJob) DB() string {
	return "qdxg"
}

func (m *EnterpriseJob) Tab() string {
	return "enterprise_job"
}

func (m *EnterpriseJob) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *EnterpriseJob) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseJob) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJob) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseJob) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeEnterpriseJob{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.EnterpriseID,
				&m.Enterprise.ID,
			},
			m.Enterprise,
		},
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.JobID,
				&mm0.ResumeID,
				&m.StudentResumes.StudentResume.ID,
			},
			m.StudentResumes,
		},
	}

}

type EnterpriseJobList struct {
	*EnterpriseJob
	List  []*EnterpriseJob
	Total int
}

func NewEnterpriseJobList(cap int) *EnterpriseJobList {
	return &EnterpriseJobList{
		NewEnterpriseJob(),
		make([]*EnterpriseJob, 0, cap),
		0,
	}
}

func (l *EnterpriseJobList) NewModel() model.Model {
	m := NewEnterpriseJob()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobList) Len() int {
	return len(l.List)
}

func NewMidStudentJobFairRead() *MidStudentJobFairRead {
	m := &MidStudentJobFairRead{}
	model.InitModel(m)
	return m
}

func (m *MidStudentJobFairRead) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairRead) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairRead) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairRead) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentJobFairRead) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairRead) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentJobFairRead) Relations() model.RelationInfoList {
	return nil

}

type MidStudentJobFairReadList struct {
	*MidStudentJobFairRead
	List  []*MidStudentJobFairRead
	Total int
}

func NewMidStudentJobFairReadList(cap int) *MidStudentJobFairReadList {
	return &MidStudentJobFairReadList{
		NewMidStudentJobFairRead(),
		make([]*MidStudentJobFairRead, 0, cap),
		0,
	}
}

func (l *MidStudentJobFairReadList) NewModel() model.Model {
	m := NewMidStudentJobFairRead()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairReadList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairReadList) Len() int {
	return len(l.List)
}

func NewMidStudentJobFairEnroll() *MidStudentJobFairEnroll {
	m := &MidStudentJobFairEnroll{}
	model.InitModel(m)
	return m
}

func (m *MidStudentJobFairEnroll) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairEnroll) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairEnroll) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairEnroll) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentJobFairEnroll) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairEnroll) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentJobFairEnroll) Relations() model.RelationInfoList {
	return nil

}

type MidStudentJobFairEnrollList struct {
	*MidStudentJobFairEnroll
	List  []*MidStudentJobFairEnroll
	Total int
}

func NewMidStudentJobFairEnrollList(cap int) *MidStudentJobFairEnrollList {
	return &MidStudentJobFairEnrollList{
		NewMidStudentJobFairEnroll(),
		make([]*MidStudentJobFairEnroll, 0, cap),
		0,
	}
}

func (l *MidStudentJobFairEnrollList) NewModel() model.Model {
	m := NewMidStudentJobFairEnroll()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairEnrollList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairEnrollList) Len() int {
	return len(l.List)
}

func NewMidStudentJobFairShare() *MidStudentJobFairShare {
	m := &MidStudentJobFairShare{}
	model.InitModel(m)
	return m
}

func (m *MidStudentJobFairShare) DB() string {
	return "qdxg"
}

func (m *MidStudentJobFairShare) Tab() string {
	return "mid_student__job_fair_read"
}

func (m *MidStudentJobFairShare) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *MidStudentJobFairShare) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentJobFairShare) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.IntelUserCode,
		&m.JobFairID,
	}
}

func (m *MidStudentJobFairShare) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentJobFairShare) Relations() model.RelationInfoList {
	return nil

}

type MidStudentJobFairShareList struct {
	*MidStudentJobFairShare
	List  []*MidStudentJobFairShare
	Total int
}

func NewMidStudentJobFairShareList(cap int) *MidStudentJobFairShareList {
	return &MidStudentJobFairShareList{
		NewMidStudentJobFairShare(),
		make([]*MidStudentJobFairShare, 0, cap),
		0,
	}
}

func (l *MidStudentJobFairShareList) NewModel() model.Model {
	m := NewMidStudentJobFairShare()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentJobFairShareList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentJobFairShareList) Len() int {
	return len(l.List)
}

func NewJobFairStatistic() *JobFairStatistic {
	m := &JobFairStatistic{}
	model.InitModel(m)
	return m
}

func (m *JobFairStatistic) DB() string {
	return "qdxg"
}

func (m *JobFairStatistic) Tab() string {
	return "job_fair_statistic"
}

func (m *JobFairStatistic) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobFairID", "JobFairID", &m.JobFairID},
		{"ReadCount", "ReadCount", &m.ReadCount},
		{"EnrollCount", "EnrollCount", &m.EnrollCount},
		{"ShareCount", "ShareCount", &m.ShareCount},
		{"CreateDate", "CreateDate", &m.CreateDate},
	}
}

func (m *JobFairStatistic) AutoIncField() model.Field {
	return &m.ID
}

func (m *JobFairStatistic) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *JobFairStatistic) UniqueKeys() []model.FieldList {
	return nil
}
func (m *JobFairStatistic) Relations() model.RelationInfoList {
	return nil

}

type JobFairStatisticList struct {
	*JobFairStatistic
	List  []*JobFairStatistic
	Total int
}

func NewJobFairStatisticList(cap int) *JobFairStatisticList {
	return &JobFairStatisticList{
		NewJobFairStatistic(),
		make([]*JobFairStatistic, 0, cap),
		0,
	}
}

func (l *JobFairStatisticList) NewModel() model.Model {
	m := NewJobFairStatistic()
	l.List = append(l.List, m)
	return m
}

func (l *JobFairStatisticList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFairStatisticList) Len() int {
	return len(l.List)
}

func NewJobFair() *JobFair {
	m := &JobFair{}
	model.InitModel(m)
	return m
}

func (m *JobFair) DB() string {
	return "qdxg"
}

func (m *JobFair) Tab() string {
	return "job_fair"
}

func (m *JobFair) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *JobFair) AutoIncField() model.Field {
	return &m.ID
}

func (m *JobFair) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *JobFair) UniqueKeys() []model.FieldList {
	return nil
}
func (m *JobFair) Relations() model.RelationInfoList {
	return nil

}

type JobFairList struct {
	*JobFair
	List  []*JobFair
	Total int
}

func NewJobFairList(cap int) *JobFairList {
	return &JobFairList{
		NewJobFair(),
		make([]*JobFair, 0, cap),
		0,
	}
}

func (l *JobFairList) NewModel() model.Model {
	m := NewJobFair()
	l.List = append(l.List, m)
	return m
}

func (l *JobFairList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFairList) Len() int {
	return len(l.List)
}

func NewJobFlag() *JobFlag {
	m := &JobFlag{}
	model.InitModel(m)
	return m
}

func (m *JobFlag) DB() string {
	return "qdxg"
}

func (m *JobFlag) Tab() string {
	return "job_fair"
}

func (m *JobFlag) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *JobFlag) AutoIncField() model.Field {
	return &m.ID
}

func (m *JobFlag) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *JobFlag) UniqueKeys() []model.FieldList {
	return nil
}
func (m *JobFlag) Relations() model.RelationInfoList {
	return nil

}

type JobFlagList struct {
	*JobFlag
	List  []*JobFlag
	Total int
}

func NewJobFlagList(cap int) *JobFlagList {
	return &JobFlagList{
		NewJobFlag(),
		make([]*JobFlag, 0, cap),
		0,
	}
}

func (l *JobFlagList) NewModel() model.Model {
	m := NewJobFlag()
	l.List = append(l.List, m)
	return m
}

func (l *JobFlagList) SetTotal(total int) {
	l.Total = total
}

func (l *JobFlagList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeLanguageSkill() *MidStudentResumeLanguageSkill {
	m := &MidStudentResumeLanguageSkill{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeLanguageSkill) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeLanguageSkill) Tab() string {
	return "mid_student_resume__language_skill"
}

func (m *MidStudentResumeLanguageSkill) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"LanguageSkillID", "LanguageSkillID", &m.LanguageSkillID},
	}
}

func (m *MidStudentResumeLanguageSkill) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeLanguageSkill) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.LanguageSkillID,
	}
}

func (m *MidStudentResumeLanguageSkill) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeLanguageSkill) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeLanguageSkillList struct {
	*MidStudentResumeLanguageSkill
	List  []*MidStudentResumeLanguageSkill
	Total int
}

func NewMidStudentResumeLanguageSkillList(cap int) *MidStudentResumeLanguageSkillList {
	return &MidStudentResumeLanguageSkillList{
		NewMidStudentResumeLanguageSkill(),
		make([]*MidStudentResumeLanguageSkill, 0, cap),
		0,
	}
}

func (l *MidStudentResumeLanguageSkillList) NewModel() model.Model {
	m := NewMidStudentResumeLanguageSkill()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeLanguageSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeLanguageSkillList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeStudentTrain() *MidStudentResumeStudentTrain {
	m := &MidStudentResumeStudentTrain{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeStudentTrain) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentTrain) Tab() string {
	return "mid_student_resume__student_train"
}

func (m *MidStudentResumeStudentTrain) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"TrainID", "TrainID", &m.TrainID},
	}
}

func (m *MidStudentResumeStudentTrain) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeStudentTrain) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.TrainID,
	}
}

func (m *MidStudentResumeStudentTrain) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeStudentTrain) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeStudentTrainList struct {
	*MidStudentResumeStudentTrain
	List  []*MidStudentResumeStudentTrain
	Total int
}

func NewMidStudentResumeStudentTrainList(cap int) *MidStudentResumeStudentTrainList {
	return &MidStudentResumeStudentTrainList{
		NewMidStudentResumeStudentTrain(),
		make([]*MidStudentResumeStudentTrain, 0, cap),
		0,
	}
}

func (l *MidStudentResumeStudentTrainList) NewModel() model.Model {
	m := NewMidStudentResumeStudentTrain()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentTrainList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentTrainList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeStudentHonor() *MidStudentResumeStudentHonor {
	m := &MidStudentResumeStudentHonor{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeStudentHonor) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentHonor) Tab() string {
	return "mid_student_resume__student_honor"
}

func (m *MidStudentResumeStudentHonor) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"HonorID", "HonorID", &m.HonorID},
	}
}

func (m *MidStudentResumeStudentHonor) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeStudentHonor) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.HonorID,
	}
}

func (m *MidStudentResumeStudentHonor) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeStudentHonor) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeStudentHonorList struct {
	*MidStudentResumeStudentHonor
	List  []*MidStudentResumeStudentHonor
	Total int
}

func NewMidStudentResumeStudentHonorList(cap int) *MidStudentResumeStudentHonorList {
	return &MidStudentResumeStudentHonorList{
		NewMidStudentResumeStudentHonor(),
		make([]*MidStudentResumeStudentHonor, 0, cap),
		0,
	}
}

func (l *MidStudentResumeStudentHonorList) NewModel() model.Model {
	m := NewMidStudentResumeStudentHonor()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentHonorList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentHonorList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeStudentExperience() *MidStudentResumeStudentExperience {
	m := &MidStudentResumeStudentExperience{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeStudentExperience) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentExperience) Tab() string {
	return "mid_student_resume__student_experience"
}

func (m *MidStudentResumeStudentExperience) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"ExperienceID", "ExperienceID", &m.ExperienceID},
	}
}

func (m *MidStudentResumeStudentExperience) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeStudentExperience) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.ExperienceID,
	}
}

func (m *MidStudentResumeStudentExperience) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeStudentExperience) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeStudentExperienceList struct {
	*MidStudentResumeStudentExperience
	List  []*MidStudentResumeStudentExperience
	Total int
}

func NewMidStudentResumeStudentExperienceList(cap int) *MidStudentResumeStudentExperienceList {
	return &MidStudentResumeStudentExperienceList{
		NewMidStudentResumeStudentExperience(),
		make([]*MidStudentResumeStudentExperience, 0, cap),
		0,
	}
}

func (l *MidStudentResumeStudentExperienceList) NewModel() model.Model {
	m := NewMidStudentResumeStudentExperience()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentExperienceList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentExperienceList) Len() int {
	return len(l.List)
}

func NewMidStudentResumeStudentSkill() *MidStudentResumeStudentSkill {
	m := &MidStudentResumeStudentSkill{}
	model.InitModel(m)
	return m
}

func (m *MidStudentResumeStudentSkill) DB() string {
	return "qdxg"
}

func (m *MidStudentResumeStudentSkill) Tab() string {
	return "mid_student_resume__student_skill"
}

func (m *MidStudentResumeStudentSkill) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"SkillID", "SkillID", &m.SkillID},
	}
}

func (m *MidStudentResumeStudentSkill) AutoIncField() model.Field {
	return nil
}

func (m *MidStudentResumeStudentSkill) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ResumeID,
		&m.SkillID,
	}
}

func (m *MidStudentResumeStudentSkill) UniqueKeys() []model.FieldList {
	return nil
}
func (m *MidStudentResumeStudentSkill) Relations() model.RelationInfoList {
	return nil

}

type MidStudentResumeStudentSkillList struct {
	*MidStudentResumeStudentSkill
	List  []*MidStudentResumeStudentSkill
	Total int
}

func NewMidStudentResumeStudentSkillList(cap int) *MidStudentResumeStudentSkillList {
	return &MidStudentResumeStudentSkillList{
		NewMidStudentResumeStudentSkill(),
		make([]*MidStudentResumeStudentSkill, 0, cap),
		0,
	}
}

func (l *MidStudentResumeStudentSkillList) NewModel() model.Model {
	m := NewMidStudentResumeStudentSkill()
	l.List = append(l.List, m)
	return m
}

func (l *MidStudentResumeStudentSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *MidStudentResumeStudentSkillList) Len() int {
	return len(l.List)
}

func NewStudentTrain() *StudentTrain {
	m := &StudentTrain{
		StudentResume: &StudentResume{},
	}
	model.InitModel(m)
	return m
}

func (m *StudentTrain) DB() string {
	return "qdxg"
}

func (m *StudentTrain) Tab() string {
	return "student_train"
}

func (m *StudentTrain) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *StudentTrain) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentTrain) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentTrain) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentTrain) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeStudentTrain{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.TrainID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
		},
	}

}

type StudentTrainList struct {
	*StudentTrain
	List  []*StudentTrain
	Total int
}

func NewStudentTrainList(cap int) *StudentTrainList {
	return &StudentTrainList{
		NewStudentTrain(),
		make([]*StudentTrain, 0, cap),
		0,
	}
}

func (l *StudentTrainList) NewModel() model.Model {
	m := NewStudentTrain()
	l.List = append(l.List, m)
	return m
}

func (l *StudentTrainList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentTrainList) Len() int {
	return len(l.List)
}

func NewStudentHonor() *StudentHonor {
	m := &StudentHonor{
		StudentResume: &StudentResume{},
	}
	model.InitModel(m)
	return m
}

func (m *StudentHonor) DB() string {
	return "qdxg"
}

func (m *StudentHonor) Tab() string {
	return "student_honor"
}

func (m *StudentHonor) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *StudentHonor) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentHonor) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentHonor) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentHonor) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeStudentHonor{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.HonorID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
		},
	}

}

type StudentHonorList struct {
	*StudentHonor
	List  []*StudentHonor
	Total int
}

func NewStudentHonorList(cap int) *StudentHonorList {
	return &StudentHonorList{
		NewStudentHonor(),
		make([]*StudentHonor, 0, cap),
		0,
	}
}

func (l *StudentHonorList) NewModel() model.Model {
	m := NewStudentHonor()
	l.List = append(l.List, m)
	return m
}

func (l *StudentHonorList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentHonorList) Len() int {
	return len(l.List)
}

func NewStudentExperience() *StudentExperience {
	m := &StudentExperience{
		StudentResume: &StudentResume{},
	}
	model.InitModel(m)
	return m
}

func (m *StudentExperience) DB() string {
	return "qdxg"
}

func (m *StudentExperience) Tab() string {
	return "student_experience"
}

func (m *StudentExperience) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
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

func (m *StudentExperience) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentExperience) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentExperience) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentExperience) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeStudentExperience{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.ExperienceID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
		},
	}

}

type StudentExperienceList struct {
	*StudentExperience
	List  []*StudentExperience
	Total int
}

func NewStudentExperienceList(cap int) *StudentExperienceList {
	return &StudentExperienceList{
		NewStudentExperience(),
		make([]*StudentExperience, 0, cap),
		0,
	}
}

func (l *StudentExperienceList) NewModel() model.Model {
	m := NewStudentExperience()
	l.List = append(l.List, m)
	return m
}

func (l *StudentExperienceList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentExperienceList) Len() int {
	return len(l.List)
}

func NewStudentSkill() *StudentSkill {
	m := &StudentSkill{
		StudentResume: &StudentResume{},
	}
	model.InitModel(m)
	return m
}

func (m *StudentSkill) DB() string {
	return "qdxg"
}

func (m *StudentSkill) Tab() string {
	return "student_skill"
}

func (m *StudentSkill) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"Name", "Name", &m.Name},
		{"Description", "Description", &m.Description},
		{"Status", "Status", &m.Status},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentSkill) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentSkill) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentSkill) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentSkill) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeStudentSkill{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.SkillID,
				&mm0.ResumeID,
				&m.StudentResume.ID,
			},
			m.StudentResume,
		},
	}

}

type StudentSkillList struct {
	*StudentSkill
	List  []*StudentSkill
	Total int
}

func NewStudentSkillList(cap int) *StudentSkillList {
	return &StudentSkillList{
		NewStudentSkill(),
		make([]*StudentSkill, 0, cap),
		0,
	}
}

func (l *StudentSkillList) NewModel() model.Model {
	m := NewStudentSkill()
	l.List = append(l.List, m)
	return m
}

func (l *StudentSkillList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentSkillList) Len() int {
	return len(l.List)
}

func NewStudentResume() *StudentResume {
	m := &StudentResume{
		StudentTrain:        &StudentTrainList{},
		StudentHonor:        &StudentHonorList{},
		StudentExperience:   &StudentExperienceList{},
		StudentSkill:        &StudentSkillList{},
		StudentLanguageType: &JobFlagList{},
	}
	model.InitModel(m)
	return m
}

func (m *StudentResume) DB() string {
	return "qdxg"
}

func (m *StudentResume) Tab() string {
	return "student_resume"
}

func (m *StudentResume) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"IntelUserCode", "IntelUserCode", &m.IntelUserCode},
		{"Introduction", "Introduction", &m.Introduction},
		{"CreateTime", "CreateTime", &m.CreateTime},
		{"UpdateTime", "UpdateTime", &m.UpdateTime},
	}
}

func (m *StudentResume) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentResume) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentResume) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentResume) Relations() model.RelationInfoList {
	mm0 := MidStudentResumeStudentTrain{}
	mm1 := MidStudentResumeStudentHonor{}
	mm2 := MidStudentResumeStudentExperience{}
	mm3 := MidStudentResumeStudentSkill{}
	return model.RelationInfoList{
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm0.ResumeID,
				&mm0.TrainID,
				&m.StudentTrain.ID,
			},
			m.StudentTrain,
		},
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm1.ResumeID,
				&mm1.HonorID,
				&m.StudentHonor.ID,
			},
			m.StudentHonor,
		},
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm2.ResumeID,
				&mm2.ExperienceID,
				&m.StudentExperience.ID,
			},
			m.StudentExperience,
		},
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&mm3.ResumeID,
				&mm3.SkillID,
				&m.StudentSkill.ID,
			},
			m.StudentSkill,
		},
		model.RelationInfo{
			model.FieldList{
				&m.ID,
				&m.StudentLanguageType.ID,
			},
			m.StudentLanguageType,
		},
	}

}

type StudentResumeList struct {
	*StudentResume
	List  []*StudentResume
	Total int
}

func NewStudentResumeList(cap int) *StudentResumeList {
	return &StudentResumeList{
		NewStudentResume(),
		make([]*StudentResume, 0, cap),
		0,
	}
}

func (l *StudentResumeList) NewModel() model.Model {
	m := NewStudentResume()
	l.List = append(l.List, m)
	return m
}

func (l *StudentResumeList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentResumeList) Len() int {
	return len(l.List)
}

func NewEnterpriseSnapshot() *EnterpriseSnapshot {
	m := &EnterpriseSnapshot{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseSnapshot) DB() string {
	return "qdxg"
}

func (m *EnterpriseSnapshot) Tab() string {
	return "enterprise_snapshot"
}

func (m *EnterpriseSnapshot) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"EnterpriseID", "EnterpriseID", &m.EnterpriseID},
		{"Info", "Info", &m.Info},
	}
}

func (m *EnterpriseSnapshot) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseSnapshot) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseSnapshot) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseSnapshot) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseSnapshotList struct {
	*EnterpriseSnapshot
	List  []*EnterpriseSnapshot
	Total int
}

func NewEnterpriseSnapshotList(cap int) *EnterpriseSnapshotList {
	return &EnterpriseSnapshotList{
		NewEnterpriseSnapshot(),
		make([]*EnterpriseSnapshot, 0, cap),
		0,
	}
}

func (l *EnterpriseSnapshotList) NewModel() model.Model {
	m := NewEnterpriseSnapshot()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseSnapshotList) Len() int {
	return len(l.List)
}

func NewEnterpriseJobSnapshot() *EnterpriseJobSnapshot {
	m := &EnterpriseJobSnapshot{}
	model.InitModel(m)
	return m
}

func (m *EnterpriseJobSnapshot) DB() string {
	return "qdxg"
}

func (m *EnterpriseJobSnapshot) Tab() string {
	return "enterprise_job_snapshot"
}

func (m *EnterpriseJobSnapshot) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"JobID", "JobID", &m.JobID},
		{"Into", "Into", &m.Into},
	}
}

func (m *EnterpriseJobSnapshot) AutoIncField() model.Field {
	return &m.ID
}

func (m *EnterpriseJobSnapshot) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *EnterpriseJobSnapshot) UniqueKeys() []model.FieldList {
	return nil
}
func (m *EnterpriseJobSnapshot) Relations() model.RelationInfoList {
	return nil

}

type EnterpriseJobSnapshotList struct {
	*EnterpriseJobSnapshot
	List  []*EnterpriseJobSnapshot
	Total int
}

func NewEnterpriseJobSnapshotList(cap int) *EnterpriseJobSnapshotList {
	return &EnterpriseJobSnapshotList{
		NewEnterpriseJobSnapshot(),
		make([]*EnterpriseJobSnapshot, 0, cap),
		0,
	}
}

func (l *EnterpriseJobSnapshotList) NewModel() model.Model {
	m := NewEnterpriseJobSnapshot()
	l.List = append(l.List, m)
	return m
}

func (l *EnterpriseJobSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *EnterpriseJobSnapshotList) Len() int {
	return len(l.List)
}

func NewStudentResumeSnapshot() *StudentResumeSnapshot {
	m := &StudentResumeSnapshot{}
	model.InitModel(m)
	return m
}

func (m *StudentResumeSnapshot) DB() string {
	return "qdxg"
}

func (m *StudentResumeSnapshot) Tab() string {
	return "student_resume_snapshot"
}

func (m *StudentResumeSnapshot) FieldInfos() model.FieldInfoList {
	return model.FieldInfoList{
		{"ID", "ID", &m.ID},
		{"ResumeID", "ResumeID", &m.ResumeID},
		{"Info", "Info", &m.Info},
	}
}

func (m *StudentResumeSnapshot) AutoIncField() model.Field {
	return &m.ID
}

func (m *StudentResumeSnapshot) PrimaryKey() model.FieldList {
	return model.FieldList{
		&m.ID,
	}
}

func (m *StudentResumeSnapshot) UniqueKeys() []model.FieldList {
	return nil
}
func (m *StudentResumeSnapshot) Relations() model.RelationInfoList {
	return nil

}

type StudentResumeSnapshotList struct {
	*StudentResumeSnapshot
	List  []*StudentResumeSnapshot
	Total int
}

func NewStudentResumeSnapshotList(cap int) *StudentResumeSnapshotList {
	return &StudentResumeSnapshotList{
		NewStudentResumeSnapshot(),
		make([]*StudentResumeSnapshot, 0, cap),
		0,
	}
}

func (l *StudentResumeSnapshotList) NewModel() model.Model {
	m := NewStudentResumeSnapshot()
	l.List = append(l.List, m)
	return m
}

func (l *StudentResumeSnapshotList) SetTotal(total int) {
	l.Total = total
}

func (l *StudentResumeSnapshotList) Len() int {
	return len(l.List)
}
