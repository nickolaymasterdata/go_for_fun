package main

import (
	"fmt"
	"math"
)

func main() {

//	fmt.Println(saveThePrisoner(46934, 543563655, 46743))

/*	fmt.Println(saveThePrisoner(530, 533048047, 529))
	fmt.Println(saveThePrisoner(436776012, 436776012, 436776011))
	fmt.Println(saveThePrisoner(999999999, 999999998, 999983945))
	fmt.Println(saveThePrisoner(59, 78693934, 2))
	fmt.Println(saveThePrisoner(49, 897910613, 48))
	fmt.Println(saveThePrisoner(124, 738996353, 2))
	fmt.Println(saveThePrisoner(999999999, 871077789, 999999998))
	fmt.Println(saveThePrisoner(2, 576581, 1))
	fmt.Println(saveThePrisoner(526663404, 801992888, 526663403))
	fmt.Println(saveThePrisoner(999999998, 999999998, 1))
	fmt.Println(saveThePrisoner(126, 859530642, 126))
	fmt.Println(saveThePrisoner(1000000000, 999999999, 1000000000))
	fmt.Println(saveThePrisoner(107, 425601402, 2))
	fmt.Println(saveThePrisoner(381, 695699141, 380))
	fmt.Println(saveThePrisoner(11, 32020900, 6))
	fmt.Println(saveThePrisoner(468840391, 468840391, 1))
	fmt.Println(saveThePrisoner(999999999, 29010, 1))
	fmt.Println(saveThePrisoner(31, 238250965, 2))
	fmt.Println(saveThePrisoner(6, 923562791, 1))
	fmt.Println(saveThePrisoner(39, 558119524, 38))
	fmt.Println(saveThePrisoner(121, 652798629, 1))
	fmt.Println(saveThePrisoner(94, 105224796, 94))
	fmt.Println(saveThePrisoner(9, 903414482, 5))
	fmt.Println(saveThePrisoner(1718761, 828441828, 1718761))
	fmt.Println(saveThePrisoner(4970962, 984250547, 4970961))
	fmt.Println(saveThePrisoner(19, 235344290, 2))
	fmt.Println(saveThePrisoner(514824323, 514824324, 514824323))
	fmt.Println(saveThePrisoner(181, 511813156, 180))
	fmt.Println(saveThePrisoner(66, 810757794, 2))
	fmt.Println(saveThePrisoner(154, 935852917, 154))
	fmt.Println(saveThePrisoner(1000000000, 999999999, 999974361))
	fmt.Println(saveThePrisoner(21, 603073253, 20))
	fmt.Println(saveThePrisoner(29, 834017184, 28))
	fmt.Println(saveThePrisoner(195446094, 586338283, 195446093))
	fmt.Println(saveThePrisoner(93995, 173193482, 93995))
	fmt.Println(saveThePrisoner(101, 143467773, 101))
	fmt.Println(saveThePrisoner(134, 677010612, 134))
	fmt.Println(saveThePrisoner(99, 741806010, 2))
	fmt.Println(saveThePrisoner(75, 129103876, 2))
	fmt.Println(saveThePrisoner(689371544, 689371544, 689370115))
	fmt.Println(saveThePrisoner(28410362, 340924345, 22721112))
	fmt.Println(saveThePrisoner(170, 780653100, 170))
	fmt.Println(saveThePrisoner(193, 945602138, 192))
	fmt.Println(saveThePrisoner(96, 23494832, 95))
	fmt.Println(saveThePrisoner(944675683, 944675683, 20312))
	fmt.Println(saveThePrisoner(27, 546238476, 26))
	fmt.Println(saveThePrisoner(76195990, 223258463, 489))
	fmt.Println(saveThePrisoner(999999999, 269208525, 1))
	fmt.Println(saveThePrisoner(108, 280122192, 108))
	fmt.Println(saveThePrisoner(16, 995404080, 15))
	fmt.Println(saveThePrisoner(50158215, 451423257, 50151646))
	fmt.Println(saveThePrisoner(4, 467711281, 4))
	fmt.Println(saveThePrisoner(145, 71654651, 144))
	fmt.Println(saveThePrisoner(1000000000, 2985, 1))
	fmt.Println(saveThePrisoner(990301380, 1, 990271854))
	fmt.Println(saveThePrisoner(999999999, 6413, 21476))
	fmt.Println(saveThePrisoner(2, 468939243, 1))
	fmt.Println(saveThePrisoner(399, 592025825, 398))
	fmt.Println(saveThePrisoner(8, 666688807, 8))
	fmt.Println(saveThePrisoner(7, 633100633, 2))
	fmt.Println(saveThePrisoner(12, 124444631, 7))
	fmt.Println(saveThePrisoner(8, 347412080, 5))
	fmt.Println(saveThePrisoner(999999999, 1, 999999998))
	fmt.Println(saveThePrisoner(42774012, 765544482, 2))
	fmt.Println(saveThePrisoner(18, 359622755, 18)) */
	fmt.Println(saveThePrisoner(198, 964246139, 2))
/*	fmt.Println(saveThePrisoner(999999999, 999999998, 999999999))
	fmt.Println(saveThePrisoner(10, 143638249, 1))
*/	fmt.Println(saveThePrisoner(1946080, 978878239, 2))/*
	fmt.Println(saveThePrisoner(1000000000, 999976501, 999990588))
	fmt.Println(saveThePrisoner(999999999, 999978713, 28209))
	fmt.Println(saveThePrisoner(433677591, 433663369, 206662538))
	fmt.Println(saveThePrisoner(999999999, 1, 1))
	fmt.Println(saveThePrisoner(2325, 562408200, 2))
	fmt.Println(saveThePrisoner(172, 456632596, 171))
	fmt.Println(saveThePrisoner(19, 563815520, 11))
	fmt.Println(saveThePrisoner(34339, 656699084, 101))
	fmt.Println(saveThePrisoner(1000000000, 999997154, 999999999))
	fmt.Println(saveThePrisoner(1000000000, 1000000000, 90143095))
	fmt.Println(saveThePrisoner(2, 213164653, 1))
	fmt.Println(saveThePrisoner(134, 644278309, 113))
	fmt.Println(saveThePrisoner(1000000000, 640282835, 2))
	fmt.Println(saveThePrisoner(1000000000, 1000000000, 999999999))
	fmt.Println(saveThePrisoner(999999999, 999999999, 999999999))
	fmt.Println(saveThePrisoner(999999999, 11132, 999999998))
	fmt.Println(saveThePrisoner(197, 190791557, 197))
	fmt.Println(saveThePrisoner(1000000000, 1000000000, 1000000000))
	*/fmt.Println(saveThePrisoner(46, 56740430, 45)) /*
	fmt.Println(saveThePrisoner(40, 277585960, 1))
	fmt.Println(saveThePrisoner(56, 306549319, 56))
	fmt.Println(saveThePrisoner(62, 803079454, 43))
	fmt.Println(saveThePrisoner(184, 834149464, 184))
	fmt.Println(saveThePrisoner(9, 526219551, 9))
	fmt.Println(saveThePrisoner(999999999, 999999999, 583101931))
	fmt.Println(saveThePrisoner(1000000000, 999999999, 999999999))
	fmt.Println(saveThePrisoner(176, 719643761, 1))
	fmt.Println(saveThePrisoner(1000000000, 999999999, 328966243))
	fmt.Println(saveThePrisoner(65, 980609890, 37))
*/
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	m = m + s - 1
	if m == n {
		return n
	}
	if m > n {
		div := int32(math.Trunc(float64(m / n)))
		m = m - (n * div)
		//	  fmt.Println(m)
	}
	if m == 0 {
		return s
	}
	return m
}
