# Go-Error

Em Go, a interface de erro (`error`) é um recurso padrão. No entanto, em projetos maiores, ela pode ser um pouco limitada. Sinto falta de recursos como:

- **Definição de contextos**: Seria interessante ter uma ferramenta para indicar o contexto onde o erro ocorreu.
- **Tratamento de erros**: O formato `err != nil` é interessante, porém pouco visual.
- **Internacionalização**: Em projetos maiores, seria útil ter uma ferramenta para traduzir os erros em diversos idiomas.
- **Retorno em JSON**: Também seria interessante ter uma ferramenta nativa que controla os erros e retorna um JSON, indicando o código de status.

Este projeto não visa substituir a interface `error` padrão do Go, mas sim estender seus recursos para quem precisar.

## Documentação

- [Instalação](./docs/installation.md)
- [Uso](./docs/usage.md)
- [Erros Vazios](./docs/blank-errors.md)
- [Verificando a Ocorrência de Erros](./docs/checking-errors.md)
- [Trabalhando com o Idioma Padrão](./docs/default-language.md)
- [Adicionando um Nome para o Código de Status](./docs/status-code-name.md)
- [Alterando Variáveis Úteis](./docs/customization.md)