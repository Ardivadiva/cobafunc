package cobafunc

import (
	"fmt"
	"testing"
	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dblisttamu",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertmonitor(t *testing.T) {
	nama := "Dipa"
	email := "ardvprw@gmail.com"
	kota := "Mojokerto"
	status := "Mahasiswi"
	hsl := Inserttamu(nama, email, kota, status)
	fmt.Println(hsl)
}

func TestGetDataNama(t *testing.T) {
	nama := "Dipa"
	dt := GetDataNama(nama)
	fmt.Println(dt)
}

func TestGetDataStatus(t *testing.T) {
	status := "Mahasiswi"
	dt := GetDataStatus(status)
	fmt.Println(dt)
}

func TestGetDataKota(t *testing.T) {
	kota := "Mojokerto"
	dt := GetDataKota(kota)
	fmt.Println(dt)
}
