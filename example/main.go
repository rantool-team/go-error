package main

import goerror "github.com/rantool-team/go-error"

func main() {
	e := goerror.CreateError("The world is ungly", "none")
	e.SetStatusCode(220)
	e.Emit()
}
