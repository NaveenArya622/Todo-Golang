package Database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DBConnection *sqlx.DB

func init() {
	//var err error
	dbs := "host=localhost port=5433 user=local password=local dbname=todo sslmode=disable"
	DB, err := sqlx.Open("postgres", dbs)
	if err != nil {
		log.Fatal("error in db connection")
	}
	DBConnection = DB

}

func CloseDatabase() error {
	return DBConnection.Close()
}
