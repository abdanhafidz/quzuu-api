package repositories

import (
	"strings"
	"time"

	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetProgress(id_event int, id_account int, id_problem_set int) (res *models.ExamProgress, progress *gorm.DB) {
	var e models.ExamProgress
	progress = db.Where("id_account = ? AND id_problem_set = ? AND id_event = ?", id_account, id_problem_set, id_event).Find(&e)
	return &e, progress
}

type QuestionsOnGoingExam struct {
	IDQuestion   uint
	Type         string //MultChoices, ShortAns, Essay, IntPuzzle, IntType
	Question     string
	Options      [][]string
	IDProblemSet uint
}

func CastAnswerFrame(id_problem_set int) (ans [][]string) {
	type Lquestions struct {
		cnt_array int
	}
	var e []Lquestions
	db.Raw("SELECT array_length(ans_key, 1) as cnt_array FROM questions WHERE id_problem_set = ? ORDER BY id_question", id_problem_set).Find(&e)
	i := 0
	for _, t := range e {
		ansString := strings.Repeat("0,", t.cnt_array-1) + "0"
		ansArray := strings.Split(ansString, ",")
		ans[i] = ansArray
		i++
	}
	return ans
}

func GetQuestions(id_problem_set int) (res []QuestionsOnGoingExam, MCQuestions *gorm.DB) {
	var e []QuestionsOnGoingExam
	MCQuestions = db.Raw("SELECT * FROM questions WHERE id_problem_set = ? ORDER BY id_question ASC", id_problem_set).Find(&e)
	return e, MCQuestions
}

func GetQuestionsReview(id_problem_set int) (res []models.Questions, MCQuestions *gorm.DB) {
	var e []models.Questions
	MCQuestions = db.Raw("SELECT * FROM questions WHERE id_problem_set = ? ORDER BY id_question ASC", id_problem_set).Find(&e)
	return e, MCQuestions
}

// func GetAnsKey(id_problem_set int) (res []string, MCQuestions *gorm.DB) {

// }

func CreateProgress(id_event int, id_account int, id_problem_set int, due time.Time, ans [][]string) (res interface{}, createProgress *gorm.DB) {
	e := &models.ExamProgress{
		IDAccount:    uint(id_account),
		IDEvent:      uint(id_event),
		IDProblemSet: uint(id_problem_set),
		CreatedAt:    time.Now(),
		DueAt:        due,
		Answers:      ans,
	}
	createProgress = db.Create(&e)
	return &e, createProgress
}

func GetResult(id_event int, id_account int, id_problem_set int) (res *models.Result, result *gorm.DB) {
	var e models.Result
	result = db.Where("id_event = ? AND id_account = ? AND id_problem_set = ?", id_event, id_account, id_problem_set).Find(&e)
	return &e, result
}

func CreateResult(id_event int, id_account int, id_problem_set int, id_progress int, data *models.Result) (res *models.Result, result *gorm.DB) {
	var e = data
	result = db.Create(e)
	return e, result
}
