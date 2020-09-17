package main

import (
	"fmt"
	"strings"
)

func main() {

	//	fmt.Println(sherlockAndAnagrams("kkkk"))
	initialize("wuhmbspjnfviogqzldrcxtaeyk")
	fmt.Println(answerQuery(4,5))
	fmt.Println(answerQuery(2,3))

}


var message string
var addChar bool
func initialize(s string) {
    // This function is called once before all queries.
 message = s
}

/*
 * Complete the 'answerQuery' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER l
 *  2. INTEGER r
 */

func answerQuery(l int32, r int32) (result int32) {
   subString:= message[l-1:r]
   addChar = false
  /*length :=  len(subString)

  for i:= 0; i <= length; i++{
      if len(subString) == 0 {
        break   
      } 
  }*/
  result = Polindrom(subString)
  if addChar == true {
	result++ 
  }
  return   result
}

func Polindrom(subStr string) (result int32){

if len(subStr) == 1 {
    return
}
if len(subStr) == 0 {
    return
}
subStrToBeChecked :=  subStr[1:]

    newSub :=  strings.Replace(subStrToBeChecked, string(subStr[0]), "", 1 )
    if newSub != subStrToBeChecked  {
       return ( Polindrom( newSub)) + 1
    } else {
		addChar = true
		return ( Polindrom( newSub ))
		
	}

    return result
}
