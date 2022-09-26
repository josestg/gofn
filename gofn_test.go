package gofn

import (
	"fmt"
	"strings"
)

func ExampleReduce() {
	sum := Reduce(0, []int{1, 2, 3, 4, 5}, func(acc int, v int) int {
		return acc + v
	})

	concat := Reduce("", []string{"a", "b", "c"}, func(acc string, v string) string {
		return acc + v
	})

	// this is equivalent to Map([]int{1, 2, 3}, func(v int) int { return v * 2 })
	twice := Reduce(make([]int, 0), []int{1, 2, 3}, func(acc []int, v int) []int {
		return append(acc, 2*v)
	})

	fmt.Println(sum)
	fmt.Println(concat)
	fmt.Println(twice)

	// Output:
	// 15
	// abc
	// [2 4 6]
}

func ExampleMap() {
	twice := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})

	uppercase := Map([]string{"a", "b", "c"}, strings.ToUpper)

	fmt.Println(twice)
	fmt.Println(uppercase)

	// Output:
	// [2 4 6]
	// [A B C]
}

func ExampleFilter() {
	odd := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 1
	})

	fmt.Println(odd)
	// Output:
	// [1 3 5]
}

func ExampleApplyOptions() {
	type Config struct {
		A int
		B string
	}

	withA := func(a int) Option[*Config] {
		return func(c *Config) {
			c.A = a
		}
	}

	withB := func(b string) Option[*Config] {
		return func(c *Config) {
			c.B = b
		}
	}

	var c Config
	ApplyOptions(&c, withA(1), withB("b"))

	fmt.Printf("%+v\n", c)
	// Output:
	// {A:1 B:b}
}

func ExampleDecorate() {
	type fn func(int int)

	trace := make([]int, 0)

	df := func(traceID int) func(fn) fn {
		return func(f fn) fn {
			return func(x int) {
				trace = append(trace, traceID)
				defer func() {
					trace = append(trace, traceID)
				}()
				f(x)
			}
		}
	}

	f := func(x int) { trace = append(trace, 0) } // the center.
	g := df(1)
	h := df(2)
	i := df(3)

	f = Decorate[fn](f, g, h, i)

	f(0)
	fmt.Println(trace)

	// Output:
	// [3 2 1 0 1 2 3]

}

func ExampleReverse() {
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5}))
	// Output:
	// [5 4 3 2 1]
}

func ExampleReversedDecorate() {
	type fn func(int int)

	trace := make([]int, 0)

	df := func(traceID int) func(fn) fn {
		return func(f fn) fn {
			return func(x int) {
				trace = append(trace, traceID)
				defer func() {
					trace = append(trace, traceID)
				}()
				f(x)
			}
		}
	}

	f := func(x int) { trace = append(trace, 0) } // the center.
	g := df(1)
	h := df(2)
	i := df(3)

	f = ReversedDecorate[fn](f, g, h, i)

	f(0)
	fmt.Println(trace)

	// Output:
	// [1 2 3 0 3 2 1]
}
