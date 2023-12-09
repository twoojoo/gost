package gost

import (
	"errors"
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
