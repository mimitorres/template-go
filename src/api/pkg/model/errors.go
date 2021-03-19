package model

import "fmt"

var (
	ErrNotFound   = fmt.Errorf("resource not found")
	ErrServer     = fmt.Errorf("server error")
	ErrBadRequest = fmt.Errorf("bad request")
)

type CauseList []interface{}

type ApiError struct {
	Status  int       `json:"status"`
	Message string    `json:"message",omitempty`
	Cause   CauseList `json:"cause"`
}
