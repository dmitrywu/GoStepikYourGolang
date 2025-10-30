package main

import (
	"fmt"
	"math/bits"
)

func main() {
	a := 0b101
	b := 0b011

	fmt.Printf("a = %b, b = %b\n", a, b)
	fmt.Printf("a & b = %d = %b\n", a&b, a&b)
	fmt.Printf("a | b = %d = %b\n", a|b, a|b)
	fmt.Printf("a ^ b = %d = %b\n", a^b, a^b)
	fmt.Printf("a &^ b = %d = %b\n", a&^b, a&^b)
	fmt.Printf("^a: = %d = %b\n", ^a, ^a)
	fmt.Printf("a << 1 = %d = %b\n", a<<1, a<<1)
	fmt.Printf("a >> 1 = %d = %b\n", a>>1, a>>1)
	bits.Len(uint(a))
}
