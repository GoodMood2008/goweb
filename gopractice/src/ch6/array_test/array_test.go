package array_test

import "testing"

func TestArrayInit(t *testing.T) {
  var arr [3]int
  t.Log(arr[1], arr[2])
  arr1 := [4]int{1,2,3,4}
  arr2 := [...]int{1,2,3,4}

  for i, e := range arr1 {
    t.Log(i, e)
  }

  for i := 0; i < len(arr2); i++ {
    t.Log(arr2[i])
  }
}

func TestArraySection(t *testing.T) {
  arr3 := [...]int{1,2,3,4}
  arr3_sec := arr3[:3]
  t.Log(arr3_sec)
}
