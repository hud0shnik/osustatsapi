package convert

import (
	"strconv"
	"strings"
)

// ToInt переводит string в int
func ToInt(s string) int {

	i, _ := strconv.Atoi(s)

	return i

}

// ToInt64 переводит string в int64
func ToInt64(s string) int64 {

	i, _ := strconv.ParseInt(s, 10, 64)

	return i

}

// ToBool переводит string в bool
func ToBool(s string) bool {

	f, _ := strconv.ParseBool(s)

	return f

}

// ToFloat64 переводит string в float64
func ToFloat64(s string) float64 {

	i, _ := strconv.ParseFloat(s, 64)

	return i

}

// ToSlice переводи string в []int. За сепаратор используется запятая
func ToSlice(s string) []int {

	result := []int{}
	sliceStr := strings.Split(s, ",")

	if len(sliceStr) == 1 && sliceStr[0] == "" {
		return nil
	}

	for _, digit := range sliceStr {
		result = append(result, ToInt(digit))
	}

	return result

}
