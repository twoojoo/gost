package gost

import (
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
