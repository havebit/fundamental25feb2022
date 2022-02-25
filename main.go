package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacciGo(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func fibonacciGo(ch chan int) {
	a, b := 0, 1

	for {
		ch <- a
		a, b = b, a+b
	}
}

func mainOgGo() {
	start := time.Now()

	wg.Add(3)
	go slowPrint("one")
	go slowPrint("two")
	go slowPrint("three")

	wg.Wait()
	fmt.Println(time.Since(start))
}

var wg sync.WaitGroup

func slowPrint(s string) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println(s)
	wg.Done()
}

func mainOfClosureFibo() {
	f := newFibonacciFunc()

	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}

func newFibonacciFunc() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func fibonacci(n int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func mainOfClosure() {
	fn := closure()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func closure() func() int {
	var c counter
	return func() int {
		defer c.inc()
		return int(c)
	}
}

type counter int

func (c *counter) inc() {
	*c++
}

type greetingFunc func(...string) string

func newGreetingFunc() greetingFunc {
	return greeting
}

func printGreeting(fn greetingFunc) {
	fmt.Println(fn("yod", "new"))
}

func mainOfStringer() {
	var i Integer = 10

	fmt.Println(i)
}

type Integer int

func (s Integer) String() string {
	return "string"
}

func mainOfSwitch() {
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("%T %d", v, v)
	case string:
		fmt.Printf("%T %s", v, v)
	default:
		fmt.Println("undefined type")
	}
}

func mainOfInterface() {
	var i interface{}

	n := 10

	i = n

	fmt.Printf("%T %v\n", i, i)

	s := "text"

	i = s
	fmt.Printf("i %T %v\n", i, i)

	s = i.(string)
	fmt.Printf("s %T %v\n", s, s)

	if n, ok := i.(string); ok {
		fmt.Printf("n %T %v\n", n, n)
	}
}

func mainOfDefer() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	if _, err := strconv.Atoi("1a"); err != nil {
		log.Panic(err)
	}

	fmt.Println("end")
}

func greeting(names ...string) string {
	fmt.Printf("%T\n", names)
	if len(names) < 1 {
		return "Hello Guy"
	}

	return fmt.Sprintf("Hello %s", names[0])
}

func mainOfMethod() {
	var i Int = 15
	fmt.Printf("%q\n", i.String())

	i.Set(16)
	fmt.Printf("%q\n", i.String())
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

func mainOfOscar() {
	f, err := os.Open("./oscar.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameCount := map[string]int{}
	for _, record := range records {
		nameCount[record[3]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name)
		}
	}
}

func main1() {
	data, err := os.ReadFile("./oscar.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

func mainOfMap() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
		"d": "durian",
		"e": "eggplant",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func couple(ss string) (r []string) {
	for s := []rune(ss + "*"); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return
}

func mainOfSlice() {
	var slice []int

	nums := [...]int{2, 4, 6, 8, 10}

	slice = nums[1:3]

	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	newSlice := append(slice, 14, 16, 18)

	nums[2] = 22
	fmt.Println(newSlice)
	fmt.Println(slice)
	fmt.Println(nums)
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
