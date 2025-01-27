package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type ProblemSetAssign struct {
	IDProblemSetAssign uint `gorm:"primaryKey" json:"id_problem_set_assign"`
	IDEvent            uint `json:"id_event"`
	IDProblemSet       uint `json:"id_problem_set"`
}

type Announcement struct {
	IDAnnouncement uint      `gorm:"primaryKey" json:"id_announcement"`
	Title          string    `json:"title"`
	CreatedAt      time.Time `json:"created_at"`
	Message        string    `json:"message"`
	Publisher      string    `json:"publisher"`
	IDEvent        uint      `json:"id_event"`
}

type Account struct {
	IDAccount   uint      `gorm:"primaryKey" json:"id_account"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber int       `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type Questions struct {
	IDQuestion   uint     `gorm:"primaryKey" json:"id_question"`
	Type         string   `json:"type"` //MultChoices, ShortAns, Essay, IntPuzzle, IntType
	Question     string   `json:"question"`
	Options      []string `gorm:"type:text[]" json:"options"`
	AnsKey       []string `gorm:"type:text[]" json:"ans_key"`
	CorrMark     float64  `json:"corr_mark"`
	IncorrMark   float64  `json:"incorr_mark"`
	NullMark     float64  `json:"null_mark"`
	IDProblemSet uint     `json:"id_problem_set"`
}

// x = 1, x = 2, x = 3, x = 4
// x++, x--, x+=1, x**, x/=1
type ProblemSet struct {
	IDProblemSet uint          `gorm:"primaryKey" json:"id_problem_set"`
	Title        string        `json:"title"`
	Duration     time.Duration `json:"duration"`
	Randomize    uint          `json:"randomize"`
	MC_Count     uint          `json:"mc_count"`
	SA_Count     uint          `json:"sa_count"`
	Essay_Count  uint          `json:"essay_count"`
}

type AccountDetails struct {
	IDDetail    uint      `gorm:"primaryKey" json:"id_detail"`
	IDAccount   uint      `json:"id_account"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Institution string    `json:"institution"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type Events struct {
	IDEvent    uint      `gorm:"primaryKey" json:"id_event"`
	Title      string    `json:"title"`
	StartEvent time.Time `json:"start_event"`
	EndEvent   time.Time `json:"end_event"`
	SID        string    `json:"sid"`
	Public     string    `json:"public"`
}

type EventAssign struct {
	IDAssign   uint      `gorm:"primaryKey" json:"id_assign"`
	IDAccount  uint      `json:"id_account"`
	IDEvent    uint      `json:"id_event"`
	AssignedAt time.Time `json:"assigned_at"`
}

type Result struct {
	IDResult      uint      `gorm:"primaryKey" json:"id_result"`
	IDAccount     uint      `json:"id_account"`
	IDEvent       uint      `json:"id_event"`
	IDProblemSet  uint      `json:"id_problem_set"`
	IDProgress    uint      `json:"id_progress"`
	FinishTime    time.Time `json:"finish_time"`
	Correct       uint      `json:"correct"`
	Incorrect     uint      `json:"incorrect"`
	Empty         uint      `json:"empty"`
	OnCorrection  uint      `json:"on_correction"`
	ManualScoring float64   `json:"manual_scoring"`
	MCScore       float64   `json:"mc_score"`
	ManualScore   float64   `json:"manual_score"`
	FinalScore    float64   `json:"final_score"`
}

type ExamProgress struct {
	IDProgress     uint      `gorm:"primaryKey" json:"id_progress"`
	IDAccount      uint      `json:"id_account"`
	IDEvent        uint      `json:"id_event"`
	IDProblemSet   uint      `json:"id_problem_set"`
	CreatedAt      time.Time `json:"created_at"`
	DueAt          time.Time `json:"due_at"`
	QuestionsOrder []string  `gorm:"type:text[]" json:"questions_order"`
	Answers        any       `gorm:"type:jsonb" json:"answers"`
}
type ExamProgress_Result struct {
	IDProgress     uint           `gorm:"primaryKey" json:"id_progress"`
	IDAccount      uint           `json:"id_account"`
	IDEvent        uint           `json:"id_event"`
	IDProblemSet   uint           `json:"id_problem_set"`
	CreatedAt      time.Time      `json:"created_at"`
	DueAt          time.Time      `json:"due_at"`
	QuestionsOrder []string       `gorm:"type:text[]" json:"questions_order"`
	Answers        postgres.Jsonb `gorm:"type:jsonb" json:"answers"`
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
