package hocon

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	t.Run("parse a config with a single value", testParse_SingleValue)
	t.Run("parse a config with a single value as an environment variable", testParse_SingleValueEnvironmentVariable)
}

func testParse_SingleValue(t *testing.T) {
	is := is.NewRelaxed(t)
	input := strings.NewReader(`foo = bar`)

	c, err := Parse(input)
	is.NoErr(err)

	value, isFound := c.Get("foo")
	is.True(isFound)
	is.Equal(value.String(), "bar")
}

func testParse_SingleValueEnvironmentVariable(t *testing.T) {
	is := is.NewRelaxed(t)
	input := strings.NewReader(`foo = ${MYAPP_HOME}`)

	c, err := Parse(input)
	is.NoErr(err)

	value, isFound := c.Get("foo")
	is.True(isFound)
	is.Equal(value.String(), "${MYAPP_HOME}")
}
