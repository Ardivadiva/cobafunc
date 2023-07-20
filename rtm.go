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
func GetDataNama(nam string) (data []Tamu) {
	user := MongoConnect("dblisttamu").Collection("datatamu")
	filter := bson.M{"nama": nam}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataNama :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataStatus(tus string) (data []Tamu) {
	user1 := MongoConnect("dblisttamu").Collection("datatamu")
	filter := bson.M{"status": tus}
	cursor, err := user1.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataStatus :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

