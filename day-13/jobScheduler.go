package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
*/

var mu sync.Mutex // concurrency lock
type scheduleNode struct {
	F    func()        // function
	D    int64         // delay
	Next *scheduleNode // next node
}

var scheduleRoot *scheduleNode
var wg sync.WaitGroup
var timer time.Time

func init() {
	timer = time.Now()
	go doJobs()
	wg.Add(1)
}

func main() {
	defer wg.Wait()
	f1 := func() { fmt.Println("hey 2") }   // 1000
	f2 := func() { fmt.Println("hey 3") }   // 1500
	f3 := func() { fmt.Println("hey 1") }   // 500
	f4 := func() { fmt.Println("hey 0") }   // 10
	f5 := func() { fmt.Println("hey 0.5") } // 100
	f6 := func() { fmt.Println("hey 4") }   // 2000

	schedule(f1, 1000)
	schedule(f2, 1500)
	schedule(f3, 500)
	schedule(f4, 10)
	schedule(f5, 100)
	schedule(f6, 2000)
}

func schedule(f func(), d int64) {
	nj := &scheduleNode{ // new job
		F: f,
		D: d + getTimeSince(),
	}
	mu.Lock()

	if scheduleRoot == nil {
		scheduleRoot = nj
	} else {
		t := scheduleRoot // target
		s := false        // skip
		for t.D < nj.D {
			if t.Next == nil {
				t.Next = nj
				s = true
				break
			} else {
				t = t.Next
			}
		}
		if !s {
			tmp := *t
			*t = *nj
			t.Next = &tmp
		}
	}
	mu.Unlock()
}

func getTimeSince() int64 {
	return time.Now().Sub(timer).Milliseconds()
}

func doJobs() {
	defer wg.Done()
	for {
		if scheduleRoot != nil {
			mu.Lock()
			t := getTimeSince() // time

			if t > scheduleRoot.D {
				scheduleRoot.F()
				scheduleRoot = scheduleRoot.Next
			}

			mu.Unlock()
		}
	}
}
