package main

func main() {
	//selectChan()
	//multiSelectChan()
	multiSelectChanV2()

}

func selectChan() {
	chan1 := make(chan int, 1)
	chan1 <- 1
	select {
	//kanala veri girişini bekletebiliyoruz.
	case c1Val := <-chan1:
		println(c1Val)
	}
}

func multiSelectChan() {
	chan1 := make(chan int, 1)
	chan1 <- 1
	chan2 := make(chan int, 2)
	chan2 <- 2

	select {
	//Kod burada sırayla okunmaz go random case seçer.
	case c1Val := <-chan1:
		println(c1Val)

	case c2Val := <-chan2:
		println(c2Val)
	}
}

//V1 sadece tek bir veri okuyordu. Biz hepsini okumamız lazımsa.
func multiSelectChanV2() {

	chan1 := make(chan int, 1)
	chan1 <- 1
	chan2 := make(chan int, 2)
	chan2 <- 2
	var done bool
	for !done {
		select {
		//Kod burada sırayla okunmaz go random case seçer.
		case c1Val := <-chan1:
			println(c1Val)

		case c2Val := <-chan2:
			println(c2Val)
		default:
			done = true
		}
	}

}
