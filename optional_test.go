package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitNil(t *testing.T) {
	o1 := Value[string]{}
	assert.Nil(t, o1.value)

	o2 := OfNil[string]()
	assert.Nil(t, o2.value)
}
func TestInitOfValue(t *testing.T) {
	o := Of("foo")
	if assert.NotNil(t, o.value) {
		assert.Equal(t, "foo", *o.value)
	}
}

func TestSetters(t *testing.T) {
	o := Value[string]{}
	o.Set("foo")
	if assert.NotNil(t, o.value) {
		assert.Equal(t, "foo", *o.value)
	}

	o.SetNil()
	assert.Nil(t, o.value)
}

func TestGetters(t *testing.T) {
	o := Of("foo")

	assert.True(t, o.HasValue())

	ok, v := o.Value()
	assert.True(t, ok)
	assert.Equal(t, "foo", v)

	o.SetNil()
	ok, v = o.Value()
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestGetterOfDefaultValue(t *testing.T) {
	o := OfNil[string]()
	assert.Equal(t, "bar", o.ValueOrDefault("bar"))

	o.Set("foo")
	assert.Equal(t, "foo", o.ValueOrDefault("bar"))
}

func TestTypeParameters(t *testing.T) {
	s := Of("foo")
	assert.Equal(t, "foo", *s.value)

	i := Of(42)
	assert.Equal(t, 42, *i.value)

	x := Of(struct{}{})
	assert.Equal(t, struct{}{}, *x.value)

}

func TestPointers(t *testing.T) {
	s := "foo"
	sPtr := &s

	o := Of(sPtr)
	assert.True(t, sPtr == *o.value)
}

func TestPanicsAsExpected(t *testing.T) {
	o := OfNil[string]()
	assert.Nil(t, o.value)

	o.Set("foo")
	if assert.NotNil(t, o.value) {
		assert.Equal(t, "foo", *o.value)
	}

	ok, v := o.Value()
	assert.Equal(t, true, ok)
	assert.Equal(t, "foo", v)

	assert.NotPanics(t, func() {
		assert.Equal(t, "foo", o.MustValue())
	})

	o.SetNil()
	assert.Nil(t, o.value)

	assert.Panics(t, func() {
		o.MustValue()
	})
}
