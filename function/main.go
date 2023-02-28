package main

import "fmt"

func main() {
	s := [](func(x int) func() int{
		return func() int{
			return x
		})
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d %p\n", i, &i)
	}
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d %p\n", i, &i)
	}
}
