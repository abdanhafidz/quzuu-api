package models

import (
	"time"
)

type ProblemSetAssign struct {
	IDProblemSetAssign uint `gorm:"primaryKey"`
	IDEvent            uint
	IDProblemSet       uint
}

type Announcement struct {
	IDAnnouncement uint `gorm:"primaryKey"`
	Title          string
	CreatedAt      time.Time
	Message        string
	Publisher      string
	IDEvent        uint
}

type Account struct {
	IDAccount   uint `gorm:"primaryKey"`
	Name        string
	Username    string
	Email       string
	Password    string
	PhoneNumber int
	CreatedAt   time.Time
	DeletedAt   time.Time
}

type MCQuestion struct {
	IDMCQuestion uint `gorm:"primaryKey"`
	Question     int64
	Opt1         string
	Opt2         string
	Opt3         string
	Opt4         string
	Opt5         string
	AnsKey       int64
	CorrMark     float64
	IncorrMark   float64
	NullMark     float64
	IDProblemSet uint
}

type ShortAnsQuestion struct {
	IDSAQuestion uint `gorm:"primaryKey"`
	Question     string
	AnsKey       string
	CorrMark     float64
	IncorrMark   float64
	NullMark     float64
	IDProblemSet uint
}

type EssayQuestion struct {
	IDEssayQuestion uint `gorm:"primaryKey"`
	Question        int64
	Mark            float64
	IDProblemSet    uint
}

type ProblemSet struct {
	IDProblemSet uint `gorm:"primaryKey"`
	Title        string
	Duration     time.Time
	Randomize    uint
}

type AccountDetails struct {
	IDDetail    uint `gorm:"primaryKey"`
	IDAccount   uint
	Province    string
	City        string
	Institution string
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Events struct {
	IDEvent    uint `gorm:"primaryKey"`
	Title      string
	StartEvent string
	EndEvent   string
	SID        string
	Public     string
}

type EventAssign struct {
	IDAssign   uint `gorm:"primaryKey"`
	IDAccount  uint
	IDEvent    uint
	AssignedAt time.Time
}

type Result struct {
	IDResult     uint `gorm:"primaryKey"`
	IDAccount    uint
	IDEvent      uint
	IDProblemSet uint
	IDProgress   uint
	FinishTime   time.Time
	CorrectMC    uint
	IncorrectMC  uint
	NullMC       uint
	CorrectSA    uint
	IncorrectSA  uint
	NullSA       uint
	EssayScoring float64
	MCScore      float64
	SAScore      float64
	EssayScore   float64
	FinalScore   float64
}

type ExamProgress struct {
	IDProgress     uint `gorm:"primaryKey"`
	IDAccount      uint
	IDEvent        uint
	IDProblemSet   uint
	CreatedAt      time.Time
	DueAt          time.Time
	QuestionsOrder string
	Answers        string
}

// Gorm table name settings
func (ProblemSetAssign) TableName() string { return "problem_sets_assign" }
func (Announcement) TableName() string     { return "announcement" }
func (Account) TableName() string          { return "account" }
func (MCQuestion) TableName() string       { return "mc_questions" }
func (ShortAnsQuestion) TableName() string { return "shortans_questions" }
func (EssayQuestion) TableName() string    { return "essay_questions" }
func (ProblemSet) TableName() string       { return "problem_sets" }
func (AccountDetails) TableName() string   { return "account_details" }
func (Events) TableName() string           { return "events" }
func (EventAssign) TableName() string      { return "event_assign" }
func (Result) TableName() string           { return "result" }
func (ExamProgress) TableName() string     { return "exam_progress" }
