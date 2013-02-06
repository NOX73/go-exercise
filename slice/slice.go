package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[0] = 1
	a[2] = 2
	a[4] = 4
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	b = b[:1]
	b[0] = 1
	printSlice("b", b)
	c := b[:2]
	c[1] = 5
	c[0] = 2
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	printSlice("d", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
