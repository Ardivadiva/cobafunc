package cobafunc

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func Inserttamu(db *mongo.Database, nama string, email string, kota string, status string) (InsertedID interface{}) {
	var datatamu Tamu
	datatamu.Nama = nama
	datatamu.Email = email
	datatamu.Kota = kota
	datatamu.Status = status
	return InsertOneDoc(db, "datatamu", datatamu)
}
func GetDatatamu(kot string) (data []Tamu) {
	user := MongoConnect("dblisttamu").Collection("datatamu")
	filter := bson.M{"kota": kot}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataTamu :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}