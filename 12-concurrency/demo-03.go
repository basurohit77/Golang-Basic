package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go f1(&wg)
		}
	*/
	wg := &sync.WaitGroup{}  //pointer
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go f1(wg)
	}
	f2()
	wg.Wait()
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)

	fmt.Println("f1 invocation completed")
	wg.Done()
}

func f2() {

	fmt.Println("f2 invoked")
}
