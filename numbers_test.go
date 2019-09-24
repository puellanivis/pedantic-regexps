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

const (
	numberUint16HexRegexString     = `^` + numberUint16Hex + `$`
	numberUint16HexFullRegexString = `^` + numberUint16HexFull + `$`
	numberUint16HexAmbiRegexString = `^` + numberUint16HexAmbi + `$`
)

func TestByteBin(t *testing.T) {
	input := numberByteBinRegexString

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

func TestByteOct(t *testing.T) {
	input := numberByteOctRegexString

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

func TestByteDec(t *testing.T) {
	input := numberByteDecRegexString

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
	tests = append(tests, test{"256", false})
	for i := 0; i < 100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%d", i), false})
	}
	for i := 0; i < 10; i++ {
		tests = append(tests, test{fmt.Sprintf("00%d", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

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

func TestByteHex(t *testing.T) {
	input := numberByteHexRegexString

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

func TestByteBinFull(t *testing.T) {
	input := numberByteBinFullRegexString

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

func TestByteOctFull(t *testing.T) {
	input := numberByteOctFullRegexString

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

func TestByteDecFull(t *testing.T) {
	input := numberByteDecFullRegexString

	r, err := syntax.Parse(input, syntax.Perl)
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

func TestByteHexFull(t *testing.T) {
	input := numberByteHexFullRegexString

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

func TestUint16Hex(t *testing.T) {
	input := numberUint16HexRegexString

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
	for i := 0; i < 65536; i++ {
		tests = append(tests, test{fmt.Sprintf("%x", i), true})
		tests = append(tests, test{fmt.Sprintf("%X", i), true})
	}
	tests = append(tests, test{"10000", false})
	for i := 0; i < 0x1000; i++ {
		tests = append(tests, test{fmt.Sprintf("0%x", i), false})
		tests = append(tests, test{fmt.Sprintf("0%X", i), false})
	}
	for i := 0; i < 0x100; i++ {
		tests = append(tests, test{fmt.Sprintf("00%x", i), false})
		tests = append(tests, test{fmt.Sprintf("00%X", i), false})
	}
	for i := 0; i < 0x10; i++ {
		tests = append(tests, test{fmt.Sprintf("000%x", i), false})
		tests = append(tests, test{fmt.Sprintf("000%X", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

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

func TestUint16HexFull(t *testing.T) {
	input := numberUint16HexFullRegexString

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
	for i := 0; i < 65536; i++ {
		tests = append(tests, test{fmt.Sprintf("%04x", i), true})
		tests = append(tests, test{fmt.Sprintf("%04X", i), true})
	}
	tests = append(tests, test{"10000", false})
	for i := 0; i < 0x1000; i++ {
		tests = append(tests, test{fmt.Sprintf("%x", i), false})
		tests = append(tests, test{fmt.Sprintf("%X", i), false})
	}
	for i := 0; i < 0x100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%x", i), false})
		tests = append(tests, test{fmt.Sprintf("0%X", i), false})
	}
	for i := 0; i < 0x10; i++ {
		tests = append(tests, test{fmt.Sprintf("00%x", i), false})
		tests = append(tests, test{fmt.Sprintf("00%X", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

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

func TestUint16HexAmbi(t *testing.T) {
	input := numberUint16HexAmbiRegexString

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
	for i := 0; i < 65536; i++ {
		tests = append(tests, test{fmt.Sprintf("%x", i), true})
		tests = append(tests, test{fmt.Sprintf("%X", i), true})
		tests = append(tests, test{fmt.Sprintf("%04x", i), true})
		tests = append(tests, test{fmt.Sprintf("%04X", i), true})
	}
	tests = append(tests, test{"10000", false})
	for i := 0; i < 0x100; i++ {
		tests = append(tests, test{fmt.Sprintf("0%x", i), false})
		tests = append(tests, test{fmt.Sprintf("0%X", i), false})
	}
	for i := 0; i < 0x10; i++ {
		tests = append(tests, test{fmt.Sprintf("00%x", i), false})
		tests = append(tests, test{fmt.Sprintf("00%X", i), false})
	}
	tests = append(tests, test{"-1", false})
	tests = append(tests, test{"alpha", false})

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
