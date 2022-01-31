package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (product *Product) ToString() string {
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

/*
func (products *Products) FilterCostlyProdoucts() Products {
	filteredProducts := Products{}
	for _, p := range *products {
		if p.Cost > 1000 {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts
}

func (products *Products) FilterStationaryProdoucts() Products {
	filteredProducts := Products{}
	for _, p := range *products {
		if p.Category == "Stationary" {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts
}

*/

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

func (products *Products) ToString() string {
	result := ""
	for _, p := range *products {
		result += p.ToString() + "\n"
	}
	return result
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
	fmt.Print(costlyProducts.ToString())

	fmt.Println("Filter Stationary products")
	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductCriteria)
	fmt.Print(stationaryProducts.ToString())

	fmt.Println("All")
	fmt.Println("Are all the products costly products ? :", products.All(costlyProductCriteria))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ? :", products.Any(costlyProductCriteria))
}

/*
Write the following function
IndexOf => return the index of the given product in the collection
Includes => return true if the given product is present in the collection else return false
Filter => return products that matches the given criteria
	Use cases:
		1. filter all costly products (cost > 1000)
		2. filter all stationary products (category == stationary)
		3. filter all costly stationary products
All => return true if ALL the products matches the given criteria else return false
	Use cases:
		1. Are all the products costly products (cost > 1000) ?
		2. Are all the products stationary products? (category == "Stationary")

Any => return true if ANY of the products matches the given criteria else returns false
	Use cases:
		1. Are there ANY costly product (cost > 1000) ?
		2. Are there ANY stationary products? (category == "Stationary")

*/
