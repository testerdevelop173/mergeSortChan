package main

func mergeSortedChannels(ch1, ch2 chan int) chan int {
}

func fillCh1(ch chan int) {
	ch <- 1
	ch <- 3
	ch <- 4
	close(ch)

}

func fillCh2(ch chan int) {
	ch <- -3
	ch <- 1
	ch <- 10
	close(ch)

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go fillCh1(ch1)
	go fillCh2(ch2)
	outMergChan := mergeSortedChannels(ch1, ch2)

}
