package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10

	fmt.Println("a + 3:", a+3)
	fmt.Println("a - 3:", a-3)
	fmt.Println("a * 3:", a*3)
	fmt.Println("a / 3:", a/3)
	fmt.Println("a % 3:", a%3) // остаток от деления

	b := 10.5
	fmt.Println(b - 3*math.Trunc(b/3)) // остаток от деления для float
	fmt.Println(math.Mod(b, 3))

	num := 5
	fmt.Println(num)
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)

	num += 3
	fmt.Println(num)
	num -= 3
	fmt.Println(num)
	num /= 3
	fmt.Println(num)
	num *= 3
	fmt.Println(num)

}
