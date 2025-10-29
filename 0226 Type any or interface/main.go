package main

import "fmt"

func main() {
	// Когда нужно положить в переменную что-то... но мы пока не знаем, что.
	// или в одну переменную разные типы данных
	// var val interface{} = "Hello, my friend!"
	// any = interface {}
	var val any = "Hello, my friend!"
	fmt.Println(val)
	val = 50
	fmt.Println(val)
}
