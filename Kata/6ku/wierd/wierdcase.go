package main

import ( "strings"
"fmt" )


func main() { 
  
  fmt.Println( toWeirdCase("tEst case") )

} 

func toWeirdCase(str string) string {
  // Your code here and happy coding!
 var result string
  isUpper := true
  for _, ch := range  str { 
  if isUpper == true && ch != ' ' {
    result = result + strings.ToUpper(string(ch))
    isUpper = false
  }  else { 
    result = result + strings.ToLower(string(ch))
    isUpper = true
    }
  } 
  return result
}