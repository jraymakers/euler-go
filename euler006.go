package main

import (
  "fmt"
  "github.com/jraymakers/euler-go/eulerlib"
)

func main() {
  n := 100
  ints := eulerlib.IntsUpTo(n)
  fmt.Printf("ints: %v\n", ints)
  squares := eulerlib.Squares(ints)
  fmt.Printf("squares: %v\n", squares)
  sumOfSquares := eulerlib.Sum(squares)
  fmt.Printf("sumOfSquares: %v\n", sumOfSquares)
  sum := eulerlib.Sum(ints)
  fmt.Printf("sum: %v\n", sum)
  squareOfSum := sum * sum
  fmt.Printf("squareOfSum: %v\n", squareOfSum)
  difference := squareOfSum - sumOfSquares
  fmt.Printf("difference: %v\n", difference)
}
