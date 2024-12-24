package main

import "fmt"

func mergeSortedChannels(ch1, ch2 chan int, len1, len2 int) chan int {
    outMergChan := make(chan int)

    go func() {
        defer close(outMergChan)

        for i := 0; i < len1+len2; i++ {
            select {
            case v, ok := <-ch1:
                if ok {
                    outMergChan <- v
                }
            case v, ok := <-ch2:
                if ok {
                    outMergChan <- v
                }
            }
        }
    }()

    return outMergChan
}
//доработка , работать с безразмерным 

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


func delDubl(inMergChan <-chan int) chan int {
    
	//outDedupedChan := make(chan int)
	
	resultMergChan :=make (chan int)

	for _, value:= range inMergChan{
		

	}
     return resultMergChan
 }



func main() {
    in1 := []int{1, 3, 4, 7, 8, 9, 12, 15}
    in2 := []int{2, 4, 5, 6, 8, 10, 12, 14}

    ch1 := fillChannel(in1)
    ch2 := fillChannel(in2)

    len1 := len(in1)
    len2 := len(in2)

	delDubl(outMergChan)

    fmt.Println("the result is:")
    mergedChan := mergeSortedChannels(ch1, ch2, len1, len2)
    for v := range mergedChan {
        fmt.Println(v)
    }
}