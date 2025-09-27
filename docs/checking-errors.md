### Verificando a Ocorrência de Erros

Assim como não recomendamos retornar `nil`, recomendamos que, para verificar um erro, você use o método `HasError()`.

**Exemplo:**

```go
if err.HasError() {
    err.Emit()
}
```