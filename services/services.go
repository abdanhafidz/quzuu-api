package services

import "github.com/quzuu-be/lib"

type ServiceResult[T any] struct {
	Result    T
	Exception lib.Exception
	Error     error
}
