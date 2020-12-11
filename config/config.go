package config

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type envConfig struct {
	Port           string `json:"port"`
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath  string `json:"public_key_path"`
	LogFileName    string `json:"log_file_name"`
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
