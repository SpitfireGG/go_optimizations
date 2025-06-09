package main

import "fmt"

const N = 10

func Repeat(str rune) []rune {

	mkstr := make([]rune, 0, N)

	func() {

		for i := 0; i < N; i++ {
			mkstr = append(mkstr, str)
		}
	}()
	return mkstr
}

func main() {

	res := Repeat('s')
	fmt.Println(string(res))

}
