package pedantic

import (
	"regexp"
	"regexp/syntax"
	"testing"
)

const emailRegexString = `^` + emailString + `$`

func TestEmail(t *testing.T) {
	input := emailRegexString

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
		// Simple 7bit ASCII:
		{"user@example.org", true},
		{"user", false},
		{"user@", false},
		{"@example.org", false},
		// Dotted forms:
		{"user.name@example.org", true},
		{".user@example.org", false},
		{"user.@example.org", false},

		// Various special characters
		{" @example.org", false},
		{"!@example.org", true},
		{"\"@example.org", false},
		{"#@example.org", true},
		{"$@example.org", true},
		{"%@example.org", true},
		{"&@example.org", true},
		{"'@example.org", true},
		{"(@example.org", false},
		{")@example.org", false},
		{"*@example.org", true},
		{"+@example.org", true},
		{",@example.org", false},
		{"-@example.org", true},
		{".@example.org", false},
		{"/@example.org", true},
		{":@example.org", false},
		{";@example.org", false},
		{"<@example.org", false},
		{"=@example.org", true},
		{">@example.org", false},
		{"?@example.org", true},
		{"@@example.org", false},
		{"[@example.org", false},
		{"\\@example.org", false},
		{"]@example.org", false},
		{"^@example.org", true},
		{"_@example.org", true},
		{"`@example.org", true},
		{"{@example.org", true},
		{"|@example.org", true},
		{"}@example.org", true},
		{"~@example.org", true},
		{"\x7f@example.org", false},

		// Basic Unicode test:
		{"\u00a0@example.org.com", true}, // UNSURE: standard says yes, but this is Unicode whitepace
		{"\u00a1@example.org.com", true},
		{"\ufffd@example.org.com", false}, // test that replacement character does not match

		// Various quote tests:
		{`""@example.org`, true},
		{`"@example.org`, false},
		{`"\""@example.org`, true},
		{`"\"@example.org`, false},
		{`"\\"@example.org`, true},
		{`"\\@example.org`, false},
		{`user@[]`, true},
		{`user@[`, false},
		{`user@]`, false},
		{`user@[\[]`, true},
		{`user@\[]`, false},
		{`user@[\[`, false},
		{`user@[\]]`, true},
		{`user@\]]`, false},
		{`user@[\]`, false},
		{`user@[\\]`, true},
		{`user@\\]`, false},
		{`user@[\\`, false},
		{`user@[example.org]`, true},

		// Test properly escaped whitespace in quotes.
		{`"\ "@example.org`, true},     // This whitespace is properly escaped.
		{`\ "@example.org`, false},     // This whitespace is properly escaped, but no start quote.
		{`"\ @example.org`, false},     // This whitespace is properly escaped, but no end quote.
		{"\"\\\t\"@example.org", true}, // This whitespace is properly escaped. Yes, this is how it is encoded.
		{"\\\t\"@example.org", false},  // This whitespace is properly escaped, but no start quote.
		{"\"\\\t@example.org", false},  // This whitespace is properly escaped, but no end quote.
		{`user@[\ ]`, true},            // This whitespace is properly escaped.
		{`user@\ ]`, false},            // This whitespace is properly escaped, but no start bracket.
		{`user@[\ `, false},            // This whitespace is properly escaped, but no end bracket.
		{"user@[\\\t]", true},          // This whitespace is properly escaped. Yes, this is how it is encoded.
		{"user@\\\t]", false},          // This whitespace is properly escaped, but no start bracket.
		{"user@[\\\t", false},          // This whitespace is properly escaped, but no end bracket.

		// Test CR and LF not allowed:
		{"\"\n\"@example.org", false},
		{"\"\r\"@example.org", false},
		{"\"\\\n\"@example.org", false}, // even escaped, not allowed
		{"\"\\\r\"@example.org", false}, // even escaped, not allowed
		{"\"\r\n\"@example.org", false},
		{"\"\n\r\"@example.org", false},
		{"user@[\n]", false},
		{"user@[\r]", false},
		{"user@[\\\n]", false}, // even escaped, not allowed
		{"user@[\\\r]", false}, // even escaped, not allowed
		{"user@[\n\r]", false},
		{"user@[\r\n]", false},

		// Test Folding-White-Space (FWS) handling (it is invisible, and not part of the address):
		{"user @example.org", false}, // CANONICALLY: user@example.org
		{" user@example.org", false}, // CANONICALLY: user@example.org
		{"user@ example.org", false}, // CANONICALLY: user@example.org
		{"user@example.org ", false}, // CANONICALLY: user@example.org
		{"\" \"@example.org", true},
		{"\"  \"@example.org", false},    // CANONICALLY: " "@example.org
		{"\" \\  \"@example.org", true},  // quoted-string is: "   "
		{"\" \\\t \"@example.org", true}, // quoted-string is: " \t "
		{"\"\t\"@example.org", false},    // CANONICALLY: " "@example.org
		{" \"user\"@example.org", false}, // CANONICALLY: "user"@example.org
		{"\" user\"@example.org", true},
		{"\"user \"@example.org", true},
		{"\" user \"@example.org", true},
		{"\"\tuser\"@example.org", false}, // CANONICALLY: " user"@example.org
		{"\"user\t\"@example.org", false}, // CANONICALLY: "user "@example.org
		{"\"  user\"@example.org", false}, // CANONICALLY: " user"@example.org
		{"\"user  \"@example.org", false}, // CANONICALLY: "user "@example.org
		{"\"user\" @example.org", false},  // CANONICALLY: "user"@example.org
		{"user@ [example.org]", false},    // CANONICALLY: user@[example.org]
		{"user@[ example.org]", true},
		{"user@[example.org ]", true},
		{"user@[\texample.org]", false},       // CANONICALLY: user@[ example.org]
		{"user@[example.org\t]", false},       // CANONICALLY: user@[example.org ]
		{"user@[  example.org]", false},       // CANONICALLY: user@[ example.org]
		{"user@[example.org  ]", false},       // CANONICALLY: user@[example.org ]
		{"user@[example.org] ", false},        // CANONICALLY: user@[example.org]
		{"\"\r\n \"@example.org", false},      // CANONICALLY: " "@example.org
		{"\" \r\n \"@example.org", false},     // CANONICALLY: " "@example.org
		{"\"\r\n user\"@example.org", false},  // CANONICALLY: " user"@example.org
		{"\" \r\n user\"@example.org", false}, // CANONICALLY: " user"@example.org
		{"\"user\r\n \"@example.org", false},  // CANONICALLY: "user "@example.org
		{"\"user \r\n \"@example.org", false}, // CANONICALLY: "user "@example.org
		{"user@\r\n [example.org]", false},    // CANONICALLY: user@[example.org]
		{"user@ \r\n [example.org]", false},   // CANONICALLY: user@[example.org]
		{"user@[\r\n example.org]", false},    // CANONICALLY: user@[ example.org]
		{"user@[ \r\n example.org]", false},   // CANONICALLY: user@[ example.org]
		{"user@[example.org\r\n ]", false},    // CANONICALLY: user@[example.org ]
		{"user@[example.org \r\n ]", false},   // CANONICALLY: user@[example.org ]
		{"user@[example.org]\r\n ", false},    // CANONICALLY: user@[example.org]
		{"user@[example.org] \r\n ", false},   // CANONICALLY: user@[example.org]
		{"user@[ \\  ]", true},                // domain-literal is: "   "
		{"user@[ \\\t ]", true},               // domain-literal is: " \t "

		// Test Comment handling (it is invisible, and not part of the address):
		{"(comment)user@example.org", false},     // CANONICALLY: user@example.org
		{"user(comment)@example.org", false},     // CANONICALLY: user@example.org
		{"user@(comment)example.org", false},     // CANONICALLY: user@example.org
		{"user@example.org(comment)", false},     // CANONICALLY: user@example.org
		{"(comment)\"user\"@example.org", false}, // CANONICALLY: "user"@example.org
		{"\"(comment)user\"@example.org", true},  // Not a comment, but part of the local-part
		{"\"user(comment)\"@example.org", true},  // Not a comment, but part of the local-part
		{"\"user\"(comment)@example.org", false}, // CANONICALLY: "user"@example.org
		{"user@(comment)[example.org]", false},   // CANONICALLY: user@[example.org]
		{"user@[(comment)example.org]", true},    // Not a comment, but part of the domain-literal
		{"user@[example.org(comment)]", true},    // Not a comment, but part of the domain-literal
		{"user@[example.org](comment)", false},   // CANONICALLY: user@[example.org]

		// From: https://youtube.com/watch?v=mrGfahzt-4Q
		// Email vs Capitalism, or, Why We Can't Have Nice Things - Dylan Beattie - NDC Oslo 2023
		{"iron.man@avengers.com", true},
		{"spider-man@avengers.com", true},
		{"t'challa@avengers.com", true},
		{"rocket+groot@avengers.com", true},
		{`"Bruce 'The Hulk' Banner"@avengers.com`, true},
		{"vision@[IPv6:2001:db8:1ff::a0b:dbd0]", true},
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
