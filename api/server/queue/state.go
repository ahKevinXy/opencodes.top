package queue

type State int

const (
	Unacked State = iota
	Acked
	Rejected
	Pushed
)
