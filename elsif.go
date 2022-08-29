package main

import "fmt"

func main() {

	nama := "bad123123"

	if nama=="bad" {
		fmt.Println("oke")
	}else if nama=="siapa"{
		fmt.Println("siapa")
	}else{
		fmt.Println("lanjut")
	}


	if len(nama)>5{
		fmt.Println("panjang cuk",len(nama))
	}else{
		fmt.Println("else",len(nama))
	}

}