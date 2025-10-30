package main

import "fmt"

const (
	execute = 0b001
	write   = 0b010
	read    = 0b100
)

// const (
// 	execute = 1 << iota // 1 0b001
// 	write               // 2 0b010
// 	read                // 4 0b100
// )

func main() {
	executeRead := execute | read
	fmt.Printf("%b\n", executeRead)

	writeRead := write | read
	fmt.Printf("%b\n", writeRead)

	permissions := 5
	fmt.Printf("Access level: %d = %b\n", permissions, permissions)

	fmt.Printf("Execute: %t\n", permissions&execute == 1)
	fmt.Printf("Read: %t\n", permissions&read == 4)
	fmt.Printf("Write: %t\n", permissions&write == 2)

	fmt.Printf("Execute: %t\n", permissions&execute > 0)
	fmt.Printf("Read: %t\n", permissions&read > 0)
	fmt.Printf("Write: %t\n", permissions&write > 0)

	num := 5
	fmt.Println(num << 2)
	fmt.Println(num >> 1)
	fmt.Println(num & 3)
	fmt.Println(num | 2)
	fmt.Println(num ^ 2)
	fmt.Println(^num)

}
