package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var group = []Person{
		Person{"Deric", 21},
		Person{"Bruno", 24},
		Person{"Bob", 36}}

	fmt.Println(group[0:2])
}
