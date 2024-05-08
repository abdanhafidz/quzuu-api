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
