package goerror

import "github.com/rantool-team/go-error/language"

func SetDefaultLanguage(NewDefaultLanguage string) {
	language.DefaultLanguage = NewDefaultLanguage
}
