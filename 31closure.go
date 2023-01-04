package main

import "fmt"

func main() {

	cuk := "halos"
	// function biasa
	sayHello := func(name string) {
		fmt.Println("halo", name)
		cuk := "halo cuk"
		fmt.Println(cuk)

	}

	sayHello("eko")

	// function dengan return
	sayGoodBye := func(name string) string {
		return "good bye " + name
	}

	fmt.Println(cuk)

	result
}
