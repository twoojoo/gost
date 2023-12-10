package gost

type OptionMatcher[T any] struct {
	onSome func(v *T) *T
	value  *T
}

func (om OptionMatcher[T]) OnNone(cb func() T) T {
	if om.value != nil {
		return *om.onSome(om.value)
	}

	return cb()
}
