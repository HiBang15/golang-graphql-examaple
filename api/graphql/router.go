package graphql

import (
	"github.com/HiBang15/golang-graphql-examaple.git/repository"
	"github.com/HiBang15/golang-graphql-examaple.git/service"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/handler"
	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	graphQL := handler.New(&handler.Config{
		Schema:           &repository.Schema,
		Pretty:           true,
		GraphiQL:         true,
	})

	r.Use(middleware.Logger)
	r.Handle("/query", graphQL)

	r.Get("/get/book/{bookname}", service.RestApiGetBook)

	return r
}
