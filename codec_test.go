package optional

import (
	"encoding/json"
	"testing"

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
