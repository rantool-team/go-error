package main

import (
	goerror "github.com/rantool-team/go-error"
	"github.com/rantool-team/go-error/language"
)

func main() {
	goerror.AddStatusCodeName("not-a-language", "FJDSJDS: ")
	e := goerror.CreateError("The world is ungly", "none")
	e.SetContext("", "", "not-a-language")
	e.SetStatusCode(220)
	e.SetLanguage(language.PORTUGUESE_STRING)
	e.AddDescriptionInOtherLanguage(language.PORTUGUESE_STRING, "portuguê")
	e.AddMessageInOtherLanguage(language.PORTUGUESE_STRING, "português")
	e.Emit()
}
