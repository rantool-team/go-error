package goerror

import errorstruct "github.com/rantool-team/go-error/error-struct"

func CreateError(messageInEnglish string, description string) Error {
	errorstruct.AddStatusCodeMultiples(StatusCodeNames)

	err := errorstruct.Error{
		Message:     messageInEnglish,
		Description: description,
	}

	err.EnableError()

	return &err
}
