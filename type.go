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

func InsertDataTamuu(db *mongo.Database, nama, email, kota string, status string) (InsertedID interface{}) {
	var datatamuu Datatamuu
	datatamuu.Nama = nama
	datatamuu.Email = email
	datatamuu.Kota = kota
	datatamuu.Status = status
	return InsertOneDoc(db, "data_tamuu", datatamuu)
}

func GetDataTamuu(kota string, db *mongo.Database, col string) (data Datatamuu) {
	act := db.Collection(col)
	filter := bson.M{"kota": kota}
	err := act.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataaccbyact: %v\n", err)
	}
	return data
}
func GetDataNama(nama string, db *mongo.Database, col string) (data Datatamuu) {
	accou := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := accou.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdatatamuu: %v\n", err)
	}
	return data
}
func DeleteDataTamuu(kota string, db *mongo.Database, col string) (data Datatamuu) {
	dct := db.Collection(col)
	filter := bson.M{"kota": kota}
	err, _ := dct.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataTamuu : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}

func DeleteDataNama(nama string, db *mongo.Database, col string) (data Datatamuu) {
	dena := db.Collection(col)
	filter := bson.M{"nama": nama}
	err, _ := dena.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataNama : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}