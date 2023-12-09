package gost

import (
	"log"
	"os"
)

type Result[T any] struct {
	value *T
	error error
}

func AsResult[T any](v T, e error) Result[T] {
	if e != nil {
		return Result[T]{value: &v}
	}

	return Result[T]{error: e}
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: &value}
}

func Error[T any](error error) Result[T] {
	return Result[T]{error: error}
}

func (o Result[T]) IsError() bool {
	return o.error != nil
}

func (o Result[T]) IsOk() bool {
	return o.value != nil
}

// Extracts the value from the Result. Panics if the Result is Error.
func (o Result[T]) Unwrap() T {
	if o.value == nil {
		panic("can't unwrap a nil pointer")
	}

	return *o.value
}

// Extracts the error from the Result. Panics if the Result is Ok.
func (o Result[T]) UnwrapError() error {
	if o.error == nil {
		panic("can't unwrap a nil error")
	}

	return o.error
}

// Extracts the value from the Result. Returns the provided value if the Result is Error.
func (o Result[T]) UnwrapOr(def T) T {
	if o.value == nil {
		return def
	}

	return *o.value
}

// Extracts the value from the Result. Returns the value returned by the provided function if the Result is Error.
func (o Result[T]) UnwrapOrElse(def func() T) T {
	if o.value == nil {
		return def()
	}

	return *o.value
}

// Extracts the value from the Result. Returns the zero value of the wrapped type if the Result is Error.
func (o Result[T]) UnwrapOrZero() T {
	if o.value == nil {
		var zero T
		return zero
	}

	return *o.value
}

// Extracts the value from the Result. Panics with the provided argument if the Result is Error.
func (o Result[T]) UnwrapOrPanic() T {
	if o.value != nil {
		return *o.value
	}

	panic(o.error)
}

// Extracts the value from the Result. Calls log.Fatal(error) if the Result is Error.
func (o Result[T]) UnwrapOrLogFatal() T {
	if o.value != nil {
		return *o.value
	}

	log.Fatal(o.error)

	var zero T
	return zero
}

// Extracts the value from the Result. Calls os.Exit(code) if the Result is Error.
func (o Result[T]) UnwrapOrExit(code int) T {
	if o.value != nil {
		return *o.value
	}

	os.Exit(code)
	var zero T
	return zero
}

// Extracts the value from the Result. Calls os.Exit([code]) if the Result is Error, getting the exit code from the provided function.
func (o Result[T]) UnwrapOrDynamicExit(code func(err error) int) T {
	if o.value != nil {
		return *o.value
	}

	os.Exit(code(o.error))
	var zero T
	return zero
}
