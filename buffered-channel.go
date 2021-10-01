package main

import "fmt"

//Buffered channelde boyut veriyoruz.
//Chan kullanılırken wg kullanılmalıdır. Çünkü bekleme şansı olmayabiliyor.
func main() {
	// bufferedChan()
	// bufferChanGoRoutine()
	bufferChanGoRoutineV2()
}

func bufferedChan() {
	bufferedChannel := make(chan int, 5)

	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	bufferedChannel <- 4
	bufferedChannel <- 5
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}

func bufferChanGoRoutine() {
	bufferedChannel := make(chan int, 5)

	go func(bufChan chan int) {
		for {
			value := <-bufChan
			fmt.Println(value)
		}
	}(bufferedChannel)
	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	bufferedChannel <- 4
	bufferedChannel <- 5
	bufferedChannel <- 6
	bufferedChannel <- 7
	bufferedChannel <- 8
	bufferedChannel <- 9
	//Halen güvensiz bir kanal çünkü ne kadar veri yazacağını kestirmek zor.
}

func bufferChanGoRoutineV2() {
	bufferedChan := make(chan int, 5)
	done := make(chan bool)
	go func(bufChan chan int, done chan bool) {
		for val := range bufChan {
			println(val)
		}
		println("kanal kapandı.")
		done <- true
	}(bufferedChan, done)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	//Fakat bu şekilde bir işlemde kanalı kapanması lazım. Yoksa hata alınacaktır.
	close(bufferedChan)
	<-done
	println("Main bitti.")
}
