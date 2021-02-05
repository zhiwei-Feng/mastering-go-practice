package main

import (
	"testing"
)

/*
goos: linux
goarch: amd64
Benchmark30fibo1-12          232           5395946 ns/op
Benchmark30fibo2-12          213           5398885 ns/op
Benchmark30fibo3-12       327279              3547 ns/op
Benchmark50fibo1-12            1        82874845600 ns/op
Benchmark50fibo2-12            1        80692751500 ns/op
Benchmark50fibo3-12       236806              4999 ns/op
PASS
ok      command-line-arguments  169.505s
*/

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}

func benchmarkfibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3(n)
	}
	result = r
}

func Benchmark30fibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}
func Benchmark30fibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}
func Benchmark30fibo3(b *testing.B) {
	benchmarkfibo3(b, 30)
}

func Benchmark50fibo1(b *testing.B) {
	benchmarkfibo1(b, 50)
}
func Benchmark50fibo2(b *testing.B) {
	benchmarkfibo2(b, 50)
}
func Benchmark50fibo3(b *testing.B) {
	benchmarkfibo3(b, 50)
}
