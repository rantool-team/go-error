package main

import (
	goerror "github.com/rantool-team/go-error"
)

func main() {
	goerror.AddStatusCodeName("not-a-language", "FJDSJDS: ")
	e := goerror.CreateError("The world is ungly", "none")
	e.SetContext("", "", "not-a-language")
	e.SetStatusCode(220)
	e.Emit()
}
