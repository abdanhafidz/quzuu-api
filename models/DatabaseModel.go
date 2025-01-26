package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
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

type Questions struct {
	IDQuestion   uint   `gorm:"primaryKey"`
	Type         string //MultChoices, ShortAns, Essay, IntPuzzle, IntType
	Question     string
	Options      []string `gorm:"type:text[]"`
	AnsKey       []string `gorm:"type:text[]"`
	CorrMark     float64
	IncorrMark   float64
	NullMark     float64
	IDProblemSet uint
}

// x = 1, x = 2, x = 3, x = 4
// x++, x--, x+=1, x**, x/=1
type ProblemSet struct {
	IDProblemSet uint `gorm:"primaryKey"`
	Title        string
	Duration     time.Duration
	Randomize    uint
	MC_Count     uint
	SA_Count     uint
	Essay_Count  uint
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
	StartEvent time.Time
	EndEvent   time.Time
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
	IDResult      uint `gorm:"primaryKey"`
	IDAccount     uint
	IDEvent       uint
	IDProblemSet  uint
	IDProgress    uint
	FinishTime    time.Time
	Correct       uint
	Incorrect     uint
	Empty         uint
	OnCorrection  uint
	ManualScoring float64
	MCScore       float64
	ManualScore   float64
	FinalScore    float64
}

type ExamProgress struct {
	IDProgress     uint `gorm:"primaryKey"`
	IDAccount      uint
	IDEvent        uint
	IDProblemSet   uint
	CreatedAt      time.Time
	DueAt          time.Time
	QuestionsOrder []string `gorm:"type:text[]"`
	Answers        any      `gorm:"type:jsonb"`
}
type JSONB []interface{}
type ExamProgress_Result struct {
	IDProgress     uint `gorm:"primaryKey"`
	IDAccount      uint
	IDEvent        uint
	IDProblemSet   uint
	CreatedAt      time.Time
	DueAt          time.Time
	QuestionsOrder []string       `gorm:"type:text[]"`
	Answers        postgres.Jsonb `gorm:"type:jsonb"`
}

// Gorm table name settings
func (ProblemSetAssign) TableName() string { return "problem_sets_assign" }
func (Announcement) TableName() string     { return "announcement" }
func (Account) TableName() string          { return "account" }
func (Questions) TableName() string        { return "questions" }
func (ProblemSet) TableName() string       { return "problem_sets" }
func (AccountDetails) TableName() string   { return "account_details" }
func (Events) TableName() string           { return "events" }
func (EventAssign) TableName() string      { return "event_assign" }
func (Result) TableName() string           { return "result" }
func (ExamProgress) TableName() string     { return "exam_progress" }
