package main

import (
	"fmt"
	"time"
)

func main() {
	evenNoCh := genEvenNos()
	for evenNo := range evenNoCh {
		fmt.Println("evenNo : ", evenNo)
	}
}

func genEvenNos() <-chan int {
	resultCh := make(chan int)
	go func() {
		var no int
		for i := 0; i < 20; i++ {
			time.Sleep(500 * time.Millisecond)
			no += 2
			fmt.Println("Generated : ", no)
			resultCh <- no
		}
		close(resultCh)
	}()
	return resultCh
}
