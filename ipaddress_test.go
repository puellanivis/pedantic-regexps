package pedantic

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"testing"
)

const ipv4NumberRegexString = `^` + ipv4Number + `$`

func TestIPv4(t *testing.T) {
	r, err := syntax.Parse(ipv4NumberRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}
	t.Log(r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%d", i), true})
	}
	tests = append(tests, test{"256", false})
	for i := 0; i < 100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%d", i), false})
	}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{fmt.Sprintf("00%d", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	IPv4 := regexp.MustCompile(ipv4NumberRegexString)

	for _, tt := range tests {
		got := IPv4.MatchString(tt.s)
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
