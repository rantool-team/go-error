# Uso

O `go-error` é um pacote que disponibiliza uma interface de erro ([`Error`](./error.go)) que pode ser facilmente assimilada à interface `error` padrão do Go. Para gerar um erro, use o método [`CreateError`](./create-error.go):

```go
err := goerror.CreateError("MENSAGEM", "DESCRICAO")
```

Isso retornará um erro que implementa a seguinte interface:

```go
type Error interface {
	Error() string // Retorna a mensagem de erro
	Emit()         // Emite o erro através de um panic

	GetContext() context.Context                                           // Retorna o contexto
	SetContext(localDescription string, description string, language string) // Altera o contexto do erro

	HasError() bool // Retorna true se houver erro e false caso contrário

	GetJSONFormatOfError() string // Retorna uma versão JSON do erro

	GetStatusCode() int               // Retorna o código de status
	SetStatusCode(statusCode int) // Altera o código de status

	GetDescription() string // Obtém a descrição no idioma correto

	AddMessageInOtherLanguage(languageName string, message string)       // Adiciona um novo idioma à mensagem de erro
	AddDescriptionInOtherLanguage(languageName string, description string) // Adiciona um novo idioma à descrição do erro

	SetLanguage(language string) // Altera para outro idioma, que deve ser demonstrado no erro
}
```

O contexto do erro é definido pela struct [`Context`](./context/context.go).