package main

import "fmt"

func main() {
	//Array
	fmt.Println("Arrays")
	//var nos [5]int
	//var nos [5]int = [5]int{4, 1, 2, 5, 3}
	nos := [5]int{4, 1, 2, 5, 3}
	fmt.Println(nos)

	//fmt.Println("Iterating an array (using index)")
	//for idx := 0; idx < len(nos); idx++ {
	//	fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	//}

	/*
		fmt.Println("Iterating an array (using range)")
		for idx, val := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, val)
		}
	*/

	fmt.Println("Iterating an array (using range)")
	for _, val := range nos {
		fmt.Printf("%d\n", val)
	}

	fmt.Println("Copying an array")
	newNos := nos
	nos[2] = 100
	fmt.Println(newNos, nos)

	fmt.Println()
	fmt.Println("Slice")
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	//var products []string
	//var products []string = make([]string, 3, 10)
	//fmt.Println(products, len(products), cap(products))
	//products[0] = "Pen"
	//products[1] = "Pencil"
	//products[2] = "Marker"
	//fmt.Println(products, len(products), cap(products))
	var products []string
	products = append(products, "Pencil", "Marker", "Duster")
	fmt.Println(products, len(products), cap(products))
	products = append(products, "Mouse")
	fmt.Println(products, len(products), cap(products))
	products = append(products, "Scribble-pad", "Stylus")
	fmt.Println(products, len(products), cap(products))
	products = append(products, "Printer")
	fmt.Println(products, len(products), cap(products))

	fmt.Println("Slicing")
	fmt.Println("Deleting an item from a slice")
	fmt.Println("Before deleting ", products)
	products = append(products[0:3], products[4:]...)
	fmt.Println("After deleting products[3]", products)
	/*
		[lo : hi] => from lo to hi-1
		[lo:] => from lo to len-1
		[:hi] => from 0 to hi-1
		[:] => copy of the slice
	*/
	fmt.Println("products[2:5] =>", products[2:5])
	fmt.Println("products[2:] =>", products[2:])
	fmt.Println("products[:5] =>", products[:5])
	//newProducts := products[:]
	//newProducts[0] = "Charger"
	newProducts := make([]string, len(products))
	copy(newProducts, products)
	newProducts = append(newProducts, "Charger")
	fmt.Println(newProducts, products)
	fmt.Println(len(newProducts), len(products))
	products = append(products, "Adapter")
	fmt.Println(newProducts, products)
	fmt.Println(len(newProducts), len(products))

	fmt.Println()

	fmt.Println("Map")
	//var productRanks map[string]int = make(map[string]int)
	/*
		productRanks := map[string]int{}
		productRanks["Pen"] = 4
	*/

	productRanks := map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 5,
	}
	productRanks["Printer"] = 3
	productRanks["Scribble-pad"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("Rank of %q is %d\n", key, val)
	}

	fmt.Println("Checking if a key exists")
	productToSearch := "Dummy"
	if val, exist := productRanks[productToSearch]; exist {
		fmt.Printf("Rank of %q is %d\n", productToSearch, val)
	} else {
		fmt.Printf("%q does not exist\n", productToSearch)
	}

	fmt.Println("Deleting from a map")
	delete(productRanks, "Pen")
	fmt.Println("productRanks after deleting Pen", productRanks)
}
