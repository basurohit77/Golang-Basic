package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-stop:
			//tick = time.Tick(0 * time.Second)
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}