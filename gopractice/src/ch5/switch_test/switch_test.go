package switch_test

import "testing"

// switch not limit in constant and int
func TestSwitch(t *testing.T) {
  for i := range []int{1,2,3,4,5} {
    switch i {
    case 0, 2:
      t.Log("Event")
    case 1, 3:
      t.Log("Odd")
    default:
      t.Log("it is not 0--3")
    }
  }
}
