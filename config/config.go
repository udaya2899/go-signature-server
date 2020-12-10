package config

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type envConfig struct {
	Port           string `json:"port"`
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath  string `json:"public_key_path"`
	PublicKey      ed25519.PublicKey
	PrivateKey     ed25519.PrivateKey
}

// Env variable has the config loaded in it on init()
var Env envConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Panicf("Cannot Load Config")
	}

	loadKey()

	log.Printf("Config file loaded successfully")
}

// loadConfig loads the config vars from $PWD/config.json
func loadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("Cannot open config.json, Err: %v", err)
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Cannot convert file to bytes, Err: %v", err)
	}

	err = json.Unmarshal(byteValue, &Env)
	if err != nil {
		return fmt.Errorf("Cannot decode config JSON, Err: %v", err)
	}

	return nil
}

func loadKey() {
	loadPrivateKey()
	loadPublicKey()
}

func loadPrivateKey() {
	bytes, err := ioutil.ReadFile(Env.PrivateKeyPath)
	if err != nil {
		log.Fatalf("Cannot load file private_key.pem, err: %v", err)
	}

	data, _ := pem.Decode(bytes)

	key, err := x509.ParsePKCS8PrivateKey(data.Bytes)
	if err != nil {
		log.Fatalf("Cannot parse private key, err: %v", err)
	}

	privKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		log.Fatalf("Private key not of type ed25519 private key")
	}

	Env.PrivateKey = privKey
}

func loadPublicKey() {
	bytes, err := ioutil.ReadFile(Env.PublicKeyPath)
	if err != nil {
		log.Fatalf("Cannot load file public_key.pem, err: %v", err)
	}

	data, _ := pem.Decode(bytes)

	key, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		log.Fatalf("Cannot parse public key, err: %v", err)
	}

	pubKey, ok := key.(ed25519.PublicKey)
	if !ok {
		log.Fatalf("Public key not of type ed25519 public key")
	}

	Env.PublicKey = pubKey
}
