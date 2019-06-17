package job

import (
	"AntLinkCampus/CampusServer/handles/job/model"
)

//db:qdxg
//tab:enterprise_account
//pk:ID
type EnterpriseAccount struct {
	model.Meta
	ID           model.Int `auto_increment:"true"`
	EnterpriseID model.Int
	Phone        model.String
	Email        model.String
	Password     model.String
	Enterprise   *EnterpriseList `rel:"EnterpriseID->ID"`
}

//db:qdxg
//tab:enterprise_revice_status
//pk:ID
type EnterpriseReviewStatus struct {
	model.Meta
	ID           model.Int `auto_increment:"true"`
	EnterpriseID model.Int
	ReviewStatus model.String
	Operator     model.String
	CreateTime   model.Datetime
	UpdateTime   model.Datetime
}

//db:qdxg
//tab:enterprise_statistic
//pk:ID
type EnterpriseStatistic struct {
	model.Meta
	ID           model.Int `auto_increment:"true"`
	EnterpriseID model.Int
	SubmitCount  model.Int
	CreateDate   model.Date
}

//db:qdxg
//tab:enterprise
//pk:ID
type Enterprise struct {
	model.Meta
	ID               model.Int `auto_increment:"true"`
	UniformCode      model.String
	Name             model.String
	RegisterCityID   model.Int
	RegisterAddress  model.String
	SectorID         model.Int
	NatureID         model.Int
	ScopeID          model.Int
	Website          model.String
	Contact          model.String
	EmployeeFromThis model.Int
	Introduction     model.String
	ZipCode          model.String
	CreateTime       model.Datetime
	UpdateTime       model.Datetime
	Account          *EnterpriseAccountList `rel:"ID->EnterpriseID"`
}

//db:qdxg
//tab:enterprise_attachment
//pk:ID
type EnterpriseAttachment struct {
	model.Meta
	ID           model.Int `auto_increment:"true"`
	EnterpriseID model.Int
	Type         model.String
	URL          model.String
	Status       model.String
	CreateTime   model.Datetime
	UpdateTime   model.Datetime
}

//db:qdxg
//tab:mid_enterprise__tag
//pk:EnterpriseID,TagID
type MidEnterpriseTag struct {
	model.Meta
	EnterpriseID model.Int
	TagID        model.Int
}

//db:qdxg
//tab:enterprise_job_static
//pk:ID
type EnterpriseJobStatistic struct {
	model.Meta
	ID          model.Int `auto_increment:"true"`
	JobID       model.Int
	SubmitCount model.Int
	CreateDate  model.Date
}

//db:qdxg
//tab:mid_student_resume__enterprise_job
//pk:ResumeID,JobID
type MidStudentResumeEnterpriseJob struct {
	model.Meta
	ResumeID     model.Int
	JobID        model.Int
	ReviewStatus model.String
	CreateTime   model.Datetime
	UpdateTime   model.Datetime
}

//db:qdxg
//tab:enterprise_job
//pk:ID
type EnterpriseJob struct {
	model.Meta
	ID              model.Int `auto_increment:"true"`
	EnterpriseID    model.Int
	Name            model.String
	CityID          model.Int
	Address         model.String
	TypeID          model.Int
	Gender          model.String
	MajorCode       model.String
	DegreeID        model.Int
	LanguageSkillID model.Int
	Description     model.String
	SalaryRangeID   model.Int
	Welfare         model.String
	Vacancies       model.Int
	ExpiredAt       model.Date
	Status          model.String
	Comment         model.String
	CreateTime      model.Datetime
	UpdateTime      model.Datetime
	Enterprise      *EnterpriseList    `rel:"EnterpriseID->ID"`
	StudentResumes  *StudentResumeList `rel:"ID->MidStudentResumeEnterpriseJob.JobID->MidEnterpriseJobStudentResume.ResumeID->StudentResume.ID"`
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairRead struct {
	model.Meta
	IntelUserCode model.String
	JobFairID     model.Int
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairEnroll struct {
	model.Meta
	IntelUserCode model.String
	JobFairID     model.Int
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairShare struct {
	model.Meta
	IntelUserCode model.String
	JobFairID     model.Int
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
}

//db:qdxg
//tab:job_fair_statistic
//pk:ID
type JobFairStatistic struct {
	model.Meta
	ID          model.Int `auto_increment:"true"`
	JobFairID   model.Int
	ReadCount   model.Int
	EnrollCount model.Int
	ShareCount  model.Int
	CreateDate  model.Date
}

//db:qdxg
//tab:job_fair
//pk:ID
type JobFair struct {
	model.Meta
	ID          model.Int `auto_increment:"true"`
	Name        model.String
	StartTime   model.Datetime
	EndTime     model.Datetime
	Description model.String
	Status      model.String
	Comment     model.String
	CreateTime  model.Datetime
	UpdateTime  model.Datetime
}

//db:qdxg
//tab:job_fair
//pk:ID
type JobFlag struct {
	model.Meta
	ID         model.Int `auto_increment:"true"`
	Name       model.String
	Type       model.String
	Value      model.Decimal
	Order      model.Int
	ParentID   model.Int
	Status     model.String
	Operator   model.String
	Comment    model.String
	CreateTime model.Datetime
	UpdateTime model.Datetime
}

//db:qdxg
//tab:mid_student_resume__language_skill
//pk:ResumeID,LanguageSkillID
type MidStudentResumeLanguageSkill struct {
	model.Meta
	ResumeID        model.Int
	LanguageSkillID model.Int
}

//db:qdxg
//tab:mid_student_resume__student_train
//pk:ResumeID,TrainID
type MidStudentResumeStudentTrain struct {
	model.Meta
	ResumeID model.Int
	TrainID  model.Int
}

//db:qdxg
//tab:mid_student_resume__student_honor
//pk:ResumeID,HonorID
type MidStudentResumeStudentHonor struct {
	model.Meta
	ResumeID model.Int
	HonorID  model.Int
}

//db:qdxg
//tab:mid_student_resume__student_experience
//pk:ResumeID,ExperienceID
type MidStudentResumeStudentExperience struct {
	model.Meta
	ResumeID     model.Int
	ExperienceID model.Int
}

//db:qdxg
//tab:mid_student_resume__student_skill
//pk:ResumeID,SkillID
type MidStudentResumeStudentSkill struct {
	model.Meta
	ResumeID model.Int
	SkillID  model.Int
}

//db:qdxg
//tab:student_train
//pk:ID
type StudentTrain struct {
	model.Meta
	ID              model.Int `auto_increment:"true"`
	IntelUserCode   model.String
	InstitutionName model.String
	StartDate       model.Date
	EndDate         model.Date
	Degree          model.String
	Major           model.String
	Description     model.String
	Status          model.String
	CreateTime      model.Datetime
	UpdateTime      model.Datetime
	StudentResume   *StudentResume `rel:"ID->MidStudentResumeStudentTrain.TrainID->MidStudentResumeStudentTrain.ResumeID->ID"`
}

//db:qdxg
//tab:student_honor
//pk:ID
type StudentHonor struct {
	model.Meta
	ID            model.Int `auto_increment:"true"`
	IntelUserCode model.String
	Name          model.String
	Description   model.String
	GrantDate     model.Date
	Status        model.String
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentHonor.HonorID->MidStudentResumeStudentHonor.ResumeID->ID"`
}

//db:qdxg
//tab:student_experience
//pk:ID
type StudentExperience struct {
	model.Meta
	ID            model.Int `auto_increment:"true"`
	IntelUserCode model.String
	CompanyName   model.String
	StartDate     model.Date
	EndDate       model.Date
	SectorID      model.Int
	Position      model.String
	Description   model.String
	Status        model.String
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentExperience.ExperienceID->MidStudentResumeStudentExperience.ResumeID->ID"`
}

//db:qdxg
//tab:student_skill
//pk:ID
type StudentSkill struct {
	model.Meta
	ID            model.Int `auto_increment:"true"`
	IntelUserCode model.String
	Name          model.String
	Description   model.String
	Status        model.String
	CreateTime    model.Datetime
	UpdateTime    model.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentSkill.SkillID->MidStudentResumeStudentSkill.ResumeID->ID"`
}

//db:qdxg
//tab:student_resume
//pk:ID
type StudentResume struct {
	model.Meta
	ID                  model.Int              `col:"ID" auto_increment:"true"`
	IntelUserCode       model.String           `col:"IntelUserCode"`
	Introduction        model.String           `col:"Introduction"`
	CreateTime          model.Datetime         `col:"CreateTime"`
	UpdateTime          model.Datetime         `col:"UpdateTime"`
	StudentTrain        *StudentTrainList      `rel:"ID->MidStudentResumeStudentTrain.ResumeID->StudentResumeStudentTrain.TrainID->ID"`
	StudentHonor        *StudentHonorList      `rel:"ID->MidStudentResumeStudentHonor.ResumeID->MidStudentResumeStudentHonor.HonorID->ID"`
	StudentExperience   *StudentExperienceList `rel:"ID->MidStudentResumeStudentExperience.ResumeID->StudentResumeStudentExperience.ExperienceID->ID"`
	StudentSkill        *StudentSkillList      `rel:"ID->MidStudentResumeStudentSkill.ResumeID->StudentResumeStudentSkill.SkillID->ID"`
	StudentLanguageType *JobFlagList           `rel:"ID->ID"`
}

//db:qdxg
//tab:enterprise_snapshot
//pk:ID
type EnterpriseSnapshot struct {
	model.Meta
	ID           model.Int `auto_increment:"true"`
	EnterpriseID model.Int
	Info         model.String
}

//db:qdxg
//tab:enterprise_job_snapshot
//pk:ID
type EnterpriseJobSnapshot struct {
	model.Meta
	ID    model.Int `auto_increment:"true"`
	JobID model.Int
	Into  model.String
}

//db:qdxg
//tab:student_resume_snapshot
//pk:ID
type StudentResumeSnapshot struct {
	model.Meta
	ID       model.Int `auto_increment:"true"`
	ResumeID model.Int
	Info     model.String
}
