package word

import (
	"fmt"
	"testing"

	"../quote"
)

func TestCount(t *testing.T) {
	count := Count("zero um dois")
	if count != 3 {
		t.Error("Got", count, "expected", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("zero zero zero um dois dois")

	for k, v := range m {
		switch k {
		case "zero":
			if v != 3 {
				t.Error("Got", v, "expected", 3)
			}
		case "um":
			if v != 1 {
				t.Error("Got", v, "expected", 1)
			}
		case "dois":
			if v != 2 {
				t.Error("Got", v, "expected", 2)
			}
		}
	}

}

func ExampleCount() {
	fmt.Println(Count("zero um dois"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
