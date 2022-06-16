package optional

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalingNull(t *testing.T) {
	o := OfNil[string]()
	j, err := o.MarshalJSON()
	if assert.NoError(t, err) {
		assert.Equal(t, `null`, string(j))
	}
}

func TestMarshalingString(t *testing.T) {
	o := Of("foo")
	j, err := o.MarshalJSON()
	if assert.NoError(t, err) {
		assert.Equal(t, `"foo"`, string(j))
	}
}

func TestMarshalingInt(t *testing.T) {
	o := Of(42)
	j, err := o.MarshalJSON()
	if assert.NoError(t, err) {
		assert.Equal(t, `42`, string(j))
	}
}

func TestUnmarshalingNull(t *testing.T) {
	j := `null`
	var o Value[string]
	if assert.NoError(t, json.Unmarshal([]byte(j), &o)) {
		assert.Nil(t, o.value)
	}
}

func TestUnmarshalingValue(t *testing.T) {
	j := `"foo"`
	var o Value[string]
	if assert.NoError(t, json.Unmarshal([]byte(j), &o)) {
		if assert.NotNil(t, o.value) {
			assert.Equal(t, "foo", *o.value)
		}
	}
}

func TestValuer(t *testing.T) {
	oi := Of(42)
	oiv, err := oi.Value()
	if assert.NoError(t, err) {
		assert.Equal(t, int64(42), oiv)
	}

	ou8 := Of(uint8(42))
	ou8v, err := ou8.Value()
	if assert.NoError(t, err) {
		assert.Equal(t, int64(42), ou8v)
	}

	ou16 := Of(uint16(42))
	ou16v, err := ou16.Value()
	if assert.NoError(t, err) {
		assert.Equal(t, int64(42), ou16v)
	}

	onil := OfNil[string]()
	onilv, err := onil.Value()
	if assert.NoError(t, err) {
		assert.Nil(t, onilv)
	}
}

func TestScannerIntegers(t *testing.T) {
	var i64 int64 = 42

	var oi Value[int]
	if assert.NoError(t, oi.Scan(i64)) {
		assert.NotNil(t, oi.value)
		assert.Equal(t, int(i64), oi.MustGet())
	}

	var oi8 Value[int8]
	if assert.NoError(t, oi8.Scan(i64)) {
		assert.NotNil(t, oi8.value)
		assert.Equal(t, int8(i64), oi8.MustGet())
	}

	var oi16 Value[int16]
	if assert.NoError(t, oi16.Scan(i64)) {
		assert.NotNil(t, oi16.value)
		assert.Equal(t, int16(i64), oi16.MustGet())
	}

	var oi32 Value[int32]
	if assert.NoError(t, oi32.Scan(i64)) {
		assert.NotNil(t, oi32.value)
		assert.Equal(t, int32(i64), oi32.MustGet())
	}
}

func TestScannerFloats(t *testing.T) {
	var f64 float64 = 42.99

	var of32 Value[float32]
	if assert.NoError(t, of32.Scan(f64)) {
		assert.NotNil(t, of32.value)
		assert.Equal(t, float32(f64), of32.MustGet())
	}

	var of64 Value[float64]
	if assert.NoError(t, of64.Scan(f64)) {
		assert.NotNil(t, of64.value)
		assert.Equal(t, f64, of64.MustGet())
	}
}

func TestScannerNils(t *testing.T) {
	var onf64 = new(Value[float64])
	if assert.NoError(t, onf64.Scan(nil)) {
		assert.Nil(t, onf64.value)
	}
	var ons Value[string]
	if assert.NoError(t, ons.Scan(nil)) {
		assert.Nil(t, ons.value)
	}
}

func TestScannerStrings(t *testing.T) {
	var os = new(Value[string])
	if assert.NoError(t, os.Scan("foo")) {
		assert.NotNil(t, os.value)
		assert.Equal(t, "foo", os.MustGet())
	}
}

func TestScannerTime(t *testing.T) {
	var otm Value[time.Time]
	tm := time.Now()
	if assert.NoError(t, otm.Scan(tm)) {
		assert.NotNil(t, otm.value)
		assert.Equal(t, tm, otm.MustGet())
	}

	var ob Value[[]byte]
	if assert.NoError(t, ob.Scan([]byte("foo"))) {
		assert.NotNil(t, ob.value)
		assert.Equal(t, []byte("foo"), ob.MustGet())
	}

	var ons Value[string]
	if assert.NoError(t, ons.Scan(nil)) {
		assert.Nil(t, ons.value)
	}
}

func TestScannerBytes(t *testing.T) {
	var ob Value[[]byte]
	if assert.NoError(t, ob.Scan([]byte("foo"))) {
		assert.NotNil(t, ob.value)
		assert.Equal(t, []byte("foo"), ob.MustGet())
	}
}

func TestScannerOverflows(t *testing.T) {
	var oi8 Value[int8]
	assert.NoError(t, oi8.Scan(int64(127)))
	assert.Error(t, oi8.Scan(int64(128)))
	assert.NoError(t, oi8.Scan(int64(-128)))
	assert.Error(t, oi8.Scan(int64(-129)))

	var ou8 Value[uint8]
	assert.NoError(t, ou8.Scan(int64(255)))
	assert.Error(t, ou8.Scan(int64(256)))
	assert.NoError(t, ou8.Scan(int64(0)))
	assert.Error(t, ou8.Scan(int64(-1)))
}

func TestMarshalingStruct(t *testing.T) {
	x := struct {
		Bar Value[string]
		Baz Value[int]
	}{
		Bar: Of("foo"),
		Baz: Of(42),
	}

	j, err := json.Marshal(x)

	if assert.NoError(t, err) {
		if assert.NotEmpty(t, j) {
			assert.Equal(t, `{"Bar":"foo","Baz":42}`, string(j))
		}
	}
}

func TestUnmarshalingStruct(t *testing.T) {
	var x struct {
		Bar Value[string]
		Baz Value[int]
	}
	j := `{
		"bar": "foo",
		"baz": 42
	}`
	if assert.NoError(t, json.Unmarshal([]byte(j), &x)) {
		if assert.NotNil(t, x.Bar.value) {
			assert.Equal(t, "foo", *x.Bar.value)
		}
		if assert.NotNil(t, x.Baz.value) {
			assert.Equal(t, 42, *x.Baz.value)
		}
	}
}
