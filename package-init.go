package main

import (
	"fmt"
	"go-dasar2/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
