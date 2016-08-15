package main

import (
  "fmt"
  "github.com/jraymakers/euler-go/eulerlib"
)

func main() {
  n := 2*1000*1000
  primes := eulerlib.PrimesUpTo(n, []int{})
  sum := eulerlib.Sum(primes)
  fmt.Printf("%v\n", sum)
}
