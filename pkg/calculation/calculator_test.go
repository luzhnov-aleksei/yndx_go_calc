package calculation

import (
	"testing"
)

func TestCalcNumbers(t *testing.T) {
	tests := []struct {
		num1, num2   float64
		operator     string
		expected     float64
		expectedErr  bool
	}{
		{1, 2, "+", 3, false},
		{5, 3, "-", 2, false},
		{4, 2, "*", 8, false},
		{6, 2, "/", 3, false},
		{6, 0, "/", 0, true}, // Ошибка деления на ноль
		{1, 1, "^", 0, true}, // Неизвестный оператор
	}

	for _, test := range tests {
		result, err := CalcNumbers(test.num1, test.num2, test.operator)
		if (err != nil) != test.expectedErr {
			t.Errorf("CalcNumbers(%f, %f, %s) error = %v, expectedErr = %v", test.num1, test.num2, test.operator, err, test.expectedErr)
		}
		if result != test.expected {
			t.Errorf("CalcNumbers(%f, %f, %s) = %f, expected = %f", test.num1, test.num2, test.operator, result, test.expected)
		}
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		expression   string
		expected     float64
		expectedErr  bool
	}{
		{"1 + 2", 3, false},
		{"5 - 3", 2, false},
		{"4 * 2", 8, false},
		{"6 / 2", 3, false},
		{"6 / 0", 0, true}, // Деление на ноль
		{"1 + 2 * 3", 7, false}, // Приоритет операторов
		{"(1 + 2) * 3", 9, false}, // Скобки
		{"2 + (3 * 4)", 14, false}, // Скобки с приоритетом
		{"1 + (2 * (3 + 4))", 15, false}, // Вложенные скобки
		{"1 + 2 ^ 3", 0, true}, // Неизвестный оператор
		{"(1 + 2", 0, true}, // Некорректное выражение
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if (err != nil) != test.expectedErr {
			t.Errorf("Calc(%s) error = %v, expectedErr = %v", test.expression, err, test.expectedErr)
		}
		if result != test.expected {
			t.Errorf("Calc(%s) = %f, expected = %f", test.expression, result, test.expected)
		}
	}
}
