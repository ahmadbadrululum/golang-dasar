package main

import "fmt"

func main() {

	person := map[string]string{
		"name":"bad",
		"add" : "demak",
	}

	person["title"]="programmer"

	fmt.Println(person["title"])

	delete(person,"title")
	
	fmt.Println(person["title"])



}