package main

import (
  "fmt"
  "github.com/jraymakers/euler-go/eulerlib"
)

func main() {
  n := 20
  ints := eulerlib.IntsUpTo(n)
  fmt.Printf("%v\n", ints)
  primes := eulerlib.PrimesUpTo(n, []int{})
  fmt.Printf("%v\n", primes)

  // for _, x := range ints {
  //   fmt.Printf("%v: %v\n", x, eulerlib.PrimeFactors(x, primes))
  // }
  combinedPfcs := []int{}
  for _, x := range ints {
    pfcs := eulerlib.PrimeFactorCounts(x, primes)
    fmt.Printf("%v: %v\n", x, pfcs)
    combinedPfcs = eulerlib.CombineFactorCounts(pfcs, combinedPfcs)
  }
  fmt.Printf("combined: %v\n", combinedPfcs)

  prod := 1
  for i, pfc := range combinedPfcs {
    prod *= eulerlib.Pow(primes[i], pfc)
  }
  fmt.Printf("%v\n", prod)
}
