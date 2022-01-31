package main

import (
	"fmt"
	"time"
)

//Share memory by communicating (using channels)
func main() {

	ch := make(chan int) //non-buffered channel
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)

}

func add(x, y int, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
}
