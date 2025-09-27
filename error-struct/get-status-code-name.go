package errorstruct

import "github.com/rantool-team/go-error/language"

var StatusCodeName language.MessageSet

func AddStatusCodeMultiples(statusCode language.MessageSet) {
	StatusCodeName = statusCode
}
