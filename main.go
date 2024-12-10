package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	// Позволяет синхронно передавать результаты вычислений от горутин в главный поток.
	results := make(chan int, len(numbers))

	// Группа для синхронизации горутин
	// гарантирует, что главный поток дождется завершения всех горутин, прежде чем продолжить выполнение.
	var syncCalculate sync.WaitGroup

	//Функция для вычисления квадратов чисел
	squareNumber := func(num int) {
		defer syncCalculate.Done()
		results <- num * num
	}

	//Добавление горутины и старт функции для вычисления
	for _, num := range numbers {
		syncCalculate.Add(1)
		go squareNumber(num)
	}

	// Закрытие горутин
	go func() {
		syncCalculate.Wait()
		close(results)
	}()

	//Суммирование результата
	sum := 0
	for square := range results {
		sum += square
	}

	// Вывод результата
	fmt.Println("Сумма квадратов:", sum)
}
