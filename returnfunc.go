package main

import "fmt"

func namaKu(nama string) string {
	if nama == "" {
		return "halo nama saya kosong"
	} else {
		return "nama saya " + nama
	}
}

func main() {
	result := namaKu("bad")

	fmt.Println(result)

	fmt.Println(namaKu(""))

}
