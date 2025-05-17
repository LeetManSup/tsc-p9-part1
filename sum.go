package calc

import "fmt"

// Sum возвращает сумму двух чисел
func Sum(a, b int) int {
	return a + b
}

// Divide выполняет деление a на b, с защитой от деления на 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
