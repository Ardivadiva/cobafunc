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
	hsl := Inserttamu(MongoConn, nama, email, kota, status)
	fmt.Println(hsl)
}

