package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.After(20 * time.Second)
	evenNoCh := genEvenNos(done)
	for evenNo := range evenNoCh {
		fmt.Println("evenNo : ", evenNo)
	}
	fmt.Println("Done")
}

func genEvenNos(done <-chan time.Time) <-chan int {
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