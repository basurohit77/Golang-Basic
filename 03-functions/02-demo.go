//anonymous functions

package main

import "fmt"

func main() {
	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		fmt.Println("add result =", x+y)
	}(100, 200)

	subResult := func(x, y int) int {
		return x - y
	}(100, 200)

	fmt.Println("subtract result =", subResult)
}
