package main

import "fmt"

func addOne(x int) int {
	return x + 1
}

func main() {
	x := 5
	x = addOne(x)
	fmt.Println(x)
}
