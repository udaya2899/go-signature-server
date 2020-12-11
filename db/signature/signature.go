package signature

import (
	"log"

	"challenge.summitto.com/udaya2899/challenge_result/db"
	badger "github.com/dgraph-io/badger/v2"
)

type ISignature interface {
	PutTransaction(key string, value string) error
	GetTransactionByID(id string) ([]byte, error)
}

type SignatureDAO struct {
	log    *log.Logger
	dbConn *badger.DB
}

func New(l *log.Logger) *SignatureDAO {
	return &SignatureDAO{
		log:    l,
		dbConn: db.DBConn,
	}
}

// PutTransaction writes the key value to the db
func (s *SignatureDAO) PutTransaction(key string, value string) error {

	err := s.dbConn.Update(func(txn *badger.Txn) error {

		err := txn.Set(
			[]byte(key),
			[]byte(value),
		)
		return err
	})

	return err
}

func (s *SignatureDAO) GetTransactionByID(id string) ([]byte, error) {
	var msg *badger.Item
	var msgVal []byte

	err := s.dbConn.View(func(txn *badger.Txn) error {
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
