package gost

import (
	"slices"
	"testing"
)

func TestOption(t *testing.T) {
	if opt := Some("value"); opt.IsSome() {
		if val := opt.Unwrap(); val != "value" {
			t.Fatal(val)
		}
	} else {
		t.Fail()
	}

	if opt := None[string](); !opt.IsNone() {
		t.Fail()
	}

	opt := None[string]()

	if val := opt.UnwrapOr("alt"); val != "alt" {
		t.Fatal(val)
	}

	if val := opt.UnwrapOrElse(func() string { return "alt" }); val != "alt" {
		t.Fatal(val)
	}

	if val := opt.UnwrapOrZero(); val != "" {
		t.Fatal(val)
	}
}

func TestAsOption(t *testing.T) {
	val := AsOption(slices.BinarySearch([]int{1, 2, 3, 4, 5}, 4)).UnwrapOr(0)

	if val != 3 {
		t.Fatal(val)
	}
}

func TestOptionMatching(t *testing.T) {
	val := AsOption(slices.BinarySearch([]int{1, 2, 3, 4, 5}, 4)).
		OnSome(func(v *int) *int { return v }).
		OnNone(func() int { return 0 })

	if val != 3 {
		t.Fatal(val)
	}
}
