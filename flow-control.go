package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// for ----------------------------------------------------
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")

	// while ----------------------------------------------------
	a := 1024
	for a > 1 {
		a /= 2
	}
	fmt.Println(a)

	// while true ----------------------------------------------------
	for {
		break
	}

	// if ----------------------------------------------------
	x := 10
	if x > 2 {
		fmt.Println("Greater than 2")
	}

	if y := 2; x%y == 0 {
		fmt.Println("Even")
	}

	// switch ----------------------------------------------------
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer ----------------------------------------------------
	fmt.Println("counting")

	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}

}
