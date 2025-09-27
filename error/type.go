package error

import "github.com/rantool-team/go-error/context"

type Error struct {
	Message     string
	Description string
	Context     context.Context
	StatusCode  int
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Emit() {

}
