package main
import ("fmt")

func main() {
  length := 0
  var tmp int
  var ans int

  for i := 1; i < 1000000; i += 2 {
    tmp = solver(i, 0)
    if tmp > length {
      length = tmp
      ans = i
    }
  }
  fmt.Printf("%v, %v\n", length, ans)
}

func solver(n int, len int) int {
  var tmp int
  if n == 1 {
    return len+1
  }
  if n%2 == 0 {
    tmp = n/2
  } else {
    tmp = 3*n+1
  }
  return solver(tmp, len+1)
}
