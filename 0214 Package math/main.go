package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Pow(2, 3))  // возведение в степень: 2 в степени 3
	fmt.Println(math.Sqrt(16))   // Квадратный корень
	fmt.Println(math.Abs(-10.5)) // абсолютное значение (без знака -)
	fmt.Println(math.Acos(1))
	fmt.Println(math.Log(5))
	fmt.Println(math.Log10(5))
	fmt.Println(math.Round(3.5)) // округление по правилам математики
	fmt.Println(math.Floor(3.7)) // округление вниз
	fmt.Println(math.Ceil(3.2))  // округление вверх
	fmt.Println(math.Trunc(3.2)) // Отбрасывание дробной части

	fmt.Println(math.Max(3.0, 4.0))

	num := 3.14159265
	fmt.Println(num)
	rounded := math.Round(num*100) / 100
	fmt.Println(rounded)

}
