package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type person struct {
	name string
	age  int
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// struct with receiver
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {

	// Hello world
	fmt.Println("Hello world")

	fmt.Println("go" + "lang")
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	var v = 1.0
	e := 5
	fmt.Println(v, e, "hi")

	const PI = 3.14
	fmt.Println(v + PI)

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// init array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// make slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// append slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	m := make(map[string]int)
	m["sam"] = 100
	x := m["sam"]
	fmt.Println(m, x)

	fmt.Println(plusPlus(2, 3, 4))

	arr := []int{3, 4, 5}
	fmt.Println(sum(arr...))

	// closure
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	f := fact(7)
	fmt.Println("Factorial: ", f)

	// struct
	p := person{"Bob", 300}
	fmt.Println(p)

	r := rect{10, 20}
	fmt.Println("Rect area: ", r.area())

	// interfaces
	cir := circle{5}
	measure(cir)

	// errors
	e12, err := f1(12)
	fmt.Println(e12, err)

	// goroutines
	fLoop("direct")
	go fLoop("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("hellooooooooo")
	fmt.Scanln()

	// channel
	number := 589
	sqrch := make(chan int)
	cubch := make(chan int)
	go square(number, sqrch)
	go cube(number, cubch)
	sumFromSquare, sumFromCube := <-sqrch, <-cubch
	fmt.Println("Square + Cube: ", sumFromCube+sumFromSquare)
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// return a function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

// error
func f1(arg int) (int, error) {
	if arg == 12 {
		return -1, errors.New("Can't work with 12")
	}
	return arg + 3, nil
}

// goroutine
func fLoop(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func square(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func cube(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
