
package main

import (
	"fmt"
//	"math"

)

func main() {

		fmt.Println(pickingNumbers([]int32{4, 6, 5 ,3, 3, 1}))
		fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
	}



func pickingNumbers(a []int32) ( res int32 ) {
    
	for _, i:= range a {
		min, max, eq := 0, 0, 0   
		   for _, j:= range a{
			   if i == j {
				   eq = eq + 1
			   } else if i + 1 == j {
					min = min + 1
			   } else if j + 1 == i { 
				   max = max + 1
			   }
		   }
			if min < max {
				   eq = eq + max
			   } else {
			   eq = eq + min
			   }
			   if res < int32(eq) {
				   res = int32(eq)
			   }
		   
	   } 
	   return res
   }