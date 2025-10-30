package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println("Start: x = ", x)

	if x > 5 {
		y := 20
		fmt.Printf("Scope test: x = %d, y = %d\n", x, y)
	}

	{
		y := 20
		x := 20
		fmt.Printf("Scope test: x = %d, y = %d\n", x, y)
	}

	fmt.Println("X at end of scope: ", x)

}
