package calculation

import (
	"fmt"
	"regexp"
	"strconv"
)

func CalcNumbers(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("ошибка: на ноль делить нельзя")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("ошибка: неизвестный оператор")
	}
}


func Calc(expression string) (float64, error) {
	var values []float64
	var operators []string

	
	calculate := func() error {
		if len(values) < 2 || len(operators) == 0 {
			return fmt.Errorf("ошибка: недостаточно операндов")
		}
		num2 := values[len(values)-1]
		num1 := values[len(values)-2]
		operator := operators[len(operators)-1]

		result, err := CalcNumbers(num1, num2, operator)
		if err != nil {
			return err
		}

		
		values = values[:len(values)-2]
		values = append(values, result)
		operators = operators[:len(operators)-1]
		return nil
	}

	
	re := regexp.MustCompile(`(\d+\.?\d*|[+\-*/()])`)
	tokens := re.FindAllString(expression, -1)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				if err := calculate(); err != nil {
					return 0, err
				}
			}
			operators = append(operators, token)

		case "(":
			operators = append(operators, token)

		case ")":
			
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				if err := calculate(); err != nil {
					return 0, err
				}
			}
			if len(operators) == 0 {
				return 0, fmt.Errorf("ошибка: нет открывающей скобки")
			}
			operators = operators[:len(operators)-1] // Удаляем открывающую скобку

		default:
			
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("ошибка: некорректный токен: %s", token)
			}
			values = append(values, num)
		}
	}

	
	for len(operators) > 0 {
		if err := calculate(); err != nil {
			return 0, err
		}
	}

	
	if len(values) != 1 {
		return 0, fmt.Errorf("ошибка: недостающие операнды")
	}

	return values[0], nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}


