package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN is the object the connections to the database */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://keneth:15052002Kj$@cluster0.ggn1w.mongodb.net/clonetwitter?retryWrites=true&w=majority")

/* ConnectDB is the function that allows me to connect the database */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

/* ConnectionCheck() is the ping to the database */
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
