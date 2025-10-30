package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("a>3 && b>3", a > 3 && b > 3)
	fmt.Println("a>6 || b>3", a > 6 || b > 3)
}
