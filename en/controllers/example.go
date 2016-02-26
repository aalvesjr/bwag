package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

type myController struct {
	appController
	*render.Render
}

func (c *myController) index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &myController{Render: render.New()}
	http.ListenAndServe(":8080", c.action(c.index))
}
