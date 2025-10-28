package main

import (
	"fmt"
	"strconv"
)

func GoodsSum(priceStr string, quantityStr string) {
	n1, _ := strconv.ParseFloat(priceStr, 64)
	n2, _ := strconv.Atoi(quantityStr)
	sum := n1 * float64(n2)
	fmt.Printf("%.2f", sum)
}

func main() {

	GoodsSum("100", "5")

	str3 := "3.14"

	num3, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("Ошибка преобразования: ", err)
	} else {
		fmt.Println("Преобразованное число (Float64): ", num3)
	}

	str := "112312341"

	num2, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Ошибка преобразования: ", err)
	} else {
		fmt.Println("Преобразованное число (int64): ", num2)
	}

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Ошибка преобразования: ", err)
	} else {
		fmt.Println("Преобразованное число: ", num)
	}

}
