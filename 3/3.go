package main

import "fmt"

func main() {
	var n, sum int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	if n < 1 {
		panic("N должно быть положительным")
	}
	for i := 1; i != n+1; i++ {
		sum += i
	}

	mean := float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее значение:", mean)
}
