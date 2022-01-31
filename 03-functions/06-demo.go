/* Higher order functions */
package main

import "fmt"

func main() {
	/* c. functions returned from other functions */
	adder := getAdder()
	fmt.Println(adder(100, 200))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
