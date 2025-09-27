package goerror

import "github.com/rantool-team/go-error/language"

var StatusCodeNames = language.MessageSet{
	English:    "Status Code: ",
	Portuguese: "Status do código: ",
}

func AddStatusCodeName(languageName string, name string) {
	StatusCodeNames.SetMessage(languageName, name)
}
