package main


import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		panic("bad arg")
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		result, err := f1(i)
		if err != nil {
			fmt.Println("f1 failed:", err )
		} else {
			fmt.Println("f1 worked:", result)
		}
	}

}
