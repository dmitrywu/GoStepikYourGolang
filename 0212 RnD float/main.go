package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	min := 20.0
	max := 50.0

	//var max, min float64 = 50, 20

	//random := rand.Float64()*(max-min) + min
	//var random float64

	random := rand.Float64()*(max-min+1) + min

	//random = rand.Float64()
	//random = random*(max-min+1) + min

	fmt.Println(random)
}
