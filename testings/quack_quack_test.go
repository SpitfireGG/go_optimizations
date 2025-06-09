package main

import (
	"fmt"
	"testing"
)

func TestMakeItQuack(t *testing.T) {

	human := Person{name: "adam", sound: ""}

	got := makeItQuack(human)
	expected := "the human says quack"

	if got != expected {
		t.Errorf("got %s but expected %s", got, expected)
	} else {
		fmt.Println("humans dont quack")
	}

}
