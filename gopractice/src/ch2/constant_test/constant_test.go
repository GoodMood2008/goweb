package constant_test
import "testing"

const (
  Monday = iota + 1
  Tuesday
  Wednesday
  Thirsday
  Friday
  Sataday
  Sunday
)

const (
  Readable = 1 << iota
  Writable
  Executable
)

func TestConstTry(t *testing.T) {
  t.Log(Monday, Tuesday)
}


func TestConstTry1(t *testing.T) {
  a := 7
  t.Log(a & Readable == Readable)
}
