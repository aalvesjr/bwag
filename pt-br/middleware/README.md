# Middleware

Se você tiver algum código que precisa ser executado a cada requisição
(request), independente da rota que o invocará, você precisará de um modo
de enfileirar os `http.Handlers` um após o outro e então executá-los em
sequência. Esta questão é resolvida de maneira elegante através do uso de
pacotes middleware. O Negroni é um pacote middleware que facilita muito o
desenvolvimento e enfileiramento de middleware, enquanto mantém intacta a
natureza combinável do ecosistema web do Go.

O Negroni já vem com alguns middlewares nativos, tais como Registro de Log,
Recuperação de Erros e Servidor de Arquivos Estáticos. Deste modo, utilizando
apenas a sua configuração inicial, o Negroni é capaz de agregar muito valor ao
seu código sem causar maiores impactos.

O exemplo abaixo mostra como usar o Negroni com o conjunto de middlewares
nativo (Middleware stack) e também como criar o seu próprio middleware
customizado (MyMiddleware).


``` go
package main

import (
    "log"
    "net/http"

    "github.com/codegangsta/negroni"
)

func main() {
    // Middleware stack
    n := negroni.New(
        negroni.NewRecovery(),
        negroni.HandlerFunc(MyMiddleware),
        negroni.NewLogger(),
        negroni.NewStatic(http.Dir("public")),
    )

    n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("Logging on the way there...")

    if r.URL.Query().Get("password") == "secret123" {
        next(rw, r)
    } else {
        http.Error(rw, "Not Authorized", 401)
    }

    log.Println("Logging on the way back...")
}
```

## Exercícios

1. Pense em algumas boas idéias para middleware e tente implementá-las usando o
Negroni.
2. Explore como o Negroni pode ser combinado com `github.com/gorilla/mux` usando
a interface `http.Handler`.
3. Experimente utilizar o Negroni para criar enfileiramentos de middleware para
grupos específicos de rotas ao invés de toda a aplicação.
