package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculate(input string) string {
	input = strings.TrimSpace(input)

	var str1, operator, str2 string
	var num int
	n, err := fmt.Sscanf(input, "%q %s %q", &str1, &operator, &str2)
	if n != 3 || err != nil {
		n, err = fmt.Sscanf(input, "%q %s %d", &str1, &operator, &num)
		if n != 3 || err != nil {
			panic("Неподходящее выражение")
		}
	}

	if len(str1) > 10 || (operator == "-" && len(str2) > 10) {
		panic("Строки должны быть длиной не более 10 символов")
	}

	switch operator {
	case "+":
		return limit(str1 + str2)
	case "-":
		return limit(subtract(str1, str2))
	case "*":
		if num < 1 || num > 10 {
			panic("Число должно быть от 1 до 10")
		}
		return limit(multiply(str1, num))
	case "/":
		if num < 1 || num > 10 {
			panic("Число должно быть от 1 до 10")
		}
		return limit(divide(str1, num))
	default:
		panic("Неподдерживаемая операция")
	}
}

func subtract(str1, str2 string) string {
	index := strings.Index(str1, str2)
	if index == -1 {
		return str1
	}
	return str1[:index] + str1[index+len(str2):]
}

func multiply(str string, n int) string {
	return strings.Repeat(str, n)
}

func divide(str string, n int) string {
	length := len(str)
	if n > length {
		return ""
	}
	return str[:length/n]
}

func limit(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := calculate(input)
	fmt.Println("Ответ: ", "\"", result, "\"")
}
