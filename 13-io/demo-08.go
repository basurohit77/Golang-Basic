package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ptrScanner := bufio.NewScanner(os.Stdin)
	for ptrScanner.Scan() {
		data := ptrScanner.Text()
		if data == "exit" {
			break
		}
		fmt.Println(ptrScanner.Text())
	}
}
