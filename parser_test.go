package gocon

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

/* TODO:
https://github.com/lightbend/config/blob/main/HOCON.md#unquoted-strings
----
{
    "foo" : { "a" : 42 },
    "foo" : { "b" : 43 }
}
should equal
{
    "foo" : { "a" : 42, "b" : 43 }
}
----
{
    "foo" : { "a" : 42 },
    "foo" : null,
    "foo" : { "b" : 43 }
}
should equal
{
    "foo" : { "b" : 43 }
}
*/

func TestParse(t *testing.T) {
	t.Run("parse a config with a single value", testParse_SingleValue)
	t.Run("parse a config with a single value as an environment variable", testParse_SingleValueEnvironmentVariable)
	t.Run("single line comments", testParse_SingleLineComment)
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

func testParse_SingleLineComment(t *testing.T) {
	is := is.NewRelaxed(t)
	input := strings.NewReader(
		`a = foo #this is a comment
		 b = bar#this is a comment
		 c = 1`)

	c, err := Parse(input)
	is.NoErr(err)

	value, isFound := c.Get("a")
	is.True(isFound)
	is.Equal(value.String(), "foo")

	value, isFound = c.Get("b")
	is.True(isFound)
	is.Equal(value.String(), "bar")

	value, isFound = c.Get("c")
	is.True(isFound)
	is.Equal(value.String(), "1")
}
