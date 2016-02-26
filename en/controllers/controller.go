package main

import "net/http"

type action func(rw http.ResponseWriter, r *http.Request) error

type appController struct{}

func (c *appController) action(a action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
