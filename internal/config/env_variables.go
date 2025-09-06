package config

import (
	"log"
	"os"
	"strconv"
)

func RedisUrl() string {
	return os.Getenv("REDIS_URL")
}

func NumberOfWorkers() int {
	workersCount, err := strconv.Atoi(os.Getenv("NUMBER_OF_WORKER"))
	if err != nil {
		log.Printf("Error when convert number of worker to int")
		workersCount = 1
	}

	return workersCount
}

func Environment() string {
	return os.Getenv("ENVIRONMENT")
}
