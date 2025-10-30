package main

import "fmt"

func main() {
	a := 5
	if a > 5 {
		fmt.Println("a больше 5")
	} else if a == 5 {
		fmt.Println("a = 5")
	} else {
		fmt.Println("a не больше 5")
	}

	b, c := 7, 1
	if b > 5 && c < 3 {
		fmt.Println("б больше 5 и ц меньше 3")
	}

	temp := -5
	if temp < 0 {
		fmt.Println("Город замерзает! Верните лето.")
	} else if temp > 35 {
		fmt.Println("Город в огне! Яичницу можно жарить на асфальте.")
	} else {
		fmt.Println("Температура в норме. Продолжаем писать код.")
	}
}
