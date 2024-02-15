package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserHandler(r chi.Router) {
	// r.Use(AuthMiddleware)
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {

		rw.Write([]byte("request from " + chi.URLParam(r, "userId")))
	})
}
