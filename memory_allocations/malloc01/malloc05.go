package main

import (
	"fmt"
	"testing"
)

func OriginalData() []int {
	slint := make([]int, 1024)
	for i := range slint {
		slint[i] = i
	}
	return slint
}

func checkEven(v int) bool {
	return v%2 == 0
}

func FilterOneAlloc(data []int) []int {

	var v = make([]int, 0, len(data)-2) // create an allocation that might trigger another resize allocation
	for _, rv := range data {
		if checkEven(rv) {
			v = append(v, rv)
		}
	}
	return v
}

// sorrrry for the wierd fucked up namespacing
func FilterNoAlloc(data []int) []int {
	var didx = 0
	for i, dataa := range data {
		if checkEven(dataa) {
			data[i] = data[didx]
			data[didx] = dataa
			didx++
		}
	}
	return data[:didx]
}

func BenchmarkOneAlloc(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterOneAlloc(OriginalData())
	}
}

func BenchmarkNoAlloc(b *testing.B) {
	b.ResetTimer()
	_ = FilterOneAlloc(OriginalData())
	for i := 0; i < b.N; i++ {
		_ = FilterNoAlloc(OriginalData())
	}
}

func main() {

	br := testing.Benchmark(BenchmarkOneAlloc)
	fmt.Println(br.MemString()) // 8192 B/op          1 allocs/op

	fmt.Println()
	sr := testing.Benchmark(BenchmarkNoAlloc)
	fmt.Println(sr.MemString()) // 0 B/op          0 allocs/op
	fmt.Println()

}
