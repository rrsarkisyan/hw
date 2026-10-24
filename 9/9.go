package main

import "fmt"

func main() {
	var (
		balance, amount float64
		action          int
	)
	fmt.Print("Начальный баланс: ")
	fmt.Scan(&balance)
	if balance < 0 {
		panic("\nБаланс не может быть отрицательным")
	}

	fmt.Println("Действия: посмотреть баланс - 1, пополнить " +
		"баланс - 2, снять средства - 3, завершение работы - 4.")

	for action != 4 {
		fmt.Print("Действие: ")
		fmt.Scan(&action)

		switch action {
		case 1:
			fmt.Printf("Баланс: %.2f\n", balance)
		case 2:
			fmt.Print("Сумма пополнения: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть положительной")
			} else {
				balance += amount
				fmt.Printf("Новый баланс: %.2f\n", balance)
			}
		case 3:
			fmt.Print("Введите сумму снятия: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть положительной")
			} else if amount > balance {
				fmt.Println("Недостаточно средств")
			} else {
				balance -= amount
				fmt.Printf("Новый баланс: %.2f\n", balance)
			}
		case 4:
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
