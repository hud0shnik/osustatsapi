package utils

import (
	"strconv"
	"strings"
)

// ---------------------- Функции поиска ------------------------

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func FindWithIndex(str, subStr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(subStr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(str[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return str[left : left+right], right + left + start
	}

	return "", start

}

// Функция поиска. Возвращает искомое значение без кавычек и индекс последнего символа
func FindStringWithIndex(str, subStr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(subStr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(str[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return strings.ReplaceAll(str[left:left+right], "\"", ""), right + left + start
	}

	return "", start

}

// Облегчённая функция поиска. Возвращает только искомое значение
func Find(str, subStr, stopChar string, start int) string {

	// Обрезка левой границы поиска
	str = str[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr)

	// Проверка на существование нужной строки
	if left != -1 {

		// Обрезка левой части
		str = str[left+len(subStr):]

		// Поиск и проверка правой границы
		right := strings.Index(str, stopChar)
		if right == -1 {
			return ""
		}

		// Обрезка правой части и вывод результата
		return str[:right]
	}

	return ""

}

// Функция поиска индекса
func Index(str, subStr string, start, end int) int {

	res := strings.Index(str[start:], subStr)

	// Проверка на существование нужной строки в диапазоне
	if res != -1 && ((end == -1) || (res+start < end)) {

		//fmt.Println(res+start, " - ", subStr)
		return res + start
	}

	//fmt.Println("index error: \t", subStr)
	return -1

}

// Функция проверки наличия подстроки
func Contains(str, subStr string, left int) bool {

	return strings.Contains(str[left:], subStr)
}

// ---------------------- Функции перевода ----------------------

func ToInt(s string) int {

	i, err := strconv.Atoi(s)
	if err != nil {
		// fmt.Println("parsing error: \t", s)
		return 0
	}

	return i

}

func ToInt64(s string) int64 {

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// fmt.Println("parsing error: \t", s)
		return 0
	}

	return i

}

func ToBool(s string) bool {

	f, err := strconv.ParseBool(s)
	if err != nil {
		// fmt.Println("parsing error: \t", s)
		return false
	}

	return f

}

func ToFloat64(s string) float64 {

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// fmt.Println("parsing error: \t", s)
		return 0
	}

	return i

}

func ToSlice(s string) []int {

	var result []int
	sliceStr := strings.Split(s, ",")

	if len(sliceStr) == 1 && sliceStr[0] == "" {
		return nil
	}

	for _, digit := range sliceStr {
		result = append(result, ToInt(digit))
	}

	return result

}
