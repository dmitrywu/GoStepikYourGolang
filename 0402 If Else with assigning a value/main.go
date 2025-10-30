package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	if num := rand.IntN(100); num > 50 { // при такой записи num доступен только в блоке if
		fmt.Printf("Number: %d\n", num)
	} else {
		fmt.Printf("Number %d too small...\n", num)
	}
}
