package utils

//import "fmt"
//import calc1 "modularity-demo/calculator"


func IsPrime(no int) bool {
	//fmt.Println(calc1.Add(10, 20))
	if no < 0 {
		return false
	}
	if no <= 3 {
		return true
	}
	for i := 2; i < no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
