package test

import (
	"sync"
	"testing"
)

var once sync.Once
var a string

func setUp() {
	a = "hello world"
}
func twoprint() {
	go doprint()
	go doprint()
}
func doprint() {
	once.Do(setUp)
	println(a)
}
func TestOnce(t *testing.T) {
	twoprint()
}
