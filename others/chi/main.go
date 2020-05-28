package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/articles/", func(r chi.Router) {
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", MyRequestHandler) // GET /articles/123
		})

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID") // from a route like /users/{userID}
	println(articleID)
	w.Write([]byte(fmt.Sprintf("hi %v", articleID)))
}
