package optional

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"
)

// MarshalJSON implements json.Marshaler interface
func (o Value[T]) MarshalJSON() ([]byte, error) {
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

// Value implements driver.Valuer interface
func (o *Value[T]) Value() (driver.Value, error) {
	if o == nil || o.value == nil {
		return nil, nil
	}

	switch val := any(*o.value).(type) {
	case driver.Valuer:
		return val.Value()
	case string, int64, float64, bool, []byte, time.Time:
		return val, nil
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case uint:
		if val > math.MaxInt64 {
			goto overflow
		}
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		if val > math.MaxInt64 {
			goto overflow
		}
		return int64(val), nil
	case float32:
		return float64(val), nil
	}

	return nil, fmt.Errorf("failed to convert optional %T value %v to supported driver.Value", o.value, o.value)

overflow:
	return nil, fmt.Errorf("overflow casting %T value %v to int64", o.value, o.value)
}

// Scan implements sql.Scanner interface
func (o *Value[T]) Scan(src any) error {
	if o == nil {
		return errors.New("cannot scan value into nil receiver")
	}

	switch val := src.(type) {
	case T:
		o.value = &val

	case nil:
		o.value = nil

	case sql.Scanner:
		return val.Scan(src)

	case int64:
		switch ov := any(&o.value).(type) {
		case **int:
			if val > math.MaxInt || val < math.MinInt {
				goto overflow
			}
			cv := int(val)
			*ov = &cv
		case **int8:
			if val > math.MaxInt8 || val < math.MinInt8 {
				goto overflow
			}
			cv := int8(val)
			*ov = &cv
		case **int16:
			if val > math.MaxInt16 || val < math.MinInt16 {
				goto overflow
			}
			cv := int16(val)
			*ov = &cv
		case **int32:
			if val > math.MaxInt32 || val < math.MinInt32 {
				goto overflow
			}
			cv := int32(val)
			*ov = &cv
		case **uint:
			if val < 0 {
				goto overflow
			}
			cv := uint(val)
			*ov = &cv
		case **uint8:
			if val < 0 || val > math.MaxUint8 {
				goto overflow
			}
			cv := uint8(val)
			*ov = &cv
		case **uint16:
			if val < 0 || val > math.MaxUint16 {
				goto overflow
			}
			cv := uint16(val)
			*ov = &cv
		case **uint32:
			if val < 0 || val > math.MaxUint32 {
				goto overflow
			}
			cv := uint32(val)
			*ov = &cv
		case **uint64:
			if val < 0 {
				goto overflow
			}
			cv := uint64(val)
			*ov = &cv
		default:
			goto unsupported
		}

	case float64:
		switch ov := any(&o.value).(type) {
		case **float32:
			if val > math.MaxFloat32 {
				goto overflow
			}
			cv := float32(val)
			*ov = &cv
		default:
			goto unsupported
		}

	default:
		goto unsupported
	}

	return nil

unsupported:
	return fmt.Errorf("failed to scan unsupported src %T into dest %T", src, o.value)

overflow:
	return fmt.Errorf("overflow casting %T value %v to %T", src, src, o.value)
}
