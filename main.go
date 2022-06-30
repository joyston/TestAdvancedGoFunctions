package main

import (
	"fmt"
	"math"
)

func dummy() {
	fmt.Println("practicing advance function")
}

func compute(fn func(x, y float64) float64) float64 {
	return fn(2, 2)
}

func returnfunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	f := dummy
	fmt.Println("Time to practice Advance function concepts")
	f()

	test := func() {
		fmt.Println("I love Golang.")
	}
	test()

	newtest := func(x int) int {
		return x * -1
	}(8)
	fmt.Println("Value is", newtest)

	advtest := func(x int) int {
		return x * -1
	}
	fmt.Println("Changin Value", advtest(9))

	hypt := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hypotinuse:", hypt(2, 3))
	fmt.Println("Computed Hypotinuse of 2,2:", compute(hypt))
	fmt.Println("Computed pow of 2,2:", compute(math.Pow))

	returnfunc("Return function concept")()
	t := returnfunc("Wowzers")
	t()
}
