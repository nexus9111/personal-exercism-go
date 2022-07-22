package scrabble

import (
	"fmt"
	"strings"
)

var letters = []struct {
	letter []string
	value  int
}{
	{[]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}, 1},
	{[]string{"D", "G"}, 2},
	{[]string{"B", "C", "M", "P"}, 3},
	{[]string{"F", "H", "V", "W", "Y"}, 4},
	{[]string{"K"}, 5},
	{[]string{"J", "X"}, 8},
	{[]string{"Q", "Z"}, 10},
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
func Score(word string) int {
	word = strings.ToUpper(word)
	result := 0
	for pos := range word {
		for _, trans := range letters {
			if contains(trans.letter, fmt.Sprintf("%c", word[pos])) {
				result += trans.value
			}
		}
	}
	return result
}
