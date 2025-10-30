package main

import "fmt"

func main() {

	var val any = true

	switch val.(type) {
	case int:
		fmt.Println("В переменной val находится тип int")
	case string:
		fmt.Println("В переменной val находится тип string")
	case float64:
		fmt.Println("В переменной val находится тип float64")
	case bool:
		fmt.Println("В переменной val находится тип bool")
	default:
		fmt.Println("В переменной val находится неизвестный тип данных.")
	}

	var x any = 0.14

	switch v := x.(type) {
	case int:
		fmt.Println("Int", v)
	case string:
		fmt.Println("String", v)
	case float64:
		fmt.Println("Float", v)
	default:
		fmt.Println("x3", v)
	}

	day := 5

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("ft:Work!")
		fallthrough
	case 6, 7:
		fmt.Println("ft:Relax...")
	default:
		fmt.Println("ft:Week not correct. Please transfer to Earth.")
	}

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Work!")
	case 6, 7:
		fmt.Println("Relax...")
	default:
		fmt.Println("Week not correct. Please transfer to Earth.")
	}

	switch {
	case day < 1:
		fmt.Println("Incorrect day")
	case day >= 1 && day <= 5:
		fmt.Println("Work!")
	default:
		fmt.Println("Relax")
	}

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Тоже хороший, но несуществующий день недели")
	}

}
