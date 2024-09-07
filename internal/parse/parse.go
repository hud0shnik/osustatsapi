package parse

import "strings"

// FindWithIndex производит поиск substr в s[start:end] и возвращает строку от конца substr до stopChar
func FindWithIndex(s, substr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	s = s[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(s, substr) + len(substr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(substr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(s[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return s[left : left+right], right + left + start
	}

	return "", start

}

// FindStringWithIndex работает как FindWithIndex, только убирает кавычки
func FindStringWithIndex(s, substr, stopChar string, start, end int) (string, int) {

	// Обрезка левой границы поиска
	s = s[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(s, substr) + len(substr)

	// Проверка на существование нужной строки и попадание в диапазон
	if left != len(substr)-1 && ((end == -1) || (left+start < end)) {

		// Поиск и проверка правой границы
		right := strings.Index(s[left:], stopChar)
		if right == -1 {
			return "", start
		}

		// Обрезка и вывод результата
		return strings.ReplaceAll(s[left:left+right], "\"", ""), right + left + start
	}

	return "", start

}

// Find работает как FindWithIndex, но не возвращает индекс
func Find(s, substr, stopChar string, start int) string {

	// Обрезка левой границы поиска
	s = s[start:]

	// Поиск индекса начала нужной строки
	left := strings.Index(s, substr)

	// Проверка на существование нужной строки
	if left != -1 {

		// Обрезка левой части
		s = s[left+len(substr):]

		// Поиск и проверка правой границы
		right := strings.Index(s, stopChar)
		if right == -1 {
			return ""
		}

		// Обрезка правой части и вывод результата
		return s[:right]
	}

	return ""

}

// Index возвращает индекс substr в рамках s[start:end]
func Index(s, substr string, start, end int) int {

	res := strings.Index(s[start:], substr)

	// Проверка на существование нужной строки в диапазоне
	if res != -1 && ((end == -1) || (res+start < end)) {

		//fmt.Println(res+start, " - ", substr)
		return res + start
	}

	//fmt.Println("index error: \t", substr)
	return -1

}

// Contains показывает существует ли substr в s
func Contains(s, substr string, left int) bool {
	return strings.Contains(s[left:], substr)
}
