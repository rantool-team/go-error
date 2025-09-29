# Lista de Erros

A lista de erros é uma forma de gerenciar erros em sua aplicação. Você pode criar uma lista de erros e obtê-los por seu ID.

## CreateErrorOnErrorsListAndReturnId

Esta função cria um erro e o adiciona à lista de erros. Ela retorna o ID do erro.

```go
func CreateErrorOnErrorsListAndReturnId(msg string, description string) int
```

### Exemplo

```go
package main

import (
    "fmt"
    "github.com/rantool-team/go-error/errors-list"
)

func main() {
    id := errorslist.CreateErrorOnErrorsListAndReturnId("Not Found", "The requested resource was not found")
    fmt.Println(id)
}
```

## CreateErrorInSetOfLanguagesOnErrorsListAndReturnId

Esta função cria um erro em um conjunto de idiomas e o adiciona à lista de erros. Ela retorna o ID do erro.

```go
func CreateErrorInSetOfLanguagesOnErrorsListAndReturnId(msgSet language.MessageSet, descriptionSet language.MessageSet) int
```

### Exemplo

```go
package main

import (
    "fmt"
    "github.com/rantool-team/go-error/errors-list"
    "github.com/rantool-team/go-error/language"
)

func main() {
    msgSet := language.NewMessageSet("Not Found")
    msgSet.Add("pt-br", "Não encontrado")

    descriptionSet := language.NewMessageSet("The requested resource was not found")
    descriptionSet.Add("pt-br", "O recurso solicitado não foi encontrado")

    id := errorslist.CreateErrorInSetOfLanguagesOnErrorsListAndReturnId(msgSet, descriptionSet)
    fmt.Println(id)
}
```

## CreateErrorOnErrorsListAndReturnIdAndError

Esta função cria um erro e o adiciona à lista de erros. Ela retorna o ID do erro e o próprio erro.

```go
func CreateErrorOnErrorsListAndReturnIdAndError(msg string, description string) (int, goerror.Error)
```

### Exemplo

```go
package main

import (
    "fmt"
    "github.com/rantool-team/go-error/errors-list"
)

func main() {
    id, err := errorslist.CreateErrorOnErrorsListAndReturnIdAndError("Not Found", "The requested resource was not found")
    fmt.Println(id, err)
}
```

## CreateErrorInSetOfLanguagesOnErrorsListAndReturnIdAndError

Esta função cria um erro em um conjunto de idiomas e o adiciona à lista de erros. Ela retorna o ID do erro e o próprio erro.

```go
func CreateErrorInSetOfLanguagesOnErrorsListAndReturnIdAndError(msgSet language.MessageSet, descriptionSet language.MessageSet) (int, goerror.Error)
```

### Exemplo

```go
package main

import (
    "fmt"
    "github.com/rantool-team/go-error/errors-list"
    "github.com/rantool-team/go-error/language"
)

func main() {
    msgSet := language.NewMessageSet("Not Found")
    msgSet.Add("pt-br", "Não encontrado")

    descriptionSet := language.NewMessageSet("The requested resource was not found")
    descriptionSet.Add("pt-br", "O recurso solicitado não foi encontrado")

    id, err := errorslist.CreateErrorInSetOfLanguagesOnErrorsListAndReturnIdAndError(msgSet, descriptionSet)
    fmt.Println(id, err)
}
```

## GetErrorOnErrorId

Esta função retorna um erro da lista de erros por seu ID.

```go
func GetErrorOnErrorId(errorId int) goerror.Error
```

### Exemplo

```go
package main

import (
    "fmt"
    "github.com/rantool-team/go-error/errors-list"
)

func main() {
    id := errorslist.CreateErrorOnErrorsListAndReturnId("Not Found", "The requested resource was not found")
    err := errorslist.GetErrorOnErrorId(id)
    fmt.Println(err)
}
```