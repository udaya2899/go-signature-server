package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

// DBConn is the connection for the database used - badgetDB
var DBConn *badger.DB

func init() {
	DBConn = initDB()
}

func initDB() *badger.DB {
	dbConn, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
