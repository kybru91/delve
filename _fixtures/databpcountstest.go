package main

import (
	"math/rand"
	"sync"
	"time"
)

var globalvar1 int

func demo(id int, wait *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		sleep := rand.Intn(10) + 1
		globalvar1 = globalvar1 + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}

	wait.Done()
}

func main() {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	wait.Add(1)
	go demo(1, wait)
	go demo(2, wait)

	wait.Wait()
}
