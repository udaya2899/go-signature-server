package main

import (
	"fmt"
	"log"
	"runtime"

	"challenge.summitto.com/udaya2899/challenge_result/config"
	"challenge.summitto.com/udaya2899/challenge_result/handler"
)

func main() {

	log.Printf("CPU Cores available in machine: %v\n", runtime.NumCPU())
	log.Printf("CPU Cores that can be used by the program: %v", maxParallelism())

	ginEngine := handler.New()
	if err := ginEngine.Run(fmt.Sprintf(":%s", config.Env.Port)); err != nil {
		log.Fatalf("Cannot start signature server, Error: %v", err)
	}

	log.Printf("Signature Server started... Listening to PORT: %v", config.Env.Port)
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
