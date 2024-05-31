package models

type QuestionsResponse struct {
	MCQuestions          interface{}
	SAQuestions          interface{}
	EssayQuestions       interface{}
	InteractiveQuestions interface{}
}

type Duration struct {
	Hour int
	Min  int
	Sec  int
}
type ExamDataResponse struct {
	Progress  interface{}
	RemTime   *Duration
	Questions interface{}
}

type EventResponse struct {
	Data           *Events
	RegisterStatus int
}
