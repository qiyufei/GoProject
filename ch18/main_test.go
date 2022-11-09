package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	n := 10
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

func TestFibonacci(t *testing.T) {
	m := map[int]int{}
	m[-1] = 0
	m[0] = 0
	m[1] = 1
	m[2] = 1
	m[3] = 2
	m[4] = 3
	m[5] = 5
	m[6] = 8
	m[7] = 13
	m[8] = 21
	m[9] = 34
	m[10] = 55
	for k, v := range m {
		re := Fibonacci(k)
		if re == v {
			t.Logf("正确：n为%d, 值为%d", k, re)
		} else {
			t.Logf("错误：预期为%d,但是计算结果为%d", v, re)
		}
	}
}
