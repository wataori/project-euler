package main
import (
  "fmt"
  "math"
)

func main()  {
  fmt.Printf("%v\n", solver(1))
}

func solver(n int) int {
  triangul := n*(n+1)/2
  count := count_divisors(2, triangul, 2)
  if count > 500 {
    return triangul
  } else {
    return solver(n+1)
  }
}

func count_divisors(n int, max int, count int) int {
  if float64(n) >= math.Sqrt(float64(max)) {
    return count
  } else if max == 1 {
    return 1
  } else if max % n == 0 {
    return count_divisors(n+1, max, count+2)
  } else {
    return count_divisors(n+1, max, count)
  }
}
