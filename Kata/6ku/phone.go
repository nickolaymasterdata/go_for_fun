package main

import (
	"fmt"
	"strconv"
)

func main() {

	phone := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	fmt.Println(phone)
}

func CreatePhoneNumber(numbers [10]uint) string {

	var Results string = "("

	for idx, number := range numbers {
		Results = Results + strconv.FormatUint(uint64(number), 10)
		if idx == 2 {

			Results = Results + ") "

		} else if idx == 5 {

			Results = Results + "-"
		} else {

		}
	}

	return Results
}
