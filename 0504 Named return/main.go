package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	helloStr := "Привет, друг!"
	bytesLength, runesLength := getFullLength(helloStr)
	fmt.Printf("Строка %q занимает %d байт и имеет %d символов.\n", helloStr, bytesLength, runesLength)
	fmt.Println(bytesLength, runesLength)
}

func getFullLength(str string) (bytes int, runes int) {
	bytes = len(str)
	runes = utf8.RuneCountInString(str)
	// Если используешь именованные возвращаемые значения, можно использовать return без параметров
	return
}
