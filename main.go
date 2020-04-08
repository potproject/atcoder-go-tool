package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d %d %d", c, b, a)
}

//-------------------------------------------
// util tools

// abs 絶対値を返す
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// sliceInt Intをスライスする
func sliceInt(k int) []int {
	s := make([]int, getDigit(k))
	count := 0
	for k > 0 {
		i := k % 10
		s[len(s)-1-count] = i
		count++
		k = k / 10
	}
	return s
}

// reverse int[]を反転
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// getDigit intの桁数を取得
func getDigit(num int) int {
	digit := 0
	for num != 0 {
		num /= 10
		digit++
	}
	return digit
}
