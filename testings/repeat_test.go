package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	got := Repeat('s')
	want := "ssssssssss"

	got01 := Repeat('a')
	want01 := "aaaaaaaaaa"

	if string(got) != want {
		t.Errorf("got %q want %q\n", got, want)
	}

	if string(got01) != want01 {
		t.Errorf("got %q want %q\n", got01, want01)
	}
}
