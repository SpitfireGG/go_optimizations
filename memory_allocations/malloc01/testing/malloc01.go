package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	t *[5]uint64
	s []byte
)

func f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = &[5]uint64{}
	}
}

func g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]byte, 32768)
	}
}

func AllocViewF() string {
	bf := testing.Benchmark(f)
	strb := strings.Builder{}
	strb.WriteString(fmt.Sprintf("%d\n", bf.AllocedBytesPerOp()))
	strbPrint := strb.String()

	return strbPrint

}

func AllocViewG() string {
	bg := testing.Benchmark(g)
	strb := strings.Builder{}
	strb.WriteString(fmt.Sprintf("%d\n", bg.AllocedBytesPerOp()))
	strbPrint := strb.String()

	return strbPrint

}
