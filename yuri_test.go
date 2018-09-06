// Copyright (c) 2018, Hayden Eskriett <hayden@eskriett.com>
// See LICENSE for licensing details

package yuri

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    string
	expected interface{}
}

func getExpected(c testCase) string {
	switch t := c.expected.(type) {
	case string:
		return t
	case bool:
		if t {
			return c.input
		}
	}
	return ""
}

var testCases = []testCase{
	// SHOULD FAIL
	{``, nil},
	{`http://`, nil},
	{`http://.`, nil},
	{`http://..`, nil},
	{`http://../`, nil},
	{`http://?`, nil},
	{`http://??`, nil},
	{`http://??/`, nil},
	{`http://#`, nil},
	{`http://##`, nil},
	{`http://##/`, nil},
	{`//`, nil},
	{`//a`, nil},
	{`///a`, nil},
	{`///`, nil},
	{`http:///a`, nil},
	{`foo.bar`, nil},
	{`foo://1234`, nil},
	{`h://foobar`, nil},
	{`http:// foo.bar`, nil},
	{`:// foo`, nil},
	{`ftps://foo.bar/`, nil},
	{`http://-error-.invalid/`, nil},
	{`http://-a.b.co`, nil},
	{`http://.www.foo.bar/`, nil},

	// SHOULD EXTRACT
	{`http://foo.bar`, true},
	{`https://foo.bar`, true},
	{`hxxp://foo.bar`, true},
	{`hxxps://foo.bar`, true},
	{`ftp://foo.bar`, true},
	{`mailto:foo@bar.baz`, true},
	{`http://foo.bar/`, true},
	{`http://foo.bar/baz_qux`, true},
	{`http://foo.bar/baz_qux/`, true},
	{`http://foo.bar/baz_qux_(quux)`, true},
	{`http://foo.bar/baz_qux_(quux)/(quuz)`, true},
	{`http://foo.bar/baz_qux_(quux)_(quuz)`, true},
	{`http://www.foo.bar/baz/?p=364`, true},
	{`https://www.foo.bar/baz/?qux=quux&quuz=100&corge`, true},
	{`http://✪foo.bar/123`, true},
	{`http://user:password@foo.bar:8080`, true},
	{`http://user:password@foo.bar:8080/`, true},
	{`http://user@foo.bar`, true},
	{`http://user@foo.bar/`, true},
	{`http://user@foo.bar:8080`, true},
	{`http://user@foo.bar:8080/`, true},
	{`http://user:password@foo.bar`, true},
	{`http://user:password@foo.bar/`, true},
	{`http://123.123.123.123/`, true},
	{`http://123.123.123.123:8080/`, true},
	{`http://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]`, true},
	{`http://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]/`, true},
	{`http://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:8080`, true},
	{`http://[2001:db8:85a3:0:0:8a2e:370:7334]:8080`, true},
	{`http://[2001:db8:85a3::8a2e:370:7334]:8080`, true},
	{`http://➡.foo.bar/☃`, true},
	{`http://☃.foo.bar`, true},
	{`http://☃.foo.bar/`, true},
	{`http://☃.☃☃/`, true},
	{`http://foo.bar/_(baz)#anchor`, true},
	{`http://foo.bar/baz_(qux)_quux#anchor`, true},
	{`http://foo.bar/utf8_snowman_(☃)_in_parens`, true},
	{`http://foo.bar/(something)?after=parens`, true},
	{`http://foo.bar.baz/qux/#&quux=quuz`, true},
	{`ftp://foo.bar/baz`, true},
	{`http://foo.bar/?q=baz%20url-encoded%20path`, true},
	{`url in string should be extracted http://foo.bar/baz`, `http://foo.bar/baz`},
	{`[http://foo.bar/baz]`, `http://foo.bar/baz`},
	{`[http://foo.bar/baz_(qux)]`, `http://foo.bar/baz_(qux)`},
	{`<http://foo.bar/baz>`, `http://foo.bar/baz`},
	{`<http://foo.bar/baz>qux`, `http://foo.bar/baz`},
	{`,http://foo.bar/baz`, `http://foo.bar/baz`},
}

func TestYankURIs(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			uris := YankURIs([]byte(c.input))
			got := ""
			if len(uris) > 0 {
				got = uris[0]
			}
			want := getExpected(c)
			if got != want {
				t.Errorf(`YankURI("%s") got "%s", expected "%s"`, c.input, got, want)
			}
		})
	}
}

func ExampleYankURIs() {
	uris := YankURIs([]byte("yuri lives at https://github.com/eskriett/yuri"))
	fmt.Printf("%v\n", uris)
	// Output: [https://github.com/eskriett/yuri]
}
