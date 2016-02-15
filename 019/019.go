package main
import "fmt"
import "time"

func main() {
  t := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
  end := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC)
  ans := 0

  for end.After(t) {
    if t.Format("Mon") == "Sun" {
      ans++
    }
    t = t.AddDate(0, 1, 0)
  }

  fmt.Println(ans)
}
