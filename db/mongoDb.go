package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/basic-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkError(err error) bool {
	if err != nil {
		log.Fatal(err)

		return true
	}

	return false
}

const mongoConnectionURI string = "mongodb://localhost:27017/my_apps_db"
const dbName string = "my_apps_db"
const collectionName string = "people"

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

// InsertPerson : Insert given Person document
func InsertPerson(insertedPerson models.Person) models.TransactionResponse {
	mongoClient := connectToMongoDB()
	peopleCollection := mongoClient.Database(dbName).Collection(collectionName)

	transactionResponse := models.TransactionResponse{IsSuccess: false}

	insertedPerson.SystemEntryDate = time.Now()
	insertedPerson.UpdatedTime = time.Now()

	res, err := peopleCollection.InsertOne(context.TODO(), insertedPerson)
	transactionResponse.IsSuccess = !checkError(err)

	if err != nil {
		transactionResponse.Message = err.Error()
	} else {
		transactionResponse.Message = fmt.Sprintf("Object Id: %s inserted !", res.InsertedID)
	}

	defer mongoClient.Disconnect(context.TODO())
	return transactionResponse
}

// UpdatePersonByID : Update given Person document by Id
func UpdatePersonByID(id string, updatedPerson models.Person) models.TransactionResponse {
	mongoClient := connectToMongoDB()
	peopleCollection := mongoClient.Database(dbName).Collection(collectionName)

	transactionResponse := models.TransactionResponse{IsSuccess: false}

	objectID, _ := primitive.ObjectIDFromHex(id)

	_, err := peopleCollection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": bson.M{
		"firstName":   updatedPerson.FirstName,
		"lastName":    updatedPerson.LastName,
		"birthdate":   updatedPerson.Birthdate,
		"email":       updatedPerson.Email,
		"updatedTime": time.Now().String()}})

	transactionResponse.IsSuccess = !checkError(err)

	if err != nil {
		transactionResponse.Message = err.Error()
	} else {
		transactionResponse.Message = fmt.Sprintf("Object Id: %s updated !", id)
	}

	defer mongoClient.Disconnect(context.TODO())
	return transactionResponse
}

func getPeopleByFilter(filter bson.M) []*models.Person {
	mongoClient := connectToMongoDB()
	cur, err := mongoClient.Database(dbName).Collection(collectionName).Find(context.TODO(), filter)
	checkError(err)

	var people []*models.Person

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
