package main

import (
	"fmt"
)

type Quacker interface {
	Quack() string
}

type Person struct {
	name  string
	sound string
}

func Quack() string {
	return "quack"
}

func (p Person) Quack() string {
	return fmt.Sprintf("the human says %s", Quack())
}

func makeItQuack(q Quacker) string {
	return q.Quack()
}
