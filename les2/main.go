package main

import "fmt"

func DuplicateEncode(word string) string {
	strToRune := []rune(word)
	count := 1
	for _, v := range strToRune {
		if v 
	}
	return string(strToRune)
}

func main() {
	str := "recede"
	fmt.Println(str)
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
