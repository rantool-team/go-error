package errorslist

import (
	goerror "github.com/rantool-team/go-error"
	"github.com/rantool-team/go-error/language"
)

func CreateErrorOnErrorsListAndReturnId(msg string, description string) int {
	id, _ := CreateErrorOnErrorsListAndReturnIdAndError(msg, description)

	return id
}

func CreateErrorInSetOfLanguagesOnErrorsListAndReturnId(msgSet language.MessageSet, descriptionSet language.MessageSet) int {
	id, _ := CreateErrorInSetOfLanguagesOnErrorsListAndReturnIdAndError(msgSet, descriptionSet)

	return id
}

func CreateErrorOnErrorsListAndReturnIdAndError(msg string, description string) (int, goerror.Error) {
	err := goerror.CreateError(msg, description)

	return addNewIdWithNewError(err), err
}

func CreateErrorInSetOfLanguagesOnErrorsListAndReturnIdAndError(msgSet language.MessageSet, descriptionSet language.MessageSet) (int, goerror.Error) {
	err := goerror.CreateErrorInMultiplesLanguage(msgSet, descriptionSet)

	return addNewIdWithNewError(err), err
}

func addNewIdWithNewError(err goerror.Error) int {
	id := getLastId()

	errorsList[id] = err

	return id
}

func getLastId() int {
	return lastId + 1
}
