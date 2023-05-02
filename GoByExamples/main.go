package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("This is a test")
	fmt.Println("The sum of 2 and 3 is", sum(2, 3))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	resetInt := intSeq()
	fmt.Println(resetInt())
	fmt.Println(intSeq()())
}

func sum(a, b int) int {
	return a + b
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
