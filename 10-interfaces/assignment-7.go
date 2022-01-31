/* Assignment:7 implement Sorting */
/*
Implement "sorting" functionality for the Products
	The user should be able to sort the products by any attribute
	Important: use the "sort.Sort()" function
*/

package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type ProductPredicate func(product Product) bool

type Products []Product

func (products *Products) IndexOf(product Product) int {
	for idx, p := range *products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) >= 0
}

func (products *Products) Filter(criteria ProductPredicate) Products {
	filteredProducts := Products{}
	for _, p := range *products {
		if criteria(p) {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts
}

func (products *Products) All(criteria ProductPredicate) bool {
	for _, p := range *products {
		if !criteria(p) {
			return false
		}
	}
	return true
}

func (products *Products) Any(criteria ProductPredicate) bool {
	for _, p := range *products {
		if criteria(p) {
			return true
		}
	}
	return false
}

func (products Products) String() string {
	result := ""
	for _, p := range products {
		//result += fmt.Sprint(p) + "\n"
		result += fmt.Sprintln(p)
	}
	return result
}

/* "sort.interface" interface implementation */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

//Sorting by Name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//Sorting by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//Sorting by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//Sorting by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func main() {
	var stove = Product{102, "Stove", 5000, 5, "Utencil"}
	var nonExistentProduct = Product{1000, "Dummy", 1000, 1000, "Dummy"}
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Index of stove = ", products.IndexOf(stove))
	fmt.Println("Index of nonExistentProduct = ", products.IndexOf(nonExistentProduct))

	fmt.Println("stove exists ? = ", products.Includes(stove))
	fmt.Println("nonExistentProduct exists ? = ", products.Includes(nonExistentProduct))

	fmt.Println("Filter")
	/*
		fmt.Println("Filter costly products")
		costlyProducts := products.FilterCostlyProdoucts()
		fmt.Print(costlyProducts.ToString())

		fmt.Println("Filter Stationary products")
		stationaryProducts := products.FilterStationaryProdoucts()
		fmt.Print(stationaryProducts.ToString())
	*/

	fmt.Println("Filter costly products")
	costlyProductCriteria := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductCriteria)
	fmt.Print(costlyProducts)

	fmt.Println("Filter Stationary products")
	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductCriteria)
	fmt.Print(stationaryProducts)

	fmt.Println("All")
	fmt.Println("Are all the products costly products ? :", products.All(costlyProductCriteria))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ? :", products.Any(costlyProductCriteria))

	fmt.Println()
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Id)")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Name)")
	sort.Sort(ByName{products})
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Cost)")
	sort.Sort(ByCost{products})
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Units)")
	sort.Sort(ByUnits{products})
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Category)")
	sort.Sort(ByCategory{products})
	fmt.Println(products)

}
