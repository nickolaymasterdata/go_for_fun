package main

import ( "fmt" 
"strings" )

func main() { 
  

	fmt.Println( ValidBraces("(({{[[]]}}))") )

} 

func ValidBraces(str string) bool {
	for i:= 0; i <= len(str); i++ {
		str = strings.ReplaceAll(str, "[]", "")
		str = strings.ReplaceAll(str, "{}", "")
		str = strings.ReplaceAll(str, "()", "")
	  }
	  return len(str) == 0
	  


/*	slice := make([]rune, len(str) + 2, len(str) + 2 )
	slice = []rune(str)
	return RecurBraces(slice) */
  }
  


  /*
  
 func RecurBraces(braces []rune ) bool {
	  if len(braces) == 1 { 
		  return false 
	}
	  tempSl := braces
	  for idx, char := range braces {
		if char == '(' && braces[idx+1] == ')' ||
		   char == '[' && braces[idx+1] == ']' ||
		   char == '{' && braces[idx+1] == '}' { 
			  copy(tempSl[idx:], tempSl[idx+2:]) // 
			  tempSl[len(tempSl)-1] = ' ' 
			  tempSl[len(tempSl)-2] = ' '
			  tempSl = tempSl[:len(tempSl)-2]
			 if len(tempSl) == 0 ||  RecurBraces( tempSl ) {
			   return true
			 }
		   }
	  }  
		return false 
	  }
  */
  