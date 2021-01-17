package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "1"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))
	// string_test
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, s := range parts {
		t.Log(s)
	}
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i)
	}
}
