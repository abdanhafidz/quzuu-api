package models

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Phone    int    `form:"phone"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type EventListRequest struct {
	PerPage    int    `form:"per_page"`
	PageNumber int    `form:"page_number"`
	Filter     string `form:"filter"`
}

type EventDetailRequest struct {
	IDEvent int `form:"id_event"`
}

type ProblemSetRequest struct {
	IDEvent int `form:"id_event"`
}

type ExamRequest struct {
	IDEvent      int `form:"id_event"`
	IDProblemSet int `form:"id_problem_set"`
}

type RegisterEventRequest struct {
	IDEvent   int    `form:"id_event"`
	EventCode string `form:"event_code"`
}

type QuestionRequest struct {
	IDProblemSet int `form:"id_problem_set"`
}

type SubmitRequest struct {
	IDEvent      int `form:"id_event"`
	IDProblemSet int `form:"id_problem_set"`
}
