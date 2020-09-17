package main

import (
	"fmt"
	"strings"
)

func main() {

	//	fmt.Println(sherlockAndAnagrams("kkkk"))
	fmt.Println(sherlockAndAnagrams("abba"))
	fmt.Println(sherlockAndAnagrams("abba"))

}

func sherlockAndAnagrams(s string) (res int32) {
	var sub, subMainString string
	len := len(s)
	for step := 1; step+1 <= len; step++ {
		for iter := 0; iter+step < len; iter++ {

			if step > len-iter {
				return
			}
			if step == 1 {
				sub = string(s[iter])
			} else {
				sub = s[iter : step+iter]
			}
			subMainString = s[1+iter : len]

			for j := 0; step+j <= len-1-iter; j++ {
				res = res + CompareStrings(subMainString[j:step+j], sub)
			}
		}

	}
	return res
}
func CompareStrings(fullString string, subString string) (res int32) {

	if fullString == subString {
		return 1
	}
	for _, char := range subString {
		if fullString == "" {
			return 0
		}
		oldString := fullString
		fullString = strings.Replace(fullString, string(char), "", 1)
		if oldString == fullString {
			return 0
		}
	}
	if fullString == "" {
		return 1
	}

	//    res = int32(len(strings.Split(fullString, subString))) - 1
	//    fmt.Println("Result of comparesson:", subString, fullString)
	return 0
}
