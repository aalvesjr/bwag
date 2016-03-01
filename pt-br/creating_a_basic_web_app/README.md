# Criando um App Web Básico

Agora que terminamos de revisar o básico do HTTP, vamos criar uma aplicação web
simples porém útil em Go.

Usando o nosso programa que serve arquivos implementado no último capítulo,
vamos implementar um gerador de Markdown usando o pacote `github.com/russross/blackfriday`.

## Form HTML

Para começar, vamos precisar de um form HTML básico para o input em markdown:

``` html
<html>
  <head>
    <link href="/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body>
    <div class="container">
      <div class="page-title">
        <h1>Gerador de Markdown</h1>
        <p class="lead">Gere o seu markdown usando Go</p>
        <hr />
      </div>

      <form action="/markdown" method="POST">
        <div class="form-group">
          <textarea class="form-control" name="body" cols="30" rows="10"></textarea>
        </div>

        <div class="form-group">
          <input type="submit" class="btn btn-primary pull-right" />
        </div>
      </form>
    </div>
    <script src="/js/bootstrap.min.js"></script>
  </body>
</html>
```

Coloque esse HTML em um arquivo chamdo `index.html` no diretório "public" da nossa aplicação
e o arquivo `bootstrap.min.css` do http://getbootstrap.com/ no diretório "public/css".
Note que o form envia os dados via HTTP POST para o endpoint "/markdown" da nossa aplicação.
Ainda não temos um handler para essa rota, então vamos criar um.

## A rota "/markdown"

O programa para tratar a rota '/markdown' e servir o arquivo `index.html` a partir do diretório
public é assim:

``` go
package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
```

Vamos dividí-lo em partes menores para analisar melhor.

``` go
http.HandleFunc("/markdown", GenerateMarkdown)
http.Handle("/", http.FileServer(http.Dir("public")))
```

Estamos usando os métodos `http.HandleFunc` e `http.Handle`
para definir rotas simples para a nossa aplicação. É importante notar que
chamar o `http.Handle` usando o padrão "/" funciona como uma rota genérica
que pega todos os requests, por isso definimos ela por último.
O `http.FileServer` retorna um `http.Handler` então usamos o `http.Handle`
para mapear um padrão em string para um handler. A forma alternativa,
`http.HandleFunc`, usa um `http.HandlerFunc` em vez de um `http.Handler`.
Pode ser mais conveniente pensar dessa forma, tratar rotas através de uma
função em vez de um objeto.

``` go
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
}
```

Nossa função GenerateMarkdown implementa a interface padrão `http.HandlerFunc`
e renderiza HTML a partir de um campo do form contendo texto formatado em markdown.
Nesse caso, o conteúdo do campo é obtido através do `r.FormValue("body")`.
É bem comum receber dados de input do objeto `http.Request` que o `http.HandlerFunc`
recebe como argumento. Alguns outros exemplos de input são os membros `r.Header`,
`r.Body` e `r.URL`.

Finalizamos o tratamento do request escrevendo no nosso `http.ResponseWriter`. Note
que não mandamos um código de resposta HTTP explicitamente. Quando escrevemos
no response sem definir um código, o pacote `net/http` assume que o código é um
`200 OK`. Isso significa que se algo tivesse dado errado, deveríamos setar o código
através do método `rw.WriterHeader()`.

``` go
http.ListenAndServe(":8080", nil)
```

A última parte do programa inicia o servidor e passamos `nil` como handler,
que assume que os requests HTTP serão gerenciados pelo `http.ServeMux` padrão
do pacote `net/http`, que é configurado usando o `http.Handle` e o `http.HandleFunc`.

Isso é tudo que é necessário para gerar markdown como um serviço em Go. É uma quantidade
surpreendentemente pequena de código para a quantidade de trabalho pesada que está
sendo feita. No próximo capítulo vamos aprender como fazer deploy da nossa aplicação
para web usando o Heroku.
