package main

import (
  "fmt"
)

func main() {
  for a := 1; a <= 1000-6; a++ {
    for b := a+1; b <= 1000-3; b++ {
      c := 1000 - a - b;
      if (a*a + b*b == c*c) {
        fmt.Printf("%v^2 + %v^2 = %v^2, abc = %v\n", a, b, c, a*b*c)
      }
    }
  }
}
