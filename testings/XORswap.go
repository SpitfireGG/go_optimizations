package main

import (
	"fmt"
	"strings"
)

func BinXorSwap(a int, b int) string {
	var temp int
	temp = a ^ b
	b = temp ^ a
	a = b ^ temp
	strb := strings.Builder{}
	strb.WriteString(fmt.Sprintf("a = %d, b = %d", a, b))
	pf := strb.String()
	return pf
}
