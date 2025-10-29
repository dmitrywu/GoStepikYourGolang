package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input string: ")
	scanner.Scan()
	fmt.Println("Your input: ", scanner.Text())
	// var firstName, lastName string
	// fmt.Print("Name?: ")
	// fmt.Scan(&firstName) // здесь именно указатель
	// fmt.Print("Family?: ")
	// fmt.Scan(&lastName)
	// fmt.Printf("Your name is: %s, family is: %s\n", firstName, lastName)
	// fmt.Print("Enter name and family (space - separator): ")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Printf("Your name is: %s, family is: %s\n", firstName, lastName)

}
