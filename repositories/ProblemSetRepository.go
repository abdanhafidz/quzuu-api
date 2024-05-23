package repositories

import (
	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetProblemSet(id_event int, id_account int) (res interface{}, problemSets *gorm.DB) {
	var e []models.ProblemSet
	problemSets = db.Raw("SELECT problem_sets.id_problemset,problem_sets.title,problem_sets.duration,problem_sets.randomize FROM problem_sets INNER JOIN problem_sets_assign ON problem_sets_assign.id_problemset = problem_sets.id_problemset WHERE problem_sets_assign.id_event = ? ", id_event).Find(&e)
	return &e, problemSets

}
