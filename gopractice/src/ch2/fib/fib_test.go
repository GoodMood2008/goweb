package fib

import "testing"
import "fmt"



func TestFibList(t *testing.T) {
  // var a int = 1
  // var b int = 1
  a, b := 1, 1
  fmt.Print(a)
  for i := 0; i < 5; i++ {
    fmt.Print(" ", b)
    tmp := a
    a = b
    b = tmp + a
  }
  fmt.Println()
}


func TestExchange(t *testing.T) {
  a, b := 4, 5
  b, a = a, b
  t.Log(a, b)
}
