package main

import "fmt"

func main() {
  
  var a string
  fmt.Println("a=", a) //

  a = "Mystring"

  var b *string

  fmt.Println("a=", a) //Mystring
  fmt.Println("b=", b) //<nil>

  b = &a
  d = &a[0]

  fmt.Println("b=", b)//0xf840028220
  fmt.Println("*b=", *b)//Mystring
}

