package config

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func loadKey() error {
	privKeyErr := loadPrivateKey()
	if privKeyErr != nil {
		return fmt.Errorf("Cannot load privateKey, Err: %v", privKeyErr)
	}
	pubKeyErr := loadPublicKey()
	if pubKeyErr != nil {
		return fmt.Errorf("Cannot load pubKey, Err: %v", pubKeyErr)
	}

	return nil
}

func loadPrivateKey() error {
	bytes, err := ioutil.ReadFile(Env.PrivateKeyPath)
	if err != nil {
		return fmt.Errorf("Cannot load file private_key.pem, err: %v", err)
	}

	data, _ := pem.Decode(bytes)

	key, err := x509.ParsePKCS8PrivateKey(data.Bytes)
	if err != nil {
		return fmt.Errorf("Cannot parse private key, err: %v", err)
	}

	privKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return fmt.Errorf("Private key not of type ed25519 private key")
	}

	Env.PrivateKey = privKey

	return nil
}

func loadPublicKey() error {
	bytes, err := ioutil.ReadFile(Env.PublicKeyPath)
	if err != nil {
		return fmt.Errorf("Cannot load file public_key.pem, err: %v", err)
	}

	data, _ := pem.Decode(bytes)

	key, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return fmt.Errorf("Cannot parse public key, err: %v", err)
	}

	pubKey, ok := key.(ed25519.PublicKey)
	if !ok {
		return fmt.Errorf("Public key not of type ed25519 public key")
	}

	Env.PublicKey = pubKey

	return nil
}
