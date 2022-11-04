package gocon

import (
	"testing"

	"github.com/matryer/is"
)

func TestType(t *testing.T) {
	is := is.New(t)
	value := newConfigUnresolvedValue("foo")
	is.Equal(value.Type(), configUnresolvedValueType)
}

// Validate the last value is returned by the String() method
func TestString(t *testing.T) {
	is := is.New(t)

	noValue := &configUnresolvedValue{value: []string{}}
	value1 := newConfigUnresolvedValue("foo")
	value2 := newConfigUnresolvedValue("foo")
	value2.addValue("bar")

	is.Equal(value1.String(), "foo")
	is.Equal(value2.String(), "bar")
	is.Equal(noValue.String(), "")
}
