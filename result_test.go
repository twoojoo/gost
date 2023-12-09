package gost

import (
	"errors"
	"os"
	"testing"
)

func TestResult(t *testing.T) {
	if res := Ok("value"); res.IsOk() {
		if val := res.Unwrap(); val != "value" {
			t.Fatal(val)
		}
	} else {
		t.Fail()
	}

	err := Error[string](errors.New("err"))
	if !err.IsError() {
		t.Fail()
	} else {
		if e := err.UnwrapError(); e.Error() != "err" {
			t.Fatal(e.Error())
		}
	}

	err = Error[string](errors.New("err"))
	if val := err.UnwrapOr("alt"); val != "alt" {
		t.Fatal(val)
	}

	if val := err.UnwrapOrElse(func() string { return "alt" }); val != "alt" {
		t.Fatal(val)
	}

	if val := err.UnwrapOrZero(); val != "" {
		t.Fatal(val)
	}
}

func TestAsResult(t *testing.T) {
	val := AsResult(os.ReadFile("non-existent-file.txt")).UnwrapOr([]byte{})

	if len(val) != 0 {
		t.Fatal(val)
	}
}
