package main

import (
	goerror "github.com/rantool-team/go-error"
	"github.com/rantool-team/go-error/language"
)

func main() {
	goerror.SetDefaultLanguage(language.PORTUGUESE_STRING)
	e := goerror.CreateError("O mundo é ...", "none")
	e.SetStatusCode(220)
	e.Emit()
}
