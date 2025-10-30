package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var weight float64
	var height float64
	var result string

	fmt.Print("Введите ваш вес (кг): ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Люди столько не весят")
		os.Exit(1)
	}

	fmt.Print("Введите ваш рост (см): ")
	_, err = fmt.Scan(&height)
	if err != nil {
		fmt.Println("Люди столько не растут")
		os.Exit(1)
	}

	imt := weight / math.Pow(height/100, 2)

	switch {
	case imt < 18.5:
		result = "Недостаточный вес"
	case imt >= 18.5 && imt < 25:
		result = "Нормальный вес"
	case imt >= 25 && imt < 30:
		result = "Избыточный вес"
	default:
		result = "Ожирение "
	}

	// Блин, забыл, что можно и так
	// switch {
	// case bmi < 18.5:
	// 	category = "Недостаточная масса тела"
	// case bmi < 25:
	// 	category = "Норма"
	// case bmi < 30:
	// 	category = "Избыточный вес"
	// default:
	// 	category = "Ожирение"
	// }

	fmt.Printf("Ваш ИМТ: %.2f\n", imt)
	fmt.Printf("Категория: %s\n", result)

}
