package model

import "fmt"

type ErrorE struct {
	Code       string
	Err        error
	Who        string
	StatusHTTP int
	Data       interface{}
	APIMessage string
	UserID     string
}

func NewError() ErrorE {
	return ErrorE{}
}

func (e *ErrorE) Error() string {
	return fmt.Sprintf("Code: %s, Err: %v, Who: %s, Status: %d, Data: %v, UserID:%s",
		e.Code,
		e.Err,
		e.Who,
		e.StatusHTTP,
		e.Data,
		e.UserID,
	)
}

func (e *ErrorE) HasCode() bool {
	return e.Code != ""
}

func (e *ErrorE) HasStatusHttp() bool {
	return e.StatusHTTP > 0
}

func (e *ErrorE) HasData() bool {
	return e.Data != nil
}
