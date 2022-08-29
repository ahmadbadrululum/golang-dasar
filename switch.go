package main

import "fmt"

func main() {

	nama := "1"
	switch {
	case len(nama)>3:
		fmt.Println("panjang")
	case len(nama)<3:
		fmt.Println("pendek")
	default:
		fmt.Println("default")
	
		
	}


}