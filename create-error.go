package goerror

import (
	errorstruct "github.com/rantool-team/go-error/error-struct"
	"github.com/rantool-team/go-error/language"
)

func CreateError(message string, description string) Error {
	errorstruct.AddStatusCodeMultiples(StatusCodeNames)

	err := errorstruct.Error{}

	err.AddMessageInOtherLanguage(language.DefaultLanguage, message)
	err.AddDescriptionInOtherLanguage(language.DefaultLanguage, description)

	err.SetLanguage(language.DefaultLanguage)

	err.EnableError()

	return &err
}
