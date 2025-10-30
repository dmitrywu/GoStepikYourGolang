package main

import (
	"fmt"
	"os"
)

func main() {
	// Никогда не начинайте сразу резко врываться в написание кода
	// Сначала очень важно точно понять, что нужно сделать
	// До конца понять, со всеми нюансами

	var input int
	var timeOfDay string

	fmt.Println(" Введите время в часах (0-23):")
	_, err := fmt.Scan(&input)

	if err != nil || input < 0 || input > 23 {
		fmt.Println("Неверно задано время.")
		os.Exit(1)
	}

	switch {
	case input >= 6 && input < 12:
		timeOfDay = "утро"
	case input >= 12 && input < 18:
		timeOfDay = "день"
	case input >= 18 && input < 23:
		timeOfDay = "вечер"
	default:
		timeOfDay = "ночь"
	}

	fmt.Printf("Сейчас %02dч. - %s\n", input, timeOfDay)
}
