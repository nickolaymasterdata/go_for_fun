
package main

import 
( 
"fmt"
"math"
 )


func main() { 
  
	
	fmt.Println(Going2((4))) //1.146651

} 


func Going(n int) float64 {
	var fact, factSum uint64 = 1, 0
	
	for i:= 1; i <= n; i ++ {
	 fact = fact * uint64(i)
	 factSum = factSum + fact
	} 

result := ( float64(float64(factSum) / float64 (fact )) )
fmt.Println(result)
result = math.Floor(result*1000000)/1000000
	return  result //( float64(float64(factSum) / float64 (fact ))
   }
   func Going2(n int) float64 {

	if n == 1 { 
		return float64(1) 
	} else {
	  return  ( 1 + ( 1 / float64(n) )  * Going2(n-1) )
	}
	}