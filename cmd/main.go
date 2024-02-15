package main

import (
	"net/http"

	"github.com/aphrem-thomas/password-manager/api/handlers"
	"github.com/go-chi/chi"
)

func main() {
	// cr, _ := services.NewAccountService(services.WithMemoryAccountRepository())
	// id := uuid.New()
	// cr.AddAccount()
	// fmt.Println(cr.GetAccount(id))

	r := chi.NewRouter()
	r.Route("/user/{userId}", handlers.UserHandler)

	http.ListenAndServe(":3000", r)

}
