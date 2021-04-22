package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	changeWordToLowerCase := strings.ToLower(word)
	strToRune := []rune(changeWordToLowerCase)
	result := ""
	countRepetInRune := 0
	for _, v := range strToRune {
		for _, val := range strToRune {
			if v == val {
				countRepetInRune++
			}
		}
		if countRepetInRune == 1 {
			result += "("
		} else {
			result += ")"
		}
		countRepetInRune = 0
	}
	return result
}

func main() {
	str := "Success"
	fmt.Println(str)
	fmt.Println(DuplicateEncode(str))
}

/*
The goal of this exercise is to convert a string to a new string where each character in the new string is
"(" if that character appears only once in the original string, or ")" if that character appears more than
 once in the original string. Ignore capitalization when determining if a character is a duplicate.

 Examples
"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))(("


*/
