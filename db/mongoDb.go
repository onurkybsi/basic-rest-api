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

const mongoConnectionURI string = "mongodb://localhost:27017/my_apps_db"

// GetAllPeople : Return all data from "people" collection of "my_apps_db" database
func GetAllPeople() []*models.Person {

	return getPeopleByFilter(bson.M{})
}

// GetPeopleByName : Return people that matched by name
func GetPeopleByName(name string) []*models.Person {
	filter := bson.M{
		"$or": []bson.M{
			bson.M{"firstName": name},
			bson.M{"lastName": name},
		},
	}

	return getPeopleByFilter(filter)
}

func getPeopleByFilter(filter bson.M) []*models.Person {
	var people []*models.Person

	mongoClient := connectToMongoDB()

	cur, err := mongoClient.Database("my_apps_db").Collection("people").Find(context.TODO(), filter)
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
