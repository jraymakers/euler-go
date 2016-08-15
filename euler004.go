package main

import (
  "fmt"
  "github.com/jraymakers/euler-go/eulerlib"
)

func largestPalindromeProduct(max, min int) int {
  maxP, p := 0, 0
  for a := max; a >= min; a-- {
    for b := max; b >= a; b-- {
      p = a * b
      if eulerlib.Palindrome(p) && p > maxP {
        maxP = p
      }
    }
  }
  return maxP
}

func main() {
  fmt.Printf("%d\n", largestPalindromeProduct(999, 900))
}
