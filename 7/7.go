package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)
	if n < 0 {
		panic("N не должно быть отрицательным")
	}
	for i := 1; i != n+1; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
