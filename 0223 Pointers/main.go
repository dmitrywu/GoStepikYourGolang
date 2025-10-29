package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Value X:", x)
	fmt.Println("Pointer on X:", &x)

	p := &x
	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)

	*p = 100
	fmt.Println("Address:", *p)
	fmt.Println("Value X:", x)

}
