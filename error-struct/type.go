package errorstruct

import (
	"strconv"

	"github.com/rantool-team/go-error/context"
	"github.com/rantool-team/go-error/language"
)

var PREFIX_ERROR_APPEAR = "\n////////////////ERROR////////////////\n"
var SUFIX_ERROR_APPEAR = "\n/////////////////////////////////////\n"

var PREFIX_CONTEXT_APPEAR = "\n---------------------------------\n"
var SUFIX_CONTEXT_APPEAR = "\n---------------------------------\n"

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
	HasLocalDefined   bool
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
	res += e.retStatusCodeIfNotZero()
	res += "\n"
	res += e.getContextString()
	res += e.GetDescription()
	res += e.getLocalDefinition()
	res += SUFIX_ERROR_APPEAR

	return res
}

func (e Error) retStatusCodeIfNotZero() string {
	if e.StatusCode != 0 {
		nameStatusCode := StatusCodeName.GetMessage(e.Context.Language)
		statusInStr := strconv.Itoa(e.GetStatusCode())
		return "\n" + nameStatusCode + statusInStr
	}

	return ""
}

func (e Error) getContextString() string {
	if e.Context.Description != "" && e.Context.LocalDescription != "" {
		res := PREFIX_CONTEXT_APPEAR
		res += e.Context.Description
		res += "\n"
		res += e.Context.LocalDescription
		res += SUFIX_CONTEXT_APPEAR
		return res
	}

	return ""
}

func (e Error) getErrorMessage() string {
	if e.HasMessageSet {
		return e.MessageSet.GetMessage(e.Context.Language)
	}

	return e.Message
}

func (e Error) GetDescription() string {
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

func (e *Error) AddMessageInOtherLanguage(languageName string, message string) {
	e.HasMessageSet = true
	e.MessageSet.SetMessage(languageName, message)
}

func (e *Error) AddDescriptionInOtherLanguage(languageName string, description string) {
	e.HasDescriptionSet = true
	e.DescriptionSet.SetMessage(languageName, description)
}

func (e *Error) SetLanguage(language string) {
	e.Context.Language = language
}

func (e *Error) SetLocal(local context.Local) {
	e.HasLocalDefined = true
	e.Context.Local = local
}
