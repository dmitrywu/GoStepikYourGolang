package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "12345"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Some error detected!", err)
	} else {
		fmt.Println("Converted number: ", num)
	}

}
