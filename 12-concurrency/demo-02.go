package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //here global variable

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go f1()
	}
	f2()

	wg.Wait()
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)

	fmt.Println("f1 invocation completed")
	wg.Done()
}

func f2() {

	fmt.Println("f2 invoked")
}
