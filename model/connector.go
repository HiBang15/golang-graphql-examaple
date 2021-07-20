package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Database struct {
	Name       string `yaml:"name"`
	Connection string `yaml:"connection"`
}
type App struct {
	AppName     string `yaml:"name"`
	Debug       bool `yaml:"debug"`
	Port        string `yaml:"port"`
	Service     string `yaml:"service"`
	Certificate string `yaml:"certificate"`
	PemKey      string `yaml:"pem_key"`
	Host        string `yaml:"host"`
	Stage       string `yaml:"stage"`
}

type Environment struct {
	App       App                 `yaml:"app"`
	Databases map[string]Database `yaml:"databases"`
	path      string
}

var MongoDb *mongo.Database

func (e *Environment) InitMongoDB() (db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(e.Databases["mongodb"].Connection)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Connect db fail with err: " + err.Error())
		return db, err
	}
	MongoDb = client.Database(e.Databases["mongodb"].Name)
	log.Println("Mongodb Ready!!!")
	return db, nil
}
