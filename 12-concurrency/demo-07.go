package main

import (
	"fmt"
	"sync"
	"time"
)

//common memory for communicating the result (from the add function)
//Communicate by sharing memory (NOT PREFERRED)
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)

}

func add(x, y int, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	result = x + y
	wg.Done()

}
