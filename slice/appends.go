package main

import "fmt"

func main() {
  fmt.Println("")

  a := [...]int{1, 2, 3}//it's array
  var b []int

  fmt.Println(a)

  printSlice("b", b)

  b = a[:2]
  printSlice("b", b)

  //b[3] = 4 //index out of range
  
  b = append(b, 4, 5, 6) // create new slice

  printSlice("b", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
