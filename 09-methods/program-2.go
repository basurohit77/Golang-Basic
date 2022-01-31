package main

import "fmt"

type myStr string   /* alias for string */

type fn func(int) bool

func main() {
	s := myStr("This is a sample string") // converting the string to myStr
	fmt.Println(s)
	fmt.Println(s.Length())
}

func (s myStr) Length() int {
	return len(s)
}
