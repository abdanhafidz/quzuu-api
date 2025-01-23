package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type EventListRequest struct {
	PerPage    int    `json:"per_page"`
	PageNumber int    `json:"page_number"`
	Filter     string `json:"filter"`
}

type EventDetailRequest struct {
	IdEvent int `json:"id_event"`
}

type ProblemSetRequest struct {
	IdEvent int `json:"id_event"`
}

type ExamRequest struct {
	IdEvent      int `json:"id_event"`
	IdProblemSet int `json:"id_problem_set"`
}

type RegisterEventRequest struct {
	IdEvent   int    `json:"id_event"`
	EventCode string `json:"event_code"`
}

type QuestionRequest struct {
	IdProblemSet int `json:"id_problem_set"`
}

type SubmitRequest struct {
	IdEvent      int `json:"id_event"`
	IdProblemSet int `json:"id_problem_set"`
}
