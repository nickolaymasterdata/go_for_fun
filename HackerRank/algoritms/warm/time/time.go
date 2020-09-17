package main

import (

    "fmt"
    "strconv"
)

func main() {
//fmt.Println(timeConversion("12:00:00PM"))
fmt.Println(timeConversion("12:00:00AM"))
fmt.Println(timeConversion("11:22:33AM"))
fmt.Println(timeConversion("11:22:33PM"))
fmt.Println(timeConversion("13:00:00PM"))
fmt.Println(timeConversion("13:22:33AM"))
}

func timeConversion(s string) string {
    hours, _ := strconv.Atoi(s[0:2])
    if s[8:10] == "PM" {
        if hours >= 12 {
            return  s[0:8]
        }
        hoursS := strconv.FormatInt( int64(12 + hours ), 10)
        return hoursS + string(s[2:8])
        } else {
            if  hours >= 12 {
                hoursS := strconv.FormatInt( int64( hours - 12 ), 10)
                return hoursS + string(s[2:8])
            } 
          return  s[0:8]

            }
        }