package main

import ( 
	"math" 
	"fmt" )

func main() { 

	fmt.Println( IterPi(0.1) )

}
	func IterPi(epsilon float64) (int, string) {
		var myPi, lorDiv float64 = 4, 4
		var lor float64 = 3
		var count int
	   for count = 1; count <= 1000000; count++ { 
		lorDiv = 4 / lor
		 if count%2 == 0 {
			  myPi = myPi + lorDiv
		 } else {
			myPi = myPi - lorDiv
		 }
	   lor = lor + 2
	   if epsilon >  math.Abs(myPi - math.Pi ) { 
		return count, fmt.Sprintf("%.10f", myPi)
	   }
	  }
	  	return 0, " "
	}