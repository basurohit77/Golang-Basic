/* error handling */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Dont try to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("Something went wront!")
		fmt.Println(err)
		return
	}
	//result, _ := divide(100, 0)
	fmt.Println("result = ", result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		//return 0, fmt.Errorf("%d cannot be divided by zero", x)
		return 0, DivideByZeroError
	}
	return x / y, nil
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		//return 0, fmt.Errorf("%d cannot be divided by zero", x)
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
