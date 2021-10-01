package main

import (
	"fmt"
	"time"
)

//channel okuma yazmanın bloking olduğu bir typedir. Channel bir yönden lock'lanan bir arraydir.
//channel ikiye ayrılır buffer ve unbuffer.
//unbuffer her okuma ve yazma işlemi blokkingdir. Eş zamanlı bir data yazıp okunur.
func main() {

	//unbuffer()
	unbufferReadWrite()
}

func unbuffer() {
	unbufferedChan := make(chan int)

	var unbufferedChan2 chan int
	fmt.Println(unbufferedChan)  //adress
	fmt.Println(unbufferedChan2) //nil

	go func(unbufChan <-chan int) {
		//Veri gelene kadar blokla.
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	unbufferedChan <- 1
}

func unbufferReadWrite() {
	unbufferedChan := make(chan int)
	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)
	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)
	//İşlem hızlı gerçekleştiği için çıktığı görme şansımız olmayabilir
	//Bu yüzden biraz cpu uyutuyoruz.
	time.Sleep(time.Second)
	fmt.Println("Selam kanal.")

}
