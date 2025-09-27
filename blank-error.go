package goerror

import errorstruct "github.com/rantool-team/go-error/error-struct"

func CreateBlankError() Error {
	err := errorstruct.Error{}

	return &err
}
