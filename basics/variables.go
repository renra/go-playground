package main

import "fmt"

var foo string
var is_used bool = false
const factor = 1

func main() {
  var foo int
  pi := 3.14 //:= for implicit type declaration, only inside functions

  fmt.Println(foo)
  fmt.Println(is_used)
  fmt.Println(pi)

  // conversions / casts
  fmt.Println(float32(1))

  fmt.Printf("Variable is_used is of type: %T", is_used)
  fmt.Println()

  fmt.Printf("Variable factor is of type: %T", factor)
  fmt.Println()
}
