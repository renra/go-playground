package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

// returning multiple values
func swap(x, y string) (string, string) {
  return y, x
}

func split(sum int) (x, y int) {
  x = sum * 10 / 18
  y = sum - x

  //naked return
  return
}

func main() {
  fmt.Println(add(10, 5))
  fmt.Println(swap("World", "Hello"))
  fmt.Println(split(29))
}
