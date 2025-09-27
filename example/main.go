package main

import (
	goerror "github.com/rantool-team/go-error"
)

func main() {
	goerror.SetDefaultLanguage("BOMDIA")
	e := goerror.CreateError("The world is ungly", "none")
	e.SetStatusCode(220)
	e.Emit()
}
