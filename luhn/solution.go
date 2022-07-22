package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	// check id only contain digits
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	if len(id) <= 1 || strings.IndexFunc(id, isNotDigit) != -1 {
		return false
	}

	id = Reverse(id)
	sumToCheck := 0

	for index, char := range id {
		intVar, err := strconv.Atoi(fmt.Sprintf("%c", char))
		if err != nil {
			return false
		}
		if index%2 != 0 {
			intVar = intVar * 2
			if intVar > 9 {
				intVar = intVar - 9
			}
			sumToCheck += intVar
		} else {
			sumToCheck += intVar
		}
	}
	return sumToCheck%10 == 0
}
