package optional

import (
	"bytes"
	"encoding/json"
)

// MarshalJSON implements json.Marshaler interface
func (o *Value[T]) MarshalJSON() ([]byte, error) {
	if o.value == nil {
		return []byte(`null`), nil
	}
	return json.Marshal(*o.value)
}

// MarshalJSON implements json.Unmarshaler interface
func (o *Value[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`null`)) {
		o.value = nil
		return nil
	}
	o.value = new(T)
	return json.Unmarshal(data, o.value)
}
