# Erros Vazios

Não aconselhamos retornar `nil` para a interface. Em vez disso, use o método [`CreateBlankError`](./blank-error.go):

```go
err := CreateBlankError()
```

Esse código retorna um erro vazio, que tem o retorno do método `HasError()` como `false`.