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
	lock := sync.Mutex{}
	raceTest := &RaceTest{}
	wg := &sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		lock.Lock()
		go increment(raceTest, wg)
		lock.Unlock()
	}
	wg.Wait()
	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup,) {
	rt.Val += 1

	wg.Done()
}
