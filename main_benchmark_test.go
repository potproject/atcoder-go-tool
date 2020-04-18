package main

import (
	"testing"
)

// ベンチマーク
// opは自分のPCの処理速度にも比例するためあくまで目安として
//

// 0.26 ns/op
func BenchmarkAbs(b *testing.B) {
	base := -123
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = abs(base)
	}
}

// 62.1 ns/op
func BenchmarkSliceInt(b *testing.B) {
	base := 1234567890
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sliceInt(base)
	}
}

// 5.07 ns/op
func BenchmarkSliceIntSum(b *testing.B) {
	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sliceIntSum(base)
	}
}

// 6.70 ns/op
func BenchmarkReverse(b *testing.B) {
	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = reverse(base)
	}
}

// 13.2 ns/op
func BenchmarkGetDigit(b *testing.B) {
	base := 1234567890
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getDigit(base)
	}
}

// 16.1 ns/op
func BenchmarkGetSumDigit(b *testing.B) {
	base := 1234567890
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getSumDigit(base)
	}
}

// 4683 ns/op
func BenchmarkTestPrime100(b *testing.B) {
	base := int64(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lastPrime(base)
	}
}

// 11.0 ns/op
func BenchmarkTestGcd(b *testing.B) {
	x := int64(240)
	y := int64(320)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gcd(x, y)
	}
}
