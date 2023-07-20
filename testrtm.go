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
	DBName:   "listtamu",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	name := "Dipa"
	notelp := "6287815683284"
	email := "ardvprw@gmail.com"
	kota := "Mojokerto"
	hasil := InsertDataListTamu(MongoConn, name, notelp, email, kota)
	fmt.Println(hasil)

}

func TestGetDatatamu(t *testing.T) {
	name := "Dipa"
	hasil := GetDataListTamu(name, MongoConn, "data_listtamu")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	kota := "Mojokerto"
	hasil := DeleteDataListTamu(kota, MongoConn, "data_listtamu")
	fmt.Println(hasil)

}