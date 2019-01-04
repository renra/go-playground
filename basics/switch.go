package main

import (
  "fmt"
  "runtime"
)

func main() {
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OSX huh?")
  case "linux":
    fmt.Println("That's the right choice")
  default:
    fmt.Println(os)
  }
}
