package repository

import (
	"context"
	"github.com/HiBang15/golang-graphql-examaple.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetBookByName(ctx context.Context, name string) (result interface{}) {
	var book model.Book

	data := model.MongoDb.Collection("booklist").FindOne(ctx, bson.M{"name": name})
	data.Decode(&book)
	return book
}

func GetBookList(ctx context.Context, limit int) (result interface{}) {
	var book model.Book
	var books []model.Book

	option := options.Find().SetLimit(int64(limit))
	cur, err := model.MongoDb.Collection("booklist").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}

	for cur.Next(ctx) {
		cur.Decode(&book)
		books = append(books, book)
	}

	return books
}

func InsertBook(ctx context.Context, book model.Book) error {
	_, err := model.MongoDb.Collection("booklist").InsertOne(ctx, book)
	return err
}

func UpdateBook(ctx context.Context, book model.Book) error {
	filter := bson.M{"name": book.Name}
	update := bson.M{"$set": book}
	upsertBool := true
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	_, err := model.MongoDb.Collection("booklist").UpdateOne(ctx, filter, update, &updateOption)
	return err
}

func DeleteBook(ctx context.Context, name string) error {
	_, err := model.MongoDb.Collection("booklist").DeleteOne(ctx, bson.M{"name": name})
	return err
}
