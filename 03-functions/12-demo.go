/* panic & recovery */

package main

import (
	"fmt"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("recovered from panic")
			fmt.Println(e)
		} else {
			fmt.Println("program exiting successfully")
		}
	}()
	result, err := divideClient(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result =", result)
}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			err = fmt.Errorf("unable to divide %d by %d", x, y)
		}
	}()
	result, _ = divide(x, y)
	return
}

//3rd party api
func divide(x, y int) (quotient int, remainder int) {
	if y == 0 {
		panic("divisor cannot be zero")
	}
	quotient = x / y
	remainder = x % y
	return
}
