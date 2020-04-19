package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush()
	a, b := nextInt(), nextInt()
	i := nextInts(a)
	fmt.Println(a, b)
	fmt.Println(i)

}

//-------------------------------------------
// I/O util tools

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func nextFloat64() float64 {
	a, _ := strconv.ParseFloat(next(), 64)
	return a
}

func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}
func nextFloats(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = nextFloat64()
	}
	return ret
}
func nextStrings(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = next()
	}
	return ret
}

func chars(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
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

// sliceInt []intをスライスする
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

// sliceIntSum []intの合計を出す
func sliceIntSum(v []int) int {
	sum := 0
	for _, x := range v {
		sum += x
	}
	return sum
}

// reverse []intを反転
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

// getSumDigit intの各位の和を返す
func getSumDigit(num int) int {
	sum := 0
	for num != 0 {
		t := num % 10
		num /= 10
		sum += t
	}
	return sum
}

// lastPrime 素数 エラトステネスの篩
func lastPrime(max int64) int64 {
	if max <= 1 {
		return 2
	}
	primes := make([]int64, 1, max)
	primesF := make([]float64, 1, max)
	primes[0] = 2
	primesF[0] = 2.0

	count := int64(1)
	for n := int64(3); ; n += 2 {
		flag := true
		f := float64(n)
		rf := math.Sqrt(f)
		for i := 1; i < len(primes); i++ {
			if primesF[i] > rf {
				break
			} else if (n % primes[i]) == 0 {
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
			primes = append(primes, n)
			primesF = append(primesF, f)
		}
	}
}

// gcd 最大公約数
func gcd(p, q int64) int64 {
	for q != 0 {
		r := p % q
		p = q
		q = r
	}
	return p
}
