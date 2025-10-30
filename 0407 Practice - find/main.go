package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var inputGoods string
	var itemName1, itemName2, itemName3 string
	var price int

	fmt.Print("Input string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка ввода: %s", err)
	}

	inputGoods = strings.ToLower(scanner.Text())

	itemName1 = "Клавиатура JZ9"
	itemName2 = "Наушники N45"
	itemName3 = "Смартфон S10"

	switch {
	case strings.Contains(strings.ToLower(itemName1), inputGoods):
		price = 19200
		fmt.Printf("Клава %d", price)
	case strings.Contains(strings.ToLower(itemName2), inputGoods):
		price = 9600
		fmt.Printf("Уши %d", price)
	case strings.Contains(strings.ToLower(itemName3), inputGoods):
		price = 55000
		fmt.Printf("Смарт %d", price)
	default:
		fmt.Printf("Товар %q не найден.", scanner.Text())
	}

}
