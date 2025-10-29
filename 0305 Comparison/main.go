package main

import "fmt"

func main() {
	a, b, c := 5, 5, 10
	fmt.Println(a == b)
	fmt.Println(b == c)

	fmt.Println(5 != 10)

	fmt.Println(10 > 5)
	fmt.Println(10 < 5)

	fmt.Println(10 <= 5)
	fmt.Println(10 >= 5)
}
