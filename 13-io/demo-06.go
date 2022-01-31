package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	ptrReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer ptrReader.Close()
	bufferedReader := bufio.NewReader(ptrReader)
	sentenceCount := 0
	for {
		sentence, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		sentenceCount++
		fmt.Printf("sentence #%d: %s\n", sentenceCount, sentence)
	}
}
