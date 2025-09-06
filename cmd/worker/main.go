package main

import (
	"log"
	"os"

	"github.com/blog-jobs/internal/config"
	"github.com/blog-jobs/internal/queues"
	"github.com/blog-jobs/internal/workers"
	"github.com/blog-jobs/schedule"
)

func main() {
	queueClient := queues.NewRedisQueue(config.RedisUrl(), "default")

	schedule := schedule.NewSchedule(queueClient)
	dispatcher := workers.NewDispatcher(queueClient, config.NumberOfWorkers())

	file, err := os.OpenFile("internal/logger/logs/"+config.Environment(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Fail to open log file: %s", err)
	}
	defer file.Close()

	log.SetOutput(file)

	go dispatcher.Run()
	schedule.Run()
}
