### Adicionando um Nome para o Código de Status

Um ponto muito importante é que, na emissão de erros, o código de status, quando diferente de 0, aparece após um texto indicativo do que se trata. Para adicionar em diferentes idiomas, use o método [`AddStatusCodeName`](./status-code-names.go):

```go
goerror.AddStatusCodeName("NOME_DO_IDIOMA", "EQUIVALENTE_DO_CÓDIGO_DE_STATUS: ")
```