package main

import (
	"fmt"
	"math"
)

func main() {
	sayHello("Dimka", "Юзверь")
	// PetBattle(3, 2)

	printNumberInfo(-4)
	printNumberInfo(0)
	printNumberInfo(12)
	printNumberInfo(16)
}

func sayHello(username, role string) {
	fmt.Printf("Привет, %s. Вы - %q\n", username, role)
}

func PetBattle(cats, dogs int) {
	switch {
	case cats == dogs:
		fmt.Printf("Ничья! Все дружат!")
	case cats > dogs:
		fmt.Printf("Котики победили со счетом %d:%d!", cats, dogs)
	default:
		fmt.Printf("Собачки победили со счетом %d:%d!", dogs, cats)
	}
}

func printNumberInfo(num int) {
	if num < 0 {
		fmt.Printf("Число %d отрицательное.\n", num)
	} else if num > 0 {
		fmt.Printf("Число %d положительное.\n", num)
	} else {
		fmt.Printf("Число равно 0.\n")
	}

	if num%2 == 0 {
		fmt.Printf("Число %d четное.\n", num)
	} else {
		fmt.Printf("Число %d нечетное.\n", num)
	}

	if num > 0 {
		numSq := math.Sqrt(float64(num))
		if math.Trunc(numSq) != numSq {
			fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, numSq)
		} else {
			fmt.Printf("Квадратный корень числа %d является целым числом и равен %.0f.\n", num, numSq)
		}
	}
}
