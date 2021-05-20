package server

import (
	"opencodes/model/db"
	"sync"
	"time"
)

var (
	Wg        *sync.WaitGroup
	QueueUser chan *db.User
	locTime   int64
	Counts    int
)

func InitUserQueue() {

	Wg = &sync.WaitGroup{}
	QueueUser = make(chan *db.User, 1000)
	locTime = time.Now().Unix() + 60
	Counts = 0
	go work()
}

func work() {

	defer func() {
		if err := recover(); err != nil {

		}
	}()

	for {
		if locTime == time.Now().Unix() {
			locTime = time.Now().Unix() + 60

			Counts = 0
		}
		select {
		case user, ok := <-QueueUser:
			if !ok {

				return
			}

			go createUser(user)
		}
	}
}

func createUser(user *db.User) {
	defer func() {
		Wg.Done()

		if err := recover(); err != nil {

		}
	}()
	err := db.CreateUser(user)
	if err != nil {

	} else {
		Counts++
	}

}
