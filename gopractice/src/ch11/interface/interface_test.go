package interface_test

import "testing"

type Programmar interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Pringln(\"hello world\")"
}

func TestClient(t *testing.T) {
	var p Programmar
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
