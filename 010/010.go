package main
import ("fmt")

func main() {
  max := 2000000;
  var primes [2000001]int;
  ans := 0;

  for i := 2; i<max+1; i+=1 {
    if primes[i] == 0 {
      for j := 2; i*j<max+1; j++ {
        primes[i*j] = 1;
      }
    }
  }

  for i := 2; i < len(primes); i++ {
    if primes[i] == 0 {
      ans += i;
    }
  }
  fmt.Printf("%v\n", ans);
}
