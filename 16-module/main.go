package main

import (
	"fmt"
	"testmodel/organization/employee"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}
//type Organization struct {
//	Name string
//	City string
//}
//type Employee struct {
//	Id   int
//	Name string
//	Org  Organization
//}

func main() {
	//var product Product
	//var product Product = Product{101, "Pen", 10, 100, "Stationary"}
	/*
		var product Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	//var product = Product{Id: 100, Name: "Pen", Cost: 10}
	product := Product{Id: 100, Name: "Pen", Cost: 10}
	/*
		var p2 = Product{Id: 100, Name: "Pen", Cost: 10}
		fmt.Println(product == p2)
		fmt.Println(ToString(p2))
	*/

	/*
		var productPtr *Product = &product
		//fmt.Println((*productPtr).Id)
		fmt.Println(productPtr.Id)
	*/

	fmt.Println(ToString(&product))

	//emp := &Employee{
	//	Id:   5001,
	//	Name: "Rohit",
	//	Org: Organization{
	//		Name: "IBM",
	//		City: "Hyderabad",
	//	},
	//}
	//emp := Employee{}

		//emp := new(Employee) //returns a pointer to the Employee object
		//emp := &Employee{}   //returns a pointer to the Employee object
		//fmt.Printf("%#v\n", *emp)
		//fmt.Println(emp.Id)

	emp1 := employee.NewEmployee(5001, "Rohit", "IBM", "Bengaluru")
	fmt.Printf("%#v\n", emp1)

	//Write the "ApplyDiscount" method to apply the given discount to the given product
	ApplyDiscount(&product, 10)
	fmt.Println(ToString(&product))
}

func ToString(product *Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, discount float64) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}