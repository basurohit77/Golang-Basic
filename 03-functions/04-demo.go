/* Higher order functions */
package main

import "fmt"

func main() {
	/* a. functions can be assigned to variabled */
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	/*
		var fn = func() {
			fmt.Println("fn invoked")
		}
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

}
