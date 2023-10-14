package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func detectNumType(input string) string {
	romanNumerals := "IVXLCDM"
	for _, char := range input {
		if strings.ContainsRune(romanNumerals, char) {
			return "roman"
		}
	}
	return "arabic"

}

func isValidValue(value int) bool {
	return value >= 1 && value <= 10
}

func romanToArabic(roman string) int {
	mappings := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	total := 0
	prevValue := 0
	for _, char := range roman {
		value := mappings[char]
		if value > prevValue {
			total += value - 2*prevValue

		} else {
			total += value
		}
		prevValue = value
	}
	return total

}

func arabicToRoman(num int) string {
	if num <= 0 {
		return ""
	}
	mappings := []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	result := ""
	for _, mapping := range mappings {
		for num >= mapping.value {
			result += mapping.symbol
			num -= mapping.value
		}
	}
	return result
}
func calculate(input string) string {

	operators := []string{"+", "-", "*", "/"}

	var op string

	for _, operator := range operators {
		if strings.Contains(input, operator) {
			op = operator
			break
		}
	}

	if op == "" {
		return "Ошибка: неверная операция"
	}

	parts := strings.Split(input, op)

	if len(parts) != 2 {
		return "Ошибка: неверный формат ввода"
	}

	a, b := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	numberType := detectNumType(a)

	if numberType != detectNumType(b) {
		return "Ошибка:числа разных типов"
	}

	var aValue, bValue int

	if numberType == "roman" {
		aValue = romanToArabic(a)
		bValue = romanToArabic(b)
	} else {
		aValue, _ = strconv.Atoi(a)
		bValue, _ = strconv.Atoi(b)

	}
	if !isValidValue(aValue) || !isValidValue(bValue) {
		return "Ошибка: числа должны быть от 1 до 10 включительно"
	}
	var result int

	switch op {
	case "+":
		result = aValue + bValue
	case "-":
		result = aValue - bValue
	case "*":
		result = aValue * bValue
	case "/":
		result = aValue / bValue

	}
	if numberType == "roman" {
		if result <= 0 {
			return "Ошибка:результат меньше или равен нулю"
		}
		return arabicToRoman(result)
	}
	return fmt.Sprintf("%d", result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" || text == "quit" {
			break
		}
		fmt.Println(calculate(text))
	}

}
