package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["mystr"] = 12

  val, res := m["noval"]

  fmt.Println("val: ", val)
  fmt.Println("res: ", res)

  val, res = m["mystr"]

  fmt.Println("val: ", val)
  fmt.Println("res: ", res)

}
