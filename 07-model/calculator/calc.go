package calculator

import "utils"

import "fmt"

var opCount int

func init() {
	fmt.Println("calculator.init invoked")
}

func GetOpCount() int {
	return opCount
}
