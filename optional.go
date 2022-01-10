package optional

type Value[T any] struct {
	value *T
}

func Of[T any](t T) Value[T] {
	return Value[T]{&t}
}

func OfNil[T any]() Value[T] {
	return Value[T]{nil}
}

func (o *Value[T]) Set(value T) {
	o.value = &value
}

func (o *Value[T]) SetNil() {
	o.value = nil
}

func (o *Value[T]) Value() (ok bool, value T) {
	if o.value != nil {
		return true, *o.value
	}
	zeroValue := new(T)
	return false, *zeroValue
}

func (o *Value[T]) HasValue() bool {
	return o.value != nil
}

func (o *Value[T]) ValueOrDefault(value T) T {
	if o.value == nil {
		return value
	}
	return *o.value
}

func (o *Value[T]) MustValue() (value T) {
	return *o.value
}
