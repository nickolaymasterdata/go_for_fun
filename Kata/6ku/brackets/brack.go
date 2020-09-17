package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(ValidParentheses(")"))
	fmt.Println(ValidParentheses(")"))
	fmt.Println(ValidParentheses("(({[]{()}}[{(())}[]]))"))

}

	func ValidParentheses(str string) bool {
		for pl, l := 0, len(str); pl != l; pl, l = l, len(str) {
			str = strings.ReplaceAll(str, "()", "")
			str = strings.ReplaceAll(str, "[]", "")
			str = strings.ReplaceAll(str, "{}", "")
		  }
		  return len(str) == 0
		}
		