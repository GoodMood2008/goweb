package operator_test

import "testing"

const (
  Readable = 1 << iota
  Writable
  Executable
)

func TestCompare(t *testing.T) {
  a := [...]int{1, 2, 3, 4}
  b := [...]int{1, 3, 2, 4}
  d := [...]int{1, 2, 3, 4}
  t.Log(a == b)
  t.Log(a == d)
}

// 按位清零
func TestBit(t *testing.T) {
  a := 7
  a = a &^ Readable
  t.Log(a)
}
