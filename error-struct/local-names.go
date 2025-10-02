package errorstruct

import (
	"fmt"

	"github.com/rantool-team/go-error/context"
	"github.com/rantool-team/go-error/language"
)

var LocalDemonstrationLineChar = func(local context.Local) language.MessageSet {
	return language.MessageSet{
		Portuguese: fmt.Sprintf("Linha Inicial: %d, Linha final: %d, Caractere inicial: %d, Caractere final: %d", local.StartLine, local.EndLine, local.StartChar, local.EndChar),
		English:    fmt.Sprintf("Start Line: %d, End line: %d, Initial char: %d, End char: %d", local.StartLine, local.EndLine, local.StartChar, local.EndChar),
	}
}

var LocalDemonstrationFileModPackWorkspace = func(local context.Local) language.MessageSet {
	return language.MessageSet{
		Portuguese: fmt.Sprintf("Arquivo: %s, Pacote: %s, Modulo: %s, Worskpace: %s", local.File, local.Package, local.Module, local.Workspace),
		English:    fmt.Sprintf("File: %s, Package: %s, Module: %s, Worskpace: %s", local.File, local.Package, local.Module, local.Workspace),
	}
}
