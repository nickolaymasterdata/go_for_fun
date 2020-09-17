package main

import (
	"fmt"
	"math"

)

func main() {

		fmt.Println(kangaroo( 0, 3, 4, 2 ))
		fmt.Println(kangaroo(0,2,5,3))
	}






func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    return recur( x1, v1, x2, v2, int32( math.Abs( float64( (x1 + v1 ) - (x2+ v2))))   )
}
func recur( x1 int32, v1 int32, x2 int32, v2 int32, delta int32 ) string {
    x1 = x1 + v1
    x2 = x2+ v2
    if x1 == x2 {
        return "YES"
    }
    newDelta := int32( math.Abs( float64( (x1 + v1 ) - (x2+ v2)))) 
    if newDelta <= delta {
        return recur( x1, v1, x2, v2, newDelta )
    }
    return "NO"
}