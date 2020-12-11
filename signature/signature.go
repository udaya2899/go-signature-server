package signature

import (
	"crypto/ed25519"
	"fmt"
	"log"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/db/signature"
	"github.com/google/uuid"
)

// Service struct is returned by the NewService function
type Service struct {
	log     *log.Logger
	queryer signature.ISignature
}

// NewService creates a new connection as well as logger entry
func NewService(logger *log.Logger) *Service {
	return &Service{
		log:     logger,
		queryer: signature.New(logger),
	}
}

// SignedMessage is the struct that is returned after signing
type SignedMessage struct {
	Message []byte
	Sign    []byte
}

// GetPublicKey returns the public key stored from the pem file loaded in config
func (s Service) GetPublicKey() (ed25519.PublicKey, error) {
	if config.Env.PublicKey == nil {
		return nil, fmt.Errorf("Public Key not found")
	}
	return config.Env.PublicKey, nil
}

// PutTransactionBlob writes transaction to the db by generating a new uuid as key
func (s Service) PutTransactionBlob(txn string) (string, error) {
	id := uuid.New().String()

	err := s.queryer.PutTransaction(id, txn)
	if err != nil {
		return "", fmt.Errorf("Cannot write transaction, err: %v", err)
	}

	return id, nil
}

// PostSignature gets the blob values of the given transactions and returns the blob as well as the signature
func (s Service) PostSignature(transactionIDs []string) (*SignedMessage, error) {
	if transactionIDs == nil {
		return nil, fmt.Errorf("No Transaction IDs received")
	}

	var values []byte

	for _, id := range transactionIDs {

		value, err := s.queryer.GetTransactionByID(id)
		if err != nil {
			return nil, fmt.Errorf("Error getting transaction by id for id: %v, err: %v", id, err)
		}

		values = append(values, value...)
	}

	s.log.Printf("Value obtained from ID: %+v", string(values))

	signature := ed25519.Sign(config.Env.PrivateKey, values)

	// For safety, verifying the self-signed message with the public key
	if ok := ed25519.Verify(config.Env.PublicKey, values, signature); !ok {
		return nil, fmt.Errorf("Cannot verify self-signed signature, something went wrong")
	}

	s.log.Printf("Signed successfully for given transactions")

	return &SignedMessage{
		Message: values,
		Sign:    signature,
	}, nil
}
