package errorslist

import (
	goerror "github.com/rantool-team/go-error"
	"github.com/rantool-team/go-error/language"
)

func CreateErrorOnErrorsListAndReturnId(msg string, description string) int {
	err := goerror.CreateError(msg, description)

	return addNewIdWithNewError(err)
}

func CreateErrorInSetOfLanguagesOnErrorsListAndReturnId(msgSet language.MessageSet, descriptionSet language.MessageSet) int {
	err := goerror.CreateErrorInMultiplesLanguage(msgSet, descriptionSet)

	return addNewIdWithNewError(err)
}

func addNewIdWithNewError(err goerror.Error) int {
	id := getLastId()

	errorsList[id] = err

	return id
}

func getLastId() int {
	return lastId + 1
}
