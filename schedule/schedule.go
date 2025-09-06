package schedule

import (
	"github.com/blog-jobs/internal/jobs"
	"github.com/blog-jobs/internal/queues"

	"time"
)

type Schedule struct {
	queue queues.Queue
}

func NewSchedule(queue queues.Queue) *Schedule {
	return &Schedule{queue: queue}
}

func (s *Schedule) Run() {
	healthcheckJob := jobs.NewHealthcheckJob("default")

	var now time.Time
	for {
		now = time.Now()

		if now.Hour() == 0 {
			s.queue.Push(healthcheckJob, "default")
		}
	}
}
