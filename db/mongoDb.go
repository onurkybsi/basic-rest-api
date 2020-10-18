package db

import (
	"context"
	"fmt"
	"log"

	"github.com/basic-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const mongoPass string = "?"

var mongoConnectionURI string = fmt.Sprintf("mongodb+srv://admin:%s@mymongodb.oiync.mongodb.net/<dbname>?retryWrites=true&w=majority", mongoPass)

// GetAllPeople : Return all data from "people" collection of "my_apps_db" database
func GetAllPeople() []*models.Person {
	var people []*models.Person

	mongoClient := connectToMongoDB()

	cur, err := mongoClient.Database("my_apps_db").Collection("people").Find(context.TODO(), bson.D{{}})
	checkError(err)

	for cur.Next(context.TODO()) {

		var person models.Person
		err := cur.Decode(&person)
		checkError(err)

		people = append(people, &person)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(context.TODO())

	return people
}

func connectToMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoConnectionURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkError(err)

	err = client.Ping(context.TODO(), nil)
	checkError(err)

	fmt.Println("Connected to MongoDB!")

	return client
}
