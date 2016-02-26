package main
import "fmt"
import "math"

func main() {
  locate := 1000000
  ans := 0
  digits := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  idxes := []int {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
  // fmt.Println(append(digits[:2], digits[3:]...))
  for idx := 1; idx <= 10; idx++ {
    fmt.Println(locate)
    fmt.Println(idx)
    i := locate%idx
    locate = locate/idx
    idxes[10-idx] = i
    fmt.Println(i)
    fmt.Println(digits)
    fmt.Println(idxes)
    fmt.Println(digits[i])
    // fmt.Println(int(math.Pow(10, float64(10-idx))))

    // ans += digits[i] * int(math.Pow(10, float64(10-idx)))
    // digits = append(digits[:i], digits[i+1:]...)
    fmt.Println("-----------------------------")
  }
fmt.Println("-----------------------------")
fmt.Println(idxes)
fmt.Println("-----------------------------")
  for idx := 0; idx < 10; idx++ {
    fmt.Println(idx)
    fmt.Println(idxes[idx])
    fmt.Println(digits)
    fmt.Println(digits[idxes[idx]])
    ans += digits[idxes[idx]] * int(math.Pow(10, float64(10-idx)))
    if idx == 0 {
      digits = digits[idx:]
    } else {
      digits = append(digits[:idxes[idx-1]], digits[idxes[idx]:]...)
    }
    fmt.Println(digits)
    fmt.Println(ans)
  }
  fmt.Println(ans)
  fmt.Println(digits[1] * int(math.Pow(10, 2)))

}
