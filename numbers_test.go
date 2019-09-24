package pedantic

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"testing"
)

const (
	numberByteBinRegexString = `^` + numberByteBin + `$`
	numberByteOctRegexString = `^` + numberByteOct + `$`
	numberByteDecRegexString = `^` + numberByteDec + `$`
	numberByteHexRegexString = `^` + numberByteHex + `$`
)

const (
	numberByteBinFullRegexString = `^` + numberByteBinFull + `$`
	numberByteOctFullRegexString = `^` + numberByteOctFull + `$`
	numberByteDecFullRegexString = `^` + numberByteDecFull + `$`
	numberByteHexFullRegexString = `^` + numberByteHexFull + `$`
)

func TestByteBin(t *testing.T) {
	r, err := syntax.Parse(numberByteBinRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteBinRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%b", i), true})
	}
	tests = append(tests, test{"100000000", false})
	for i := 0; i < 128; i++ {
		tests = append(tests, test{fmt.Sprintf("0%b", i), false})
	}
	for i := 0; i < 64; i++ {
		tests = append(tests, test{fmt.Sprintf("00%b", i), false})
	}
	for i := 0; i < 32; i++ {
		tests = append(tests, test{fmt.Sprintf("000%b", i), false})
	}
	for i := 0; i < 16; i++ {
		tests = append(tests, test{fmt.Sprintf("0000%b", i), false})
	}
	for i := 0; i < 8; i++ {
		tests = append(tests, test{fmt.Sprintf("00000%b", i), false})
	}
	for i := 0; i < 4; i++ {
		tests = append(tests, test{fmt.Sprintf("000000%b", i), false})
	}
	for i := 0; i < 2; i++ {
		tests = append(tests, test{fmt.Sprintf("0000000%b", i), false})
	}
	tests = append(tests, test{"00000000", false})
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteBin := regexp.MustCompile(numberByteBinRegexString)

	for _, tt := range tests {
		got := NumberByteBin.MatchString(tt.s)
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

func TestByteOct(t *testing.T) {
	r, err := syntax.Parse(numberByteOctRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteOctRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%o", i), true})
	}
	tests = append(tests, test{"400", false})
	for i := 0; i < 0100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%o", i), false})
	}
	for i := 0; i < 010; i++ {
		tests = append(tests, test{fmt.Sprintf("00%o", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteOct := regexp.MustCompile(numberByteOctRegexString)

	for _, tt := range tests {
		got := NumberByteOct.MatchString(tt.s)
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

func TestByteDec(t *testing.T) {
	r, err := syntax.Parse(numberByteDecRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteDecRegexString)
	t.Log("simplify:", r.Simplify())

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

	NumberByteDec := regexp.MustCompile(numberByteDecRegexString)

	for _, tt := range tests {
		got := NumberByteDec.MatchString(tt.s)
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

func TestByteHex(t *testing.T) {
	r, err := syntax.Parse(numberByteHexRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteHexRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%x", i), true})
		tests = append(tests, test{fmt.Sprintf("%X", i), true})
	}
	tests = append(tests, test{"100", false})
	for i := 0; i < 0x10; i++ {
		tests = append(tests, test{fmt.Sprintf("0%x", i), false})
		tests = append(tests, test{fmt.Sprintf("0%X", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteHex := regexp.MustCompile(numberByteHexRegexString)

	for _, tt := range tests {
		got := NumberByteHex.MatchString(tt.s)
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

func TestByteBinFull(t *testing.T) {
	r, err := syntax.Parse(numberByteBinFullRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteBinFullRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%08b", i), true})
	}
	tests = append(tests, test{"100000000", false})
	for i := 0; i < 128; i++ {
		tests = append(tests, test{fmt.Sprintf("%b", i), false})
	}
	for i := 0; i < 64; i++ {
		tests = append(tests, test{fmt.Sprintf("0%b", i), false})
	}
	for i := 0; i < 32; i++ {
		tests = append(tests, test{fmt.Sprintf("00%b", i), false})
	}
	for i := 0; i < 16; i++ {
		tests = append(tests, test{fmt.Sprintf("000%b", i), false})
	}
	for i := 0; i < 8; i++ {
		tests = append(tests, test{fmt.Sprintf("0000%b", i), false})
	}
	for i := 0; i < 4; i++ {
		tests = append(tests, test{fmt.Sprintf("00000%b", i), false})
	}
	for i := 0; i < 2; i++ {
		tests = append(tests, test{fmt.Sprintf("000000%b", i), false})
	}
	tests = append(tests, test{"0", false})
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteBinFull := regexp.MustCompile(numberByteBinFullRegexString)

	for _, tt := range tests {
		got := NumberByteBinFull.MatchString(tt.s)
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

func TestByteOctFull(t *testing.T) {
	r, err := syntax.Parse(numberByteOctFullRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteOctFullRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%03o", i), true})
	}
	tests = append(tests, test{"400", false})
	for i := 0; i < 0100; i++ {
		tests = append(tests, test{fmt.Sprintf("%o", i), false})
	}
	for i := 0; i < 010; i++ {
		tests = append(tests, test{fmt.Sprintf("0%o", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteOctFull := regexp.MustCompile(numberByteOctFullRegexString)

	for _, tt := range tests {
		got := NumberByteOctFull.MatchString(tt.s)
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

func TestByteDecFull(t *testing.T) {
	r, err := syntax.Parse(numberByteDecFullRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteDecRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%03d", i), true})
	}
	tests = append(tests, test{"256", false})
	for i := 0; i < 100; i++ {
		tests = append(tests, test{fmt.Sprintf("%d", i), false})
	}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{fmt.Sprintf("0%d", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteDecFull := regexp.MustCompile(numberByteDecFullRegexString)

	for _, tt := range tests {
		got := NumberByteDecFull.MatchString(tt.s)
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

func TestByteHexFull(t *testing.T) {
	r, err := syntax.Parse(numberByteHexFullRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", numberByteHexFullRegexString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		s     string
		match bool
	}

	var tests []test
	for i := 0; i < 256; i++ {
		tests = append(tests, test{fmt.Sprintf("%02x", i), true})
		tests = append(tests, test{fmt.Sprintf("%02X", i), true})
	}
	tests = append(tests, test{"100", false})
	for i := 0; i < 0x10; i++ {
		tests = append(tests, test{fmt.Sprintf("%x", i), false})
		tests = append(tests, test{fmt.Sprintf("%X", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

	NumberByteHexFull := regexp.MustCompile(numberByteHexFullRegexString)

	for _, tt := range tests {
		got := NumberByteHexFull.MatchString(tt.s)
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
