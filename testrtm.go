package cobafunc

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dblisttamu",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Nama := "Dipa"
	Email := "ardvprw@gmail.com"
	Kota := "Mojokerto"
	Status := "Mahasiswi"
	hasil := InsertDataTamuu(MongoConn, Nama, Email, Kota, Status)
	fmt.Println(hasil)

}

func TestGetDataTamuu(t *testing.T) {
	Nama := "Dipa"
	hasil := GetDataTamuu(Nama, MongoConn, "data_tamuu")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Status := "Mahasiswi"
	hasil := DeleteDataTamuu(Status, MongoConn, "data_tamuu")
	fmt.Println(hasil)

}