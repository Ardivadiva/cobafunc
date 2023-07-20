package cobafunc

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataListTamu(db *mongo.Database, name string, notelp string, email string, kota string) (InsertedID interface{}) {
	var datalisttamu Listtamu
	datalisttamu.Name = name
	datalisttamu.Notelp = notelp
	datalisttamu.Email = email
	datalisttamu.Kota = kota
	return InsertOneDoc(db, "datalisttamu", datalisttamu)
}

func GetDataListTamu(kota string, db *mongo.Database, col string) (data Listtamu) {
	act := db.Collection(col)
	filter := bson.M{"kota": kota}
	err := act.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataaccbyact: %v\n", err)
	}
	return data
}
func GetDataName(Name string, db *mongo.Database, col string) (data Listtamu) {
	accou := db.Collection(col)
	filter := bson.M{"name": Name}
	err := accou.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdatalisttamu: %v\n", err)
	}
	return data
}
func DeleteDataListTamu(kota string, db *mongo.Database, col string) (data Listtamu) {
	dct := db.Collection(col)
	filter := bson.M{"kota": kota}
	err, _ := dct.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataListTamu : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}

func DeleteDataName(Name string, db *mongo.Database, col string) (data Listtamu) {
	dena := db.Collection(col)
	filter := bson.M{"name": Name}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataName : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}