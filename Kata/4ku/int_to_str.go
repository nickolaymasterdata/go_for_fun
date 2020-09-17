package main

import (
	"fmt"
  "math"
)

func main() {

	fmt.Println(NextBigger(4863242620))
	fmt.Println(NextBigger(441))
	fmt.Println(NextBigger(144))
	fmt.Println(NextBigger(9))
	fmt.Println(NextBigger(8655543333200))
}

func NextBigger(n int) int {
	slice := intToSlice(n, []int{})
	slice = SortSlice(slice)
	newN := sliceToInt(slice)

	if newN <= n  {
		return -1
	}
	return newN
}

func intToSlice(n int, slice []int) []int {
	if n != 0 {
		i := n % 10
    slice =  append([]int{i}, slice...)
		return intToSlice(n/10, slice)
	}
	return slice
}

func SortSlice(slice []int) []int {
	var temp int = 0
	for i := len(slice) - 1; i >= 0; i-- {
		if temp < slice[i] {
			temp = slice[i]
		} else if temp > slice[i] {
			slice[i+1] = slice[i]
			slice[i] = temp
			return slice
		}
	}
	return slice
}

func sliceToInt(slice []int) int {
	var newN int
	for i := 0; i<= len(slice) - 1; i++ {
		newN = newN + slice[i]*(int(math.Pow(float64(10), float64(len(slice) - i - 1 ))))
	}
	return newN
}