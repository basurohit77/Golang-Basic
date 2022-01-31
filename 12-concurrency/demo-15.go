package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "data - 1"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "data - 2"
	}()
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()
	var stop bool = false
	for !stop {
		select {
		case d1 := <-ch1:
			fmt.Println("Data from ch1 - ", d1)
		case d2 := <-ch2:
			fmt.Println("Data from ch2 - ", d2)
		case <-done:
			fmt.Println("Done")
			stop = true
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("no channel operations were successful")
		}
	}

}
