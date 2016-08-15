package eulerlib

import (
  // "fmt"
  "strconv"
)

func CanDivide(x int, ys []int) bool {
  for _, y := range ys {
    if x % y == 0 {
      return true
    }
  }
  return false
}

func IntsUpTo(n int) []int {
  ints := []int{}
  for x := 1; x <= n; x++ {
    ints = append(ints, x)
  }
  return ints
}

func Max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func Palindrome(x int) bool {
  runes := []rune(strconv.Itoa(x))
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    if runes[i] != runes[j] {
      return false
    }
  }
  return true
}

func Pow(b, e int) int {
  prod := 1
  for i := 0; i < e; i++ {
    prod *= b
  }
  return prod
}

func Product(xs []int) int {
  prod := 1
  for _, x := range xs {
    prod *= x
  }
  return prod
}

func Squares(xs []int) []int {
  squares := make([]int, len(xs))
  for i, x := range xs {
    squares[i] = x * x
  }
  return squares
}

func Sum(xs []int) int {
  sum := 0
  for _, x := range xs {
    sum += x
  }
  return sum
}

// Primes

func LargestKnownPrime(primes []int) int {
  numPrimes := len(primes)
  if numPrimes > 0 {
    return primes[numPrimes-1]
  } else {
    return 2
  }
}

func AppendNextPrime(primes []int) []int {
  for x := LargestKnownPrime(primes); ; x++ {
    if !CanDivide(x, primes) {
      // fmt.Printf("new prime: %v\n", x)
      return append(primes, x);
    }
  }
}

func PrimesUpTo(n int, primes []int) []int {
  for x := 2; x <= n; x++ {
    if !CanDivide(x, primes) {
      // fmt.Printf("new prime: %v\n", x)
      primes = append(primes, x);
    }
  }
  return primes
}

func NthPrime(n int) int {
  primes := []int{}
  for len(primes) < n {
    primes = AppendNextPrime(primes)
  }
  return primes[n-1]
}

func PrimeFactors(n int, primes []int) []int {
  factors := []int{}
  i := 0
  p := primes[i]
  for n > 1 {
    if n % p == 0 {
      n = n / p
      factors = append(factors, p)
    } else {
      i = i+1
      if i >= len(primes) {
        primes = AppendNextPrime(primes)
      }
      p = primes[i]
    }
  }
  return factors
}

func PrimeFactorCounts(n int, primes []int) []int {
  factorCounts := []int{0}
  i := 0
  p := primes[i]
  for n > 1 {
    if n % p == 0 {
      n = n / p
      factorCounts[i]++
    } else {
      i = i+1
      factorCounts = append(factorCounts, 0)
      if i >= len(primes) {
        primes = AppendNextPrime(primes)
      }
      p = primes[i]
    }
  }
  return factorCounts
}

func CombineFactorCounts(xs []int, ys []int) []int {
  xslen := len(xs)
  yslen := len(ys)
  if xslen > yslen {
    xs, ys = ys, xs
    xslen, yslen = yslen, xslen
  }
  combined := make([]int, yslen)
  i := 0
  for ; i < xslen; i++ {
    combined[i] = Max(xs[i], ys[i])
  }
  for ; i < yslen; i++ {
    combined[i] = ys[i]
  }
  return combined
}
