package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		panic("n должно быть больше 1")
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("Число не является ни простым, ни составным")
		return
	}
	if isPrime(n) {
		fmt.Println("Простое число")
	} else {
		fmt.Println("Составное число")
	}
}
