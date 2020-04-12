package main

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	e1 := 123
	a1 := abs(123)
	if a1 != e1 {
		t.Errorf("invalid a1:%v e1:%v", a1, e1)
	}
	e2 := 123
	a2 := abs(-123)
	if a2 != e2 {
		t.Errorf("invalid a2:%v e2:%v", a2, e2)
	}
}

func TestSliceInt(t *testing.T) {
	i := 1234567890
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a := sliceInt(i)
	if !reflect.DeepEqual(a, e) {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}

func TestSliceIntSum(t *testing.T) {
	e := 45
	a := sliceIntSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	if a != e {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	e := reverse([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	if !reflect.DeepEqual(a, e) {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}
func TestGetDigit(t *testing.T) {
	e := 10
	a := getDigit(1234567890)
	if a != e {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}

func TestGetSumDigit(t *testing.T) {
	e := 45
	a := getSumDigit(1234567890)
	if a != e {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}

func TestLastPrime(t *testing.T) {
	e := int64(97)
	a := lastPrime(25)
	if a != e {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
	e2 := int64(2)
	a2 := lastPrime(1)
	if a2 != e2 {
		t.Errorf("invalid a2:%v e2:%v", a2, e2)
	}
}

func TestGcd(t *testing.T) {
	e := int64(4)
	a := gcd(20, 32)
	if a != e {
		t.Errorf("invalid a:%v e:%v", a, e)
	}
}
