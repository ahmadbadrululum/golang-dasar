package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var bad Customer
	// bad.Name = "bad"
	// bad.Address = "Indonesia"
	// bad.Age = 30

	bad.sayHi("Joko")
	// bad.sayHuuu()

	//fmt.Println(bad.Name)
	//fmt.Println(bad.Address)
	//fmt.Println(bad.Age)
	//
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}
