/* panic & recovery */

package main

import "fmt"

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
	result := divide(100, 0)
	e := recover()
	fmt.Println("result =", result)
}

func divide(x, y int) (result int) {
	result = x / y
	return
}
