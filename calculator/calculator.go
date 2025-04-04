package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(expression string) (float64, error) {
	fmt.Printf("Вычисление выражения: %q\n", expression)

	// 1. Токенизация
	fmt.Println("\n=== Этап 1: Токенизация ===")
	tokens, err := tokenize(expression)
	if err != nil {
		fmt.Println("Ошибка токенизации:", err)
		return 0, err
	}
	fmt.Printf("Токены: %v\n", tokens)

	// 2. Преобразование в ОПН
	fmt.Println("\n=== Этап 2: Преобразование в ОПН ===")
	rpn, err := shuntingYard(tokens)
	if err != nil {
		fmt.Println("Ошибка преобразования в ОПН:", err)
		return 0, err
	}
	fmt.Printf("ОПН: %v\n", rpn)

	// 3. Вычисление ОПН
	fmt.Println("\n=== Этап 3: Вычисление ОПН ===")
	result, err := evalRPN(rpn)
	if err != nil {
		fmt.Println("Ошибка вычисления:", err)
		return 0, err
	}
	fmt.Printf("Промежуточные результаты: %v\n", rpn)

	fmt.Println("\n=== Результат ===")
	fmt.Printf("%q = %v\n", expression, result)
	return result, nil
}

// Разбиение строки на токены (числа, операторы, скобки)
func tokenize(expr string) ([]string, error) {
	var tokens []string
	var buffer strings.Builder

	for _, ch := range expr {
		if unicode.IsSpace(ch) {
			continue
		}

		if unicode.IsDigit(ch) || ch == '.' {
			buffer.WriteRune(ch)
		} else {
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, buffer.String())
	}

	return tokens, nil
}

func shuntingYard(tokens []string) ([]string, error) {
	var output []string
	stack := Stack[string]{}

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {

		if _, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, token)
			continue
		}

		switch token {
		case "(":
			stack.Push(token)
		case ")":
			foundMatching := false
			for {
				top, ok := stack.Pop()
				if !ok {
					return nil, fmt.Errorf("непарная скобка")
				}
				if top == "(" {
					foundMatching = true
					break
				}
				output = append(output, top)
			}
			if !foundMatching {
				return nil, fmt.Errorf("непарная скобка")
			}
		default:
			if op, isOp := precedence[token]; isOp {
				for {
					top, ok := stack.Pop()
					if !ok {
						break
					}
					if top == "(" {
						stack.Push(top)
						break
					}
					if precedence[top] >= op {
						output = append(output, top)
					} else {
						stack.Push(top)
						break
					}
				}
				stack.Push(token)
			} else {
				return nil, fmt.Errorf("неизвестный токен: %s", token)
			}
		}
	}

	// Обработка оставшихся операторов
	for {
		top, ok := stack.Pop()
		if !ok {
			break
		}
		if top == "(" {
			return nil, fmt.Errorf("непарная скобка")
		}
		output = append(output, top)
	}

	return output, nil
}

func evalRPN(rpn []string) (float64, error) {
	stack := Stack[float64]{}

	for _, token := range rpn {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(num)
		} else {
			// Получаем два верхних значения из стека
			b, ok := stack.Pop()
			if !ok {
				return 0, fmt.Errorf("недостаточно операндов")
			}
			a, ok := stack.Pop()
			if !ok {
				return 0, fmt.Errorf("недостаточно операндов")
			}

			var res float64
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				res = a / b
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}
			stack.Push(res)
		}
	}

	result, ok := stack.Pop()
	if !ok {
		return 0, fmt.Errorf("ошибка вычисления: пустой стек")
	}
	if _, ok := stack.Pop(); ok {
		return 0, fmt.Errorf("ошибка вычисления: в стеке остались лишние элементы")
	}

	return result, nil
}
