package error

import (
	"github.com/rantool-team/go-error/context"
	"github.com/rantool-team/go-error/language"
)

var PREFIX_ERROR_APPEAR = "\n////////////////ERROR////////////////\n"
var SUFIX_ERROR_APPEAR = "\n/////////////////////////////////////\n"

type Error struct {
	hasError          bool
	Message           string
	HasMessageSet     bool
	MessageSet        language.MessageSet
	Description       string
	HasDescriptionSet bool
	DescriptionSet    language.MessageSet
	Context           context.Context
	StatusCode        int
}

func (e Error) Error() string {
	return e.getErrorMessage()
}

func (e Error) Emit() {
	panic(e.montarErrorFormat())
}

func (e Error) montarErrorFormat() string {
	res := PREFIX_ERROR_APPEAR
	res += e.getErrorMessage()
	res += "\n"
	res += e.getDescription()
	res += SUFIX_ERROR_APPEAR

	return res
}

func (e Error) getErrorMessage() string {
	if e.HasMessageSet {
		return e.MessageSet.GetMessage(e.Context.Language)
	}

	return e.Message
}

func (e Error) getDescription() string {
	if e.HasDescriptionSet {
		return e.DescriptionSet.GetMessage(e.Context.Language)
	}

	return e.Description
}

func (e Error) GetContext() context.Context {
	return e.Context
}

func (e *Error) SetContext(localDescription string, description string, language string) {
	e.Context = context.Context{
		Description:      description,
		LocalDescription: localDescription,
		Language:         language,
	}
}

func (e *Error) HasError() bool {
	return e.hasError
}

func (e *Error) RemoveError() {
	e.hasError = false
}

func (e *Error) EnableError() {
	e.hasError = true
}

func (e Error) GetStatusCode() int {
	return e.StatusCode
}

func (e *Error) SetStatusCode(statusCode int) {
	e.StatusCode = statusCode
}
