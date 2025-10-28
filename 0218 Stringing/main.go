package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "12345"
	i := 555
	s1 := strconv.Itoa(i)
	fmt.Println(s1)
	fmt.Println(strconv.FormatInt(int64(i), 2))
	fmt.Println(strconv.FormatInt(int64(i), 16))

	number := 10.5254
	fmt.Println(strconv.FormatFloat(number, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(number, 'e', 2, 64))
	fmt.Println(strconv.FormatFloat(number, 'x', 0, 64))

	str2 := fmt.Sprintf("Привет, вот число: %v", number)
	fmt.Println(str2)

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Some error detected!", err)
	} else {
		fmt.Println("Converted number: ", num)
	}

}
