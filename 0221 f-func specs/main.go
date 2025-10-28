package main

import "fmt"

func Presenter(num int) {
	fmt.Printf("Запись числа %d в разных системах счисления:\n", num)
	fmt.Printf("Десятичная:%d\n", num)
	fmt.Printf("Двоичная:%b\n", num)
	fmt.Printf("Восьмеричная:%o\n", num)
	fmt.Printf("Шестнадцатеричная:%X\n", num)
}

func main() {
	fmt.Printf("This is %s\n", "string")
	fmt.Printf("This is int - %d\n", 5)
	fmt.Printf("This is float - %f\n", 51.1232)
	fmt.Printf("This is float - %.2f\n", 51.1232)
	fmt.Printf("This is float - %.0f\n", 51.1232)
	fmt.Printf("This is boolean - %t\\%t\n", true, false)
	fmt.Printf("Percent %d%%\n", 32)
	fmt.Printf("Binary: %b\n", 47)
	fmt.Printf("Binary: %#b\n", 47)
	fmt.Printf("Octa: %o\n", 47)
	fmt.Printf("Dec: %d\n", 47)
	fmt.Printf("Hex: %x\n", 47)
	fmt.Printf("Hex: %#x\n", 47)
	fmt.Printf("Hex: %#X\n", 47)
	fmt.Printf("Any value: %v", 0x1E)
}
