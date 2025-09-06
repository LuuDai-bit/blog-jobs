package jobs

type Job interface {
	Handle() error
	ToMap() (map[string]string, error)
}
