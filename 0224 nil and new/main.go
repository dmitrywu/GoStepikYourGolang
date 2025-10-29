package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)
	// разыменование - *p - dereference

	x := new(int)
	// new инициализирует нулевое значение (в зависимости от типа данных)
	// и возвращает указатель
	fmt.Println(x)
	fmt.Println(*x)
	*x = 50
	fmt.Println(*x)
}
