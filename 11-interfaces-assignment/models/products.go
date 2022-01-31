package models

import (
	"fmt"
	"sort"
)

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
		result += fmt.Sprint(p) + "\n"
	}
	return result
}

/* "sort.interface" interface implementation */

/* func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

//Sorting by Name
type byName struct {
	Products
}

func (byName byName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//Sorting by Cost
type byCost struct {
	Products
}

func (byCost byCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//Sorting by Units
type byUnits struct {
	Products
}

func (byUnits byUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//Sorting by Category
type byCategory struct {
	Products
}

func (byCategory byCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(byName{products})
	case "Units":
		sort.Sort(byUnits{products})
	case "Cost":
		sort.Sort(byCost{products})
	case "Category":
		sort.Sort(byCategory{products})
	}
}
*/
func (products Products) Sort(attrName string) {
	var IdComparer = func(i, j int) bool {
		return products[i].Id < products[j].Id
	}

	var NameComparer = func(i, j int) bool {
		return products[i].Name < products[j].Name
	}

	var CostComparer = func(i, j int) bool {
		return products[i].Cost < products[j].Cost
	}

	var UnitsComparer = func(i, j int) bool {
		return products[i].Units < products[j].Units
	}

	var CategoryComparer = func(i, j int) bool {
		return products[i].Category < products[j].Category
	}

	switch attrName {
	case "Id":
		sort.Slice(products, IdComparer)
	case "Name":
		sort.Slice(products, NameComparer)
	case "Units":
		sort.Slice(products, UnitsComparer)
	case "Cost":
		sort.Slice(products, CostComparer)
	case "Category":
		sort.Slice(products, CategoryComparer)
	}

}
