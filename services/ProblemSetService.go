package services

import (
	"errors"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/repositories"
)

func ProblemSetService(id_event int, id_account int) (data interface{}, status string, err error) {
	data, problemSets := repositories.GetProblemSet(id_event, id_account)
	statusProblemSet, errProblemSet := middleware.RecordCheck(problemSets)
	statusAssign, errAssign := EventRoleCheck(id_event, id_account)
	err = errors.Join(errProblemSet, errAssign)
	if statusAssign == "unauthorized" {
		return false, "unauthorized", err
	} else {
		return data, statusProblemSet, err
	}

}
