package main

import "fmt"

func main() {

	// len, cap, append, make, copy

	var values = [...]string{"jan","feb","mar", "apr","mei"}


	var slice1 = values[2:5]
	fmt.Println(slice1);
	fmt.Println(cap(values));

	// append(values, "halo")
}