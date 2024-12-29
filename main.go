package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите выражение:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	result := calculate(input)
	fmt.Println(result)
}

func calculate(input string) string {
	// Удаляем пробелы
	// input = strings.TrimSpace(input)

	// Проверяем на наличие операций
	if strings.Contains(input, "+") {
		return addition(input)
	} else if strings.Contains(input, "-") {
		return subtraction(input)
	} else if strings.Contains(input, "*") {
		return multiplication(input)
	} else if strings.Contains(input, "/") {
		return division(input)
	}

	panic("Недопустимая операция")
}

func addition(input string) string {
	parts := strings.Split(input, "+")
	if len(parts) != 2 {
		panic("Неверный формат сложения")
	}

	str1 := strings.Trim(parts[0], " ")
	str2 := strings.Trim(parts[1], " ")

	return str1 + str2
}

func subtraction(input string) string {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		panic("Неверный формат вычитания")
	}

	str1 := strings.Trim(parts[0], " ")
	str2 := strings.Trim(parts[1], " ")

	if !strings.Contains(str1, str2) {
		return str1
	}

	return strings.ReplaceAll(str1, str2, "")
}

func multiplication(input string) string {
	parts := strings.Split(input, "*")
	if len(parts) != 2 {
		panic("Неверный формат умножения")
	}

	str1 := strings.Trim(parts[0], " ")
	numStr := strings.Trim(parts[1], " ")

	n, err := strconv.Atoi(numStr)
	if err != nil || n < 1 || n > 10 {
		panic("Неподходящее число для умножения")
	}

	return strings.Repeat(str1, n)
}

func division(input string) string {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		panic("Неверный формат деления")
	}

	str1 := strings.Trim(parts[0], " ")
	numStr := strings.Trim(parts[1], " ")

	n, err := strconv.Atoi(numStr)
	if err != nil || n < 1 || n > 10 {
		panic("Неподходящее число для деления")
	}

	length := len(str1)
	dividedLength := length / n
	result := str1[:dividedLength]

	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
