package main

import (
	"fmt"
	"sync"
)

/* var mutex sync.Mutex
var opCount int */

type OpCount struct {  // Make the share variable as a struct with mutex
	count int
	sync.Mutex
}

var opCount OpCount

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i, wg)
	}
	wg.Wait()
	fmt.Println("OpCount = ", opCount.count)
}

func add(x, y int, wg *sync.WaitGroup) {
	opCount.Lock()
	{
		opCount.count++
	}
	opCount.Unlock()
	//fmt.Println(x + y)
	wg.Done()

}
