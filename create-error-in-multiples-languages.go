package goerror

import "github.com/rantool-team/go-error/language"

func CreateErrorInMultiplesLanguage(messagesSet language.MessageSet, descriptionSet language.MessageSet) Error {
	err := CreateError(messagesSet.GetMessageInDefaultLanguage(), descriptionSet.GetMessageInDefaultLanguage())
	messagesSet.AppendOnAllLanguages(err.AddMessageInOtherLanguage)
	descriptionSet.AppendOnAllLanguages(err.AddDescriptionInOtherLanguage)
	return err
}
