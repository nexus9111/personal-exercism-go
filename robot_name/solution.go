package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var usedName []string

// Define the Robot type here.
type Robot struct {
	name string
}

func contain(s string) bool {
	for _, x := range usedName {
		if x == s {
			return true
		}
	}
	return false
}

func generateRadomName() string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < 2; i++ {
		result += fmt.Sprintf("%c", 'A'+rune(rand.Intn(26)))
	}
	for i := 0; i < 3; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}
	if contain(result) {
		return generateRadomName()
	}
	usedName = append(usedName, result)
	return result
}

func (r *Robot) Name() (string, error) {
	if len(r.name) != 5 {
		r.name = generateRadomName()
	}
	return r.name, nil

}

func (r *Robot) Reset() {
	for i, v := range usedName {
		if v == r.name {
			usedName = append(usedName[:i], usedName[i+1:]...)
			break
		}
	}
	r.name = generateRadomName()
}
