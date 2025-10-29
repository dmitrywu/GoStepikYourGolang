package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	random := math.Floor((rand.Float64()*100)*10) / 10
	random = 0.1
	fmt.Printf("Исходное число: %.1f\n", random)
	fmt.Printf("%f\n", math.Mod(random, 2))
	fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random*1.1)
	// fmt.Printf("Исходное число является четным: %t\n", int(math.Trunc(random))%2 == 0)
	//if int(math.Mod(random, 2)) != 0
	//fmt.Printf("%f\n", math.Trunc(random))
	//fmt.Printf("%d\n", int(random))
	//fmt.Printf("Исходное число является четным: %t\n", int(math.Mod(random, 2)) == 0 && float64(int(random))-random == 0)
	//fmt.Printf("Исходное число является четным: %t\n", int(random)&1 == 0 && int(random) != 0)
	fmt.Printf("Исходное число является четным: %t\n", int(random)&1 == 0 && float64(int(random))-random == 0)
	//fmt.Printf("Исходное число является четным: %t\n", math.Mod(random, 1) == 0 && math.Mod(random, 2) == 0)
	//fmt.Printf("Исходное число является четным: %t\n", math.Mod(random, 2) == 0)
	fmt.Printf("Предпоследняя цифра целой части исходного числа: %d\n", int((math.Trunc(random)/100)*10))
}
