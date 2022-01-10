package optional

// Value wraps an optional value of type T
type Value[T any] struct {
	value *T
}

// Of instantiates a new optional value with value t
func Of[T any](t T) Value[T] {
	return Value[T]{&t}
}

// OfNil instantiates a new optional value of type T with value set to nil
func OfNil[T any]() Value[T] {
	return Value[T]{nil}
}

// Set sets the underlying optional value to the provided non-nil value
func (o *Value[T]) Set(value T) {
	o.value = &value
}

// SetNil sets the underlying optional value to nil
func (o *Value[T]) SetNil() {
	o.value = nil
}

// Get returns true if optional value is non-nil, along with the underlying value
func (o *Value[T]) Get() (ok bool, value T) {
	if o.value != nil {
		return true, *o.value
	}
	zeroValue := new(T)
	return false, *zeroValue
}

// ValueOrDefault returns the underlying non-nil value or the provided fallback value
func (o *Value[T]) ValueOrDefault(value T) T {
	if o.value == nil {
		return value
	}
	return *o.value
}

// MustGet returns the underlying value or PANICS if value is nil
func (o *Value[T]) MustGet() (value T) {
	return *o.value
}

// HasValue returns true if the underlying value is non-nil
func (o *Value[T]) HasValue() bool {
	return o.value != nil
}

// IsNil returns true if the underlying value is nil
func (o *Value[T]) IsNil() bool {
	return o.value == nil
}
