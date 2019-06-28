package model

import "github.com/wangjun861205/nborm"

//db:qdxg
//tab:enterprise_account
//pk:ID
type EnterpriseAccount struct {
	nborm.Meta
	ID           nborm.Int `auto_increment:"true"`
	EnterpriseID nborm.Int
	Phone        nborm.String
	Email        nborm.String
	Password     nborm.String
	Enterprise   *EnterpriseList `rel:"EnterpriseID->ID(Name='xxx')"`
}

//db:qdxg
//tab:enterprise_revice_status
//pk:ID
type EnterpriseReviewStatus struct {
	nborm.Meta
	ID           nborm.Int `auto_increment:"true"`
	EnterpriseID nborm.Int
	ReviewStatus nborm.String
	Operator     nborm.String
	CreateTime   nborm.Datetime
	UpdateTime   nborm.Datetime
}

//db:qdxg
//tab:enterprise_statistic
//pk:ID
type EnterpriseStatistic struct {
	nborm.Meta
	ID           nborm.Int `auto_increment:"true"`
	EnterpriseID nborm.Int
	SubmitCount  nborm.Int
	CreateDate   nborm.Date
}

//db:qdxg
//tab:enterprise
//pk:ID
type Enterprise struct {
	nborm.Meta
	ID               nborm.Int `auto_increment:"true"`
	UniformCode      nborm.String
	Name             nborm.String
	RegisterCityID   nborm.Int
	RegisterAddress  nborm.String
	SectorID         nborm.Int
	NatureID         nborm.Int
	ScopeID          nborm.Int
	Website          nborm.String
	Contact          nborm.String
	EmployeeFromThis nborm.Int
	Introduction     nborm.String
	ZipCode          nborm.String
	CreateTime       nborm.Datetime
	UpdateTime       nborm.Datetime
	Account          *EnterpriseAccountList `rel:"ID->EnterpriseID"`
}

//db:qdxg
//tab:enterprise_attachment
//pk:ID
type EnterpriseAttachment struct {
	nborm.Meta
	ID           nborm.Int `auto_increment:"true"`
	EnterpriseID nborm.Int
	Type         nborm.String
	URL          nborm.String
	Status       nborm.String
	CreateTime   nborm.Datetime
	UpdateTime   nborm.Datetime
}

//db:qdxg
//tab:mid_enterprise__tag
//pk:EnterpriseID,TagID
type MidEnterpriseTag struct {
	nborm.Meta
	EnterpriseID nborm.Int
	TagID        nborm.Int
}

//db:qdxg
//tab:enterprise_job_static
//pk:ID
type EnterpriseJobStatistic struct {
	nborm.Meta
	ID          nborm.Int `auto_increment:"true"`
	JobID       nborm.Int
	SubmitCount nborm.Int
	CreateDate  nborm.Date
}

//db:qdxg
//tab:mid_student_resume__enterprise_job
//pk:ResumeID,JobID
type MidStudentResumeEnterpriseJob struct {
	nborm.Meta
	ResumeID     nborm.Int
	JobID        nborm.Int
	ReviewStatus nborm.String
	CreateTime   nborm.Datetime
	UpdateTime   nborm.Datetime
}

//db:qdxg
//tab:enterprise_job
//pk:ID
type EnterpriseJob struct {
	nborm.Meta
	ID              nborm.Int `auto_increment:"true"`
	EnterpriseID    nborm.Int
	Name            nborm.String
	CityID          nborm.Int
	Address         nborm.String
	TypeID          nborm.Int
	Gender          nborm.String
	MajorCode       nborm.String
	DegreeID        nborm.Int
	LanguageSkillID nborm.Int
	Description     nborm.String
	SalaryRangeID   nborm.Int
	Welfare         nborm.String
	Vacancies       nborm.Int
	ExpiredAt       nborm.Datetime
	Status          nborm.String
	Comment         nborm.String
	CreateTime      nborm.Datetime
	UpdateTime      nborm.Datetime
	Enterprise      *EnterpriseList    `rel:"EnterpriseID->ID"`
	StudentResumes  *StudentResumeList `rel:"ID->MidStudentResumeEnterpriseJob(ReviewStatus=1).JobID->MidEnterpriseJobStudentResume.ResumeID->ID"`
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairRead struct {
	nborm.Meta
	IntelUserCode nborm.String
	JobFairID     nborm.Int
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairEnroll struct {
	nborm.Meta
	IntelUserCode nborm.String
	JobFairID     nborm.Int
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
}

//db:qdxg
//tab:mid_student__job_fair_read
//pk:IntelUserCode,JobFairID
type MidStudentJobFairShare struct {
	nborm.Meta
	IntelUserCode nborm.String
	JobFairID     nborm.Int
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
}

//db:qdxg
//tab:job_fair_statistic
//pk:ID
type JobFairStatistic struct {
	nborm.Meta
	ID          nborm.Int `auto_increment:"true"`
	JobFairID   nborm.Int
	ReadCount   nborm.Int
	EnrollCount nborm.Int
	ShareCount  nborm.Int
	CreateDate  nborm.Date
}

//db:qdxg
//tab:job_fair
//pk:ID
type JobFair struct {
	nborm.Meta
	ID          nborm.Int `auto_increment:"true"`
	Name        nborm.String
	StartTime   nborm.Datetime
	EndTime     nborm.Datetime
	Description nborm.String
	Status      nborm.String
	Comment     nborm.String
	CreateTime  nborm.Datetime
	UpdateTime  nborm.Datetime
}

//db:qdxg
//tab:job_fair
//pk:ID
type JobFlag struct {
	nborm.Meta
	ID         nborm.Int `auto_increment:"true"`
	Name       nborm.String
	Type       nborm.String
	Value      nborm.Decimal
	Order      nborm.Int
	ParentID   nborm.Int
	Status     nborm.String
	Operator   nborm.String
	Comment    nborm.String
	CreateTime nborm.Datetime
	UpdateTime nborm.Datetime
}

//db:qdxg
//tab:mid_student_resume__language_skill
//pk:ResumeID,LanguageSkillID
type MidStudentResumeLanguageSkill struct {
	nborm.Meta
	ResumeID        nborm.Int
	LanguageSkillID nborm.Int
}

//db:qdxg
//tab:mid_student_resume__student_train
//pk:ResumeID,TrainID
type MidStudentResumeStudentTrain struct {
	nborm.Meta
	ResumeID nborm.Int
	TrainID  nborm.Int
}

//db:qdxg
//tab:mid_student_resume__student_honor
//pk:ResumeID,HonorID
type MidStudentResumeStudentHonor struct {
	nborm.Meta
	ResumeID nborm.Int
	HonorID  nborm.Int
}

//db:qdxg
//tab:mid_student_resume__student_experience
//pk:ResumeID,ExperienceID
type MidStudentResumeStudentExperience struct {
	nborm.Meta
	ResumeID     nborm.Int
	ExperienceID nborm.Int
}

//db:qdxg
//tab:mid_student_resume__student_skill
//pk:ResumeID,SkillID
type MidStudentResumeStudentSkill struct {
	nborm.Meta
	ResumeID nborm.Int
	SkillID  nborm.Int
}

//db:qdxg
//tab:student_train
//pk:ID
type StudentTrain struct {
	nborm.Meta
	ID              nborm.Int `auto_increment:"true"`
	IntelUserCode   nborm.String
	InstitutionName nborm.String
	StartDate       nborm.Date
	EndDate         nborm.Date
	Degree          nborm.String
	Major           nborm.String
	Description     nborm.String
	Status          nborm.String
	CreateTime      nborm.Datetime
	UpdateTime      nborm.Datetime
	StudentResume   *StudentResume `rel:"ID->MidStudentResumeStudentTrain.TrainID->MidStudentResumeStudentTrain.ResumeID->ID"`
}

//db:qdxg
//tab:student_honor
//pk:ID
type StudentHonor struct {
	nborm.Meta
	ID            nborm.Int `auto_increment:"true"`
	IntelUserCode nborm.String
	Name          nborm.String
	Description   nborm.String
	GrantDate     nborm.Date
	Status        nborm.String
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentHonor.HonorID->MidStudentResumeStudentHonor.ResumeID->ID"`
}

//db:qdxg
//tab:student_experience
//pk:ID
type StudentExperience struct {
	nborm.Meta
	ID            nborm.Int `auto_increment:"true"`
	IntelUserCode nborm.String
	CompanyName   nborm.String
	StartDate     nborm.Date
	EndDate       nborm.Date
	SectorID      nborm.Int
	Position      nborm.String
	Description   nborm.String
	Status        nborm.String
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentExperience.ExperienceID->MidStudentResumeStudentExperience.ResumeID->ID"`
}

//db:qdxg
//tab:student_skill
//pk:ID
type StudentSkill struct {
	nborm.Meta
	ID            nborm.Int `auto_increment:"true"`
	IntelUserCode nborm.String
	Name          nborm.String
	Description   nborm.String
	Status        nborm.String
	CreateTime    nborm.Datetime
	UpdateTime    nborm.Datetime
	StudentResume *StudentResume `rel:"ID->MidStudentResumeStudentSkill.SkillID->MidStudentResumeStudentSkill.ResumeID->ID"`
}

//db:qdxg
//tab:student_resume
//pk:ID
type StudentResume struct {
	nborm.Meta
	ID                  nborm.Int              `col:"ID" auto_increment:"true"`
	IntelUserCode       nborm.String           `col:"IntelUserCode"`
	Introduction        nborm.String           `col:"Introduction"`
	CreateTime          nborm.Datetime         `col:"CreateTime"`
	UpdateTime          nborm.Datetime         `col:"UpdateTime"`
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
	nborm.Meta
	ID           nborm.Int `auto_increment:"true"`
	EnterpriseID nborm.Int
	Info         nborm.String
}

//db:qdxg
//tab:enterprise_job_snapshot
//pk:ID
type EnterpriseJobSnapshot struct {
	nborm.Meta
	ID    nborm.Int `auto_increment:"true"`
	JobID nborm.Int
	Into  nborm.String
}

//db:qdxg
//tab:student_resume_snapshot
//pk:ID
type StudentResumeSnapshot struct {
	nborm.Meta
	ID       nborm.Int `auto_increment:"true"`
	ResumeID nborm.Int
	Info     nborm.String
}
