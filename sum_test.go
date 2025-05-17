package calc

import "testing"

// Тестируем обычную сумму
func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sum(2, 3) = %d; want %d", result, expected)
	}
}

// Тестируем деление с результатом
func TestDivide(t *testing.T) {
	result, err := Divide(6, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 3
	if result != expected {
		t.Errorf("Divide(6, 2) = %d; want %d", result, expected)
	}
}

// Тестируем деление на ноль
func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected error on division by zero, got nil")
	}
}
