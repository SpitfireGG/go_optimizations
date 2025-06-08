package main

import (
	"fmt"
	"strings"
	"testing"
)

func getData() [][]int {
	return [][]int{
		{1, 2},
		{9, 10, 11},
		{6, 2, 3, 7},
		{11, 5, 7, 12, 16},
		{8, 5, 6},
	}
}

var sb strings.Builder

const format = "Allocate from %v to %v (when append slice#%v).\n"

func MergeWithOneLoop(data [][]int) []int {

	var oldCap int
	var r []int
	for i, s := range data {
		r = append(r, s...)
		if oldCap == cap(r) {
			continue
		}
		sb.WriteString(fmt.Sprintf(format, oldCap, cap(r), i))
		oldCap = cap(r)
	}
	sb.String()
	return r
}

func MergeWithOneLoopBr(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeWithOneLoop(getData())
	}

}

func main() {
	rs := testing.Benchmark(MergeWithOneLoopBr)
	fmt.Println(rs.MemString())
	fmt.Println(rs.MemAllocs)
}

