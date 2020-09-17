package main

import 
( "strings" 
"fmt"
"strconv" )


func main() { 
  

	fmt.Println( is_valid_ip("12.255.56.22.22") )

} 

func is_valid_ip(ip string) bool {

for idx, mask := range strings.Split(ip, "." ) { 

  i, err := strconv.Atoi(mask)
  if ( i < 0 && i > 255 ) || idx > 3 || err != nil { 
		return false
	//	break
    }
  }
	return true
}