package queues

import (
	"github.com/blog-jobs/internal/tools"
	"github.com/go-redis/redis/v7"

	"encoding/json"

	"github.com/blog-jobs/internal/jobs"
)

type RedisQueue struct {
	client *redis.Client
	key    string
}

func NewRedisQueue(addr, key string) *RedisQueue {
	opt, _ := redis.ParseURL(addr)
	rdb := redis.NewClient(opt)

	return &RedisQueue{client: rdb, key: key}
}

func (rq *RedisQueue) Push(job jobs.Job, queue string) error {
	queueKey := tools.JobQueueKey(queue)
	jobMap, err := job.ToMap()
	if err != nil {
		return err
	}

	jobMapJsonBytes, err := json.Marshal(jobMap)
	if err != nil {
		return err
	}

	jobMapJson := string(jobMapJsonBytes)

	return rq.client.LPush(queueKey, jobMapJson).Err()
}

func (rq *RedisQueue) Pop() (jobs.Job, error) {
	queueKey := tools.JobQueueKey("default")
	jobStr, err := rq.client.LPop(queueKey).Result()

	if err != nil {
		return nil, err
	}

	var jobMap map[string]string
	json.Unmarshal([]byte(jobStr), &jobMap)

	job := buildJob(jobMap)

	return job, err
}

func buildJob(jobMap map[string]string) jobs.Job {
	if jobMap == nil {
		return nil
	}

	var job jobs.Job
	switch jobMap["ExecutedJobName"] {
	case "HealthcheckJob":
		job = jobs.NewHealthcheckJob(jobMap["PriorityType"])
	default:
		job = nil
	}

	return job
}
