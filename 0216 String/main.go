package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello\n\tmy\n\tfriend!\\ & \"aliens\" \uFAFA"
	str2 := `
	This
	is 
	test
	backtick
	`
	// Экранирование
	// \n (\r - for windows)
	// \t
	// \"
	// \u (unicode)

	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(utf8.RuneCountInString(str))

	// len - для строк по умолчанию выводит количество байт, а не символов

	// " " - для строк
	// ' ' - для рун
	// ` ` - "бэктики" - всё, что напишешь, будет выведено именно так, как написано

}
