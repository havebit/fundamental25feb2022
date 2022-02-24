package main

import (
	"fmt"

	"gitlab.com/cjexpress/tildi/signac/learngo/customer"
)

func main() {
	fmt.Println(customer.Title)
}

func mainOfPonter() {
	var p *int
	fmt.Println(p)

	fmt.Println(p == nil)

	i := 42
	p = &i

	fmt.Println(p, &i)

	fmt.Println(*p, i)

	*p = 43
	fmt.Println(*p, i)

	i = 44
	fmt.Println(*p, i)
}

func prime(n int) {
	for i := 2; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}
}

func IsCorrect() bool {
	return true
}

func swap(a, b int) (int, int) {
	return b, a
}

func add(a, b int) int {
	return a + b
}
