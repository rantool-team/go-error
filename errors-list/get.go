package errorslist

import goerror "github.com/rantool-team/go-error"

func GetErrorOnErrorId(errorId int) goerror.Error {
	erro, ok := errorsList[errorId]
	if !ok {
		panicBecauseDontExistId(errorId)
	}

	return erro
}

func panicBecauseDontExistId(id int) {
	panic("You selected an error that not exists. The id " + string(id) + " don't corresponds to no error")
}
