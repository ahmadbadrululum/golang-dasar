package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30)
	fmt.Println(total)

	// slice := []int{20,30,50}
	// total2 := sumAll(slice ...)
	// fmt.Println(total2+total)
}
