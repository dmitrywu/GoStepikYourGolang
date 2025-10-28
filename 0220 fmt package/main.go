package main

import "fmt"

func WeatherToday(city string, temp int, weather string) {
	fmt.Printf("В городе %s температура %d°C, %s.", city, temp, weather)
}

func main() {
	fmt.Print("Hello\n")
	fmt.Println("Hello,", "Go!")
	fmt.Printf("Hello, %s\n", "name")

	name := "Dmitry"
	age := 48
	fmt.Printf("My name is %s. I'm %d old.\n", name, age)
	out := fmt.Sprintf("My name is %s. I'm %d old.\n", name, age)
	fmt.Println(out)
}
