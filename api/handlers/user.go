package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type NewUserInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserHandler(r chi.Router) {
	// r.Use(AuthMiddleware)
	r.Get("/{userId}", func(rw http.ResponseWriter, r *http.Request) {

		rw.Write([]byte("request from " + chi.URLParam(r, "userId")))
	})

	r.Post("/", func(rw http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user NewUserInfo
		if err := decoder.Decode(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
		}
		fmt.Println(user)
	})
}
