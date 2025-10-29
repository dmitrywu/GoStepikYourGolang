package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello, "
	str2 := "world!"
	sumStr := str1 + str2
	fmt.Println(sumStr)

	resStr := fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(resStr)

	resultJoin := strings.Join([]string{str1, str2}, "")
	fmt.Println(resultJoin)

	var buffer strings.Builder

	buffer.WriteString(str1)
	buffer.WriteString(str2)
	buffer.WriteString("!!!!!!!!!!!!")
	fmt.Println(buffer.String())
}
