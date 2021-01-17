package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

// 有点像AOP
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn1(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarPara(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear")
	}()
	t.Log("go")
}
