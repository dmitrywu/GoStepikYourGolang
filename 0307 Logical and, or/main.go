package main

import "fmt"

func main() {
	a := 5
	b := 10

	result := a > 3 || (b > 10 && a < 3)
	// у && приоритет выше и оно выполнится первым
	// если логическое выражение длинное, имеет смысл расставлять скобки

	fmt.Println(result)

	age := 15
	role := "user"
	status := "active"
	fmt.Println((role == "admin" || role == "moderator") || (role == "user" && status == "active" && age >= 18))

}
