package foobar

import "strconv"

func Say(n int) string {
	if n == 5 {
		return "Bar"
	}
	if n == 4 {
		return strconv.Itoa(n)
	}
	if n == 3 {
		return "Foo"
	}
	if n == 2 {
		return strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
