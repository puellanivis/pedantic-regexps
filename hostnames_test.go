package pedantic

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"testing"
)

const hostnameRegexString = `^` + hostnameString + `$`

func TestHostname(t *testing.T) {
	input := hostnameRegexString

	r, err := syntax.Parse(input, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", input)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%d", i), true})
	}
	tests = append(tests, test{"256", true})
	for i := 0; i < 100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%d", i), true})
	}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{fmt.Sprintf("00%d", i), true})
	}
	tests = append(tests, test{"-1", false})

	tests = append(tests, []test{
		{"example", true},
		{"example.org", true},
		{"example.tld", true},
		{"a-a.example", true},
		{"a--a.example", true},
		{"0-0.example", true},
		{"a-.example", false},
		{"-a.example", false},
		{"\xa0.example", false},   // basic unicode test
		{"\ufffd.example", false}, // ensure replacement character is not matched
	}...)

	re := regexp.MustCompile(input)

	for _, tt := range tests {
		got := re.MatchString(tt.s)
		if got != tt.match {
			switch tt.match {
			case true:
				t.Errorf("expected %q to match, but it did not", tt.s)
			case false:
				t.Errorf("expected %q to not match, but it did", tt.s)
			}
		}
	}
}
