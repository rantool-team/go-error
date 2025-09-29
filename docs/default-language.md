# Trabalhando com o Idioma Padrão

Você pode definir o idioma padrão do seu projeto, que será usado caso o contexto não tenha sido alterado. Para alterar o idioma padrão, use a função [`SetDefaultLanguage`](./set-language-default.go):

```go
goerror.SetDefaultLanguage("NOVO_IDIOMA")
```

Para garantir que você esteja escrevendo corretamente, recomendo usar o subpacote [`language`](./language/languages.go), que possui um enumerado de constantes indicando os idiomas. Para obter a constante correta, escolha as que terminam com `_STRING`.