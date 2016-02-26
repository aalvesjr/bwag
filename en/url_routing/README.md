# URL Routing

For some simple applications, the default `http.ServeMux` can take you pretty
far. If you need more power in how you parse URL endpoints and route them to
the proper handler, you may need to pull in a third party routing framework.
For this tutorial, we will use the popular
`github.com/julienschmidt/httprouter` library as our router.
`github.com/julienschmidt/httprouter` is a great choice for a router as it is a
very simple implementation with one of the best performance benchmarks out of
all the third party Go routers.

In this example, we will create some routing for a RESTful resource called
"posts". Below we define mechanisms to view index, show, create, update,
destroy, and edit posts.

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", homeHandler)

	// Posts collection
	r.GET("/posts", postsIndexHandler)
	r.POST("/posts", postsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", postShowHandler)
	r.PUT("/posts/:id", postUpdateHandler)
	r.GET("/posts/:id/edit", postEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func postsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

func postsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

func postShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

func postUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

func postDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

func postEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
```

## Exercises

1. Explore the documentation for `github.com/julienschmidt/httprouter`.
2. Find out how well `github.com/julienschmidt/httprouter` plays nicely with existing `http.Handler`s like `http.FileServer`
3. `httprouter` has a very simple interface. Explore what kind of abstractions can be built on top of this fast router to make building things like RESTful routing easier.
