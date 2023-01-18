package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "jateng", "indonesia"}
	address2 := &address1

	address2.Province = "Mana"

	*address2 = Address{"Surabaya", "jatim", "indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address4 := new(Address)
	address4.City = "jakarta"
	address4.Province = "jatim"
	address4.Country = "Indonesia"
	// address4.= "Indonesia"
	fmt.Println(address4)

}
