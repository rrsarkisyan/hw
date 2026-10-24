package main

import "fmt"

func main() {
	// В задании не указаны ограничения на значения во вводе, поэтому использован тип с плавающей запятой
	var x, y float64
	var s string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&s)

	if y == 0 && s == "/" {
		panic("Разделить на ноль нельзя")
	}

	fmt.Print("Результат: ")

	switch s {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	default:
		fmt.Println("Неизвестная операция")
	}
}
