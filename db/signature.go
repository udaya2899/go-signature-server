package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

// PutTransaction writes
func PutTransaction(key string, value string) error {

	err := DBConn.Update(func(txn *badger.Txn) error {

		err := txn.Set(
			[]byte(key),
			[]byte(value),
		)
		return err
	})

	return err
}

func GetTransactionByID(id string) ([]byte, error) {
	var msg *badger.Item
	var msgVal []byte

	err := DBConn.View(func(txn *badger.Txn) error {
		var err error
		msg, err = txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = msg.Value(func(v []byte) error {

			log.Printf("key=%s, value=%s\n", msg.Key(), v)

			msgVal = append([]byte{}, v...)

			return nil
		})

		return err
	})

	return msgVal, err
}
