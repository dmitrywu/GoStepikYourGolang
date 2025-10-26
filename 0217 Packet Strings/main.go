package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "                     Hello!         "
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))

	str = "Hello from Golang!"
	upperStr := strings.ToUpper(str)
	fmt.Println(upperStr)
	lowStr := strings.ToLower(str)
	fmt.Println(lowStr)

	words := strings.Split(str, " ")
	fmt.Println(words)

	fmt.Println(strings.Contains(str, "fr"))
	fmt.Println(strings.Contains(str, "FR"))

	fmt.Println(strings.HasPrefix(str, "Hel"))
	fmt.Println(strings.HasPrefix(str, "HEl"))

	fmt.Println(strings.HasSuffix(str, "."))
	fmt.Println(strings.HasSuffix(str, "!"))

	fmt.Println("Fold: ", strings.EqualFold(str, "Hello FROM GoLang!"))

}
