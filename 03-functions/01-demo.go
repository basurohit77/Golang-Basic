package main

import "fmt"

var counter int

func main() {
	fmt.Println(isPrime(100))
	fmt.Println(add(100, 200))
	increment()
	fmt.Printf("counter = %d\n", counter)
	increment()
	fmt.Printf("counter = %d\n", counter)
	increment()
	fmt.Printf("counter = %d\n", counter)
	increment()
	fmt.Printf("counter = %d\n", counter)
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("quotient = %d, remainder = %d\n", quotient, remainder)
	*/
	quotient, _ := divide(100, 7)
	fmt.Printf("quotient = %d\n", quotient)
}

func increment() {
	counter++
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	/*
		var result int
		result = x + y
	*/
	result := x + y
	return result
}

/* func add(x, y int) (result int) {
	result = x + y
	return
} */

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
