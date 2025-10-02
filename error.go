package goerror

import "github.com/rantool-team/go-error/context"

type Error interface {
	Error() string
	Emit()

	GetContext() context.Context
	SetContext(localDescription string, description string, language string)

	HasError() bool

	GetJSONFormatOfError() string

	GetStatusCode() int
	SetStatusCode(statusCode int)

	GetDescription() string

	AddMessageInOtherLanguage(languageName string, message string)
	AddDescriptionInOtherLanguage(languageName string, description string)

	SetLanguage(language string)

	SetLocal(local context.Local)
	GetJsonLocal() string
	GetLocal() context.Local
}
