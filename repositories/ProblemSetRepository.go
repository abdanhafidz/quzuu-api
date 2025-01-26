package repositories

import (
	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetProblemSet(id_event int, id_account int) (res *[]models.ProblemSet, problemSets *gorm.DB) {
	var e []models.ProblemSet
	problemSets = db.Raw("SELECT problem_sets.id_problem_set,problem_sets.title,problem_sets.duration,problem_sets.randomize FROM problem_sets INNER JOIN problem_sets_assign ON problem_sets_assign.id_problem_set = problem_sets.id_problem_set WHERE problem_sets_assign.id_event = ? ", id_event).Find(&e)

	return &e, problemSets
}

func GetProblemSetDetail(id_problem_set int) (res *models.ProblemSet, problemSets *gorm.DB) {
	var e models.ProblemSet
	problemSets = db.Where("id_problem_set = ?", id_problem_set).Find(&e)
	return &e, problemSets
}
