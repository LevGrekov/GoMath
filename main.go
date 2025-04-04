package main

import (
	"bufio"
	"fmt"
	"matrix_example/calculator"
	"matrix_example/matrix"
	"os"
	"strings"
)

func main() {
	// Создаем две матрицы
	a := matrix.New(2, 3)
	b := matrix.New(3, 2)

	// Заполняем матрицу A
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 3)
	a.Set(1, 0, 4)
	a.Set(1, 1, 5)
	a.Set(1, 2, 6)

	// Заполняем матрицу B
	b.Set(0, 0, 7)
	b.Set(0, 1, 8)
	b.Set(1, 0, 9)
	b.Set(1, 1, 10)
	b.Set(2, 0, 11)
	b.Set(2, 1, 12)

	// Перемножаем матрицы
	result, err := a.Multiply(b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Выводим результат
	fmt.Println("Result of multiplication:")
	for i := 0; i < result.Rows(); i++ {
		for j := 0; j < result.Cols(); j++ {
			val, _ := result.Get(i, j)
			fmt.Printf("%6.1f", val)
		}
		fmt.Println()
	}

	// Генерируем случайную матрицу
	randomMatrix := matrix.GenerateRandom(3, 3)
	fmt.Println("\nRandom matrix:")
	for i := 0; i < randomMatrix.Rows(); i++ {
		for j := 0; j < randomMatrix.Cols(); j++ {
			val, _ := randomMatrix.Get(i, j)
			fmt.Printf("%6.1f", val)
		}
		fmt.Println()
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (или 'exit' для выхода): ")
		expr, _ := reader.ReadString('\n')
		expr = strings.TrimSpace(expr)

		if expr == "exit" {
			break
		}

		result, err := calculator.Calculate(expr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("Результат: %v\n", result)
	}

	fmt.Println("Калькулятор завершил работу")
}
