package main

import (
	"fmt"
	"testing"
	"unsafe"
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

func main() {
	fmt.Println(unsafe.Sizeof(*t)) // uint64 -> typically 8 bytes on 64-bits ( 8 * 5 )bytes
	bf := testing.Benchmark(f)
	fmt.Println(bf.AllocedBytesPerOp())
	fmt.Println(bf.MemString()) // 48 B/op          1 allocs/op
	fmt.Println()
	bg := testing.Benchmark(g)
	fmt.Println(bg.AllocedBytesPerOp())
	fmt.Println(bg.MemString()) // 32768 B/op          1 allocs/op

}
