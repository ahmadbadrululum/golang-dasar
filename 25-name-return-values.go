package main

import "fmt"

func getFullName2() (firstName string, lastName string) {
	firstName = "Eko"
	lastName = "bad"
	return
}

func main() {

	a, b := getFullName2()

	fmt.Println(a)
	fmt.Println(b)
}
