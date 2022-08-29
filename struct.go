


// struct contoh prototype data 


// function yg memanggil dirinya sendiri
package main

import "fmt"


type Costumer struct {
	Name, Address string
	Age int	
}

func main() {

	bad := Costumer{
		Name : "bad",
		Address : "denmark",
		Age : 20,
	}

	fmt.Println(bad)
}