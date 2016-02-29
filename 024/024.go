package main
import "fmt"
import "math"

func main() {
  location := 1000000
  ans := 0
  digits := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  idxes := []int {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

  for idx := 1; idx <= 10; idx++ {
    i := location%idx
    location = location/idx
    idxes[10-idx] = i
  }
  for idx := 0; idx < 10; idx++ {
    ans += digits[idxes[idx]] * int(math.Pow(10, float64(9-idx)))
    digits = append(digits[:idxes[idx]], digits[idxes[idx]+1:]...)
  }
  fmt.Println(ans)
}
