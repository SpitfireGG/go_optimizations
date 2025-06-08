package main

import (
	"fmt"
	"testing"
)

func getData() [][]int {
	return [][]int{
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6, 7},
	}
}

func mergeWithOneLoop(data ...[]int) []int {

	var d []int
	for _, dt := range data {
		d = append(d, dt...)
	}
	return d
}

func mergeWithTwoLoops(data ...[]int) []int {
	var n int = 0
	for _, cc := range data {
		n += len(cc)
	}

	var d []int = make([]int, 0, n)
	for _, c := range data {
		d = append(d, c...)
	}
	return d
}

func BenchmarkOneLoop(b *testing.B) {
	data := getData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeWithOneLoop(data...)
	}
}

func BenchmarkTwoLoop(b *testing.B) {
	data := getData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mergeWithTwoLoops(data...)
	}
}

func main() {

	brOL := testing.Benchmark(BenchmarkOneLoop)
	brTL := testing.Benchmark(BenchmarkTwoLoop)

	fmt.Println(brOL.MemString()) // 736 B/op          5 allocs/op
	fmt.Println(brTL.MemString()) // 224 B/op          1 allocs/op
}
