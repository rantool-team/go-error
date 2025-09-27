package goerror

import "github.com/rantool-team/go-error/context"

type Error interface {
	Error() string
	Emit()
	EmitInLanguage()

	GetContext(localDescription string, description string) context.Context
	SetContext()

	HasError() bool

	GetJSONFormatOfError() string

	GetStatusCode() int
	SetStatusCode(statusCode int)

	GetDescription() string
}
