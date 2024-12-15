package main

func mergeSortedChannels(ch1, ch2 chan int) chan int {
}

func fillChannel(input []int) chan int {
	out := make(chan int)

	go func() {
		for _, value := range input {
			out <- value
		}
		close(out)
	}()

	return out
}

func main() {
	in1 := []int{1, 3, 4, 5}
	in2 := []int{-3, 1, 10}

	ch1 := fillChannel(in1)
	ch2 := fillChannel(in2)
	outMergChan := mergeSortedChannels(ch1, ch2)

}
