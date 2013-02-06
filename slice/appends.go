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

  c := &b
  d := &b[1]

  fmt.Println("----------") 
  fmt.Println("&c=", &c)    //0xf840039120
  fmt.Println("c=", c)      //&[1 2]
  fmt.Println("&d=", &d)    //0xf840039128
  fmt.Println("d=", d)      //0xf840028234
  fmt.Println("----------") 
  
  b = append(b, 4, 5, 6) // create new slice

  printSlice("b", b)
  fmt.Println("c=", c)
  printSlice("*c", *c)
}

func printSlice(s string, x *T) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
