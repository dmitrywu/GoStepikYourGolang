package main

import "fmt"

func main() {
	var r1 rune = 'A'
	fmt.Printf("Idx: %d, symbol: %c\n", r1, r1)

	r2 := '\u1234'
	fmt.Printf("Idx: %d, symbol: %c\n", r2, r2)

	r3 := rune(128512)
	fmt.Printf("Idx: %d, symbol: %c\n", r3, r3)

	str := "Hello"
	fmt.Printf("Number: %d, symbol: %c\n", str[0], str[0])

	strRu := "Привет"
	fmt.Printf("Number: %d, symbol: %c\n", strRu[0], strRu[0])
	fmt.Printf("Number: %d, symbol: %c\n", strRu[1], strRu[1])

	runes := []rune(strRu)
	fmt.Println(runes)
	fmt.Printf("Number: %d, symbol: %c\n", runes[0], runes[0])

}
