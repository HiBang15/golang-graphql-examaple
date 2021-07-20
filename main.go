package main

import (
	"github.com/HiBang15/golang-graphql-examaple.git/api/graphql"
	"github.com/HiBang15/golang-graphql-examaple.git/model"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/url"
)

func main() {
	routes := chi.NewRouter()
	r := graphql.RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := model.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
