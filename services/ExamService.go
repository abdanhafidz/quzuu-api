package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

type ExamTimer struct {
	DueTime  time.Time
	Duration time.Duration
}

func GetExamTimer(id_problem_set int) ExamTimer {
	problemSet, _ := repositories.GetProblemSetDetail(id_problem_set)
	CurrentTime := time.Now()
	dueTime := CurrentTime.Add(problemSet.Duration)

	return ExamTimer{
		DueTime:  dueTime,
		Duration: problemSet.Duration,
	}
}

func CheckUserExam(id_event int, id_account int, id_problem_set int) (data bool, status string, err error) {
	_, dbResult := repositories.GetResult(id_event, id_account, id_problem_set)
	countResult, errResult := middleware.RecordCheck(dbResult)
	if countResult != "no-record" {
		return false, "exam-submitted", errResult
	}
	CurrentTime := time.Now()
	progressCreated, progressRow := repositories.GetProgress(id_event, id_account, id_problem_set)
	statusProgress, _ := middleware.RecordCheck(progressRow)
	var examTimer ExamTimer
	if statusProgress == "no-record" {
		examTimer = GetExamTimer(id_problem_set)
	} else {
		examTimer.DueTime = progressCreated.DueAt
	}
	if CurrentTime.After(examTimer.DueTime) {
		return false, "exam-finished", errResult
	}
	return true, "ready", nil
}
func CheckStatusExam(id_event int, id_account int, id_problem_set int) (data bool, status string, err error) {
	statusRole, errRole := EventRoleCheck(id_event, id_account)
	fmt.Println("status role", statusRole)
	if statusRole == "unauthorized" {
		return false, "unauthorized", errRole
	} else {
		eventStatus, _, errEventStatus := GetEventStatus(id_event, id_account)
		fmt.Println("Event status", eventStatus)
		err = errors.Join(errRole, errEventStatus)
		if (statusRole == "unregistered" || statusRole == "registered") && eventStatus == "notStart" {
			// Sudah daftar atau belum daftar tapi eventBelum Mulai ya gaboleh bikin progress / lihat soal
			return false, "notStart", errEventStatus
		} else if statusRole == "unregistered" && eventStatus == "OnGoing" {
			// Belum daftar ya suruh daftar dulu
			return false, "unregistered", errEventStatus
		} else if statusRole == "registered" && eventStatus == "OnGoing" {
			return true, "ready", errEventStatus
			// Keregister, OnGoing, waktu dia abis
		} else if statusRole == "registered" && eventStatus == "finish" {
			return false, "Time-Out", err
		} else {
			return false, eventStatus, err
		}
	}

}
func ExamService(id_event int, id_account int, id_problem_set int) (data interface{}, status string, err error) {
	CekStatus, ExamStatus, CekExamErr := CheckStatusExam(id_event, id_account, id_problem_set)
	CekUserExam, UserExamStatus, UserExamErr := CheckUserExam(id_event, id_account, id_problem_set)
	fmt.Println("Exam Status", ExamStatus, "User Status", UserExamStatus)
	fmt.Println("Boolean Exam Status", CekStatus, "Boolean UserExam Status", CekUserExam)
	if CekStatus && CekUserExam {
		// Udah daftar dan event OnGoing
		progress, progressRow := repositories.GetProgress(id_event, id_account, id_problem_set)
		statusProgress, _ := middleware.RecordCheck(progressRow)
		status = statusProgress
		QuestionsData, _ := repositories.GetQuestions(id_problem_set)
		ansArray := repositories.CastAnswerFrame(id_problem_set)
		fmt.Println(statusProgress, progressRow.RowsAffected)
		if status == "no-record" {
			// Kalau belum mulai mengerjakan / first assignment
			examTimer := GetExamTimer(id_problem_set)

			// var convertedDueTime time.Time
			// convertedDueTime = dueTime.(time.Time)
			_, progressRowCreated := repositories.CreateProgress(id_event, id_account, id_problem_set, examTimer.DueTime, ansArray)
			progressCreated, _ := repositories.GetProgress(id_event, id_account, id_problem_set)
			statusProgressCreated, errProgressCreated := middleware.RecordCheck(progressRowCreated)
			err = errors.Join(err, errProgressCreated)
			if statusProgressCreated == "ok" {
				data = models.ExamDataResponse{Progress: progressCreated, Questions: &QuestionsData,
					RemTime: &models.Duration{
						Hour: int(examTimer.Duration.Hours()),
						Min:  int(examTimer.Duration.Minutes()),
						Sec:  int(examTimer.Duration.Seconds())}}
				return data, statusProgressCreated, err
			}
		} else {
			// kalau udah mulai ngerjain sebelumnya, lanjutin dari waktu sebelumnya
			currentTime := time.Now()
			hr, min, sec := middleware.DiffTime(currentTime, progress.DueAt)
			if hr > 0 {
				hr = 0
			}
			if min > 0 {
				min = 0
			}
			if sec > 0 {
				sec = 0
			}
			data = models.ExamDataResponse{Progress: progress, Questions: &QuestionsData, RemTime: &models.Duration{Hour: hr * -1, Min: min * -1, Sec: sec * -1}}
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

func addCorrect(Correct *int) {
	*Correct++
}

func addIncorrect(InCorrect *int) {
	*InCorrect++
}

func addEmpty(Empty *int) {
	*Empty++
}

func addScore(Score *float64, weight float64) {
	*Score += weight
}

type weight struct {
	c_weight float64
	i_weight float64
	e_weight float64
}

func (result *result) CheckAnswer(answer string, answer_key string, q_type string, weight weight) {
	if answer == answer_key && q_type != "essay" {
		addCorrect(&result.Correct)
		addScore(&result.Score, weight.c_weight)
	} else if answer == "0" || answer == "null" {
		addEmpty(&result.Empty)
		addScore(&result.Score, weight.e_weight)
	} else if answer != answer_key && q_type != "essay" {
		addIncorrect(&result.Incorrect)
		addScore(&result.Score, weight.i_weight)
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
		userAnswers, _ := userProgress.Answers.Value()
		for _, UserAnswer := range userAnswers.([]any) {
			j := 0
			convertedUserAnswer := UserAnswer.([]any)
			ans_len := len(convertedUserAnswer)
			for _, ParseAnswer := range convertedUserAnswer {
				var Weight weight
				Weight.c_weight = float64(Questions[i].CorrMark) / float64(ans_len)
				Weight.i_weight = float64(Questions[i].IncorrMark) / float64(ans_len)
				Weight.e_weight = float64(Questions[i].NullMark) / float64(ans_len)
				var res = &Result
				res.CheckAnswer(ParseAnswer.(string), Questions[i].AnsKey[j], Questions[i].Type, Weight)
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
	// Waktu Mulai , Due Time
	// Due Time - Waktu Mulai
	return CekStatus, ExamStatus, CekExamErr
}
