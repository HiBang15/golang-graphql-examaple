package service

import (
	"context"
	"github.com/HiBang15/golang-graphql-examaple.git/model"
	"github.com/HiBang15/golang-graphql-examaple.git/payload"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)
/* Rest API */
func RestApiGetBook(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	books := &model.Book{}
	bookName := chi.URLParam(r, "bookname")

	cur, err := model.MongoDb.Collection("booklist").Find(ctx, bson.M{"name": bookName})
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for cur.Next(ctx) {
		cur.Decode(&books)

	}

	if s := books; s != nil {
		payload.HttpResponseSuccess(w, r, books)
		return
	}
	return
}
