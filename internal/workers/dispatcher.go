package workers

import (
	"log"
	"time"

	"github.com/blog-jobs/internal/queues"
)

type Dispatcher struct {
	queue   queues.Queue
	workers int
}

func NewDispatcher(q queues.Queue, workers int) *Dispatcher {
	return &Dispatcher{queue: q, workers: workers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.workers; i++ {
		go d.StartWorker(i)
	}

	select {}
}

func (d *Dispatcher) StartWorker(id int) {
	for {
		// Handle job for default queue
		job, err := d.queue.Pop()
		sleepTime := 5

		if err != nil {
			log.Printf("Worker %d: error fetching job: %v", id, err)
			sleepTime = 10
			time.Sleep(time.Duration(sleepTime) * time.Second)
			continue
		}

		if err := job.Handle(); err != nil {
			log.Printf("Worker %d: job failed: %v", id, err)
			sleepTime = 1
			time.Sleep(time.Duration(sleepTime) * time.Second)
			continue
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
