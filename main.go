package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := "ab12 3333 ddddd ddd"
	test2 := "+7 (904) 675 60-79"
	fmt.Println(normalize(test1))
	fmt.Println(normalize(test2))
}

func normalize(phone string) string {
	digits := "1234567890"
	var ret string
	for _, ch := range phone {
		if strings.Contains(digits, string(ch)) {
			ret += string(ch)
		}
	}
	return ret
}
