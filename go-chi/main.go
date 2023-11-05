package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// Delete リクエストの場合
	r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is Delete request")
	})

	// パスパラメータを受け取る場合
	r.Get("/{userId}/detail", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		fmt.Fprintf(w, "userId is %s", userId)
	})

	http.ListenAndServe(":3000", r)
}
