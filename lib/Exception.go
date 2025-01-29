package lib

type Exception struct {
	Unauthorized             bool
	DataNotFound             bool
	DataDuplicate            bool
	ExamSubmitted            bool
	ExamTimeOut              bool
	EventTimeOut             bool
	EventNotStart            bool
	EventOnGoing             bool
	UserNotRegisteredToEvent bool
	InvalidEventCode         bool
	Message                  string
}

var Exceptions = Exception{}
