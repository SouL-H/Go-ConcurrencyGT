package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Farklı goroutinlerin aynı hafızaya erişmeye çalışmasına Race-Condition denir.
func main() {

	raceExampleMutex()
	//raceExampleAtomic()
}

func raceExampleMutex() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	//Aynı anda erişimleri engellemek için mutexle bu hatayı önleyebiliriz.
	lock := sync.Mutex{}
	val := 0

	go func() {
		for i := 0; i < 1000000; i++ {
			//Memory kitleyip açıyoruz tek bir yer erişiyor bu şekilde.
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()
		wg.Wait()
	fmt.Println(val)

}

//2.yöntem Atomic ile çözülür
func raceExampleAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}
