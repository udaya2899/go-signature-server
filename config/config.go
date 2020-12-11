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
	LogFilePath    string `json:"log_file_path"`
	PublicKey      ed25519.PublicKey
	PrivateKey     ed25519.PrivateKey
}

// Env variable has the config loaded in it on init()
var Env envConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Panicf("Cannot Load Config, Err: %v", err)
	}

	err = loadKey()
	if err != nil {
		log.Panicf("Cannot load keys from config path, Err: %v", err)
	}

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
