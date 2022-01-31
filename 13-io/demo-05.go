package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteArray, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray))
}
