package pedantic

import (
	"regexp"
	"regexp/syntax"
	"testing"
	"unicode/utf16"
)

func TestUnicode(t *testing.T) {
	unicodeNonASCIIString := `^[` + unicodeNonASCII + `]$`

	r, err := syntax.Parse(unicodeNonASCIIString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log("input:", unicodeNonASCIIString)
	t.Log("simplify:", r.Simplify())

	type test struct {
		r     []rune
		match bool
	}

	tests := []test{
		// Test unicode codepoints.
		{[]rune{0x80}, false},   // 0x0080: technically it should, but no.
		{[]rune{0xa0}, true},    // 0x00A0
		{[]rune{0xd7ff}, true},  // 0xd7ff
		{[]rune{0xd800}, false}, // 0xd800

		{[]rune{0xd800, 0xdc00}, true},  // high surrogate with low surrogate pair
		{[]rune{0xdbff, 0xdffd}, true},  // high surrogate with low surrogate pair
		{[]rune{0xdbff, 0xdffe}, false}, // valid surrogate pair, but noncharacter
		{[]rune{0xdbff, 0xdfff}, false}, // valid surrogate pair, but noncharacter

		{[]rune{0xdfff}, false}, // 0xdfff
		{[]rune{0xe000}, true},  // 0xe000
		{[]rune{0xfdcf}, true},  // 0xfdcf
		{[]rune{0xfdd0}, false}, // 0xfdd0

		{[]rune{0xfdef}, false}, // 0xfdef
		{[]rune{0xfdf0}, true},  // 0xfdf0
		{[]rune{0xfffc}, true},  // 0xfffc
		{[]rune{0xfffd}, false}, // 0xfffd

		{[]rune{0xffff}, false},  // 0xffff
		{[]rune{0x10000}, true},  // 0x010000
		{[]rune{0x1fffd}, true},  // 0x01fffd
		{[]rune{0x1fffe}, false}, // 0x01fffe
		{[]rune{0x1ffff}, false}, // 0x01ffff

		{[]rune{0x20000}, true},  // 0x020000
		{[]rune{0x2fffd}, true},  // 0x02fffd
		{[]rune{0x2fffe}, false}, // 0x02fffe
		{[]rune{0x2ffff}, false}, // 0x02ffff

		{[]rune{0x30000}, true},  // 0x030000
		{[]rune{0x3fffd}, true},  // 0x03fffd
		{[]rune{0x3fffe}, false}, // 0x03fffe
		{[]rune{0x3ffff}, false}, // 0x03ffff

		{[]rune{0x40000}, true},  // 0x040000
		{[]rune{0x4fffd}, true},  // 0x04fffd
		{[]rune{0x4fffe}, false}, // 0x04fffe
		{[]rune{0x4ffff}, false}, // 0x04ffff

		{[]rune{0x50000}, true},  // 0x050000
		{[]rune{0x5fffd}, true},  // 0x05fffd
		{[]rune{0x5fffe}, false}, // 0x05fffe
		{[]rune{0x5ffff}, false}, // 0x05ffff

		{[]rune{0x60000}, true},  // 0x060000
		{[]rune{0x6fffd}, true},  // 0x06fffd
		{[]rune{0x6fffe}, false}, // 0x06fffe
		{[]rune{0x6ffff}, false}, // 0x06ffff

		{[]rune{0x70000}, true},  // 0x070000
		{[]rune{0x7fffd}, true},  // 0x07fffd
		{[]rune{0x7fffe}, false}, // 0x07fffe
		{[]rune{0x7ffff}, false}, // 0x07ffff

		{[]rune{0x80000}, true},  // 0x080000
		{[]rune{0x8fffd}, true},  // 0x08fffd
		{[]rune{0x8fffe}, false}, // 0x08fffe
		{[]rune{0x8ffff}, false}, // 0x08ffff

		{[]rune{0x90000}, true},  // 0x090000
		{[]rune{0x9fffd}, true},  // 0x09fffd
		{[]rune{0x9fffe}, false}, // 0x09fffe
		{[]rune{0x9ffff}, false}, // 0x09ffff

		{[]rune{0xa0000}, true},  // 0x0a0000
		{[]rune{0xafffd}, true},  // 0x0afffd
		{[]rune{0xafffe}, false}, // 0x0afffe
		{[]rune{0xaffff}, false}, // 0x0affff

		{[]rune{0xb0000}, true},  // 0x0b0000
		{[]rune{0xbfffd}, true},  // 0x0bfffd
		{[]rune{0xbfffe}, false}, // 0x0bfffe
		{[]rune{0xbffff}, false}, // 0x0bffff

		{[]rune{0xc0000}, true},  // 0x0c0000
		{[]rune{0xcfffd}, true},  // 0x0cfffd
		{[]rune{0xcfffe}, false}, // 0x0cfffe
		{[]rune{0xcffff}, false}, // 0x0cffff

		{[]rune{0xd0000}, true},  // 0x0d0000
		{[]rune{0xdfffd}, true},  // 0x0dfffd
		{[]rune{0xdfffe}, false}, // 0x0dfffe
		{[]rune{0xdffff}, false}, // 0x0dffff

		{[]rune{0xe0000}, true},  // 0x0e0000
		{[]rune{0xefffd}, true},  // 0x0efffd
		{[]rune{0xefffe}, false}, // 0x0efffe
		{[]rune{0xeffff}, false}, // 0x0effff

		{[]rune{0xf0000}, true},  // 0x0f0000
		{[]rune{0xffffd}, true},  // 0x0ffffd
		{[]rune{0xffffe}, false}, // 0x0ffffe
		{[]rune{0xfffff}, false}, // 0x0fffff

		{[]rune{0x100000}, true},  // 0x100000
		{[]rune{0x10fffd}, true},  // 0x10fffd
		{[]rune{0x10fffe}, false}, // 0x10fffe
		{[]rune{0x10ffff}, false}, // 0x10ffff

		{[]rune{0x110000}, false}, // 0x110000
	}

	unicodeNonASCIIRegex := regexp.MustCompile(unicodeNonASCIIString)
	for _, tt := range tests {
		s := string(tt.r)

		if len(tt.r) > 1 {
			if !utf16.IsSurrogate(tt.r[0]) || !utf16.IsSurrogate(tt.r[1]) {
				t.Fatal("test pair is not composed of surrogates")
			}

			t.Log("surrogate test")
			s = string(utf16.DecodeRune(tt.r[0], tt.r[1]))
		}

		got := unicodeNonASCIIRegex.MatchString(s)
		if got != tt.match {
			switch tt.match {
			case true:
				t.Errorf("expected %x (%q) to match, but it did not", tt.r, s)
			case false:
				t.Errorf("expected %x (%q) to not match, but it did", tt.r, s)
			}
		}
	}
}
