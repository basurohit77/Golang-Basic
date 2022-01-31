/* Closures */
package main

import "fmt"

func main() {

	fn()
	fn()
	fmt.Println("main completed")

}

func getIncrement() func() int {
	var counter int
	increment := func() int {
		counter++
		return counter
	}
	return increment
}

func fn() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}
