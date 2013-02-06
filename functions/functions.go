package main

import "fmt"

func main() {
  
  MyPrint("Hello")
  MyPrint("Hello", "My", "Dear", "Friend")
}

func MyPrint(first string, tail ...interface{}) {
  fmt.Println(first)
  fmt.Println(tail...)
}
