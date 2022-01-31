/* Higher order functions */

package main

import "fmt"

func main() {
	/* b. Functions can be passed as arguments to other functions */

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(multiply, 100, 200)
		logOperation(divide, 100, 200)
	*/
	loggerAdd := logOperation(add)
	loggerAdd(100, 200)
}

/*
on invoking the add / subtract / multiply / divide operation
expected Result:
	Before invocation
	result = 100
	After invocation
*/

/*
func logAdd(x, y int) {
	fmt.Println("Before invocation")
	result := add(x, y)
	fmt.Println("result = ", result)
	fmt.Println("After invocation")
}

func logSubtract(x, y int) {
	fmt.Println("Before invocation")
	result := subtract(x, y)
	fmt.Println("result = ", result)
	fmt.Println("After invocation")
}
*/

func logOperation(oper func(int, int) int) func(int, int) {
	decoratedOperation := func(x, y int) {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("result = ", result)
		fmt.Println("After invocation")
	}
	return decoratedOperation
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
