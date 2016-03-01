# Introdução

Bem vindo ao **Criando Aplicações Web com Go**! Se você esta lendo isto então você acabou de começar sua jornada de iniciante a profissional. É sério, programação WEB em Go é tão divertido e fácil que você nem vai perceber quanta informação você está aprendendo ao longo do caminho!

Tenha em mente que ainda há partes deste livro que estão incompletas e precisam de um pouco de amor. A beleza da publicação de código aberto é que eu posso te dar um livro incompleto e ainda sim ter valor para você.

Antes de entrar em todos os pormenores, vamos começar com algumas regras básicas:

## Pré-requisitos
Para manter este tutorial pequeno e focado, Eu estou assumindo que você está preparado das seguintes formas:

1. Você tem instalado a [Linguagem de Programação Go (en)](https://golang.org).
2. Você tem configurado o `GOPATH`, seguindo o tutorial [Como Escrever Código Go (en)](https://golang.org/doc/code.html#Organization).
3. Você está um pouco familiarizado com os conceitos básicos de Go. ([Um Passeio por Go](http://go-tour-br.appspot.com) é um bom lugar para começar)
4. Você tem instalado todos os [pacotes necessários](#required-packages)
5. Você tem instalado o [Cinto de Ferramentas do Heroku ](https://toolbelt.heroku.com/)
6. Você tem uma conta [Heroku](https://id.heroku.com/signup)

## Pacotes Necessários

Na maior parte do tutorial nós iremos usar os pacotes da biblioteca padrão para construir nossas aplicações web. Algumas lições como Banco de Dados, Middleware e Roteamento de URL vão exigir pacotes de terceiros. Aqui está uma lista de todos os pacotes go que você vai precisar instalar antes de começar:

Nome | Caminho a Importar | Descrição
---- | ----------- | -----------
[httprouter](https://github.com/julienschmidt/httprouter) | github.com/julienschmidt/httprouter | Roteador de requisições HTTP de alta performance que escala bem
[Negroni](https://github.com/codegangsta/negroni) | github.com/codegangsta/negroni | Middleware HTTP idiomático
[Black Friday](https://github.com/russross/blackfriday) | github.com/russross/blackfriday | Processador de markdown
[Render](https://github.com/unrolled/render/tree/v1) | gopkg.in/unrolled/render.v1 | Renderização fácil para JSON, XML, e HTML
[SQLite3](https://github.com/mattn/go-sqlite3) | github.com/mattn/go-sqlite3 | Driver sqlite3 para go

Você pode instalar (ou atualizar) esses pacotes executando o comando a seguir no seu terminal

``` bash
go get -u <import_path>
```

Por exemplo, se você deseja instalar o Negroni, o comando seria:

``` bash
go get -u github.com/codegangsta/negroni
```
