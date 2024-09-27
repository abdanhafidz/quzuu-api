package services

import (
	"errors"
	"time"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func CheckUserExam(id_event int, id_account int, id_problem_set int) (data bool, status string, err error) {
	_, dbResult := repositories.GetResult(id_event, id_account, id_problem_set)
	countResult, errResult := middleware.RecordCheck(dbResult)
	if countResult != "no-record" {
		return false, "exam-submitted", errResult
	}
	problemSet, _ := repositories.GetProblemSetDetail(id_problem_set)
	CurrentTime := time.Now()
	var hr int
	var min int
	var sec int
	go func() {
		hr = problemSet.Duration.Hour()
		min = problemSet.Duration.Minute()
		sec = problemSet.Duration.Second()
	}()
	dueTime := CurrentTime.Add(time.Hour*time.Duration(hr) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
	if CurrentTime.After(dueTime) {
		return false, "exam-finished", errResult
	}
	return true, "ready", nil
}
func CheckStatusExam(id_event int, id_account int, id_problem_set int) (data bool, status string, err error) {
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
			return true, "ready", errEventStatus
		} else if statusRole == "registered" && eventStatus == "finish" {
			return false, "Time-Out", err
		}
	}
	return false, "error", errRole
}
func ExamService(id_event int, id_account int, id_problem_set int) (data interface{}, status string, err error) {
	CekStatus, ExamStatus, CekExamErr := CheckStatusExam(id_event, id_account, id_problem_set)
	CekUserExam, UserExamStatus, UserExamErr := CheckUserExam(id_event, id_account, id_problem_set)
	if CekStatus && CekUserExam {
		// Udah daftar dan event OnGoing
		progress, progressRow := repositories.GetProgress(id_event, id_account, id_problem_set)
		statusProgress, _ := middleware.RecordCheck(progressRow)
		status = statusProgress
		QuestionsData, _ := repositories.GetQuestions(id_problem_set)
		ansArray := repositories.CastAnswerFrame(id_problem_set)
		if status == "no-record" {
			// Kalau belum mulai mengerjakan / first assignment
			problemSet, _ := repositories.GetProblemSetDetail(id_problem_set)
			currentTime := time.Now()
			var hr int
			var min int
			var sec int
			go func() {
				hr = problemSet.Duration.Hour()
				min = problemSet.Duration.Minute()
				sec = problemSet.Duration.Second()
			}()
			dueTime := currentTime.Add(time.Hour*time.Duration(hr) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
			progressCreated, progressRowCreated := repositories.CreateProgress(id_event, id_account, id_problem_set, dueTime, ansArray)
			statusProgressCreated, errProgressCreated := middleware.RecordCheck(progressRowCreated)
			err = errors.Join(err, errProgressCreated)
			if statusProgressCreated == "ok" {
				data = models.ExamDataResponse{Progress: progressCreated, Questions: &QuestionsData, RemTime: &models.Duration{Hour: hr, Min: min, Sec: sec}}
				return data, statusProgressCreated, err
			}
		} else {
			// kalau udah mulai ngerjain sebelumnya, lanjutin dari waktu sebelumnya
			currentTime := time.Now()
			hr, min, sec := middleware.DiffTime(currentTime, progress.DueAt)
			data = models.ExamDataResponse{Progress: progress, Questions: &QuestionsData, RemTime: &models.Duration{Hour: hr, Min: min, Sec: sec}}
			return data, "ok", err
		}
	} else if UserExamStatus == "exam-submitted" {
		data = false
		status = "exam-submitted"
		err = UserExamErr
		return data, status, err
	} else if UserExamStatus == "exam-finished" {
		data = false
		status = "exam-finished"
		err = UserExamErr
		return data, status, err
	} else {
		data = false
		status = ExamStatus
		err = CekExamErr
		return data, status, err
	}
	return data, status, err
}

type result struct {
	Correct   int
	Incorrect int
	Empty     int
	Score     float64
}

func AddCorrect(Correct *int) {
	*Correct++
}

func AddIncorrect(InCorrect *int) {
	*InCorrect++
}

func AddEmpty(Empty *int) {
	*Empty++
}

func AddScore(Score *float64, weight float64) {
	*Score += weight
}

type weight struct {
	c_weight float64
	i_weight float64
	e_weight float64
}

func (result *result) CheckAnswer(answer string, answer_key string, q_type string, weight weight) {
	if answer == answer_key && q_type != "essay" {
		AddCorrect(&result.Correct)
		AddScore(&result.Score, weight.c_weight)
	} else if answer == "0" || answer == "null" {
		AddEmpty(&result.Empty)
		AddScore(&result.Score, weight.e_weight)
	} else if answer != answer_key && q_type != "essay" {
		AddIncorrect(&result.Incorrect)
		AddScore(&result.Score, weight.i_weight)
	}
}
func SubmitExamService(id_event int, id_account int, id_problem_set int) (data interface{}, status string, err error) {
	CekStatus, ExamStatus, CekExamErr := CheckStatusExam(id_event, id_account, id_problem_set)
	CekUserExam, UserExamStatus, UserExamErr := CheckUserExam(id_event, id_account, id_problem_set)
	if CekStatus && CekUserExam {
		var Result result
		Questions, _ := repositories.GetQuestionsReview(id_problem_set)
		userProgress, _ := repositories.GetProgress(id_event, id_account, id_problem_set)
		i := 0
		for _, UserAnswer := range userProgress.Answers {
			j := 0
			ans_len := len(UserAnswer)
			for _, ParseAnswer := range UserAnswer {
				var Weight weight
				Weight.c_weight = float64(Questions[i].CorrMark) / float64(ans_len)
				Weight.i_weight = float64(Questions[i].IncorrMark) / float64(ans_len)
				Weight.e_weight = float64(Questions[i].NullMark) / float64(ans_len)
				var res = &Result
				res.CheckAnswer(ParseAnswer, Questions[i].AnsKey[j], Questions[i].Type, Weight)
				j++
			}
			i++
		}
		return Result, "ok", err
	} else if UserExamStatus == "exam-submitted" {
		data = false
		status = "exam-submitted"
		err = UserExamErr
		return data, status, err
	} else if UserExamStatus == "exam-finished" {
		data = false
		status = "exam-finished"
		err = UserExamErr
		return data, status, err
	}
	return CekStatus, ExamStatus, CekExamErr
}
