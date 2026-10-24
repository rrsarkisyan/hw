package main

import "fmt"

func main() {
	var n, current, pos, neg, zero int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)
	if n < 0 {
		panic("N не должно быть отрицательным")
	}
	for i := 0; i != n; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&current)
		if current > 0 {
			pos++
		} else if current < 0 {
			neg++
		} else {
			zero++
		}
	}
	fmt.Println("Положительных чисел:", pos)
	fmt.Println("Отрицательных чисел:", neg)
	fmt.Println("Нулей:", zero)
}
