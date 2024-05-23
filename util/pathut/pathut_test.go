package pathut

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	testCases := []struct {
		path     string
		expected []string
	}{
		{
			path:     "C:/Program Files\\Git\\bin\\git.exe",
			expected: []string{"C:", "Program Files", "Git", "bin", "git.exe"},
		},
		{
			path:     "/usr/local/bin/kubectl",
			expected: []string{"usr", "local", "bin", "kubectl"},
		},
		{
			path:     "foo/bar\\baz",
			expected: []string{"foo", "bar", "baz"},
		},
	}
	for _, tc := range testCases {
		actual := SplitPath(tc.path)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestJoinPath(t *testing.T) {
	type args struct {
		elem []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{[]string{}}, ""},
		{"one empty", args{[]string{""}}, ""},
		{"two empty", args{[]string{"", ""}}, ""},
		{"three empty", args{[]string{"", "", ""}}, ""},
		{"one", args{[]string{"a"}}, "a"},
		{"two", args{[]string{"a", "b"}}, filepath.Join("a", "b")},
		{"three", args{[]string{"a", "b", "c"}}, filepath.Join("a", "b", "c")},
		{"four", args{[]string{"a", "b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with empty in the middle", args{[]string{"a", "", "b"}}, filepath.Join("a", "b")},
		{"with empty at the end", args{[]string{"a", "b", ""}}, filepath.Join("a", "b")},
		{"with empty at the beginning", args{[]string{"", "a", "b"}}, filepath.Join("a", "b")},
		{"with backslash", args{[]string{"a\\b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with forwardslash", args{[]string{"a/b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with backslash at the end", args{[]string{"a\\b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with forwardslash at the end", args{[]string{"a/b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with backslash at the beginning", args{[]string{"a\\b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"with forwardslash at the beginning", args{[]string{"a/b", "c", "d"}}, filepath.Join("a", "b", "c", "d")},
		{"test prefix", args{[]string{"/a/b", "c", "d"}}, filepath.Join("/a", "b", "c", "d")},
		{"test prefix2", args{[]string{"\\a/b", "c", "d"}}, filepath.Join("/a", "b", "c", "d")},
		{"test prefix", args{[]string{"/a/b", "c", "d"}}, filepath.Join("/a", "b", "c", "d")},
		{"test prefix2", args{[]string{"\\a/b", "c", "d"}}, filepath.Join("/a", "b", "c", "d")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filepath.Join(tt.args.elem...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinPath2(t *testing.T) {
	fmt.Println(filepath.Join("D/a", "b/v\\d", "c", "d"))
	path := JoinPath("/a/b", "c", "d")
	t.Log(path)
}
