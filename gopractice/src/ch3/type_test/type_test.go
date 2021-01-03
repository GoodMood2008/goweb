package type_test

import "testing"

// type预定义
//math.MaxInt64
//math.MaxFloat64
//math.MaxUnit32

type MyInt int64
func TestImplicit(t *testing.T) {
  var a int = 1
  var b int64
  // b = a not supportimplicit transformation
  // c = a not support implicit transformaion
  b = int64(a)

  //var c MyInt
  // c = MyInt(a)
  t.Log(a, b)
}

// not support pointer arithmetics
func Testpoint(t *testing.T) {
  a := 1
  aPtr := &a
  //aPtr = aPtr + 1
  t.Log(a, aPtr)
  t.Logf("%T %T", a, aPtr)
}


func TestString(t *testing.T) {
  var a string
  t.Log(len(a))
}
