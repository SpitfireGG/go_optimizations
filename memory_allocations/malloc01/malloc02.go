package main

import (
	"fmt"
	"testing"
)

var (
	bs = []byte{32: 'b'} // 33 bytes
	st string
)

func concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st = string(bs) + string(bs)
		// processing:  33 +  33 + 66 total still compiler and environment dependent
	}
}

func main() {
	br := testing.Benchmark(concat)
	fmt.Println(br.AllocedBytesPerOp())
	fmt.Println(br.AllocsPerOp())
	fmt.Println(br.MemString())
}

