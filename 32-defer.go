package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApp()
}
