package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	_, lastName2 := getFullName()
	fmt.Println(lastName2)

}
