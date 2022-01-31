package dummy

import "fmt"
import "modularity-demo/calculator/utils"

func init() {
	fmt.Println("dummy.init triggered")
	fmt.Printf("Is 94 a prime number ? : %t\n", utils.IsPrime(94))
}
