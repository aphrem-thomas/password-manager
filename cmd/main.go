package main

import (
	"net/http"

	"github.com/aphrem-thomas/password-manager/api/handlers"
	"github.com/aphrem-thomas/password-manager/middlewares"
	"github.com/go-chi/chi"
)

func main() {
	// cr, _ := services.NewAccountService(services.WithMemoryAccountRepository())
	// id := uuid.New()
	// cr.AddAccount()
	// fmt.Println(cr.GetAccount(id))

	r := chi.NewRouter()
	r.Use(middlewares.TestMiddleware)
	r.Route("/user", handlers.UserHandler)

	http.ListenAndServe(":3000", r)

}
