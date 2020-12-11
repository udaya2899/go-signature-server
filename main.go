package main

import (
	"log"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/handler"
	"challenge.summitto.com/udaya2899/challenge_result/logger"
)

func main() {

	logger.Logger.Printf("Initializing Signature Server")

	ginEngine := handler.New()
	if err := ginEngine.Run(config.Env.Port); err != nil {
		log.Fatalf("Cannot start signature server, Error: %v", err)
	}

	logger.Logger.Printf("Signature Server started... Listening to PORT: %v", config.Env.Port)
}
