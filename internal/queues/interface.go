package queues

import "github.com/blog-jobs/internal/jobs"

type Queue interface {
	Push(job jobs.Job, queue string) error
	Pop() (jobs.Job, error)
}
