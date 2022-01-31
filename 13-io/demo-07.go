package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	ptrReader, err := os.Open("data2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer ptrReader.Close()

	ptrScanner := bufio.NewScanner(ptrReader)  // \n as deliminator and scanner pointer
	lineCount := 0
	for ptrScanner.Scan() {
		lineCount++
		fmt.Printf("line #%d : %s\n", lineCount, ptrScanner.Text())
	}
}
