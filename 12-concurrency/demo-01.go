package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	f2()

	/*
		var input string
		fmt.Scanln(&input)
	*/

	time.Sleep(1 * time.Millisecond)
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
