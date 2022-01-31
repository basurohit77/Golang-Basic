package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter the name :")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Hi %s, Have a nice day!\n", name)

	var no int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d\n", &no)
	if no%2 == 0 {
		fmt.Printf("%d is even\n", no)
	} else {
		fmt.Printf("%d is odd\n", no)
	}

	var n1, n2 int
	fmt.Println("Enter two numbers")
	fmt.Scanf("%d,%d\n", &n1, &n2)
	fmt.Println(n1 + n2)
}
