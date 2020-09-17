package main

import (
	"fmt"
)

func main() {

	fmt.Println(birthday([]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1},18 , 7))
}

func birthday(s []int32, d int32, m int32) (result int32) {
	sum := int32(0)
	for idx, r := range s { 
		sum = sum + r
		if int32(idx) == m - 1 && sum == d {
			result++
		} else  if int32(idx) >= m {
			 sum = sum - s[int32(idx) - ( m -1 ) ] 
			 if sum == d {
				   result++
			 }
		}
	
	}
return result
}