package model

import "errors"

type Error struct {
	Err           error
	IsResponseMsg bool
}

func NewError(err string, arg ...bool) *Error {
	e := &Error{Err: errors.New(err)}
	if len(arg) > 0 {
		e.IsResponseMsg = arg[0]
	}
	return e
}
