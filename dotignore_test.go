package dotignore

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {

	type TestCase struct {
		path    string
		pattern string
		exp     bool
	}

	cases := []TestCase{
		// empty
		{
			path:    "",
			pattern: "",
			exp:     true,
		},
		{
			path:    "path",
			pattern: "",
			exp:     true,
		},
		{
			path:    "/path",
			pattern: "",
			exp:     false,
		},
		{
			path:    "path/",
			pattern: "",
			exp:     false,
		},
		{
			path:    "/path/",
			pattern: "",
			exp:     false,
		},
		// "file"
		{
			path:    "",
			pattern: "path",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "path",
			exp:     true,
		},
		{
			path:    "/path",
			pattern: "path",
			exp:     true,
		},
		{
			path:    "path/",
			pattern: "path",
			exp:     false,
		},
		{
			path:    "/path/",
			pattern: "path",
			exp:     false,
		},
		// "absolute file"
		{
			path:    "",
			pattern: "/path",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "/path",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "/path",
			exp:     true,
		},
		{
			path:    "path/",
			pattern: "/path",
			exp:     false,
		},
		{
			path:    "/path/",
			pattern: "/path",
			exp:     false,
		},
		// "path"
		{
			path:    "",
			pattern: "path/",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "path/",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "path/",
			exp:     false,
		},
		{
			path:    "path/",
			pattern: "path/",
			exp:     true,
		},
		{
			path:    "/path/",
			pattern: "path/",
			exp:     true,
		},
		// "absolute path"
		{
			path:    "",
			pattern: "/path/",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "/path/",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "/path/",
			exp:     false,
		},
		{
			path:    "path/",
			pattern: "/path/",
			exp:     false,
		},
		{
			path:    "/path/",
			pattern: "/path/",
			exp:     true,
		},
		// "wildcard"
		{
			path:    "",
			pattern: "*",
			exp:     true,
		},
		{
			path:    "path",
			pattern: "*",
			exp:     true,
		},
		{
			path:    "/path",
			pattern: "*",
			exp:     true,
		},
		{
			path:    "path/",
			pattern: "*",
			exp:     true,
		},
		{
			path:    "/path/",
			pattern: "*",
			exp:     true,
		},
		// "wildcard"
		{
			path:    "",
			pattern: "/*",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "/*",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "/*",
			exp:     true,
		},
		{
			path:    "path/",
			pattern: "/*",
			exp:     true,
		},
		{
			path:    "/path/",
			pattern: "/*",
			exp:     true,
		},
		// **/path (absolute file)
		{
			path:    "",
			pattern: "**/path",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "**/path",
			exp:     true,
		},
		{
			path:    "/path",
			pattern: "**/path",
			exp:     true,
		},
		{
			path:    "path/",
			pattern: "**/path",
			exp:     false,
		},
		{
			path:    "/path/",
			pattern: "**/path",
			exp:     false,
		},
		// **/path/ (absolute path)
		{
			path:    "",
			pattern: "**/path/",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "**/path/",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "**/path/",
			exp:     false,
		},
		{
			path:    "path/",
			pattern: "**/path/",
			exp:     true,
		},
		{
			path:    "/path/",
			pattern: "**/path/",
			exp:     true,
		},
		// path/** (path)
		{
			path:    "",
			pattern: "path/**",
			exp:     false,
		},
		{
			path:    "path",
			pattern: "path/**",
			exp:     false,
		},
		{
			path:    "/path",
			pattern: "path/**",
			exp:     false,
		},
		{
			path:    "path/",
			pattern: "path/**",
			exp:     true,
		},
		{
			path:    "/path/",
			pattern: "path/**",
			exp:     true,
		},
		// a/**/b (path and file)
		{
			path:    "",
			pattern: "a/**/b",
			exp:     false,
		},
		{
			path:    "a/b",
			pattern: "a/**/b",
			exp:     true,
		},
		{
			path:    "a/x/b",
			pattern: "a/**/b",
			exp:     true,
		},
		{
			path:    "a/x/y/b",
			pattern: "a/**/b",
			exp:     true,
		},
		{
			path:    "a/x/y/b",
			pattern: "a/x/b",
			exp:     false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			parts := strings.Split(tc.path, "/")
			patterns := strings.Split(tc.pattern, "/")
			res := compare(parts, patterns)
			assert.Equal(t, tc.exp, res)
		})
	}
}
