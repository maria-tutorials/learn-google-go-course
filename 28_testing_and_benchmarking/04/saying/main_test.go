package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Buenos dias Chico"
	s := Greet("Chico")
	if s != expected {
		t.Error("got", s, "expected", expected)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Chica"))
	// Output:
	// Buenos dias Chica
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Perro")
	}
}
