package main

import "time"

//Birden fazla chan input geliyorsa tek bir chan ile input vermek için birleştirmek için kullanılan yapıdır.

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	outChan := make(chan int, 10000)

	done := make(chan int)
	//Bütün işlemi yöneten channel
	go func() {
		for {
			println("çalışma bekliyor.")
			val, open := <-outChan
			if !open {
				break
			}
			println(val)
		}
		done <- 1
	}()

	go func() {
		for {
			c1 <- 1
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			c2 <- 2
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			c3 <- 3
			time.Sleep(time.Second)
		}
	}()
		//3 channeli al ve tek channel çıktı ver.
	fanIn([]chan int{c1,c2,c3},outChan)
	<-done
}

func fanIn(inChans []chan int,outputChan chan int){
	for _,ch :=range inChans{
		go func(c chan int){
			for{
				val,open := <-c
				if !open{
					break
				}
				outputChan<-val
			}
		}(ch)
	}
}
