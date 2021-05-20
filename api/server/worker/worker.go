package worker

import (
	"fmt"
	orm "opencodes/pkg/database"
)

type Workers struct {
	retries int
	types   string
	status  int
	params  interface{}
}

var (
	//Wg          *sync.WaitGroup
	QueueWorker chan *Workers
)

func NewWorkers(r int, ty string, s int) *Workers {
	return &Workers{
		retries: r,
		types:   ty,
		status:  s,
		params:  nil,
	}
}

func WorkSetUp() {

	//Wg = &sync.WaitGroup{}
	QueueWorker = make(chan *Workers, 1000000)

	go work()
}

func work() {
	defer func() {
		if err := recover(); err != nil {

		}
	}()

	for {

		select {
		case work, ok := <-QueueWorker:
			if !ok {

				return
			}
			err := orm.Eloquent.Exec("select * from t_user limit 1").Error
			fmt.Println(err)
			DoWork(work)
		}
	}
}

func DoWork(w *Workers) {
	if w.retries == 3 {
		return
	}

}
