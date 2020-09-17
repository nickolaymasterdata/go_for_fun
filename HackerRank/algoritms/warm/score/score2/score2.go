package main

import (
	"fmt"
)

func main() {
	//	var i int32
//	var result []int32
	/*	for i = 99999832; i >= 1; i--{
		result = append(result, i)
	}*/
//	fmt.Println(climbingLeaderboard(result, result))
	fmt.Println(climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}))
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var result []int32
	var max, min int32
//	var j, k int32
//	var isAdded bool
	sortedScore := deleteDuplicates(scores)
	length := int32(len(scores))
	max, min = scores[1], scores[ length -1 ]
	for _, i := range alice {
		/*	isAdded = false
			for j = k; j >= 0; j-- {
				if sortedScore[j] == i {
					k = j
					result = append(result, j+1)
					isAdded = true
					break
				} else if sortedScore[j] > i {
					result = append(result, j+2)
					isAdded = true
					k = j
					break
				}
			}
			if isAdded == false {
				result = append(result, 1) */
		score, scoreExist := sortedScore[i]
		if scoreExist == false {
			if max < i  {
				result = append(result, 1)
			} else if min > i  {
				result = append(result, sortedScore[min] + 1)
			}
		} else { 
			result = append(result, score)
		}
	}
	return result
}

func deleteDuplicates(toBeChecked []int32) ( sorted map[int32]int32 ) {
	var oldVal, count  int32

	sorted = make(map[int32]int32, len(toBeChecked) )
	count = 1
	for _, i := range toBeChecked {
		if i != oldVal {
			sorted[i] = count
			if oldVal != 0 {
				for k := oldVal-1; k > i; k-- {
					sorted[k] = (count )
				}
			}
			oldVal = i
			count++
		}
	}
	return sorted
}

/*func findNumb(currentScore int32, Scorelist []int32, length int32) int32 {
	var j int32
	for j = length - 1; j >= 0; j-- {
		if Scorelist[j] == currentScore {
			return j + 1

		} else if Scorelist[j] > currentScore {
			return j + 2
		}
	}
	return 1
}*/
 