package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c *Calculator) evaluate(expression string) (interface{}, error) {
	tokens := strings.Fields(expression)
	if len(tokens) != 3 {
		return nil, fmt.Errorf("invalid expression format")
	}

	val1, err := strconv.Atoi(tokens[0])
	if err != nil {
		roman1, err := RomanToArabic(tokens[0])
		if err != nil {
			return nil, fmt.Errorf("invalid number format: %s", tokens[0])
		}
		val1 = roman1
	}

	val2, err := strconv.Atoi(tokens[2])
	if err != nil {
		roman2, err := RomanToArabic(tokens[2])
		if err != nil {
			return nil, fmt.Errorf("invalid number format: %s", tokens[2])
		}
		val2 = roman2
	}

	if (val1 < 1 || val1 > 10) || (val2 < 1 || val2 > 10) {
		return nil, fmt.Errorf("the entered number must be from 1 to 10 inclusive")
	}

	isArabic1 := isArabic(tokens[0])
	isArabic2 := isArabic(tokens[2])

	if (isArabic1 && !isArabic2) || (!isArabic1 && isArabic2) {
		return nil, fmt.Errorf("either only Arabic or Roman numerals are allowed")
	}

	operator := tokens[1]
	result := 0
	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 == 0 {
			return nil, fmt.Errorf("division by zero is not allowed")
		}
		result = val1 / val2
	default:
		return nil, fmt.Errorf("invalid operation: %s", operator)
	}

	if !isArabic1 {
		if romanResult, ok := toRoman(result); ok {
			return romanResult, nil
		}
	}

	return result, nil
}

func isArabic(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func (c *Calculator) parseOperand(operand string) (int, error) {
	arabic, err := strconv.Atoi(operand)
	if err == nil {
		return arabic, nil
	}

	roman, err := RomanToArabic(operand)
	if err != nil {
		return 0, fmt.Errorf("invalid number format: %s", operand)
	}

	return roman, nil
}

func RomanToArabic(roman string) (int, error) {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10}
	result := 0
	for k := range roman {
		if k < len(roman)-1 && rMap[roman[k:k+1]] < rMap[roman[k+1:k+2]] {
			result -= rMap[roman[k:k+1]]
		} else {
			result += rMap[roman[k:k+1]]
		}
	}
	return result, nil
}

func toRoman(arabic int) (string, bool) {
	if arabic <= 0 {
		return "", false
	}

	var result strings.Builder

	for _, numeral := range []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	} {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), true
}

func main() {
	calculator := &Calculator{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the expression")

	for scanner.Scan() {
		expression := scanner.Text()

		if expression == "0" {
			fmt.Println("Goodbye.")
			break
		}

		result, err := calculator.evaluate(expression)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}

		fmt.Println("Enter the next expression (or '0' for exit):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error input:", err)
	}
}
