package main

import (
	"fmt"
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	// Your code here and happy coding!
	sliceStr := strings.Split(str, " ")
	//fmt.Println(sliceStr)
	strToRune := ""
	for _, val := range sliceStr {
		valAsRune := []rune(val)
		for i := 0; i < len(valAsRune); i++ {
			if i%2 == 0 {
				strToRune += string(unicode.ToUpper(valAsRune[i]))
			} else {
				strToRune += string(unicode.ToLower(valAsRune[i]))
			}
		}
		strToRune += " "
	}
	return strings.TrimRight(strToRune, " ")
}

/*
func toWeirdCase(str string) string {
	output := ""

	index := 0
	for _, c := range str {
		if c == ' ' {
			index = 0
			output += " "
			continue
		}
		if index%2 == 0 {
			output += string(unicode.ToUpper(c))
		} else {
			output += string(unicode.ToLower(c))
		}
		index++
	}

	return output
}*/

func main() {
	str := "Hi Lola"
	fmt.Println(toWeirdCase(str))
}
