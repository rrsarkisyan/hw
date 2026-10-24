package main

import "fmt"

func main() {
	var n, current, min, max int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	if n < 1 {
		panic("N должно быть положительным")
	}
	fmt.Print("Введите число: ")
	fmt.Scan(&current)

	min, max = current, current

	for i := 0; i != n-1; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&current)
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
	}
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
}
