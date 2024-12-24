package main

import "fmt"

func mergeSortedChannels(ch1, ch2 chan int) chan int {
	outMergChan := make(chan int)

	go func() {
		defer close(outMergChan)

		go func() {
			for v := range ch1 {
				outMergChan <- v
			}
		}()

		go func() {
			for v := range ch2 {
				outMergChan <- v
			}
		}()
	}()

	return outMergChan
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
	in1 := []int{1, 3, 4, 7, 8, 22, 90, 5}
	in2 := []int{-3, 1, 10, 1, 12, 12, 90}

	ch1 := fillChannel(in1)
	ch2 := fillChannel(in2)

	fmt.Println("the result is:", mergeSortedChannels(ch1, ch2))

}

// требуется: объединить 2 канала ch1, ch2 и
//убрать дублирование. Чтобы каждый элемнт был уникальным
