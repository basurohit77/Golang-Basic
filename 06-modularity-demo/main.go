package main

/*
import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Printf("Op Count = %d\n", calculator.GetOpCount())
	fmt.Println(utils.IsPrime(97))
}
*/


import (
	"fmt"
	"github.com/fatih/color"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
	_ "modularity-demo/dummy"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Printf("Op Count = %d\n", calc.GetOpCount())
	color.Blue("Is 97 a prime number ? : %t\n", utils.IsPrime(97))
}


/* import (
	"fmt"
	. "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(Add(100, 200))
	fmt.Println(Subtract(100, 200))
	fmt.Printf("Op Count = %d\n", GetOpCount())
	fmt.Println(utils.IsPrime(97))
} */
