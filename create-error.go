package goerror

import (
	errorstruct "github.com/rantool-team/go-error/error-struct"
	"github.com/rantool-team/go-error/language"
)

func CreateError(messageInEnglish string, description string) Error {
	errorstruct.AddStatusCodeMultiples(StatusCodeNames)

	err := errorstruct.Error{
		Message:     messageInEnglish,
		Description: description,
	}

	err.SetLanguage(language.DefaultLanguage)

	err.EnableError()

	return &err
}
