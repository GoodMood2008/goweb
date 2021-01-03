package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

// 判定某个值是否存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	if _, ok := m1[1]; ok {
		t.Log("1 is exist")
	} else {
		t.Log("1 is not exist")
	}
}

// 访问整个map
func TestTraveMap(t *testing.T) {
	m1 := map[int]int{4: 16, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
