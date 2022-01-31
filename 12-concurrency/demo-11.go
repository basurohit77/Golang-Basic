package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Assignment:
		Print "Hello" & "World" simultaneously
		IMPORTANT : Do NOT move the loop outside the "print" function
*/

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan bool,1)
	ch2 := make(chan bool)
	go print("Hello", ch1, ch2, wg)
	go print("World", ch2, ch1, wg) // here when the last world is printed, ch1 <- true, now go scheduler find no
	                                  // go subroutine to handel it, Hence ch1 is a buffer channel
	ch1 <- true
	wg.Wait()
	fmt.Println("End of main!")
}

func print(s string, ch1, ch2 chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
		ch2 <- true
	}
	wg.Done()
}
