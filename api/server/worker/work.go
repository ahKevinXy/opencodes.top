package worker

type Worker interface {
	Do()
	Retry()
	Success()
	Errors()
	Del()
}
