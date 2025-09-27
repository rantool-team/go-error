# error-go
Em GO existe já por padrão a interface error. O problema, na minha opnião, é que ela é muito limitada em projetos maiores. Por exemplo, para mim falta recursos como:
- Definição dos contexto: Seria interessante ter uma ferramenta de contexto, para indicar o contexto onde erro aparece.
- Error handling: O formato err != nil, é interessante, porém pouco visual
- Talvez seja inutil, mas em projetos maiores, seria interessante ter uma ferramenta para pluralizar os erros em diversas línguas. 
- Retorno em JSON dos erros, seria também interessante, já ter nativamente, uma ferramente que controla os erros e retorna um JSON sobre o assunto, e ainda podendo indicando o status code. 
De fato, esses problemas, são muito específicos. E não cabe a este projeto substituir a interface error, mas sim, aumentar os recursos dela, caso alguém precise.
## Documentação
### Instalação
Execute o comando:
```bash
go get github.com/rantool-team/go-error
```
### Uso
O goerror, é um pacote. A qual disponibiliza uma interface de erro(Error) a qual pode ser tranquilamente assimiliada a interface error do Go padrão. Para gerar um erro, use o método CreateError:
```go
err := goerror.CreateError("MENSSAGEM", "DESCRICAO")
```
Isso retornará um erro, que comprimenta a seguinte interface:
```go
type Error interface {
	Error() string // Retorna a menssagem de erro
	Emit() // Emite o erro através de um panic

	GetContext() context.Context // Retorna o contexto
	SetContext(localDescription string, description string, language string) // Altera o contexto do erro

	HasError() bool // Retorna true se tem erro e false se não.

	GetJSONFormatOfError() string // Retorna uma versão JSON do erro.

	GetStatusCode() int // Retorna o status code
	SetStatusCode(statusCode int) // Altera o status code

	GetDescription() string // Obtem a descrição na língua correta

	AddMessageInOtherLanguage(languageName string, message string) // Adiciona uma língua nova a menssagem do erro
	AddDescriptionInOtherLanguage(languageName string, description string) // Adiciona uma língua nova a descrição do erro

	SetLanguage(language string) // Altera para uma outra língua, que deve ser demonstrada no erro
}
```
## Erros vazios
Não aconselheamos retornar nil para a interface. O mais interessante, seria executar o método:
```go
err := CreateBlankError()
```
Esse código retorna um erro vázio, que tem o retorno do método HasError() como false.
## Verificando se aconteceu um erro 
Assim como não recomandos que retorne nil, recomendamos também que para verificar um erro, use o método HasError().
Exemplo:
```go
if err.HasError() {
    err.Emit()
}
```
## Trabalhando com a língua padrão
Você pode definir a língua padrão do seu projeto, a qual, caso não tenha sido alterado o contexto, será a língua usada.
Para alterar a língua padrão basta:
```go
goerror.SetDefaultLanguage("NOVA_LINGUA")
```
Para garantir que esteja escrevendo corretamente, recomendo usar o sub pacote, language, que possui um enumerado de constantes indicando as línguas. Para pegar a constante corretamente escolha as que terminam com _STRING. 
## Adicionando um nome para o status code
Um ponto muito importante, é que na emição de erros o status code quando é diferente de 0, aparece após um texto indicativo do que se trata. Para adicionar em diferentes línguas, use o método:
```go
goerror.AddStatusCodeName("NOME_DA_LINGUA", "EQUIVALENTE_DO_STATUS_CODE: ")
```
## Alterando variáveis uteis
Você pode alterar as seguintes variáveis, para aumentar a personalização(para alterar você precisar usar o sub pacote errorstruct("github.com/rantool-team/go-error/error-struct")):
```go
import 

var PREFIX_JSON = "" // Prefix passando para encodificador do JSON.
var IDENT_JSON = "  " // Identação passado para para o encodificador do JSON.

var PREFIX_ERROR_APPEAR = "\n////////////////ERROR////////////////\n" // Prefix do erro
var SUFIX_ERROR_APPEAR = "\n/////////////////////////////////////\n" // O que aparece depois do erro

var PREFIX_CONTEXT_APPEAR = "---------------------------------" // O que aparece antes do contexto quando existe
var SUFIX_CONTEXT_APPEAR = "---------------------------------" // O que aparece depois do contexto do erro
```
