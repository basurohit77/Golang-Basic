package main

import (
	"fmt"
	"interfaces-demo/models"
)

func main() {
	var stove = models.Product{102, "Stove", 5000, 5, "Utencil"}
	var nonExistentProduct = models.Product{1000, "Dummy", 1000, 1000, "Dummy"}
	products := models.Products{
		models.Product{105, "Pen", 5, 50, "Stationary"},
		models.Product{107, "Pencil", 2, 100, "Stationary"},
		models.Product{103, "Marker", 50, 20, "Utencil"},
		models.Product{102, "Stove", 5000, 5, "Utencil"},
		models.Product{101, "Kettle", 2500, 10, "Utencil"},
		models.Product{104, "Scribble Pad", 20, 20, "Stationary"},
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
	costlyProductCriteria := func(p models.Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductCriteria)
	fmt.Print(costlyProducts)

	fmt.Println("Filter Stationary products")
	stationaryProductCriteria := func(p models.Product) bool {
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
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Name)")
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Cost)")
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Units)")
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default Sort (by Category)")
	products.Sort("Category")
	fmt.Println(products)

}
