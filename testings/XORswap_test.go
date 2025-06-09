package main

import "testing"

func TestBinXorSwap(t *testing.T) {
	got := BinXorSwap(10, 20)
	expected := "a = 10, b = 20"

	if expected != got {
		t.Errorf("got %s but expected %s\n", got, expected)
	}
}
