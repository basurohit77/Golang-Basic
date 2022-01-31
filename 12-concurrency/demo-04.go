package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var opCount int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i, wg)
	}
	wg.Wait()
	fmt.Println("OpCount = ", opCount)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	{
		opCount++  // opCount is a share variable
	}
	mutex.Unlock()
	fmt.Println(x + y)
	wg.Done()
}
