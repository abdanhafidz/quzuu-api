package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    int    `json:"phone" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EventListRequest struct {
	PerPage    int    `json:"per_page" binding:"required"`
	PageNumber int    `json:"page_number" binding:"required"`
	Filter     string `json:"filter"` // Optional, tidak diberi `required` jika tidak wajib.
}

type EventDetailRequest struct {
	IdEvent int `json:"id_event" binding:"required"`
}

type ProblemSetRequest struct {
	IdEvent int `json:"id_event" binding:"required"`
}

type ExamRequest struct {
	IdEvent      int `json:"id_event" binding:"required"`
	IdProblemSet int `json:"id_problem_set" binding:"required"`
}

type RegisterEventRequest struct {
	IdEvent   int    `json:"id_event" binding:"required"`
	EventCode string `json:"event_code" binding:"required"`
}

type QuestionRequest struct {
	IdProblemSet int `json:"id_problem_set" binding:"required"`
}

type SubmitRequest struct {
	IdEvent      int `json:"id_event" binding:"required"`
	IdProblemSet int `json:"id_problem_set" binding:"required"`
}
