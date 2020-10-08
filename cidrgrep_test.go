package cidrgrep_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tomdoherty/cidrgrep"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		input, cidr, want string
	}{
		{input: "10.1.1.100\n", cidr: "10.1.1.1/24", want: "10.1.1.100"},
		{input: "10.1.1.100\n127.0.0.1\n", cidr: "10.1.1.0/24", want: "10.1.1.100"},
		{input: "no place like 127.0.0.1\n", cidr: "127.0.0.1/32", want: "no place like 127.0.0.1"},

		{input: "10.1.1.100\n", cidr: "10.1.2.0/24", want: ""},
		{input: "555.1.1.999\n", cidr: "10.1.2.0/24", want: ""},
	}

	t.Parallel()

	for _, tc := range testCases {
		output := bytes.Buffer{}
		input := strings.NewReader(tc.input)
		want := tc.want

		cidrgrep.Filter(input, &output, tc.cidr)

		got := output.String()

		if want != got {
			t.Errorf("want: %q, got %q", want, got)
		}
	}
}
