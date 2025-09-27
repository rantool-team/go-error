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
```
go get github.com/rantool-team/go-error
```
### Uso
