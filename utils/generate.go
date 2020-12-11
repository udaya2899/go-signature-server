package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	generateAndWriteKeyPair()
}

func generateAndWriteKeyPair() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Cannot generate ed25519 keypair, err: %v", err)
	}

	// Create pem directory if not exists
	if _, err := os.Stat("pem"); os.IsNotExist(err) {
		os.Mkdir("pem", os.ModeDir)
	}

	writePrivateKey(privateKey)
	writePublicKey(publicKey)
}

func writePrivateKey(privateKey ed25519.PrivateKey) {
	pemPrivateFile, err := os.Create("pem/private_key.pem")
	if err != nil {
		log.Fatalf("Cannot create pem file: %v", err)
	}

	defer pemPrivateFile.Close()

	key, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Cannot marshal private key, Err: %v", err)
	}

	pemPrivateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: key,
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		log.Fatalf("Cannot encode PEM block to file, err: %v", err)
	}
}

func writePublicKey(publicKey ed25519.PublicKey) {
	pemPublicFile, err := os.Create("pem/public_key.pem")
	if err != nil {
		log.Fatalf("Cannot create pem file: %v", err)
	}

	defer pemPublicFile.Close()

	pubKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("Cannot marshal public key, Err: %v", err)
	}

	pemPublicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKey,
	}

	err = pem.Encode(pemPublicFile, pemPublicBlock)
	if err != nil {
		log.Fatalf("Cannot encode PEM block to file, err: %v", err)
	}
}
