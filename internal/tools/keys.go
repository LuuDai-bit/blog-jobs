package tools

func JobQueueKey(name string) string {
	if name == "" {
		name = "default"
	}

	return "job-queue:" + name
}
