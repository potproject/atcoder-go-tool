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
	base := 12344567890
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getDigit(base)
	}
}
