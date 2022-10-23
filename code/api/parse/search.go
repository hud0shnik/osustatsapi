package parse

import (
	"strings"
)

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func findWithIndex(str, subStr, stopChar string, start int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Проверка на существование нужной строки
	if strings.Contains(str, subStr) {

		// Поиск индекса начала нужной строки
		left := strings.Index(str, subStr) + len(subStr)

		// Поиск правой границы
		right := left + strings.Index(str[left:], stopChar)

		// Обрезка и вывод результата
		return str[left:right], right + start
	}

	// Вывод ненайденных значений для тестов
	// fmt.Println("error foundn't \t", subStr, "-")

	return "", start
}

// Облегчённая функция поиска. Возвращает только искомое значение
func find(str, subStr, stopChar string, start int) string {

	str = str[start:]
	left := strings.Index(str, subStr)

	// Проверка на существование нужной строки
	if left != -1 {

		// Обрезка левой части
		str = str[left+len(subStr):]

		// Обрезка правой части и вывод результата
		return str[:strings.Index(str, stopChar)]
	}

	return ""
}

// Функция поиска индекса
func index(str, subStr string, start int) int {

	res := strings.Index(str[start:], subStr)

	// Проверка на существование нужной строки
	if res == -1 {

		//fmt.Println("index error: \t", subStr)

		return -1
	}

	//fmt.Println(res+start, " - ", subStr)
	return res + start
}

// Функция проверки наличия подстроки
func contains(str, subStr string, left int) bool {

	return strings.Contains(str[left:], subStr)
}