package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://nicolas:nicolas@cluster0.ihehzmm.mongodb.net/?retryWrites=true&w=majority")

/*Conexion a la DB*/
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
	log.Println("Conexion Exitosa con la DB")
	return client
}

/*Ping a la DB*/
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
