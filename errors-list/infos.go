package errorslist

import goerror "github.com/rantool-team/go-error"

var lastId = 0

var errorsList = make(map[int]goerror.Error)
