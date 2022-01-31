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
	tick := time.Tick(500 * time.Millisecond)
	go func() {
		var no int
		for {
			select {
			case <-tick:
				resultCh <- no
				no += 2
			case <-done:
				close(resultCh)
				return
			}
		}
	}()
	return resultCh
}
