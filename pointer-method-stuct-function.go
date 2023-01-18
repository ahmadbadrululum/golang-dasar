package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Halcuk(belakang string) {
	man.Name = "Mr. " + man.Name + " nama belakang " + belakang
}

func main() {
	bad := Man{"bads"}
	bad.Halcuk("ulum")

	fmt.Println(bad)
}
