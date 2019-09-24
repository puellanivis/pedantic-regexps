package pedantic

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"strings"
	"testing"
)

const ipv6AddressRegexString = `^` + ipv6Address + `$`

func binToIPV6Addr(i int) string {
	s := fmt.Sprintf("%08b", i)

	switch {
	case strings.Contains(s, "0000000"):
		s = strings.Replace(s, "0000000", "!", 1)
	case strings.Contains(s, "000000"):
		s = strings.Replace(s, "000000", "!", 1)
	case strings.Contains(s, "00000"):
		s = strings.Replace(s, "00000", "!", 1)
	case strings.Contains(s, "0000"):
		s = strings.Replace(s, "0000", "!", 1)
	case strings.Contains(s, "000"):
		s = strings.Replace(s, "000", "!", 1)
	case strings.Contains(s, "00"):
		s = strings.Replace(s, "00", "!", 1)
	}

	if !strings.Contains(s, "!") {
		return strings.Join(strings.Split(s, ""), ":")
	}

	lr := strings.SplitN(s, "!", 2)
	l, r := strings.Split(lr[0], ""), strings.Split(lr[1], "")

	return strings.Join(l, ":") + "::" + strings.Join(r, ":")
}

func TestIPV6(t *testing.T) {
	input := ipv6AddressRegexString

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

	tests := []test{
		{"::", true}, // ???
		{"::0", false},
	}

	for i := 1; i < 0x100; i++ {
		tests = append(tests, test{binToIPV6Addr(i), true})
	}

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
