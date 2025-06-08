package main

import "testing"

const N = 1000

type Book struct {
	Name  string
	IBSN  uint8
	Pages uint8
}

//go:noinline
func createBook1Block(n int) []*Book {

	books := make([]Book, n)
	ptr_books := make([]*Book, n)

	for i := range ptr_books {
		ptr_books[i] = &books[i]
	}
	return *(&ptr_books)
}

//go:noinline

func createSmallMemChunks(n int) []*Book {

	books := make([]*Book, n)
	for i := range books {
		books[i] = new(Book)
	}
	return books
}

func Benchmark1Block(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = createBook1Block(N)
	}
}

func BenchmarkSmallChunks(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = createSmallMemChunks(N)
	}
}

func main() {

	rs := testing.Benchmark(Benchmark1Block)
	println(rs.MemString()) // 32768 B/op          2 allocs/op

	bs := testing.Benchmark(BenchmarkSmallChunks)
	println(bs.MemString()) // 32192 B/op       1001 allocs/op
}
