package repositories

import (
	"time"

	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetProgress(id_event int, id_account int, id_problem_set int) (res *models.ExamProgress, progress *gorm.DB) {
	var e models.ExamProgress
	progress = db.Where("id_account = ? AND id_problem_set = ? AND id_event = ?", id_account, id_problem_set, id_event).Find(&e)
	return &e, progress
}

type MCOnGoingExam struct {
	IDMCQuestion uint
	Question     int64
	Opt1         string
	Opt2         string
	Opt3         string
	Opt4         string
	Opt5         string
	IDProblemSet uint
}
type SAOnGoingExam struct {
	IDSAQuestion uint
	Question     string
	IDProblemSet uint
}

type EssayOnGoingExam struct {
	IDEssayQuestion uint `gorm:"primaryKey"`
	Question        int64
	IDProblemSet    uint
}

type InteractiveOnGoingExam struct {
	IDIntvQuestion uint
	Question       string
	IDProblemSet   uint
}

func GetMCQuestions(id_problem_set int, finished bool) (res interface{}, MCQuestions *gorm.DB) {

	if !finished {
		var e []MCOnGoingExam
		MCQuestions = db.Table("mc_questions").Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	} else {
		var e models.MCQuestion
		MCQuestions = db.Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	}
}

func GetSAQuestions(id_problem_set int, finished bool) (res interface{}, MCQuestions *gorm.DB) {

	if !finished {
		var e []SAOnGoingExam
		MCQuestions = db.Table("sa_quetsions").Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	} else {
		var e []models.ShortAnsQuestion
		MCQuestions = db.Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	}
}

func GetEssayQuestions(id_problem_set int, finished bool) (res interface{}, MCQuestions *gorm.DB) {

	if !finished {
		var e []EssayOnGoingExam
		MCQuestions = db.Table("essay_questions").Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	} else {
		var e []models.EssayQuestion
		MCQuestions = db.Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	}
}

func GetInteractiveQuestions(id_problem_set int, finished bool) (res interface{}, MCQuestions *gorm.DB) {

	if !finished {
		var e []InteractiveOnGoingExam
		MCQuestions = db.Table("interactive_questions").Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	} else {
		var e []models.InteractiveQuestion
		MCQuestions = db.Where("id_problem_set = ? ", id_problem_set).Find(&e)
		return &e, MCQuestions
	}
}
func CreateProgress(id_event int, id_account int, id_problem_set int, due time.Time) (res interface{}, createProgress *gorm.DB) {
	e := &models.ExamProgress{
		IDAccount:    uint(id_account),
		IDEvent:      uint(id_event),
		IDProblemSet: uint(id_problem_set),
		CreatedAt:    time.Now(),
		DueAt:        due,
	}
	createProgress = db.Create(&e)
	return &e, createProgress
}

func GetResult(id_event int, id_account int, id_problem_set int) (res *models.Result, result *gorm.DB) {
	var e models.Result
	result = db.Where("id_event = ? AND id_account = ? AND id_problem_set = ?", id_event, id_account, id_problem_set).Find(&e)
	return &e, result
}

// func CreateResult(id_event int, id_account int, id_problem_set int, id_progress int) (res *models.Result, result *gorm.DB) {
// 	var e = &models.Result{}
// 	result = db.Create(&e)
// 	return &e, result
// }
