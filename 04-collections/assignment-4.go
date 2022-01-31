package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type operation func(int, int) int
type operationsCollection map[int]operation

/* var operations operationsCollection = operationsCollection{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
} */

var operations operationsCollection

func main() {
	operations = make(operationsCollection)
	operations[1] = add
	operations[2] = subtract
	operations[3] = multiply
	operations[4] = divide

	for {
		userChoice, err := getUserChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		n1, n2, err := getOperands()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := performOperation(userChoice, n1, n2)
		fmt.Println("result = ", result)
	}
}

func performOperation(userChoice, n1, n2 int) (result int) {
	if operation, exists := operations[userChoice]; exists {
		result = operation(n1, n2)
		return
	}
	return
}

func getOperands() (n1, n2 int, err error) {
	fmt.Println("Enter the No1 :")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		err = errors.New("Invalid choice")
		return // no input
	}
	_, err = fmt.Sscanf(scanner.Text(), "%d", &n1)
	if err != nil {
		return
	}
	fmt.Println("Enter the No2 :")
	if !scanner.Scan() {
		err = errors.New("Invalid choice")
		return // no input
	}
	_, err = fmt.Sscanf(scanner.Text(), "%d", &n2)
	return
}

func getUserChoice() (userChoice int, err error) {
	fmt.Println("Enter the choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		err = errors.New("Invalid choice")
		return // no input
	}
	_, err = fmt.Sscanf(scanner.Text(), "%d", &userChoice)
	return
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
