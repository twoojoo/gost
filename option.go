package gost

import (
	"log"
	"os"
)

type Option[T any] struct {
	value *T
}

func AsOption[T any](v T, ok bool) Option[T] {
	if ok {
		return Option[T]{value: &v}
	}

	return Option[T]{}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Extracts the value from the Option. Panics if the Option is None.
func (o Option[T]) Unwrap() T {
	if o.value == nil {
		panic("can't unwrap a nil pointer")
	}

	return *o.value
}

// Extracts the value from the Option. Returns the provided value if the Option is None.
func (o Option[T]) UnwrapOr(def T) T {
	if o.value == nil {
		return def
	}

	return *o.value
}

// Extracts the value from the Option. Returns the value returned by the provided function if the Option is None.
func (o Option[T]) UnwrapOrElse(def func() T) T {
	if o.value == nil {
		return def()
	}

	return *o.value
}

// Extracts the value from the Option. Returns the zero value of the wrapped type if the Option is None.
func (o Option[T]) UnwrapOrZero() T {
	if o.value == nil {
		return Zero[T]()
	}

	return *o.value
}

// Extracts the value from the Option. Panics with the provided argument if the Option is None.
func (o Option[T]) UnwrapOrPanic(v any) T {
	if o.value != nil {
		return *o.value
	}

	panic(v)
}

// Extracts the value from the Option. Calls log.Fatal(v...) if the Option is None.
func (o Option[T]) UnwrapOrLogFatal(v ...any) T {
	if o.value != nil {
		return *o.value
	}

	log.Fatal(v...)
	return Zero[T]()
}

// Extracts the value from the Option. Calls os.Exit(code) if the Option is None.
func (o Option[T]) UnwrapOrExit(code int) T {
	if o.value != nil {
		return *o.value
	}

	os.Exit(code)
	return Zero[T]()
}
