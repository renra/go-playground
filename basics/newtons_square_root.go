package main

import "fmt"

func Sqrt(input float64) float64 {
  output := 1.0

  for i := 1; i < 10; i++ {
    new_output = output - (output * output - input) / (2 * output)
    fmt.Println(output)
  }

  return output
}

func main() {
  fmt.Println(Sqrt(4))
}
