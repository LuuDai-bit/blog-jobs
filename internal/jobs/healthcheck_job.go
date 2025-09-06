package jobs

import (
	"fmt"
	"time"
)

type HealthcheckJob struct {
	Job
	Id             string
	PriorityType   string
	ExecutedJobName string
}

func NewHealthcheckJob(priorityType string) *HealthcheckJob {
	if priorityType == "" {
		priorityType = "default"
	}

	return &HealthcheckJob{Id: time.Now().String(),
		PriorityType:   priorityType,
		ExecutedJobName: "HealthcheckJob"}
}

func (h *HealthcheckJob) Handle() error {
	fmt.Printf("The background job system is running")

	return nil
}

func (h *HealthcheckJob) ToMap() (map[string]string, error) {
	jobMap := map[string]string{
		"Id": h.Id,
		"PriorityType": h.PriorityType,
		"ExecutedJobName": h.ExecutedJobName,
	}

	return jobMap, nil
}
