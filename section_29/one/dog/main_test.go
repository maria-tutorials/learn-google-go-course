package dog

import (
	"fmt"
	"testing"
)

var expected int = 70

func TestYears(t *testing.T) {
	y := Years(10)
	if y != expected {
		t.Error("Got", y, "expected", expected)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(10)
	if y != 70 {
		t.Error("Got", y, "expected", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
