package main

import "fmt"

func halo(name string) string{
	return "halo cuk "+ name
}


func main() {
	say := halo
	result := say("jancuk")
	fmt.Println(result)
}
