package main

import (
	"HeadFirst/ex.struct/example1/magazine"
	"fmt"
)

func main() {

	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}

	subscriber := magazine.Subscriber{Name: "Vladislav", Rate: 4.99, Active: true}
	subscriber.Address = address
	fmt.Println(subscriber.Name)
	fmt.Println(subscriber.Rate)
	fmt.Println(subscriber.Active)
	fmt.Println(subscriber.Address)

	employee := magazine.Employee{Name: "Joy Carr", Salary: 60000}
	employee.Street = "123 Oak St"
	employee.City = "Oklahoma"
	employee.State = "NE"
	employee.PostalCode = "68111"
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println(employee.Street)
	fmt.Println(employee.City)
	fmt.Println(employee.State)
	fmt.Println(employee.PostalCode)

}
