package pedantic

import (
	"regexp"
	"regexp/syntax"
	"testing"
)

const emailRegexString = `^` + emailString + `$`

func TestEmail(t *testing.T) {
	r, err := syntax.Parse(emailRegexString, syntax.Perl)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}

	t.Log(r.Simplify())

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
	}

	Email := regexp.MustCompile(emailRegexString)

	for _, tt := range tests {
		got := Email.MatchString(tt.s)
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
