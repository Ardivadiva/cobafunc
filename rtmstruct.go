package cobafunc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct {
	ID         			primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nama				string             `bson:"nama" json:"nama"`
	Email				string             `bson:"email" json:"email"`
	Kota 				string             `bson:"kota" json:"kota"`
	Status 				string             `bson:"status" json:"status"`
}