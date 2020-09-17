package main

import (
	"fmt"
	"math"
)

func main() {

//	fmt.Println(queensAttack(4, 3, 4, 4, [][]int32{}))
	//	fmt.Println(queensAttack(0, 2, 5, 3))

	fmt.Println(queensAttack(88587, 3, 20001, 20003, [][]int32{{20001,20002}, {20001, 20004}, {20000, 20003}, {20002, 20003}, {20000 ,20004},{20000 ,20002},{20002, 20004},	{20002, 20002},	{564, 323 } }))
	fmt.Println(queensAttack(5, 3, 4, 3, [][]int32{{4, 5}, {5, 5}, {4, 2}, {2, 3}}))	
	/*	4 3
		5 5
		4 2
		2 3 
20001,20002
20001, 20004
20000, 20003
20002, 20003
20000 ,20004
20000 ,20002
20002, 20004
20002, 20002
564, 323
		*/
}

func queensAttack(n int32, k int32, x_q int32, y_q int32, obstacles [][]int32) int32 {
	var upRight, downLeft, downRight, upLeft int32

	right := n - x_q
	left := x_q - 1
	up := n - y_q
	down := y_q - 1

	upRight = returnLowest(up, right)
	downLeft = returnLowest(down, left)
	downRight = returnLowest(down, right)
	upLeft = returnLowest(up, left)

	for _, i := range obstacles {
		if i[0] == x_q {
			if i[1] < y_q { //
				down = returnLowest(down, y_q-i[1]-1)
			} else {
				up = returnLowest(down, i[1]-y_q-1)
			}

		} else if i[1] == y_q {
			if i[0] < x_q {
				left = returnLowest(left, x_q-i[0]-1)
			} else {
				right = returnLowest(right, i[0]-x_q-1)
			}

			//|x_q - x_o| = |x_q - y_o|
		} else if math.Abs(float64(x_q-i[0])) == math.Abs(float64(y_q-i[1])) {
			if i[0] < x_q && i[1] < y_q {
				downLeft = returnLowest(downLeft, returnLowest(x_q-i[0]-1, y_q-i[1]-1))
			} else if i[0] < x_q && i[1] > y_q {
				downRight = returnLowest(downRight, returnLowest(x_q - i[0]-1, i[1] - y_q-1))
			} else if i[0] > x_q && i[1] < y_q {
				upLeft = returnLowest(upLeft, returnLowest( i[0] - x_q-1, y_q - i[1]-1))
			} else if i[0] > x_q && i[1] > y_q {
				upRight = returnLowest(upRight, returnLowest(i[0]-x_q-1, i[1]-y_q-1))
			}
		}
	}
	/*   for i:= int32(-1); i <= 1; i++ {
		for  j:= int32(-1); j <= 1; j++ {
				if !(i == 0 && j == 0) {
				totalMoves = totalMoves +  QueenCalc(n, x_q, y_q, obstacles, i, j, 0)
			}
		}
	} */
	return right + left + up + down + upRight + downLeft + downRight + upLeft
}

func returnLowest(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func returnHighest(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

/*	var totalMoves int32 = 0
	for i := int32(-1); i <= 1; i++ {
		for j := int32(-1); j <= 1; j++ {
			if !(i == 0 && j == 0) {
				totalMoves = totalMoves + QueenCalc(n, r_q, c_q, obstacles, i, j, 0)
			}
		}
	}
	return totalMoves
}*/
/*func QueenCalc(n int32, x_q int32, y_q int32, obstacles [][]int32, x int32, y int32, iterator int32) int32 {
	x_q = x_q + x
	y_q = y_q + y
	iterator = iterator + 1
	for _, i := range obstacles {
		if i[0] == x_q && i[1] == y_q {
			return iterator - 1
		}
	}
	if x_q > n || y_q > n || x_q < 1 || y_q < 1 {
		return iterator - 1
	}
	return QueenCalc(n, x_q, y_q, obstacles, x, y, iterator)
}*/
