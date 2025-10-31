package main

import (
	"fmt"
	"math/rand/v2"
	"unicode/utf8"
)

func main() {
	result := Sum(2, 5)
	fmt.Println(result)

	helloStr := "Привет, друг!"
	bytesLength, runesLength := getFullLength(helloStr)
	fmt.Printf("Строка %q занимает %d байт и имеет %d символов.\n", helloStr, bytesLength, runesLength)
	fmt.Println(bytesLength, runesLength)
}

func Sum(n1, n2 int) int {
	return n1 + n2
}

func getFullLength(str string) (int, int) {
	bytes := len(str)
	runes := utf8.RuneCountInString(str)
	return bytes, runes
}

func generateCompliment(name string) string {
	// "Ты великолепен, !"
	// "У тебя потрясающая улыбка, !"
	// "Ты вдохновляешь, !"
	num := rand.UintN(3)
	switch num {
	case 0:
		return fmt.Sprintf("Ты великолепен, %s!", name)
	case 1:
		return fmt.Sprintf("У тебя потрясающая улыбка, %s!", name)
	default:
		return fmt.Sprintf("Ты вдохновляешь, %s!", name)
	}
}
