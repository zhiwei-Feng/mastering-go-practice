package main

import (
	"testing"
)

/*
goos: linux
goarch: amd64
Benchmark30fibo1-12          212           5358539 ns/op               0 B/op          0 allocs/op
Benchmark30fibo2-12          229           5344920 ns/op               0 B/op          0 allocs/op
Benchmark30fibo3-12       356478              3461 ns/op            2236 B/op          6 allocs/op
Benchmark50fibo1-12            1        81971515000 ns/op              0 B/op          0 allocs/op
Benchmark50fibo2-12            1        77388359500 ns/op              0 B/op          0 allocs/op
Benchmark50fibo3-12       235618              5053 ns/op            2481 B/op         10 allocs/op
PASS
ok      command-line-arguments  165.342s
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
