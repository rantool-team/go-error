### Alterando Variáveis Úteis

Você pode alterar as seguintes variáveis para aumentar a personalização (para alterar, você precisa usar o subpacote [`error-struct`](./error-struct/)):

```go
import "github.com/rantool-team/go-error/error-struct"

var PREFIX_JSON = "" // Prefixo passado para o codificador JSON
var IDENT_JSON = "  " // Indentação passada para o codificador JSON

var PREFIX_ERROR_APPEAR = "
////////////////ERROR////////////////
" // Prefixo do erro
var SUFIX_ERROR_APPEAR = "
/////////////////////////////////////
" // O que aparece depois do erro

var PREFIX_CONTEXT_APPEAR = "
---------------------------------
" // O que aparece antes do contexto, quando existe
var SUFIX_CONTEXT_APPEAR = "
---------------------------------
" // O que aparece depois do contexto do erro
```