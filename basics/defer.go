package main

import "fmt"

func main() {
  defer fmt.Println("Deferred message")
  fmt.Println("Undeferred message")
}
