package main

import (
	"fmt"
	"math"
)

func main() {


	fmt.Println(  beautifulDays(20, 23, 6)  )

}

func beautifulDays(i int32, j int32, k int32) int32 {
	var reverse_val  int32
	var res  int32
	for l := i; l <= j; l++ {
	   reverse_val = 0
		temp_int := l
		for temp_int > 0 {
			remainder := temp_int % 10
			reverse_val *= 10
			reverse_val += remainder 
			temp_int /= 10
			}
		if  math.Mod( float64(l - reverse_val), float64(k) ) == 0 {
			res++
			}
		}
		return res
	
	}