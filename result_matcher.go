package gost

type ResultMatcher[T any] struct {
	onOk  func(v *T) *T
	value *T
	error error
}

func (rm ResultMatcher[T]) OnError(cb func(e error) T) T {
	if rm.value != nil {
		return *rm.value
	}

	return cb(rm.error)
}
