package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "                     Hello!         "
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str)) // убрать пробелы слева и справа

	str = "Hello from Golang!"
	upperStr := strings.ToUpper(str) // верхний регистр
	fmt.Println(upperStr)
	lowStr := strings.ToLower(str) // нижний регистр
	fmt.Println(lowStr)

	words := strings.Split(str, " ") // разбить строку по (в данном случае) пробелам в слайс
	fmt.Println(words)

	fmt.Println(strings.Contains(str, "fr")) // содержит ли строка символы ... (регистрозависимо)
	fmt.Println(strings.Contains(str, "FR"))

	fmt.Println(strings.HasPrefix(str, "Hel")) // начинается ли с ... (регистрозависимо)
	fmt.Println(strings.HasPrefix(str, "HEl"))

	fmt.Println(strings.HasSuffix(str, ".")) // заканчивается ли ... (регистрозависимо)
	fmt.Println(strings.HasSuffix(str, "!"))

	// сравнить строки без учёта регистра
	fmt.Println("Fold: ", strings.EqualFold(str, "Hello FROM GoLang!"))

}
