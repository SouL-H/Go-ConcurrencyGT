package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

//Race-condition durumu var.
func main() {

	raceTest := &RaceTest{}
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		lock.Lock()
		go increment(raceTest, wg, lock)
		lock.Unlock()
	}

	wg.Wait()
	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, lk *sync.Mutex) {
	lk.Lock()
	rt.Val += 1
	lk.Unlock()
	wg.Done()
}
