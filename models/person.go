package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person struct
type Person struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate"`
	Email     string             `bson:"email" json:"email"`
}
