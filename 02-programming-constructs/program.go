package main

import "fmt"

func main() {
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is even\n", no)
		} else {
			fmt.Printf("%d is odd\n", no)
		}
	*/

	if no, d := 21, 2; no%d == 0 {
		fmt.Printf("%d is even\n", no)
		fmt.Println(d)
	} else {
		fmt.Printf("%d is odd\n", no)
	}

	//for
	//variation - 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %v\n", sum)

	//variation - 2 (using like a while loop)
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %v\n", numSum)

	//variation - 3 (infinite loop)
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Printf("x = %v\n", x)

	//switch construct
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 	=> "Excellent"
		otherwise 	=> "Invalid score"
	*/

	score := 7
	/*
		switch score {
		case 0:
		case 1:
		case 2:
		case 3:
			fmt.Println("Terrible")

		case 4:
		case 5:
		case 6:
		case 7:
			fmt.Println("Not Bad")

		case 8:
		case 9:
			fmt.Println("Good")

		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/
	/*
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not Bad")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/
	switch {
	case score >= 0 && score <= 3:
		fmt.Println("Terrible")
	case score >= 4 && score <= 7:
		fmt.Println("Not Bad")
	case score == 8 || score == 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	//falthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	default:
		fmt.Println("default")
	}
}
