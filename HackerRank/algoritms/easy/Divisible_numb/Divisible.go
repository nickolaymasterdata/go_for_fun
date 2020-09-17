package main

import (
	"fmt"
)

func main() {

	fmt.Println(divisibleSumPairs(0, 3, []int32{1, 3, 2, 6, 1, 2}))
}
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {

	var res int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if ((ar[i]+ar[j])%k) == 0 {
				res = res + 1
			}
		}
	}
	return res
}
