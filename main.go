package main

import (
	"log"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/handler"
)

func main() {

	ginEngine := handler.New()
	if err := ginEngine.Run(config.Env.Port); err != nil {
		log.Fatalf("Cannot start signature server, Error: %v", err)
	}

	log.Printf("Signature Server started... Listening to PORT: %v", config.Env.Port)
}
