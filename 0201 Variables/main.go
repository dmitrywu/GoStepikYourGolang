package main

import "fmt"

var (
	a = 5
	b = 2
	c int
)

const (
	j = 1
	k = 2.4
	l // В константах, если не указано значение, берётся из предыдущего
)

func main() {
	// var age int = 30
	// var age = 30
	age := 30
	age -= 5

	// var x, y int = 1, 2

	x, y := 1, 2

	const MY_VAR = 3.14

	fmt.Println(age, x, y)
	fmt.Println(a, b, c)
	fmt.Println(MY_VAR)
	fmt.Println(j, k, l)

}
