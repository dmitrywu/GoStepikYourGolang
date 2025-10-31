package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	isSecured := securePassword("Asd334vbsdf")
	fmt.Println(isSecured)

}

func securePassword(pass string) bool {
	if utf8.RuneCountInString(pass) >= 6 && pass != strings.ToLower(pass) && pass != strings.ToUpper(pass) {

	}

}
