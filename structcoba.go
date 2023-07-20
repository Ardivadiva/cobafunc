package cobafunc

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accounts struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama       string             `bson:"name,omitempty" json:"name,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Kota     string             `bson:"kota,omitempty" json:"kota,omitempty"`
	Status string             `bson:"status,omitempty" json:"status,omitempty"`
}