package main

import (
  "fmt"
  "github.com/jraymakers/euler-go/eulerlib"
)

func largestPrimeFactor(x int) int {
  p := 2
  primes := []int{2}
  primeFactors := []int{}
  for x > 1 {
    if (x % p == 0) {
      if len(primeFactors) == 0 || primeFactors[len(primeFactors)-1] != p {
        primeFactors = append(primeFactors, p)
      }
      x = x / p
    } else {
      p++
      for eulerlib.CanDivide(p, primes) {
        p++
      }
      primes = append(primes, p)
    }
  }
  return primeFactors[len(primeFactors)-1]
}

func main() {
  fmt.Printf("%d\n", largestPrimeFactor(600851475143))
}
