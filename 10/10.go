package main

import "fmt"

func main() {
	var n, current, even, odd int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	if n <= 0 {
		panic("N должно быть положительным")
	}
	fmt.Print("Введите число: ")
	fmt.Scan(&current)

	min, max, sum := current, current, current
	if current%2 == 0 {
		even++
	} else {
		odd++
	}

	for i := 0; i != n-1; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&current)
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
		sum += current
		if current%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", float64(sum)/float64(n))
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Количество чётных чисел:", even)
	fmt.Println("Количество нечётных чисел:", odd)
}
