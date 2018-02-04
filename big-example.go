package main

import (
	"fmt"
	"math"
	"time"
)

func bigPrint(value string) {
	fmt.Printf("\n\n\n%s\n\n", value)
}

func xRange() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func plusMinus(a, b int) (d int, c int) {
	d = a + b
	c = a - b
	return
}

func sum(nums ...int) (total int) {
	for _, v := range nums {
		total += v
	}
	return
}

func addOne(valueptr *int) {
	*valueptr++
}

type person struct {
	name string
	age  int
}

func (pers *person) firstLetter() string {
	return string(pers.name[0])
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func process(ch chan string) {
	time.Sleep(5000 * time.Millisecond)
	ch <- "process successful"
}

// MyConst Comment to document that const
const MyConst string = "a constant"

func main() {
	bigPrint("VARIABLE DECLERATION ####################################################")
	inferredType := "infer the type"
	fmt.Println(MyConst, inferredType)

	var variable1, variable2 int = 1, 2
	fmt.Println(variable1, variable2)

	bigInt := 3e11
	fmt.Println(bigInt, int64(bigInt))

	bigPrint("FOR LOOP ####################################################")
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\n", i)
		}
	}

	bigPrint("ARRAY ####################################################")
	var myArray = [5]float64{1, 2, 3, math.Floor(math.Exp(4))}
	fmt.Println(myArray)

	bigPrint("SLICE ####################################################")
	mySlice := make([]string, 5)
	mySlice[0] = "a"
	mySlice[1] = "b"
	mySlice[2] = "c"
	mySlice[3] = "d"
	mySlice[4] = "e"
	mySlice = append(mySlice, "f")
	fmt.Println(mySlice)
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[2:5])

	bigPrint("MAP ####################################################")
	myMap := map[string]int{"foo": 3, "bar": 4}
	myMap["key1"] = 1
	myMap["key2"] = 2
	delete(myMap, "key1")
	fmt.Println(myMap)

	bigPrint("RANGE ####################################################")
	for i, v := range myArray {
		fmt.Printf("myArray[%d] = %.2f\n", i, v)
	}

	for k, v := range myMap {
		fmt.Printf("%s -> %d\n", k, v)
	}

	bigPrint("FUNCTIONS ####################################################")
	plus, minus := plusMinus(4, 3)
	fmt.Println(plus, minus)

	bigPrint("VARIADIC FUNCTIONS ####################################################")
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...))

	bigPrint("CLOSURES ####################################################")
	nextInt := xRange()
	fmt.Println(nextInt(), nextInt(), nextInt())

	bigPrint("POINTER ####################################################")
	myPointedVar := 1
	addOne(&myPointedVar)
	addOne(&myPointedVar)
	fmt.Println(myPointedVar)

	bigPrint("STRUCT ####################################################")

	p := person{"Bob", 25}
	fmt.Println(p, p.name, p.age)

	bigPrint("METHODS ####################################################")
	fmt.Println(p.firstLetter())
	pPointer := &p
	fmt.Println(pPointer.firstLetter())

	bigPrint("INTERFACE ####################################################")
	rect := rect{3, 4}
	measure(rect)

	bigPrint("GOROUTINES ####################################################")
	go say("world")
	say("hello")

	bigPrint("CHANNELS ####################################################")
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	bigPrint("SELECT ####################################################")
	start := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received after 1 second: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received after 2 seconds: ", msg2)
		}
	}

	fmt.Printf("but took %s\n", time.Since(start))

	bigPrint("DEFAULT CHANNEL ####################################################")
	ch := make(chan string)
	go process(ch)
loop:
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			break loop
		default:
			fmt.Println("no value received")
		}
	}

	bigPrint("RANGE AND CLOSE CHANNEL ####################################################")
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
