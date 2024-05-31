package services

import (
	"errors"
	"time"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func ExamService(id_event int, id_account int, id_problem_set int) (data interface{}, status string, err error) {
	statusRole, errRole := EventRoleCheck(id_event, id_account)

	if statusRole == "unauthorized" {
		return false, "unauthorized", errRole
	} else {
		_, eventStatus, errEventStatus := GetEventStatus(id_event, id_account)
		err = errors.Join(errRole, errEventStatus)
		if (statusRole == "unregistered" || statusRole == "registered") && eventStatus == "notStart" {
			// Sudah daftar atau belum daftar tapi eventBelum Mulai ya gaboleh bikin progress / lihat soal
			return false, "notStart", errEventStatus
		} else if statusRole == "unregistered" && eventStatus == "onGoing" {
			// Belum daftar ya suruh daftar dulu
			return false, "unregistered", errEventStatus
		} else if statusRole == "registered" && eventStatus == "onGoing" {
			// Udah daftar dan event OnGoing
			progress, progressRow := repositories.GetProgress(id_event, id_account, id_problem_set)
			status, _ := middleware.RecordCheck(progressRow)
			var Questions *models.QuestionsResponse
			go func() {
				Questions.MCQuestions, _ = repositories.GetMCQuestions(id_problem_set, false)
				Questions.SAQuestions, _ = repositories.GetSAQuestions(id_problem_set, false)
				Questions.EssayQuestions, _ = repositories.GetEssayQuestions(id_problem_set, false)
				Questions.InteractiveQuestions, _ = repositories.GetInteractiveQuestions(id_problem_set, false)
			}()
			if status == "no-record" {
				// Kalau belum mulai mengerjakan / first assignment
				problemSet, _ := repositories.GetProblemSetDetail(id_problem_set)
				currentTime := time.Now()
				hr := problemSet.Duration.Hour()
				min := problemSet.Duration.Minute()
				sec := problemSet.Duration.Second()
				dueTime := currentTime.Add(time.Hour*time.Duration(hr) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
				progressCreated, progressRowCreated := repositories.CreateProgress(id_event, id_account, id_problem_set, dueTime)
				statusProgressCreated, errProgressCreated := middleware.RecordCheck(progressRowCreated)
				err = errors.Join(err, errProgressCreated)
				if statusProgressCreated == "ok" {
					data = models.ExamDataResponse{Progress: progressCreated, Questions: &Questions, RemTime: &models.Duration{Hour: hr, Min: min, Sec: sec}}
					return data, statusProgressCreated, err
				}

			} else {
				// kalau udah mulai ngerjain sebelumnya, lanjutin dari waktu sebelumnya
				currentTime := time.Now()
				hr, min, sec := middleware.DiffTime(currentTime, progress.DueAt)
				data = models.ExamDataResponse{Progress: progress, Questions: &Questions, RemTime: &models.Duration{Hour: hr, Min: min, Sec: sec}}
				return data, "ok", err
			}
		} else {
			return
		}
	}
	return data, status, err
}

// func SubmitExamService(id_event int, id_account int, id_problem_set int) (data interface{}, status string, err error) {

// }
