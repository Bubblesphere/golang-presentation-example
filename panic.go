package main

import "fmt"

func f1(arg int) int {
	if arg == 42 {
		panic("bad arg")
	}
	return arg + 3
}

func main() {
	for _, i := range []int{7, 42} {
		fmt.Println("f1 worked:", f1(i))
	}

}
