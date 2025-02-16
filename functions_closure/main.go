package main

import "fmt"

func do() func() int {
	a, b := 0, 1
	return func() int {
		y := a
		a, b = b, a+b
		return y
	}
}

func main() {
	f := do()
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", f())
	}

	fmt.Printf("\n")
	g := do()

	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", g())
	}

	fmt.Println("")
}
