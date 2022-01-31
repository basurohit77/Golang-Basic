package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	evenNoCh := genEvenNos(done)
	go func() {
		time.Sleep(20 * time.Second)
		done <- true
	}()

	for evenNo := range evenNoCh {
		fmt.Println("evenNo : ", evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos(done chan bool) <-chan int {
	resultCh := make(chan int)
	go func() {
		var no int
		for {
			select {
			case resultCh <- no:
				time.Sleep(500 * time.Millisecond)
				no += 2

			case <-done:
				close(resultCh)
				return
			}
		}
	}()
	return resultCh
}
