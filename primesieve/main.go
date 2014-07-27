package main

import "fmt"
import "flag"

func generate(ch chan<- int) {
  for i := 2; ;i++ {
    ch <- i
  }
}

func filter(in <-chan int, out chan<- int, prime int)  {
  for {
    i := <- in
    if i % prime != 0 {
      out <- i
    }
  }
}

func main()  {
  count := flag.Int("n", 100, "Number of primes to generate")
  flag.Parse()

  ch := make(chan int)

  go generate(ch)

  for j:= 0; j < *count; j++ {
    prime := <-ch
    fmt.Println(prime)
    ch1 := make(chan int)
    go filter(ch, ch1, prime)
    ch = ch1
  }
}
