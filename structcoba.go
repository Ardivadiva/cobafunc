package cobafunc
import "go.mongodb.org/mongo-driver/bson/primitive"

type Listtamu struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Notelp string `bson:"notelp,omitempty" json:"notelp,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Kota string `bson:"kota,omitempty" json:"kota,omitempty"`
}
