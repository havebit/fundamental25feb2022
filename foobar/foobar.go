package foobar

import (
	"fmt"
	"strconv"
)

type Intner interface {
	Intn(n int) int
}

type IntnFunc func(int) int

func (f IntnFunc) Intn(n int) int {
	return f(n)
}

func RandomSay(r Intner) string {
	// src := rand.NewSource(time.Now().Unix())
	// r := rand.New(src)

	return fmt.Sprintf("%s-%s-%s-%s", Say(r.Intn(9)+1), Say(r.Intn(9)+1), Say(r.Intn(9)+1), Say(r.Intn(9)+1))
}

func SayAny(i interface{}) string {
	return "Foo"
}

func Say(n int) string {
	if n == 5 {
		return "Bar"
	}

	if n == 3 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
