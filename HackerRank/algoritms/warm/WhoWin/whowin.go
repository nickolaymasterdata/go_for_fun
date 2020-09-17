package main

import (

    "fmt"
)

func main() {
//fmt.Println(timeConversion("12:00:00PM"))
fmt.Println(counterGame( 132 ) )
fmt.Println(counterGame( 6 ) )
fmt.Println(counterGame( 1 ) )
fmt.Println(counterGame( 2 ) )
}


func counterGame(n int64) string {
	count := whoWin(n, 0)
	 if count %2 == 0   {
		 return("Richard")//Louise and Richard
	 }
	 return "Louise"
 }
	 func whoWin( n int64, c int ) int { 
		if n == int64(1) {
            //    fmt.Println("n == 1", maxVal, n )
                return c
            }
        var maxVal int64 = 1
        for i:= 0; i <= 64; i++ {
            if maxVal < n {
                maxVal = maxVal * 2
            } else if maxVal == n {
                c++ 
               return whoWin( maxVal / 2, c )
            } else if maxVal > n {
                c++
             return  whoWin(n - ( maxVal / 2), c )
            }  
    } 
        return 0
        
    }