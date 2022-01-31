package main

import (
	"fmt"
	"sync"
	"time"
)

//Share memory by communicating (using channels)
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //non-buffered channel
	//wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	//wg.Wait()
	fmt.Println(result)

}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
	//wg.Done()

}
